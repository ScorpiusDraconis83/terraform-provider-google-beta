---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigee/Envgroup.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Apigee"
description: |-
  An `Environment group` in Apigee.
---

# google_apigee_envgroup

An `Environment group` in Apigee.


To get more information about Envgroup, see:

* [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.envgroups/create)
* How-to Guides
    * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)

## Example Usage - Apigee Environment Group Basic


```hcl
data "google_client_config" "current" {}

resource "google_compute_network" "apigee_network" {
  name = "apigee-network"
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = data.google_client_config.current.project
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [google_service_networking_connection.apigee_vpc_connection]
}

resource "google_apigee_envgroup" "env_grp" {
  name      = "my-envgroup"
  hostnames = ["abc.foo.com"]
  org_id    = google_apigee_organization.apigee_org.id
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  The resource ID of the environment group.

* `org_id` -
  (Required)
  The Apigee Organization associated with the Apigee environment group,
  in the format `organizations/{{org_name}}`.


* `hostnames` -
  (Optional)
  Hostnames of the environment group.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{org_id}}/envgroups/{{name}}`


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 30 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 30 minutes.

## Import


Envgroup can be imported using any of these accepted formats:

* `{{org_id}}/envgroups/{{name}}`
* `{{org_id}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Envgroup using one of the formats above. For example:

```tf
import {
  id = "{{org_id}}/envgroups/{{name}}"
  to = google_apigee_envgroup.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Envgroup can be imported using one of the formats above. For example:

```
$ terraform import google_apigee_envgroup.default {{org_id}}/envgroups/{{name}}
$ terraform import google_apigee_envgroup.default {{org_id}}/{{name}}
```
