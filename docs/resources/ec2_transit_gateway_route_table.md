
---
page_title: "awscc_ec2_transit_gateway_route_table Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::TransitGatewayRouteTable
---

# awscc_ec2_transit_gateway_route_table (Resource)

Resource Type definition for AWS::EC2::TransitGatewayRouteTable

## Example Usage

### Create Transit Gateway Route Table

To create a Transit Gateway Route Table, attach it to an existing Transit Gateway and tag it for easier resource management and tracking.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create Transit Gateway first as it's required
resource "awscc_ec2_transit_gateway" "example" {
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Transit Gateway Route Table
resource "awscc_ec2_transit_gateway_route_table" "example" {
  transit_gateway_id = awscc_ec2_transit_gateway.example.id

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `transit_gateway_id` (String) The ID of the transit gateway.

### Optional

- `tags` (Attributes List) Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `transit_gateway_route_table_id` (String) Transit Gateway Route Table primary identifier

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key of the associated tag key-value pair
- `value` (String) The value of the associated tag key-value pair

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ec2_transit_gateway_route_table.example
  id = "transit_gateway_route_table_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_transit_gateway_route_table.example "transit_gateway_route_table_id"
```
