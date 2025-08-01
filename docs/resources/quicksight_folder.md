
---
page_title: "awscc_quicksight_folder Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of the AWS::QuickSight::Folder Resource Type.
---

# awscc_quicksight_folder (Resource)

Definition of the AWS::QuickSight::Folder Resource Type.

## Example Usage

```terraform
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "awscc_quicksight_folder" "example" {
  aws_account_id = data.aws_caller_identity.current.account_id
  folder_id      = "analytics-team-folder"
  name           = "example"
  folder_type    = "SHARED"
  sharing_model  = "ACCOUNT"

  # Grant permissions to users and groups
  permissions = [
    {
      principal = "arn:aws:quicksight:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:user/default/analytics-admin"
      actions = [
        "quicksight:CreateFolder",
        "quicksight:DescribeFolder",
        "quicksight:UpdateFolder",
        "quicksight:DeleteFolder",
        "quicksight:CreateFolderMembership",
        "quicksight:DeleteFolderMembership",
        "quicksight:DescribeFolderPermissions",
        "quicksight:UpdateFolderPermissions"
      ]
    },
    {
      principal = "arn:aws:quicksight:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:group/default/analytics-team"
      actions = [
        "quicksight:DescribeFolder",
        "quicksight:CreateFolderMembership"
      ]
    }
  ]

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `aws_account_id` (String)
- `folder_id` (String)
- `folder_type` (String)
- `name` (String)
- `parent_folder_arn` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--permissions))
- `sharing_model` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) <p>The Amazon Resource Name (ARN) for the folder.</p>
- `created_time` (String) <p>The time that the folder was created.</p>
- `id` (String) Uniquely identifies the resource.
- `last_updated_time` (String) <p>The time that the folder was last updated.</p>

<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Optional:

- `actions` (List of String) <p>The IAM action to grant or revoke permissions on.</p>
- `principal` (String) <p>The Amazon Resource Name (ARN) of the principal. This can be one of the
            following:</p>
         <ul>
            <li>
               <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>
            </li>
            <li>
               <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>
            </li>
            <li>
               <p>The ARN of an Amazon Web Services account root: This is an IAM ARN rather than a QuickSight
                    ARN. Use this option only to share resources (templates) across Amazon Web Services accounts.
                    (This is less common.) </p>
            </li>
         </ul>


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) <p>Tag key.</p>
- `value` (String) <p>Tag value.</p>

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_quicksight_folder.example
  id = "aws_account_id|folder_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_quicksight_folder.example "aws_account_id|folder_id"
```
