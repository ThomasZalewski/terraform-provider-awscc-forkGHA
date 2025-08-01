
---
page_title: "awscc_networkmanager_connect_attachment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  AWS::NetworkManager::ConnectAttachment Resource Type Definition
---

# awscc_networkmanager_connect_attachment (Resource)

AWS::NetworkManager::ConnectAttachment Resource Type Definition

## Example Usage

### Core Network Connect Attachment with GRE Protocol

To create a connect attachment that links a transport attachment to a core network using GRE protocol, configure the core network ID, transport attachment ID, and edge location in your preferred AWS region.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
data "aws_region" "current" {}

# Example Connect Attachment - This is an example template, not intended for direct use
# since it requires existing core network, transport attachment, and VPC resources
resource "awscc_networkmanager_connect_attachment" "example" {
  # The ID of the core network to attach to - Replace with your actual core network ID
  core_network_id = "core-network-xxxx"

  # The ID of existing transport attachment (e.g., VPC attachment) - Replace with actual attachment ID
  transport_attachment_id = "attachment-xxxx"

  # The edge location for the attachment (AWS Region)
  edge_location = data.aws_region.current.name

  # Configuration options for the connect attachment
  options = {
    protocol = "GRE" # The tunnel protocol for the connect attachment
  }

  # Optional tags for the attachment
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `core_network_id` (String) ID of the CoreNetwork that the attachment will be attached to.
- `edge_location` (String) Edge location of the attachment.
- `options` (Attributes) Protocol options for connect attachment (see [below for nested schema](#nestedatt--options))
- `transport_attachment_id` (String) Id of transport attachment

### Optional

- `network_function_group_name` (String) The name of the network function group attachment.
- `proposed_network_function_group_change` (Attributes) The attachment to move from one network function group to another. (see [below for nested schema](#nestedatt--proposed_network_function_group_change))
- `proposed_segment_change` (Attributes) The attachment to move from one segment to another. (see [below for nested schema](#nestedatt--proposed_segment_change))
- `tags` (Attributes Set) Tags for the attachment. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `attachment_id` (String) The ID of the attachment.
- `attachment_policy_rule_number` (Number) The policy rule number associated with the attachment.
- `attachment_type` (String) The type of attachment.
- `core_network_arn` (String) The ARN of a core network.
- `created_at` (String) Creation time of the attachment.
- `id` (String) Uniquely identifies the resource.
- `owner_account_id` (String) The ID of the attachment account owner.
- `resource_arn` (String) The attachment resource ARN.
- `segment_name` (String) The name of the segment attachment.
- `state` (String) State of the attachment.
- `updated_at` (String) Last update time of the attachment.

<a id="nestedatt--options"></a>
### Nested Schema for `options`

Optional:

- `protocol` (String) Tunnel protocol for connect attachment


<a id="nestedatt--proposed_network_function_group_change"></a>
### Nested Schema for `proposed_network_function_group_change`

Optional:

- `attachment_policy_rule_number` (Number) The rule number in the policy document that applies to this change.
- `network_function_group_name` (String) The name of the network function group to change.
- `tags` (Attributes Set) The key-value tags that changed for the network function group. (see [below for nested schema](#nestedatt--proposed_network_function_group_change--tags))

<a id="nestedatt--proposed_network_function_group_change--tags"></a>
### Nested Schema for `proposed_network_function_group_change.tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.



<a id="nestedatt--proposed_segment_change"></a>
### Nested Schema for `proposed_segment_change`

Optional:

- `attachment_policy_rule_number` (Number) The rule number in the policy document that applies to this change.
- `segment_name` (String) The name of the segment to change.
- `tags` (Attributes Set) The list of key-value tags that changed for the segment. (see [below for nested schema](#nestedatt--proposed_segment_change--tags))

<a id="nestedatt--proposed_segment_change--tags"></a>
### Nested Schema for `proposed_segment_change.tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.



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
  to = awscc_networkmanager_connect_attachment.example
  id = "attachment_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_networkmanager_connect_attachment.example "attachment_id"
```
