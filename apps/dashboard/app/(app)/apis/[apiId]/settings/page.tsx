import { CopyButton } from "@/components/dashboard/copy-button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Code } from "@/components/ui/code";
import { getTenantId } from "@/lib/auth";
import { and, db, eq, isNull, schema, sql } from "@/lib/db";
import { notFound, redirect } from "next/navigation";
import { DeleteApi } from "./delete-api";
import { UpdateApiName } from "./update-api-name";
import { UpdateIpWhitelist } from "./update-ip-whitelist";

export const dynamic = "force-dynamic";

type Props = {
  params: {
    apiId: string;
  };
};

export default async function SettingsPage(props: Props) {
  const tenantId = getTenantId();

  const workspace = await db.query.workspaces.findFirst({
    where: (table, { and, eq, isNull }) =>
      and(eq(table.tenantId, tenantId), isNull(table.deletedAt)),
    with: {
      apis: {
        where: eq(schema.apis.id, props.params.apiId),
      },
    },
  });
  if (!workspace || workspace.tenantId !== tenantId) {
    return redirect("/new");
  }
  const api = workspace.apis.find((api) => api.id === props.params.apiId);
  if (!api) {
    return notFound();
  }
  const keys = await db
    .select({ count: sql<string>`count(*)` })
    .from(schema.keys)
    .where(and(eq(schema.keys.keyAuthId, api.keyAuthId!), isNull(schema.keys.deletedAt)))
    .then((rows) => Number.parseInt(rows.at(0)?.count ?? "0"));

  return (
    <div className="flex flex-col gap-8 mb-20 ">
      <UpdateApiName api={api} />
      <UpdateIpWhitelist api={api} workspace={workspace} />
      <Card>
        <CardHeader>
          <CardTitle>Api ID</CardTitle>
          <CardDescription>This is your api id. It's used in some API calls.</CardDescription>
        </CardHeader>
        <CardContent>
          <Code className="flex items-center justify-between w-full h-8 max-w-sm gap-4">
            <pre>{api.id}</pre>
            <div className="flex items-start justify-between gap-4">
              <CopyButton value={api.id} />
            </div>
          </Code>
        </CardContent>
      </Card>
      <DeleteApi api={api} keys={keys} />
    </div>
  );
}
