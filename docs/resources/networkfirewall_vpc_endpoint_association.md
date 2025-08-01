
---
page_title: "awscc_networkfirewall_vpc_endpoint_association Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource type definition for AWS::NetworkFirewall::VpcEndpointAssociation
---

# awscc_networkfirewall_vpc_endpoint_association (Resource)

Resource type definition for AWS::NetworkFirewall::VpcEndpointAssociation

## Example Usage

### Network Firewall VPC Association

Associates a Network Firewall with a VPC endpoint by configuring the subnet mapping and IP address type, enabling the firewall to protect network traffic in the specified VPC subnet.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Example of Network Firewall VPC Endpoint Association configuration
resource "awscc_networkfirewall_vpc_endpoint_association" "example" {
  # The ARN of an existing Network Firewall
  firewall_arn = "arn:aws:network-firewall:us-west-2:123456789012:firewall/example-firewall"

  # The ID of an existing VPC
  vpc_id = "vpc-1234567890abcdef0"

  # The subnet mapping configuration
  subnet_mapping = {
    subnet_id       = "subnet-1234567890abcdef0"
    ip_address_type = "IPV4"
  }

  description = "Example VPC endpoint association"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `firewall_arn` (String) A resource ARN.
- `subnet_mapping` (Attributes) (see [below for nested schema](#nestedatt--subnet_mapping))
- `vpc_id` (String)

### Optional

- `description` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `endpoint_id` (String) An endpoint Id.
- `id` (String) Uniquely identifies the resource.
- `vpc_endpoint_association_arn` (String) A resource ARN.
- `vpc_endpoint_association_id` (String)

<a id="nestedatt--subnet_mapping"></a>
### Nested Schema for `subnet_mapping`

Required:

- `subnet_id` (String) A SubnetId.

Optional:

- `ip_address_type` (String) A IPAddressType


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
  to = awscc_networkfirewall_vpc_endpoint_association.example
  id = "vpc_endpoint_association_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_networkfirewall_vpc_endpoint_association.example "vpc_endpoint_association_arn"
```
