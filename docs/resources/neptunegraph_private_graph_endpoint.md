
---
page_title: "awscc_neptunegraph_private_graph_endpoint Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::NeptuneGraph::PrivateGraphEndpoint resource creates an Amazon NeptuneGraph PrivateGraphEndpoint.
---

# awscc_neptunegraph_private_graph_endpoint (Resource)

The AWS::NeptuneGraph::PrivateGraphEndpoint resource creates an Amazon NeptuneGraph PrivateGraphEndpoint.

## Example Usage

### Create Private Endpoint for Neptune Graph

Creates a private graph endpoint for Neptune Graph with VPC configuration, connecting through specified subnets and secured by a custom security group that allows inbound traffic on port 8182.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
data "aws_region" "current" {}

# Example VPC for Neptune Graph
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "neptune-graph-vpc"
  }]
}

# Example Subnet 1
resource "awscc_ec2_subnet" "example1" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "neptune-graph-subnet-1"
  }]
}

# Example Subnet 2
resource "awscc_ec2_subnet" "example2" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Name"
    value = "neptune-graph-subnet-2"
  }]
}

# Security Group for Neptune Graph
resource "awscc_ec2_security_group" "example" {
  group_description = "Security group for Neptune Graph"
  vpc_id            = awscc_ec2_vpc.example.id
  security_group_ingress = [{
    ip_protocol = "tcp"
    from_port   = 8182
    to_port     = 8182
    cidr_ip     = "10.0.0.0/16"
  }]
  tags = [{
    key   = "Name"
    value = "neptune-graph-sg"
  }]
}

# Example Graph resource
resource "awscc_neptunegraph_graph" "example" {
  vector_search_configuration = {
    vector_search_dimension = 1536
  }
  provisioned_memory = 2
}

# Private Graph Endpoint
resource "awscc_neptunegraph_private_graph_endpoint" "example" {
  graph_identifier   = awscc_neptunegraph_graph.example.id
  vpc_id             = awscc_ec2_vpc.example.id
  subnet_ids         = [awscc_ec2_subnet.example1.id, awscc_ec2_subnet.example2.id]
  security_group_ids = [awscc_ec2_security_group.example.id]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `graph_identifier` (String) The auto-generated Graph Id assigned by the service.
- `vpc_id` (String) The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.

### Optional

- `security_group_ids` (List of String) The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
- `subnet_ids` (List of String) The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `private_graph_endpoint_identifier` (String) PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.

 For example, if GraphIdentifier is `g-12a3bcdef4` and VpcId is `vpc-0a12bc34567de8f90`, the generated PrivateGraphEndpointIdentifier will be `g-12a3bcdef4_vpc-0a12bc34567de8f90`
- `vpc_endpoint_id` (String) VPC endpoint that provides a private connection between the Graph and specified VPC.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_neptunegraph_private_graph_endpoint.example
  id = "private_graph_endpoint_identifier"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_neptunegraph_private_graph_endpoint.example "private_graph_endpoint_identifier"
```
