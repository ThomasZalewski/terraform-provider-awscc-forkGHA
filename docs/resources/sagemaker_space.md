
---
page_title: "awscc_sagemaker_space Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::SageMaker::Space
---

# awscc_sagemaker_space (Resource)

Resource Type definition for AWS::SageMaker::Space

## Example Usage

### SageMaker Space with Domain Configuration

Creates a SageMaker Space within a SageMaker Domain, configured with IAM authentication, default VPC networking, and customized Jupyter server settings.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
data "aws_region" "current" {}

# First, create a SageMaker domain
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type = "Service"
      identifiers = [
        "sagemaker.amazonaws.com"
      ]
    }
  }
}

resource "awscc_iam_role" "sagemaker_execution_role" {
  role_name                   = "sagemaker-space-example-role"
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json
  managed_policy_arns         = ["arn:aws:iam::aws:policy/AmazonSageMakerFullAccess"]
}

# Create a SageMaker domain first
resource "awscc_sagemaker_domain" "example" {
  domain_name = "example-domain"
  auth_mode   = "IAM"
  vpc_id      = aws_default_vpc.default.id
  subnet_ids  = [aws_default_subnet.default.id]

  default_user_settings = {
    execution_role = awscc_iam_role.sagemaker_execution_role.arn
  }

  default_space_settings = {
    execution_role = awscc_iam_role.sagemaker_execution_role.arn
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Use default VPC and subnet for simplicity
resource "aws_default_vpc" "default" {}

resource "aws_default_subnet" "default" {
  availability_zone = "${data.aws_region.current.name}a"
}

# Create a SageMaker space
resource "awscc_sagemaker_space" "example" {
  domain_id  = awscc_sagemaker_domain.example.id
  space_name = "example-space"

  space_display_name = "Example Space"

  space_settings = {
    jupyter_server_app_settings = {
      default_resource_spec = {
        instance_type = "system"
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain_id` (String) The ID of the associated Domain.
- `space_name` (String) A name for the Space.

### Optional

- `ownership_settings` (Attributes) (see [below for nested schema](#nestedatt--ownership_settings))
- `space_display_name` (String)
- `space_settings` (Attributes) A collection of settings. (see [below for nested schema](#nestedatt--space_settings))
- `space_sharing_settings` (Attributes) (see [below for nested schema](#nestedatt--space_sharing_settings))
- `tags` (Attributes List) A list of tags to apply to the space. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `space_arn` (String) The space Amazon Resource Name (ARN).
- `url` (String)

<a id="nestedatt--ownership_settings"></a>
### Nested Schema for `ownership_settings`

Optional:

- `owner_user_profile_name` (String)


<a id="nestedatt--space_settings"></a>
### Nested Schema for `space_settings`

Optional:

- `app_type` (String)
- `code_editor_app_settings` (Attributes) The CodeEditor app settings. (see [below for nested schema](#nestedatt--space_settings--code_editor_app_settings))
- `custom_file_systems` (Attributes List) (see [below for nested schema](#nestedatt--space_settings--custom_file_systems))
- `jupyter_lab_app_settings` (Attributes) The JupyterLab app settings. (see [below for nested schema](#nestedatt--space_settings--jupyter_lab_app_settings))
- `jupyter_server_app_settings` (Attributes) The Jupyter server's app settings. (see [below for nested schema](#nestedatt--space_settings--jupyter_server_app_settings))
- `kernel_gateway_app_settings` (Attributes) The kernel gateway app settings. (see [below for nested schema](#nestedatt--space_settings--kernel_gateway_app_settings))
- `remote_access` (String) This is a flag used to indicate if remote access is enabled.
- `space_managed_resources` (String) This is a flag used to indicate if space managed resources needs to be created.
- `space_storage_settings` (Attributes) Default storage settings for a space. (see [below for nested schema](#nestedatt--space_settings--space_storage_settings))

<a id="nestedatt--space_settings--code_editor_app_settings"></a>
### Nested Schema for `space_settings.code_editor_app_settings`

Optional:

- `app_lifecycle_management` (Attributes) (see [below for nested schema](#nestedatt--space_settings--code_editor_app_settings--app_lifecycle_management))
- `default_resource_spec` (Attributes) (see [below for nested schema](#nestedatt--space_settings--code_editor_app_settings--default_resource_spec))

<a id="nestedatt--space_settings--code_editor_app_settings--app_lifecycle_management"></a>
### Nested Schema for `space_settings.code_editor_app_settings.app_lifecycle_management`

Optional:

- `idle_settings` (Attributes) (see [below for nested schema](#nestedatt--space_settings--code_editor_app_settings--app_lifecycle_management--idle_settings))

<a id="nestedatt--space_settings--code_editor_app_settings--app_lifecycle_management--idle_settings"></a>
### Nested Schema for `space_settings.code_editor_app_settings.app_lifecycle_management.idle_settings`

Optional:

- `idle_timeout_in_minutes` (Number) The space idle timeout value set in minutes



<a id="nestedatt--space_settings--code_editor_app_settings--default_resource_spec"></a>
### Nested Schema for `space_settings.code_editor_app_settings.default_resource_spec`

Optional:

- `instance_type` (String) The instance type that the image version runs on.
- `lifecycle_config_arn` (String) The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
- `sage_maker_image_arn` (String) The ARN of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The ARN of the image version created on the instance.



<a id="nestedatt--space_settings--custom_file_systems"></a>
### Nested Schema for `space_settings.custom_file_systems`

Optional:

- `efs_file_system` (Attributes) (see [below for nested schema](#nestedatt--space_settings--custom_file_systems--efs_file_system))
- `fsx_lustre_file_system` (Attributes) (see [below for nested schema](#nestedatt--space_settings--custom_file_systems--fsx_lustre_file_system))
- `s3_file_system` (Attributes) (see [below for nested schema](#nestedatt--space_settings--custom_file_systems--s3_file_system))

<a id="nestedatt--space_settings--custom_file_systems--efs_file_system"></a>
### Nested Schema for `space_settings.custom_file_systems.efs_file_system`

Optional:

- `file_system_id` (String)


<a id="nestedatt--space_settings--custom_file_systems--fsx_lustre_file_system"></a>
### Nested Schema for `space_settings.custom_file_systems.fsx_lustre_file_system`

Optional:

- `file_system_id` (String)


<a id="nestedatt--space_settings--custom_file_systems--s3_file_system"></a>
### Nested Schema for `space_settings.custom_file_systems.s3_file_system`

Optional:

- `s3_uri` (String)



<a id="nestedatt--space_settings--jupyter_lab_app_settings"></a>
### Nested Schema for `space_settings.jupyter_lab_app_settings`

Optional:

- `app_lifecycle_management` (Attributes) (see [below for nested schema](#nestedatt--space_settings--jupyter_lab_app_settings--app_lifecycle_management))
- `code_repositories` (Attributes List) A list of CodeRepositories available for use with JupyterLab apps. (see [below for nested schema](#nestedatt--space_settings--jupyter_lab_app_settings--code_repositories))
- `default_resource_spec` (Attributes) (see [below for nested schema](#nestedatt--space_settings--jupyter_lab_app_settings--default_resource_spec))

<a id="nestedatt--space_settings--jupyter_lab_app_settings--app_lifecycle_management"></a>
### Nested Schema for `space_settings.jupyter_lab_app_settings.app_lifecycle_management`

Optional:

- `idle_settings` (Attributes) (see [below for nested schema](#nestedatt--space_settings--jupyter_lab_app_settings--app_lifecycle_management--idle_settings))

<a id="nestedatt--space_settings--jupyter_lab_app_settings--app_lifecycle_management--idle_settings"></a>
### Nested Schema for `space_settings.jupyter_lab_app_settings.app_lifecycle_management.idle_settings`

Optional:

- `idle_timeout_in_minutes` (Number) The space idle timeout value set in minutes



<a id="nestedatt--space_settings--jupyter_lab_app_settings--code_repositories"></a>
### Nested Schema for `space_settings.jupyter_lab_app_settings.code_repositories`

Optional:

- `repository_url` (String) A CodeRepository (valid URL) to be used within Jupyter's Git extension.


<a id="nestedatt--space_settings--jupyter_lab_app_settings--default_resource_spec"></a>
### Nested Schema for `space_settings.jupyter_lab_app_settings.default_resource_spec`

Optional:

- `instance_type` (String) The instance type that the image version runs on.
- `lifecycle_config_arn` (String) The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
- `sage_maker_image_arn` (String) The ARN of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The ARN of the image version created on the instance.



<a id="nestedatt--space_settings--jupyter_server_app_settings"></a>
### Nested Schema for `space_settings.jupyter_server_app_settings`

Optional:

- `default_resource_spec` (Attributes) (see [below for nested schema](#nestedatt--space_settings--jupyter_server_app_settings--default_resource_spec))
- `lifecycle_config_arns` (List of String) A list of LifecycleConfigArns available for use with JupyterServer apps.

<a id="nestedatt--space_settings--jupyter_server_app_settings--default_resource_spec"></a>
### Nested Schema for `space_settings.jupyter_server_app_settings.default_resource_spec`

Optional:

- `instance_type` (String) The instance type that the image version runs on.
- `lifecycle_config_arn` (String) The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
- `sage_maker_image_arn` (String) The ARN of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The ARN of the image version created on the instance.



<a id="nestedatt--space_settings--kernel_gateway_app_settings"></a>
### Nested Schema for `space_settings.kernel_gateway_app_settings`

Optional:

- `custom_images` (Attributes List) A list of custom SageMaker images that are configured to run as a KernelGateway app. (see [below for nested schema](#nestedatt--space_settings--kernel_gateway_app_settings--custom_images))
- `default_resource_spec` (Attributes) The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app. (see [below for nested schema](#nestedatt--space_settings--kernel_gateway_app_settings--default_resource_spec))
- `lifecycle_config_arns` (List of String) A list of LifecycleConfigArns available for use with KernelGateway apps.

<a id="nestedatt--space_settings--kernel_gateway_app_settings--custom_images"></a>
### Nested Schema for `space_settings.kernel_gateway_app_settings.custom_images`

Optional:

- `app_image_config_name` (String) The Name of the AppImageConfig.
- `image_name` (String) The name of the CustomImage. Must be unique to your account.
- `image_version_number` (Number) The version number of the CustomImage.


<a id="nestedatt--space_settings--kernel_gateway_app_settings--default_resource_spec"></a>
### Nested Schema for `space_settings.kernel_gateway_app_settings.default_resource_spec`

Optional:

- `instance_type` (String) The instance type that the image version runs on.
- `lifecycle_config_arn` (String) The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
- `sage_maker_image_arn` (String) The ARN of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The ARN of the image version created on the instance.



<a id="nestedatt--space_settings--space_storage_settings"></a>
### Nested Schema for `space_settings.space_storage_settings`

Optional:

- `ebs_storage_settings` (Attributes) Properties related to the space's Amazon Elastic Block Store volume. (see [below for nested schema](#nestedatt--space_settings--space_storage_settings--ebs_storage_settings))

<a id="nestedatt--space_settings--space_storage_settings--ebs_storage_settings"></a>
### Nested Schema for `space_settings.space_storage_settings.ebs_storage_settings`

Optional:

- `ebs_volume_size_in_gb` (Number) Size of the Amazon EBS volume in Gb




<a id="nestedatt--space_sharing_settings"></a>
### Nested Schema for `space_sharing_settings`

Optional:

- `sharing_type` (String)


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
  to = awscc_sagemaker_space.example
  id = "domain_id|space_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_sagemaker_space.example "domain_id|space_name"
```
