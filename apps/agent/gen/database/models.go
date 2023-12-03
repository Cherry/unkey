// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type ApisAuthType string

const (
	ApisAuthTypeKey ApisAuthType = "key"
	ApisAuthTypeJwt ApisAuthType = "jwt"
)

func (e *ApisAuthType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ApisAuthType(s)
	case string:
		*e = ApisAuthType(s)
	default:
		return fmt.Errorf("unsupported scan type for ApisAuthType: %T", src)
	}
	return nil
}

type NullApisAuthType struct {
	ApisAuthType ApisAuthType
	Valid        bool // Valid is true if ApisAuthType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullApisAuthType) Scan(value interface{}) error {
	if value == nil {
		ns.ApisAuthType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ApisAuthType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullApisAuthType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ApisAuthType), nil
}

type AuditLogsAction string

const (
	AuditLogsActionCreated AuditLogsAction = "created"
	AuditLogsActionUpdated AuditLogsAction = "updated"
	AuditLogsActionDeleted AuditLogsAction = "deleted"
)

func (e *AuditLogsAction) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuditLogsAction(s)
	case string:
		*e = AuditLogsAction(s)
	default:
		return fmt.Errorf("unsupported scan type for AuditLogsAction: %T", src)
	}
	return nil
}

type NullAuditLogsAction struct {
	AuditLogsAction AuditLogsAction
	Valid           bool // Valid is true if AuditLogsAction is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuditLogsAction) Scan(value interface{}) error {
	if value == nil {
		ns.AuditLogsAction, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuditLogsAction.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuditLogsAction) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuditLogsAction), nil
}

type AuditLogsActorType string

const (
	AuditLogsActorTypeUser AuditLogsActorType = "user"
	AuditLogsActorTypeKey  AuditLogsActorType = "key"
)

func (e *AuditLogsActorType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuditLogsActorType(s)
	case string:
		*e = AuditLogsActorType(s)
	default:
		return fmt.Errorf("unsupported scan type for AuditLogsActorType: %T", src)
	}
	return nil
}

type NullAuditLogsActorType struct {
	AuditLogsActorType AuditLogsActorType
	Valid              bool // Valid is true if AuditLogsActorType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuditLogsActorType) Scan(value interface{}) error {
	if value == nil {
		ns.AuditLogsActorType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuditLogsActorType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuditLogsActorType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuditLogsActorType), nil
}

type AuditLogsResourceType string

const (
	AuditLogsResourceTypeKey       AuditLogsResourceType = "key"
	AuditLogsResourceTypeApi       AuditLogsResourceType = "api"
	AuditLogsResourceTypeWorkspace AuditLogsResourceType = "workspace"
)

func (e *AuditLogsResourceType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuditLogsResourceType(s)
	case string:
		*e = AuditLogsResourceType(s)
	default:
		return fmt.Errorf("unsupported scan type for AuditLogsResourceType: %T", src)
	}
	return nil
}

type NullAuditLogsResourceType struct {
	AuditLogsResourceType AuditLogsResourceType
	Valid                 bool // Valid is true if AuditLogsResourceType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuditLogsResourceType) Scan(value interface{}) error {
	if value == nil {
		ns.AuditLogsResourceType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuditLogsResourceType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuditLogsResourceType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuditLogsResourceType), nil
}

type VercelBindingsEnvironment string

const (
	VercelBindingsEnvironmentDevelopment VercelBindingsEnvironment = "development"
	VercelBindingsEnvironmentPreview     VercelBindingsEnvironment = "preview"
	VercelBindingsEnvironmentProduction  VercelBindingsEnvironment = "production"
)

func (e *VercelBindingsEnvironment) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VercelBindingsEnvironment(s)
	case string:
		*e = VercelBindingsEnvironment(s)
	default:
		return fmt.Errorf("unsupported scan type for VercelBindingsEnvironment: %T", src)
	}
	return nil
}

type NullVercelBindingsEnvironment struct {
	VercelBindingsEnvironment VercelBindingsEnvironment
	Valid                     bool // Valid is true if VercelBindingsEnvironment is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVercelBindingsEnvironment) Scan(value interface{}) error {
	if value == nil {
		ns.VercelBindingsEnvironment, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VercelBindingsEnvironment.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVercelBindingsEnvironment) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VercelBindingsEnvironment), nil
}

type VercelBindingsResourceType string

const (
	VercelBindingsResourceTypeRootKey VercelBindingsResourceType = "rootKey"
	VercelBindingsResourceTypeApiId   VercelBindingsResourceType = "apiId"
)

func (e *VercelBindingsResourceType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VercelBindingsResourceType(s)
	case string:
		*e = VercelBindingsResourceType(s)
	default:
		return fmt.Errorf("unsupported scan type for VercelBindingsResourceType: %T", src)
	}
	return nil
}

type NullVercelBindingsResourceType struct {
	VercelBindingsResourceType VercelBindingsResourceType
	Valid                      bool // Valid is true if VercelBindingsResourceType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVercelBindingsResourceType) Scan(value interface{}) error {
	if value == nil {
		ns.VercelBindingsResourceType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VercelBindingsResourceType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVercelBindingsResourceType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VercelBindingsResourceType), nil
}

type WorkspacesPlan string

const (
	WorkspacesPlanFree       WorkspacesPlan = "free"
	WorkspacesPlanPro        WorkspacesPlan = "pro"
	WorkspacesPlanEnterprise WorkspacesPlan = "enterprise"
)

func (e *WorkspacesPlan) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkspacesPlan(s)
	case string:
		*e = WorkspacesPlan(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkspacesPlan: %T", src)
	}
	return nil
}

type NullWorkspacesPlan struct {
	WorkspacesPlan WorkspacesPlan
	Valid          bool // Valid is true if WorkspacesPlan is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWorkspacesPlan) Scan(value interface{}) error {
	if value == nil {
		ns.WorkspacesPlan, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WorkspacesPlan.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWorkspacesPlan) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WorkspacesPlan), nil
}

type Api struct {
	ID          string
	Name        string
	WorkspaceID string
	IpWhitelist sql.NullString
	AuthType    NullApisAuthType
	KeyAuthID   sql.NullString
}

type AuditLog struct {
	ID           string
	WorkspaceID  string
	Action       AuditLogsAction
	Description  string
	Time         time.Time
	ActorType    AuditLogsActorType
	ActorID      string
	ResourceType AuditLogsResourceType
	ResourceID   string
	Tags         json.RawMessage
}

type AuditLogChange struct {
	AuditLogID string
	Field      string
	Old        sql.NullString
	New        sql.NullString
}

type Key struct {
	ID                      string
	Hash                    string
	Start                   string
	OwnerID                 sql.NullString
	Meta                    sql.NullString
	CreatedAt               time.Time
	Expires                 sql.NullTime
	RatelimitType           sql.NullString
	RatelimitLimit          sql.NullInt32
	RatelimitRefillRate     sql.NullInt32
	RatelimitRefillInterval sql.NullInt32
	WorkspaceID             string
	ForWorkspaceID          sql.NullString
	Name                    sql.NullString
	RemainingRequests       sql.NullInt32
	KeyAuthID               string
	TotalUses               sql.NullInt64
	DeletedAt               sql.NullTime
}

type KeyAuth struct {
	ID          string
	WorkspaceID string
}

type VercelBinding struct {
	ID            string
	IntegrationID string
	WorkspaceID   string
	ProjectID     string
	Environment   VercelBindingsEnvironment
	ResourceID    string
	ResourceType  VercelBindingsResourceType
	VercelEnvID   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	LastEditedBy  string
}

type VercelIntegration struct {
	ID          string
	WorkspaceID string
	TeamID      sql.NullString
	AccessToken string
}

type Workspace struct {
	ID                    string
	Name                  string
	Slug                  sql.NullString
	TenantID              string
	StripeCustomerID      sql.NullString
	StripeSubscriptionID  sql.NullString
	Plan                  NullWorkspacesPlan
	QuotaMaxActiveKeys    sql.NullInt32
	UsageActiveKeys       sql.NullInt32
	QuotaMaxVerifications sql.NullInt32
	UsageVerifications    sql.NullInt32
	LastUsageUpdate       sql.NullTime
	BillingPeriodStart    sql.NullTime
	BillingPeriodEnd      sql.NullTime
	TrialEnds             sql.NullTime
	Features              json.RawMessage
	BetaFeatures          json.RawMessage
	Subscriptions         json.RawMessage
	PlanLockedUntil       sql.NullTime
	PlanChanged           sql.NullTime
}
