
---
page_title: "awscc_appsync_function_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  An example resource schema demonstrating some basic constructs and validation rules.
---

# awscc_appsync_function_configuration (Resource)

An example resource schema demonstrating some basic constructs and validation rules.

## Example Usage

### Create AppSync Function with DynamoDB Integration

Creates an AppSync function configuration with JavaScript runtime that integrates with DynamoDB, including all necessary supporting resources such as IAM roles, DynamoDB table, and AppSync API.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS region
data "aws_region" "current" {}

# Create AppSync API (no AWSCC equivalent yet)
resource "aws_appsync_graphql_api" "example" {
  name                = "example-api"
  authentication_type = "API_KEY"
  schema              = <<EOF
type Query {
  getExample(id: ID!): Example
}

type Example {
  id: ID!
  value: String
}
EOF

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Example DynamoDB table for the data source
resource "awscc_dynamodb_table" "example" {
  attribute_definitions = [
    {
      attribute_name = "id"
      attribute_type = "S"
    }
  ]
  key_schema = [
    {
      attribute_name = "id"
      key_type       = "HASH"
    }
  ]
  table_name   = "example-table"
  billing_mode = "PAY_PER_REQUEST"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for AppSync to access DynamoDB
resource "awscc_iam_role" "appsync_datasource_role" {
  role_name = "example-appsync-dynamodb-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "appsync.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "appsync_dynamodb_policy" {
  policy_name = "DynamoDBAccess"
  role_name   = awscc_iam_role.appsync_datasource_role.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "dynamodb:DeleteItem",
          "dynamodb:UpdateItem",
          "dynamodb:Query",
          "dynamodb:Scan"
        ]
        Resource = [
          awscc_dynamodb_table.example.arn,
          "${awscc_dynamodb_table.example.arn}/index/*"
        ]
      }
    ]
  })
}

# Create AppSync datasource
resource "aws_appsync_datasource" "example" {
  api_id = aws_appsync_graphql_api.example.id
  name   = "example_datasource"
  type   = "AMAZON_DYNAMODB"

  dynamodb_config {
    table_name = awscc_dynamodb_table.example.table_name
    region     = data.aws_region.current.name
  }

  service_role_arn = awscc_iam_role.appsync_datasource_role.arn
}

# Create AppSync function configuration using AWSCC
resource "awscc_appsync_function_configuration" "example" {
  api_id           = aws_appsync_graphql_api.example.id
  data_source_name = aws_appsync_datasource.example.name
  name             = "exampleFunction"
  description      = "Example AppSync Function"

  runtime = {
    name            = "APPSYNC_JS"
    runtime_version = "1.0.0"
  }

  code = <<EOF
export function request(ctx) {
  return {
    operation: 'GetItem',
    key: util.dynamodb.toMapValues({id: ctx.args.id})
  };
}

export function response(ctx) {
  return ctx.result;
}
EOF

  function_version = "2018-05-29"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_id` (String) The AWS AppSync GraphQL API that you want to attach using this function.
- `data_source_name` (String) The name of data source this function will attach.
- `name` (String) The name of the function.

### Optional

- `code` (String) The resolver code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
- `code_s3_location` (String) The Amazon S3 endpoint (where the code is located??).
- `description` (String) The function description.
- `function_version` (String) The version of the request mapping template. Currently, only the 2018-05-29 version of the template is supported.
- `max_batch_size` (Number) The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
- `request_mapping_template` (String) The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
- `request_mapping_template_s3_location` (String) Describes a Sync configuration for a resolver. Contains information on which Conflict Detection, as well as Resolution strategy, should be performed when the resolver is invoked.
- `response_mapping_template` (String) The Function response mapping template.
- `response_mapping_template_s3_location` (String) The location of a response mapping template in an Amazon S3 bucket. Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
- `runtime` (Attributes) Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. (see [below for nested schema](#nestedatt--runtime))
- `sync_config` (Attributes) Describes a Sync configuration for a resolver. Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked. (see [below for nested schema](#nestedatt--sync_config))

### Read-Only

- `function_arn` (String) The ARN for the function generated by the service
- `function_id` (String) The unique identifier for the function generated by the service
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--runtime"></a>
### Nested Schema for `runtime`

Optional:

- `name` (String) The name of the runtime to use. Currently, the only allowed value is APPSYNC_JS.
- `runtime_version` (String) The version of the runtime to use. Currently, the only allowed version is 1.0.0.


<a id="nestedatt--sync_config"></a>
### Nested Schema for `sync_config`

Optional:

- `conflict_detection` (String) The Conflict Detection strategy to use.
- `conflict_handler` (String) The Conflict Resolution strategy to perform in the event of a conflict.
- `lambda_conflict_handler_config` (Attributes) The LambdaConflictHandlerConfig when configuring LAMBDA as the Conflict Handler. (see [below for nested schema](#nestedatt--sync_config--lambda_conflict_handler_config))

<a id="nestedatt--sync_config--lambda_conflict_handler_config"></a>
### Nested Schema for `sync_config.lambda_conflict_handler_config`

Optional:

- `lambda_conflict_handler_arn` (String) The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_appsync_function_configuration.example
  id = "function_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_appsync_function_configuration.example "function_arn"
```
