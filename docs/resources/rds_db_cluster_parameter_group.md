---
page_title: "awscc_rds_db_cluster_parameter_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::RDS::DBClusterParameterGroup resource creates a new Amazon RDS DB cluster parameter group.
  For information about configuring parameters for Amazon Aurora DB clusters, see Working with parameter groups https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html in the Amazon Aurora User Guide.
  If you apply a parameter group to a DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.
  If you apply a change to parameter group associated with a stopped DB cluster, then the updated stack waits until the DB cluster is started.
---

# awscc_rds_db_cluster_parameter_group (Resource)

The ``AWS::RDS::DBClusterParameterGroup`` resource creates a new Amazon RDS DB cluster parameter group.
 For information about configuring parameters for Amazon Aurora DB clusters, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.
  If you apply a parameter group to a DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.
 If you apply a change to parameter group associated with a stopped DB cluster, then the updated stack waits until the DB cluster is started.

## Example Usage

### Basic example
To create a RDS db cluster parameter group
```terraform
resource "awscc_rds_db_cluster_parameter_group" "this" {
  db_cluster_parameter_group_name = "rds-db-cluster-pg"
  description                     = "RDS DB cluster parameter group"
  family                          = "aurora5.6"

  parameters = {
    character_set_server = "utf8"
    character_set_client = "latin2"
    time_zone            = "US/Eastern"
  }

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description for the DB cluster parameter group.
- `family` (String) The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.
  *Aurora MySQL* 
 Example: ``aurora-mysql5.7``, ``aurora-mysql8.0``
  *Aurora PostgreSQL* 
 Example: ``aurora-postgresql14``
  *RDS for MySQL* 
 Example: ``mysql8.0``
  *RDS for PostgreSQL* 
 Example: ``postgres13``
 To list all of the available parameter group families for a DB engine, use the following command:
  ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine <engine>`` 
 For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:
  ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine aurora-postgresql`` 
  The output contains duplicates.
  The following are the valid DB engine values:
  +   ``aurora-mysql`` 
  +   ``aurora-postgresql`` 
  +   ``mysql`` 
  +   ``postgres``
- `parameters` (String) Provides a list of parameters for the DB cluster parameter group.

### Optional

- `db_cluster_parameter_group_name` (String) The name of the DB cluster parameter group.
 Constraints:
  +  Must not match the name of an existing DB cluster parameter group.
  
  This value is stored as a lowercase string.
- `tags` (Attributes List) Tags to assign to the DB cluster parameter group. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").
- `value` (String) A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_rds_db_cluster_parameter_group.example
  id = "db_cluster_parameter_group_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_rds_db_cluster_parameter_group.example "db_cluster_parameter_group_name"
```
