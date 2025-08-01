---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_billing_billing_view Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A billing view is a container of cost & usage metadata.
---

# awscc_billing_billing_view (Resource)

A billing view is a container of cost & usage metadata.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `source_views` (Set of String) An array of strings that define the billing view's source.

### Optional

- `data_filter_expression` (Attributes) (see [below for nested schema](#nestedatt--data_filter_expression))
- `description` (String)
- `tags` (Attributes Set) An array of key-value pairs associated to the billing view being created. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `billing_view_type` (String)
- `created_at` (Number) The time when the billing view was created.
- `id` (String) Uniquely identifies the resource.
- `owner_account_id` (String)
- `updated_at` (Number) The time when the billing view was last updated.

<a id="nestedatt--data_filter_expression"></a>
### Nested Schema for `data_filter_expression`

Optional:

- `dimensions` (Attributes) (see [below for nested schema](#nestedatt--data_filter_expression--dimensions))
- `tags` (Attributes) (see [below for nested schema](#nestedatt--data_filter_expression--tags))

<a id="nestedatt--data_filter_expression--dimensions"></a>
### Nested Schema for `data_filter_expression.dimensions`

Optional:

- `key` (String)
- `values` (List of String)


<a id="nestedatt--data_filter_expression--tags"></a>
### Nested Schema for `data_filter_expression.tags`

Optional:

- `key` (String)
- `values` (List of String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_billing_billing_view.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_billing_billing_view.example "arn"
```
