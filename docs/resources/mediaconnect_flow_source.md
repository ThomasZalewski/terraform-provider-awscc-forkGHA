
---
page_title: "awscc_mediaconnect_flow_source Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::MediaConnect::FlowSource
---

# awscc_mediaconnect_flow_source (Resource)

Resource schema for AWS::MediaConnect::FlowSource

## Example Usage

### Configure MediaConnect Flow Source with Encryption

To create an encrypted MediaConnect Flow source with Zixi-push protocol, including static key encryption configuration and CIDR whitelist for secure content ingestion.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create a basic MediaConnect Flow first
resource "awscc_mediaconnect_flow" "example" {
  name              = "example-flow"
  availability_zone = "${data.aws_region.current.name}a"

  # Source configuration as per schema
  source = {
    description = "Example Source"
    name        = "example_source"
    protocol    = "zixi-push"
  }
}

# Create Flow Source
resource "awscc_mediaconnect_flow_source" "example" {
  flow_arn       = awscc_mediaconnect_flow.example.flow_arn
  name           = "example_source"
  description    = "Example Flow Source"
  protocol       = "zixi-push"
  whitelist_cidr = "10.0.0.0/16"
  ingest_port    = 2088

  # Example of static key encryption
  decryption = {
    algorithm  = "aes256"
    role_arn   = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/MediaConnectAccessRole"
    secret_arn = "arn:aws:secretsmanager:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:secret:example-key"
    key_type   = "static-key"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.
- `name` (String) The name of the source.

### Optional

- `decryption` (Attributes) The type of encryption that is used on the content ingested from this source. (see [below for nested schema](#nestedatt--decryption))
- `entitlement_arn` (String) The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
- `flow_arn` (String) The ARN of the flow.
- `gateway_bridge_source` (Attributes) The source configuration for cloud flows receiving a stream from a bridge. (see [below for nested schema](#nestedatt--gateway_bridge_source))
- `ingest_port` (Number) The port that the flow will be listening on for incoming content.
- `max_bitrate` (Number) The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
- `max_latency` (Number) The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
- `min_latency` (Number) The minimum latency in milliseconds.
- `protocol` (String) The protocol that is used by the source.
- `sender_control_port` (Number) The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
- `sender_ip_address` (String) The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
- `source_listener_address` (String) Source IP or domain name for SRT-caller protocol.
- `source_listener_port` (Number) Source port for SRT-caller protocol.
- `stream_id` (String) The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
- `vpc_interface_name` (String) The name of the VPC Interface this Source is configured with.
- `whitelist_cidr` (String) The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `ingest_ip` (String) The IP address that the flow will be listening on for incoming content.
- `source_arn` (String) The ARN of the source.
- `source_ingest_port` (String) The port that the flow will be listening on for incoming content.(ReadOnly)

<a id="nestedatt--decryption"></a>
### Nested Schema for `decryption`

Optional:

- `algorithm` (String) The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
- `constant_initialization_vector` (String) A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.
- `device_id` (String) The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.
- `key_type` (String) The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
- `region` (String) The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.
- `resource_id` (String) An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.
- `role_arn` (String) The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).
- `secret_arn` (String) The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.
- `url` (String) The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.


<a id="nestedatt--gateway_bridge_source"></a>
### Nested Schema for `gateway_bridge_source`

Optional:

- `bridge_arn` (String) The ARN of the bridge feeding this flow.
- `vpc_interface_attachment` (Attributes) The name of the VPC interface attachment to use for this bridge source. (see [below for nested schema](#nestedatt--gateway_bridge_source--vpc_interface_attachment))

<a id="nestedatt--gateway_bridge_source--vpc_interface_attachment"></a>
### Nested Schema for `gateway_bridge_source.vpc_interface_attachment`

Optional:

- `vpc_interface_name` (String) The name of the VPC interface to use for this resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_mediaconnect_flow_source.example
  id = "source_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_mediaconnect_flow_source.example "source_arn"
```
