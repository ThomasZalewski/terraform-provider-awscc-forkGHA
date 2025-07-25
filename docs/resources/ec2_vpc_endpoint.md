---
page_title: "awscc_ec2_vpc_endpoint Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS, an MKT Partner, or another AWS accounts in your organization. For more information, see the User Guide https://docs.aws.amazon.com/vpc/latest/privatelink/.
  An endpoint of type Interface establishes connections between the subnets in your VPC and an AWS-service, your own service, or a service hosted by another AWS-account. With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
  An endpoint of type gateway serves as a target for a route in your route table for traffic destined for S3 or DDB. You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to S3, see Why can't I connect to an S3 bucket using a gateway VPC endpoint? https://docs.aws.amazon.com/premiumsupport/knowledge-center/connect-s3-vpc-endpoint
  An endpoint of type GatewayLoadBalancer provides private connectivity between your VPC and virtual appliances from a service provider.
---

# awscc_ec2_vpc_endpoint (Resource)

Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS, an MKT Partner, or another AWS accounts in your organization. For more information, see the [User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/).
 An endpoint of type ``Interface`` establishes connections between the subnets in your VPC and an AWS-service, your own service, or a service hosted by another AWS-account. With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
 An endpoint of type ``gateway`` serves as a target for a route in your route table for traffic destined for S3 or DDB. You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to S3, see [Why can't I connect to an S3 bucket using a gateway VPC endpoint?](https://docs.aws.amazon.com/premiumsupport/knowledge-center/connect-s3-vpc-endpoint)
 An endpoint of type ``GatewayLoadBalancer`` provides private connectivity between your VPC and virtual appliances from a service provider.

## Example Usage

To create a VPC endpoint for S3
```terraform
#Basic

resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "awscc_ec2_vpc_endpoint" "s3" {
  vpc_id       = awscc_ec2_vpc.main.id
  service_name = "com.amazonaws.us-west-2.s3"
}
```

### Interface Endpoint Type
To create a VPC Endpoint with Interface type
```terraform
#Interface Endpoint Type

resource "awscc_ec2_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
}

resource "aws_security_group" "sg1" {
  name        = "allow_tls"
  description = "Allow TLS inbound traffic"
  vpc_id      = awscc_ec2_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = [awscc_ec2_vpc.main.cidr_block]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "allow_tls"
  }
}

resource "awscc_ec2_vpc_endpoint" "ec2" {
  vpc_id            = awscc_ec2_vpc.main.id
  service_name      = "com.amazonaws.us-west-2.ec2"
  vpc_endpoint_type = "Interface"

  security_group_ids = [
    aws_security_group.sg1.id,
  ]

  private_dns_enabled = true
}
```

### Gateway Load Balancer Endpoint Type
To create a VPC Endpoint with Gateway LB
```terraform
data "aws_caller_identity" "current" {}

resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "awscc_ec2_subnet" "main" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-1c"
}

resource "aws_internet_gateway" "ig" {
  vpc_id = awscc_ec2_vpc.main.id
}

resource "aws_lb" "test" {
  name               = "test-lb-tf"
  load_balancer_type = "gateway"
  subnets            = [awscc_ec2_subnet.main.id]
}

resource "aws_vpc_endpoint_service" "example" {
  acceptance_required        = false
  gateway_load_balancer_arns = [aws_lb.test.arn]
}

resource "awscc_ec2_vpc_endpoint" "example" {
  service_name      = aws_vpc_endpoint_service.example.service_name
  vpc_endpoint_type = aws_vpc_endpoint_service.example.service_type
  vpc_id            = awscc_ec2_vpc.main.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `vpc_id` (String) The ID of the VPC.

### Optional

- `dns_options` (Attributes) Describes the DNS options for an endpoint. (see [below for nested schema](#nestedatt--dns_options))
- `ip_address_type` (String) The supported IP address types.
- `policy_document` (String) An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
 For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. For example, if you have a JSON policy, you can convert it to YAML before including it in the YAML template, and CFNlong converts the policy to JSON format before calling the API actions for privatelink. Alternatively, you can include the JSON directly in the YAML, as shown in the following ``Properties`` section:
 ``Properties: VpcEndpointType: 'Interface' ServiceName: !Sub 'com.amazonaws.${AWS::Region}.logs' PolicyDocument: '{ "Version":"2012-10-17", "Statement": [{ "Effect":"Allow", "Principal":"*", "Action":["logs:Describe*","logs:Get*","logs:List*","logs:FilterLogEvents"], "Resource":"*" }] }'``
- `private_dns_enabled` (Boolean) Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
 To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
 This property is supported only for interface endpoints.
 Default: ``false``
- `resource_configuration_arn` (String) The Amazon Resource Name (ARN) of the resource configuration.
- `route_table_ids` (Set of String) The IDs of the route tables. Routing is supported only for gateway endpoints.
- `security_group_ids` (Set of String) The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
- `service_name` (String) The name of the endpoint service.
- `service_network_arn` (String) The Amazon Resource Name (ARN) of the service network.
- `service_region` (String) Describes a Region.
- `subnet_ids` (Set of String) The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
- `tags` (Attributes List) The tags to associate with the endpoint. (see [below for nested schema](#nestedatt--tags))
- `vpc_endpoint_type` (String) The type of endpoint.
 Default: Gateway

### Read-Only

- `creation_timestamp` (String)
- `dns_entries` (List of String)
- `id` (String) Uniquely identifies the resource.
- `network_interface_ids` (List of String)
- `vpc_endpoint_id` (String)

<a id="nestedatt--dns_options"></a>
### Nested Schema for `dns_options`

Optional:

- `dns_record_ip_type` (String) The DNS records created for the endpoint.
- `private_dns_only_for_inbound_resolver_endpoint` (String) Indicates whether to enable private DNS only for inbound endpoints. This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key of the tag.
 Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with ``aws:``.
- `value` (String) The value of the tag.
 Constraints: Tag values are case-sensitive and accept a maximum of 256 Unicode characters.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ec2_vpc_endpoint.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_vpc_endpoint.example "id"
```