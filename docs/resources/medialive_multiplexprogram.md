---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_medialive_multiplexprogram Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::MediaLive::Multiplexprogram
---

# awscc_medialive_multiplexprogram (Resource)

Resource schema for AWS::MediaLive::Multiplexprogram



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `multiplex_id` (String) The ID of the multiplex that the program belongs to.
- `multiplex_program_settings` (Attributes) The settings for this multiplex program. (see [below for nested schema](#nestedatt--multiplex_program_settings))
- `packet_identifiers_map` (Attributes) The packet identifier map for this multiplex program. (see [below for nested schema](#nestedatt--packet_identifiers_map))
- `pipeline_details` (Attributes List) Contains information about the current sources for the specified program in the specified multiplex. Keep in mind that each multiplex pipeline connects to both pipelines in a given source channel (the channel identified by the program). But only one of those channel pipelines is ever active at one time. (see [below for nested schema](#nestedatt--pipeline_details))
- `preferred_channel_pipeline` (String) The settings for this multiplex program.
- `program_name` (String) The name of the multiplex program.

### Read-Only

- `channel_id` (String) The MediaLive channel associated with the program.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--multiplex_program_settings"></a>
### Nested Schema for `multiplex_program_settings`

Optional:

- `preferred_channel_pipeline` (String) Indicates which pipeline is preferred by the multiplex for program ingest.
If set to \"PIPELINE_0\" or \"PIPELINE_1\" and an unhealthy ingest causes the multiplex to switch to the non-preferred pipeline,
it will switch back once that ingest is healthy again. If set to \"CURRENTLY_ACTIVE\",
it will not switch back to the other pipeline based on it recovering to a healthy state,
it will only switch if the active pipeline becomes unhealthy.
- `program_number` (Number) Unique program number.
- `service_descriptor` (Attributes) Transport stream service descriptor configuration for the Multiplex program. (see [below for nested schema](#nestedatt--multiplex_program_settings--service_descriptor))
- `video_settings` (Attributes) Program video settings configuration. (see [below for nested schema](#nestedatt--multiplex_program_settings--video_settings))

<a id="nestedatt--multiplex_program_settings--service_descriptor"></a>
### Nested Schema for `multiplex_program_settings.service_descriptor`

Optional:

- `provider_name` (String) Name of the provider.
- `service_name` (String) Name of the service.


<a id="nestedatt--multiplex_program_settings--video_settings"></a>
### Nested Schema for `multiplex_program_settings.video_settings`

Optional:

- `constant_bitrate` (Number) The constant bitrate configuration for the video encode.
When this field is defined, StatmuxSettings must be undefined.
- `statmux_settings` (Attributes) Statmux rate control settings.
When this field is defined, ConstantBitrate must be undefined. (see [below for nested schema](#nestedatt--multiplex_program_settings--video_settings--statmux_settings))

<a id="nestedatt--multiplex_program_settings--video_settings--statmux_settings"></a>
### Nested Schema for `multiplex_program_settings.video_settings.statmux_settings`

Optional:

- `maximum_bitrate` (Number) Maximum statmux bitrate.
- `minimum_bitrate` (Number) Minimum statmux bitrate.
- `priority` (Number) The purpose of the priority is to use a combination of the\nmultiplex rate control algorithm and the QVBR capability of the\nencoder to prioritize the video quality of some channels in a\nmultiplex over others.  Channels that have a higher priority will\nget higher video quality at the expense of the video quality of\nother channels in the multiplex with lower priority.




<a id="nestedatt--packet_identifiers_map"></a>
### Nested Schema for `packet_identifiers_map`

Optional:

- `audio_pids` (List of Number)
- `dvb_sub_pids` (List of Number)
- `dvb_teletext_pid` (Number)
- `etv_platform_pid` (Number)
- `etv_signal_pid` (Number)
- `klv_data_pids` (List of Number)
- `pcr_pid` (Number)
- `pmt_pid` (Number)
- `private_metadata_pid` (Number)
- `scte_27_pids` (List of Number)
- `scte_35_pid` (Number)
- `timed_metadata_pid` (Number)
- `video_pid` (Number)


<a id="nestedatt--pipeline_details"></a>
### Nested Schema for `pipeline_details`

Optional:

- `active_channel_pipeline` (String) Identifies the channel pipeline that is currently active for the pipeline (identified by PipelineId) in the multiplex.
- `pipeline_id` (String) Identifies a specific pipeline in the multiplex.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_medialive_multiplexprogram.example
  id = "program_name|multiplex_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_medialive_multiplexprogram.example "program_name|multiplex_id"
```
