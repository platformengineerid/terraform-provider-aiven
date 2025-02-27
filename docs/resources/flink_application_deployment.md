---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_flink_application_deployment Resource - terraform-provider-aiven"
subcategory: ""
description: |-
  The Flink Application Deployment resource allows the creation and management of Aiven Flink Application Deployments.
---

# aiven_flink_application_deployment (Resource)

The Flink Application Deployment resource allows the creation and management of Aiven Flink Application Deployments.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_id` (String) Application ID
- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `service_name` (String) Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `version_id` (String) ApplicationVersion ID

### Optional

- `parallelism` (Number) Flink Job parallelism
- `restart_enabled` (Boolean) Specifies whether a Flink Job is restarted in case it fails
- `starting_savepoint` (String) Job savepoint
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `created_at` (String) Application deployment creation time
- `created_by` (String) Application deployment creator
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `default` (String)
- `delete` (String)
- `read` (String)
- `update` (String)
