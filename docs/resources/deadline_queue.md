---
page_title: "awscc_deadline_queue Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Deadline::Queue Resource Type
---

# awscc_deadline_queue (Resource)

Definition of AWS::Deadline::Queue Resource Type

## Example Usage

```terraform
# Create S3 bucket for job attachments
resource "awscc_s3_bucket" "example" {
  bucket_name = "deadline-job-attachments-${random_id.bucket_suffix.hex}"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Generate random suffix for bucket name uniqueness
resource "random_id" "bucket_suffix" {
  byte_length = 4
}

resource "awscc_deadline_farm" "example" {
  display_name = "ExampleRenderFarm"
  description  = "Example Deadline Farm for queue demonstration"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create storage profiles for different operating systems
resource "awscc_deadline_storage_profile" "linux_storage" {
  display_name = "Linux Shared Storage"
  farm_id      = awscc_deadline_farm.example.farm_id
  os_family    = "LINUX"

  file_system_locations = [{
    name = "shared storage"
    path = "/mnt/shared"
    type = "SHARED"
    }, {
    name = "render assets"
    path = "/mnt/assets"
    type = "SHARED"
  }]
}

resource "awscc_deadline_storage_profile" "windows_storage" {
  display_name = "Windows Shared Storage"
  farm_id      = awscc_deadline_farm.example.farm_id
  os_family    = "WINDOWS"

  file_system_locations = [{
    name = "shared storage"
    path = "Z:\\"
    type = "SHARED"
    }, {
    name = "render assets"
    path = "Y:\\"
    type = "SHARED"
  }]
}

# Create an advanced Deadline Queue with job attachment settings
resource "awscc_deadline_queue" "example" {
  display_name          = "AdvancedRenderQueue"
  description           = "Advanced render queue with S3 job attachments and custom settings"
  farm_id               = awscc_deadline_farm.example.farm_id
  default_budget_action = "STOP_SCHEDULING_AND_COMPLETE_TASKS"

  # Configure job attachment settings for S3
  job_attachment_settings = {
    s3_bucket_name = awscc_s3_bucket.example.bucket_name
    root_prefix    = "job-attachments/"
  }

  # Configure job run-as user settings for POSIX systems
  job_run_as_user = {
    run_as = "QUEUE_CONFIGURED_USER"
    posix = {
      user  = "deadline-worker"
      group = "deadline-group"
    }
  }

  # Specify allowed storage profile IDs (dynamically referenced)
  allowed_storage_profile_ids = [
    awscc_deadline_storage_profile.linux_storage.storage_profile_id,
    awscc_deadline_storage_profile.windows_storage.storage_profile_id
  ]

  # Specify required file system location names
  required_file_system_location_names = [
    "shared storage",
    "render assets"
  ]

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String)
- `farm_id` (String)

### Optional

- `allowed_storage_profile_ids` (List of String)
- `default_budget_action` (String)
- `description` (String)
- `job_attachment_settings` (Attributes) (see [below for nested schema](#nestedatt--job_attachment_settings))
- `job_run_as_user` (Attributes) (see [below for nested schema](#nestedatt--job_run_as_user))
- `required_file_system_location_names` (List of String)
- `role_arn` (String)
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `queue_id` (String)

<a id="nestedatt--job_attachment_settings"></a>
### Nested Schema for `job_attachment_settings`

Optional:

- `root_prefix` (String)
- `s3_bucket_name` (String)


<a id="nestedatt--job_run_as_user"></a>
### Nested Schema for `job_run_as_user`

Optional:

- `posix` (Attributes) (see [below for nested schema](#nestedatt--job_run_as_user--posix))
- `run_as` (String)
- `windows` (Attributes) (see [below for nested schema](#nestedatt--job_run_as_user--windows))

<a id="nestedatt--job_run_as_user--posix"></a>
### Nested Schema for `job_run_as_user.posix`

Optional:

- `group` (String)
- `user` (String)


<a id="nestedatt--job_run_as_user--windows"></a>
### Nested Schema for `job_run_as_user.windows`

Optional:

- `password_arn` (String)
- `user` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_deadline_queue.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_deadline_queue.example "arn"
```