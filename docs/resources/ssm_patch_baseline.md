
---
page_title: "awscc_ssm_patch_baseline Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::SSM::PatchBaseline
---

# awscc_ssm_patch_baseline (Resource)

Resource Type definition for AWS::SSM::PatchBaseline

## Example Usage

### Windows Server Patch Baseline Configuration

Creates a high-compliance Windows Server patch baseline that automatically approves critical and important security updates after 7 days for Windows Server 2019.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create a patch baseline for Windows Server
resource "awscc_ssm_patch_baseline" "windows_baseline" {
  name        = "WindowsServerPatchBaseline"
  description = "Patch baseline for Windows Server"

  operating_system                  = "WINDOWS"
  approved_patches_compliance_level = "HIGH"

  # Global filters for the patch baseline
  global_filters = {
    patch_filters = [
      {
        key    = "PRODUCT"
        values = ["WindowsServer2019"]
      },
      {
        key    = "CLASSIFICATION"
        values = ["CriticalUpdates", "SecurityUpdates"]
      }
    ]
  }

  # Approval rules for patches
  approval_rules = {
    patch_rules = [
      {
        approve_after_days = 7
        compliance_level   = "HIGH"

        patch_filter_group = {
          patch_filters = [
            {
              key    = "MSRC_SEVERITY"
              values = ["Critical", "Important"]
            },
            {
              key    = "PATCH_SET"
              values = ["OS"]
            }
          ]
        }
      }
    ]
  }

  # Tag the resource
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the patch baseline.

### Optional

- `approval_rules` (Attributes) A set of rules defining the approval rules for a patch baseline. (see [below for nested schema](#nestedatt--approval_rules))
- `approved_patches` (List of String) A list of explicitly approved patches for the baseline.
- `approved_patches_compliance_level` (String) Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. The default value is UNSPECIFIED.
- `approved_patches_enable_non_security` (Boolean) Indicates whether the list of approved patches includes non-security updates that should be applied to the instances. The default value is 'false'. Applies to Linux instances only.
- `available_security_updates_compliance_status` (String) The compliance status for vendor recommended security updates that are not approved by this patch baseline.
- `default_baseline` (Boolean) Set the baseline as default baseline. Only registering to default patch baseline is allowed.
- `description` (String) The description of the patch baseline.
- `global_filters` (Attributes) A set of global filters used to include patches in the baseline. (see [below for nested schema](#nestedatt--global_filters))
- `operating_system` (String) Defines the operating system the patch baseline applies to. The Default value is WINDOWS.
- `patch_groups` (List of String) PatchGroups is used to associate instances with a specific patch baseline
- `rejected_patches` (List of String) A list of explicitly rejected patches for the baseline.
- `rejected_patches_action` (String) The action for Patch Manager to take on patches included in the RejectedPackages list.
- `sources` (Attributes List) Information about the patches to use to update the instances, including target operating systems and source repository. Applies to Linux instances only. (see [below for nested schema](#nestedatt--sources))
- `tags` (Attributes List) Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `patch_baseline_id` (String) The ID of the patch baseline.

<a id="nestedatt--approval_rules"></a>
### Nested Schema for `approval_rules`

Optional:

- `patch_rules` (Attributes List) (see [below for nested schema](#nestedatt--approval_rules--patch_rules))

<a id="nestedatt--approval_rules--patch_rules"></a>
### Nested Schema for `approval_rules.patch_rules`

Optional:

- `approve_after_days` (Number)
- `approve_until_date` (String)
- `compliance_level` (String)
- `enable_non_security` (Boolean)
- `patch_filter_group` (Attributes) The patch filter group that defines the criteria for the rule. (see [below for nested schema](#nestedatt--approval_rules--patch_rules--patch_filter_group))

<a id="nestedatt--approval_rules--patch_rules--patch_filter_group"></a>
### Nested Schema for `approval_rules.patch_rules.patch_filter_group`

Optional:

- `patch_filters` (Attributes List) (see [below for nested schema](#nestedatt--approval_rules--patch_rules--patch_filter_group--patch_filters))

<a id="nestedatt--approval_rules--patch_rules--patch_filter_group--patch_filters"></a>
### Nested Schema for `approval_rules.patch_rules.patch_filter_group.patch_filters`

Optional:

- `key` (String)
- `values` (List of String)





<a id="nestedatt--global_filters"></a>
### Nested Schema for `global_filters`

Optional:

- `patch_filters` (Attributes List) (see [below for nested schema](#nestedatt--global_filters--patch_filters))

<a id="nestedatt--global_filters--patch_filters"></a>
### Nested Schema for `global_filters.patch_filters`

Optional:

- `key` (String)
- `values` (List of String)



<a id="nestedatt--sources"></a>
### Nested Schema for `sources`

Optional:

- `configuration` (String)
- `name` (String)
- `products` (List of String)


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
  to = awscc_ssm_patch_baseline.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ssm_patch_baseline.example "id"
```
