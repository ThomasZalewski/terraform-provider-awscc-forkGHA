
---
page_title: "awscc_ivs_public_key Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::IVS::PublicKey
---

# awscc_ivs_public_key (Resource)

Resource Type definition for AWS::IVS::PublicKey

## Example Usage

### IVS Public Key Creation

Creates an IVS (Interactive Video Service) public key resource that can be used for stream authentication, featuring a name, RSA public key material in PEM format, and appropriate tagging.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create IVS Public Key resource
resource "awscc_ivs_public_key" "example" {
  name = "example-ivs-public-key"
  # This is an example RSA public key, you should replace it with your actual public key
  public_key_material = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAjeMQF6KuSCiiWF3Owc5C\nKq3DC3hSIgdaeBUAL5qQvRLaQ4/XEktOzucM64ueUxE8Fa6wITEWKHLT2B1Tc0Ni\nrCcATZqJB5xVcB5AMyGLb5H6HrVuPRiuf9ewXHbk+8FvhPe9cjWki5QV7ERm0Z6z\nM4RXBvhECRsxYt9bluyfod6MRQRlST/L13pkB6mYhxqZWA2t+r+04hdK6EP20MvG\nSVVzXKD2+Gtg7ZVBlH5bzU7pQc6w5jJr0hppAHY8gnHR31twhH92qpAIHjSYPfHg\nJqXzYYHlR5XQPvmEXbyHKryF2G0E8Su0XQqGOBa0bWpjEje1f9tD/vkAEE1jnR47\nKwIDAQAB\n-----END PUBLIC KEY-----"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `name` (String) Name of the public key to be imported. The value does not need to be unique.
- `public_key_material` (String) The public portion of a customer-generated key pair. This field is required to create the AWS::IVS::PublicKey resource.
- `tags` (Attributes Set) A list of key-value pairs that contain metadata for the asset model. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) Key-pair identifier.
- `fingerprint` (String) Key-pair identifier.
- `id` (String) Uniquely identifies the resource.

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
  to = awscc_ivs_public_key.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ivs_public_key.example "arn"
```
