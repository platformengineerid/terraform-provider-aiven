---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_opensearch_security_plugin_config Resource - terraform-provider-aiven"
subcategory: ""
description: |-
  The OpenSearch Security Plugin Config resource allows the creation and management of AivenOpenSearch Security Plugin config.
---

# aiven_opensearch_security_plugin_config (Resource)

The OpenSearch Security Plugin Config resource allows the creation and management of AivenOpenSearch Security Plugin config.

## Example Usage

```terraform
data "aiven_project" "foo" {
  project = "example_project"
}

resource "aiven_opensearch" "bar" {
  project                 = data.aiven_project.foo.project
  cloud_name              = "google-europe-west1"
  plan                    = "startup-4"
  service_name            = "example_service_name"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"
}

resource "aiven_opensearch_user" "foo" {
  service_name = aiven_opensearch.bar.service_name
  project      = data.aiven_project.foo.project
  username     = "user-example"
}

resource "aiven_opensearch_security_config" "foo" {
  project        = data.aiven_project.foo.project
  service_name   = aiven_opensearch.bar.service_name
  admin_password = "ThisIsATest123^=^"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `admin_password` (String, Sensitive) The password for the os-sec-admin user.
- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `service_name` (String) Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `admin_enabled` (Boolean) Whether the os-sec-admin user is enabled. This indicates whether the user management with the security plugin is enabled. This is always true when the os-sec-admin password was set at least once.
- `available` (Boolean) Whether the security plugin is available. This is always true for recently created services.
- `enabled` (Boolean) Whether the security plugin is enabled. This is always true for recently created services.
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `default` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import aiven_opensearch_security_plugin_config.foo project/service_name
```
