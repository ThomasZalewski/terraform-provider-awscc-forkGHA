---
page_title: "awscc_oam_sink Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Oam::Sink
---

# awscc_oam_sink (Resource)

Resource Type definition for AWS::Oam::Sink

## Example Usage

### All Accounts Organization 

Sample sink to connect that permits links to all accounts in an organization

```terraform
data "aws_organizations_organization" "example" {}

resource "awscc_oam_sink" "example" {
  name = "SampleSink"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect    = "Allow"
      Principal = "*"
      Resource  = "*"
      Action    = ["oam:CreateLink", "oam:UpdateLink"]
      Condition = {
        StringEquals = {
          "aws:PrincipalOrgID" = data.aws_organizations_organization.example.id
        }
        "ForAllValues:StringEquals" = {
          "oam:ResourceTypes" = [
            "AWS::CloudWatch::Metric",
            "AWS::Logs::LogGroup"
          ]
        }
      }
    }]
  })
}
```

### Individual Account

Sample sink that permits a link to an individual account

```terraform
resource "awscc_oam_sink" "example" {
  name = "SampleSink"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect   = "Allow"
      Resource = "*"
      Action = [
        "oam:CreateLink",
        "oam:UpdateLink"
      ]
      Principal = {
        AWS = ["1111111111111"]
      }
      Condition = {
        "ForAllValues:StringEquals" : {
          "oam:ResourceTypes" : [
            "AWS::CloudWatch::Metric",
            "AWS::Logs::LogGroup",
            "AWS::XRay::Trace"
          ]
        }
      }
    }]
  })
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the ObservabilityAccessManager Sink.

### Optional

- `policy` (String) The policy of this ObservabilityAccessManager Sink.
- `tags` (Map of String) Tags to apply to the sink

### Read-Only

- `arn` (String) The Amazon resource name (ARN) of the ObservabilityAccessManager Sink
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_oam_sink.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_oam_sink.example "arn"
```