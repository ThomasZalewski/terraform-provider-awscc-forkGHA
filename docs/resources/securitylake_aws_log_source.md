---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_securitylake_aws_log_source Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::SecurityLake::AwsLogSource
---

# awscc_securitylake_aws_log_source (Resource)

Resource Type definition for AWS::SecurityLake::AwsLogSource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `data_lake_arn` (String) The ARN for the data lake.
- `source_name` (String) The name for a AWS source. This must be a Regionally unique value.
- `source_version` (String) The version for a AWS source. This must be a Regionally unique value.

### Optional

- `accounts` (Set of String) AWS account where you want to collect logs from.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_securitylake_aws_log_source.example
  id = "source_name|source_version"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_securitylake_aws_log_source.example "source_name|source_version"
```
