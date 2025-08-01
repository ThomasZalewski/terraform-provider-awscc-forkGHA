---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_rds_db_proxy_target_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::RDS::DBProxyTargetGroup
---

# awscc_rds_db_proxy_target_group (Resource)

Resource schema for AWS::RDS::DBProxyTargetGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `db_proxy_name` (String) The identifier for the proxy.
- `target_group_name` (String) The identifier for the DBProxyTargetGroup

### Optional

- `connection_pool_configuration_info` (Attributes) (see [below for nested schema](#nestedatt--connection_pool_configuration_info))
- `db_cluster_identifiers` (List of String)
- `db_instance_identifiers` (List of String)

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `target_group_arn` (String) The Amazon Resource Name (ARN) representing the target group.

<a id="nestedatt--connection_pool_configuration_info"></a>
### Nested Schema for `connection_pool_configuration_info`

Optional:

- `connection_borrow_timeout` (Number) The number of seconds for a proxy to wait for a connection to become available in the connection pool.
- `init_query` (String) One or more SQL statements for the proxy to run when opening each new database connection.
- `max_connections_percent` (Number) The maximum size of the connection pool for each target in a target group.
- `max_idle_connections_percent` (Number) Controls how actively the proxy closes idle database connections in the connection pool.
- `session_pinning_filters` (List of String) Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_rds_db_proxy_target_group.example
  id = "target_group_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_rds_db_proxy_target_group.example "target_group_arn"
```
