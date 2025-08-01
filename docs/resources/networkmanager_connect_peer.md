
---
page_title: "awscc_networkmanager_connect_peer Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  AWS::NetworkManager::ConnectPeer Resource Type Definition.
---

# awscc_networkmanager_connect_peer (Resource)

AWS::NetworkManager::ConnectPeer Resource Type Definition.

## Example Usage

### Configure Network Manager Connect Peer with BGP

Creates an AWS Network Manager Connect Peer with BGP configuration that establishes connectivity between VPC and Core Network using GRE protocol, including all necessary networking components such as VPC, subnet, and network attachments.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create VPC and Subnet for the Connect Peer
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = [{
    key   = "Name"
    value = "connect-peer-example-vpc"
  }]
}

resource "awscc_ec2_subnet" "example" {
  vpc_id                  = awscc_ec2_vpc.example.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "${data.aws_region.current.name}a"
  map_public_ip_on_launch = false

  tags = [{
    key   = "Name"
    value = "connect-peer-example-subnet"
  }]
}

# Create the Global Network and Core Network
resource "awscc_networkmanager_global_network" "example" {
  description = "Example Global Network for Connect Peer"

  tags = [{
    key   = "Name"
    value = "connect-peer-example-global-network"
  }]
}

# Core Network policy document
data "aws_iam_policy_document" "core_network_policy" {
  statement {
    effect = "Allow"
    actions = [
      "network-manager:*",
      "ec2:DescribeVpcs",
      "ec2:DescribeSubnets"
    ]
    resources = ["*"]
  }
}

resource "awscc_networkmanager_core_network" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  policy_document   = jsonencode(jsondecode(data.aws_iam_policy_document.core_network_policy.json))
  description       = "Example Core Network for Connect Peer"

  tags = [{
    key   = "Name"
    value = "connect-peer-example-core-network"
  }]
}

# Create VPC attachment
resource "awscc_networkmanager_vpc_attachment" "example" {
  core_network_id = awscc_networkmanager_core_network.example.core_network_id
  vpc_arn         = "arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:vpc/${awscc_ec2_vpc.example.id}"
  subnet_arns     = ["arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:subnet/${awscc_ec2_subnet.example.id}"]

  tags = [{
    key   = "Name"
    value = "connect-peer-example-vpc-attachment"
  }]
}

# Create Connect Attachment
resource "awscc_networkmanager_connect_attachment" "example" {
  core_network_id         = awscc_networkmanager_core_network.example.core_network_id
  transport_attachment_id = awscc_networkmanager_vpc_attachment.example.attachment_id
  edge_location           = data.aws_region.current.name
  options = {
    protocol = "GRE"
  }

  tags = [{
    key   = "Name"
    value = "connect-peer-example-connect-attachment"
  }]
}

# Create Connect Peer
resource "awscc_networkmanager_connect_peer" "example" {
  connect_attachment_id = awscc_networkmanager_connect_attachment.example.attachment_id
  peer_address          = "10.0.0.1"
  bgp_options = {
    peer_asn = 65000
  }
  inside_cidr_blocks = ["169.254.6.0/29"]
  subnet_arn         = "arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:subnet/${awscc_ec2_subnet.example.id}"

  tags = [{
    key   = "Name"
    value = "connect-peer-example"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connect_attachment_id` (String) The ID of the attachment to connect.
- `peer_address` (String) The IP address of the Connect peer.

### Optional

- `bgp_options` (Attributes) Bgp options for connect peer. (see [below for nested schema](#nestedatt--bgp_options))
- `core_network_address` (String) The IP address of a core network.
- `inside_cidr_blocks` (List of String) The inside IP addresses used for a Connect peer configuration.
- `subnet_arn` (String) The subnet ARN for the connect peer.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `configuration` (Attributes) Configuration of the connect peer. (see [below for nested schema](#nestedatt--configuration))
- `connect_peer_id` (String) The ID of the Connect peer.
- `core_network_id` (String) The ID of the core network.
- `created_at` (String) Connect peer creation time.
- `edge_location` (String) The Connect peer Regions where edges are located.
- `id` (String) Uniquely identifies the resource.
- `state` (String) State of the connect peer.

<a id="nestedatt--bgp_options"></a>
### Nested Schema for `bgp_options`

Optional:

- `peer_asn` (Number)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `bgp_configurations` (Attributes List) (see [below for nested schema](#nestedatt--configuration--bgp_configurations))
- `core_network_address` (String) The IP address of a core network.
- `inside_cidr_blocks` (List of String) The inside IP addresses used for a Connect peer configuration.
- `peer_address` (String) The IP address of the Connect peer.
- `protocol` (String) The protocol used for a Connect peer configuration.

<a id="nestedatt--configuration--bgp_configurations"></a>
### Nested Schema for `configuration.bgp_configurations`

Read-Only:

- `core_network_address` (String) The address of a core network.
- `core_network_asn` (Number) The ASN of the Coret Network.
- `peer_address` (String) The address of a core network Connect peer.
- `peer_asn` (Number) The ASN of the Connect peer.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_networkmanager_connect_peer.example
  id = "connect_peer_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_networkmanager_connect_peer.example "connect_peer_id"
```
