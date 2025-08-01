---
page_title: "awscc_ec2_prefix_list Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema of AWS::EC2::PrefixList Type
---

# awscc_ec2_prefix_list (Resource)

Resource schema of AWS::EC2::PrefixList Type

## Example Usage

### Create prefix list with IPv4 addresses
```terraform
resource "awscc_ec2_prefix_list" "example_prefix_list" {
  address_family   = "IPv4"
  max_entries      = 5
  prefix_list_name = "example-ipv4-prefix-list"
  entries = [
    {
      cidr        = "10.10.0.0/16"
      description = "example network"
    },
    {
      cidr        = "192.168.2.8/32"
      description = "example host"
    }
  ]
}
```

### Create prefix list with IPv6 addresses
```terraform
resource "awscc_ec2_prefix_list" "example_ipv6_prefix_list" {
  address_family   = "IPv6"
  max_entries      = 5
  prefix_list_name = "example-ipv6-prefix-list"
  entries = [
    {
      cidr        = "2001:db8::/32"
      description = "example network"
    },
    {
      cidr        = "2001:db8:abcd:0012::0/128"
      description = "example host"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address_family` (String) Ip Version of Prefix List.
- `prefix_list_name` (String) Name of Prefix List.

### Optional

- `entries` (Attributes List) Entries of Prefix List. (see [below for nested schema](#nestedatt--entries))
- `max_entries` (Number) Max Entries of Prefix List.
- `tags` (Attributes List) Tags for Prefix List (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the Prefix List.
- `id` (String) Uniquely identifies the resource.
- `owner_id` (String) Owner Id of Prefix List.
- `prefix_list_id` (String) Id of Prefix List.
- `version` (Number) Version of Prefix List.

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Optional:

- `cidr` (String)
- `description` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ec2_prefix_list.example
  id = "prefix_list_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_prefix_list.example "prefix_list_id"
```