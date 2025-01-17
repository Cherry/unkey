package ratelimit

import (
	"time"

	ratelimitv1 "github.com/unkeyed/unkey/apps/agent/gen/proto/ratelimit/v1"
	"github.com/unkeyed/unkey/apps/agent/pkg/batch"
	"github.com/unkeyed/unkey/apps/agent/pkg/cluster"
	"github.com/unkeyed/unkey/apps/agent/pkg/logging"
	"github.com/unkeyed/unkey/apps/agent/pkg/metrics"
	"github.com/unkeyed/unkey/apps/agent/pkg/prometheus"
	"github.com/unkeyed/unkey/apps/agent/pkg/ratelimit"
	"github.com/unkeyed/unkey/apps/agent/pkg/repeat"
)

type service struct {
	logger      logging.Logger
	ratelimiter ratelimit.Ratelimiter
	cluster     cluster.Cluster

	batcher            *batch.BatchProcessor[*ratelimitv1.PushPullEvent]
	syncBuffer         chan syncWithOriginRequest
	metrics            metrics.Metrics
	consistencyChecker *consistencyChecker
}

type Config struct {
	Logger  logging.Logger
	Metrics metrics.Metrics
	Cluster cluster.Cluster
}

func New(cfg Config) (Service, error) {
	aggregateMaxBufferSize := 100000

	s := &service{
		logger:             cfg.Logger,
		ratelimiter:        ratelimit.NewFixedWindow(cfg.Logger.With().Str("ratelimiter", "fixedWindow").Logger()),
		cluster:            cfg.Cluster,
		metrics:            cfg.Metrics,
		consistencyChecker: newConsistencyChecker(cfg.Logger),
		syncBuffer:         make(chan syncWithOriginRequest, 1000),
	}

	if cfg.Cluster != nil {

		s.batcher = batch.New(batch.Config[*ratelimitv1.PushPullEvent]{
			BatchSize:     50,
			BufferSize:    aggregateMaxBufferSize,
			FlushInterval: time.Millisecond * 100,
			Flush:         s.aggregateByOrigin,
			Consumers:     1,
		})

		// Process the individual requests to the origin and update local state
		// We're using 32 goroutines to parallelise the network requests'
		for range 32 {
			go func() {
				for req := range s.syncBuffer {
					s.syncWithOrigin(req)
				}
			}()
		}

		repeat.Every(time.Second, func() {
			prometheus.ChannelBuffer.With(map[string]string{
				"id": "pushpull.aggregateByOrigin",
			}).Set(float64(s.batcher.Size()) / float64(aggregateMaxBufferSize))

			prometheus.ChannelBuffer.With(map[string]string{
				"id": "pushpull.syncWithOrigin",
			}).Set(float64(len(s.syncBuffer)) / float64(cap(s.syncBuffer)))

		})

	}

	return s, nil
}
