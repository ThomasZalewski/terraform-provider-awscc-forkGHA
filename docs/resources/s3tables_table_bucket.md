
---
page_title: "awscc_s3tables_table_bucket Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Creates an Amazon S3 Tables table bucket in the same AWS Region where you create the AWS CloudFormation stack.
---

# awscc_s3tables_table_bucket (Resource)

Creates an Amazon S3 Tables table bucket in the same AWS Region where you create the AWS CloudFormation stack.

## Example Usage

### S3Tables Table Bucket with Automated Cleanup

Creates an S3Tables table bucket that automatically removes unreferenced files after 7 days and noncurrent files after 30 days, with the bucket name incorporating the AWS account ID for uniqueness.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Data sources to get AWS account ID
data "aws_caller_identity" "current" {}

resource "awscc_s3tables_table_bucket" "example" {
  table_bucket_name = "my-s3tables-table-bucket-${data.aws_caller_identity.current.account_id}"

  unreferenced_file_removal = {
    status            = "Enabled"
    unreferenced_days = 7
    noncurrent_days   = 30
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `table_bucket_name` (String) A name for the table bucket.

### Optional

- `encryption_configuration` (Attributes) Specifies encryption settings for the table bucket (see [below for nested schema](#nestedatt--encryption_configuration))
- `unreferenced_file_removal` (Attributes) Settings governing the Unreferenced File Removal maintenance action. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots. (see [below for nested schema](#nestedatt--unreferenced_file_removal))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `table_bucket_arn` (String) The Amazon Resource Name (ARN) of the specified table bucket.

<a id="nestedatt--encryption_configuration"></a>
### Nested Schema for `encryption_configuration`

Optional:

- `kms_key_arn` (String) ARN of the KMS key to use for encryption
- `sse_algorithm` (String) Server-side encryption algorithm


<a id="nestedatt--unreferenced_file_removal"></a>
### Nested Schema for `unreferenced_file_removal`

Optional:

- `noncurrent_days` (Number) S3 permanently deletes noncurrent objects after the number of days specified by the NoncurrentDays property.
- `status` (String) Indicates whether the Unreferenced File Removal maintenance action is enabled.
- `unreferenced_days` (Number) For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_s3tables_table_bucket.example
  id = "table_bucket_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_s3tables_table_bucket.example "table_bucket_arn"
```
