---
page_title: "awscc_apigatewayv2_domain_name Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::ApiGatewayV2::DomainName resource specifies a custom domain name for your API in Amazon API Gateway (API Gateway).
  You can use a custom domain name to provide a URL that's more intuitive and easier to recall. For more information about using custom domain names, see Set up Custom Domain Name for an API in API Gateway https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html in the API Gateway Developer Guide.
---

# awscc_apigatewayv2_domain_name (Resource)

The ``AWS::ApiGatewayV2::DomainName`` resource specifies a custom domain name for your API in Amazon API Gateway (API Gateway). 
 You can use a custom domain name to provide a URL that's more intuitive and easier to recall. For more information about using custom domain names, see [Set up Custom Domain Name for an API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html) in the *API Gateway Developer Guide*.

## Example Usage

### Basic

```terraform
resource "awscc_apigatewayv2_domain_name" "example" {
  domain_name = "example.com"
  domain_name_configurations = [{
    certificate_arn = var.acm_certificate_arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain_name` (String) The custom domain name for your API in Amazon API Gateway. Uppercase letters and the underscore (``_``) character are not supported.

### Optional

- `domain_name_configurations` (Attributes List) The domain name configurations. (see [below for nested schema](#nestedatt--domain_name_configurations))
- `mutual_tls_authentication` (Attributes) The mutual TLS authentication configuration for a custom domain name. (see [below for nested schema](#nestedatt--mutual_tls_authentication))
- `routing_mode` (String)
- `tags` (Map of String) The collection of tags associated with a domain name.

### Read-Only

- `domain_name_arn` (String)
- `id` (String) Uniquely identifies the resource.
- `regional_domain_name` (String)
- `regional_hosted_zone_id` (String)

<a id="nestedatt--domain_name_configurations"></a>
### Nested Schema for `domain_name_configurations`

Optional:

- `certificate_arn` (String) An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name. AWS Certificate Manager is the only supported source.
- `certificate_name` (String) The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.
- `endpoint_type` (String) The endpoint type.
- `ip_address_type` (String)
- `ownership_verification_certificate_arn` (String) The Amazon resource name (ARN) for the public certificate issued by ACMlong. This ARN is used to validate custom domain ownership. It's required only if you configure mutual TLS and use either an ACM-imported or a private CA certificate ARN as the regionalCertificateArn.
- `security_policy` (String) The Transport Layer Security (TLS) version of the security policy for this domain name. The valid values are ``TLS_1_0`` and ``TLS_1_2``.


<a id="nestedatt--mutual_tls_authentication"></a>
### Nested Schema for `mutual_tls_authentication`

Optional:

- `truststore_uri` (String) An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, ``s3://bucket-name/key-name``. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version. To update the truststore, you must have permissions to access the S3 object.
- `truststore_version` (String) The version of the S3 object that contains your truststore. To specify a version, you must have versioning enabled for the S3 bucket.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_apigatewayv2_domain_name.example
  id = "domain_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_apigatewayv2_domain_name.example "domain_name"
```
