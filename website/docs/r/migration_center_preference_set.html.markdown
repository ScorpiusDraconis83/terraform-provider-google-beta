---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/migrationcenter/PreferenceSet.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Migration Center"
description: |-
  Manages the PreferenceSet resource.
---

# google_migration_center_preference_set

Manages the PreferenceSet resource.


To get more information about PreferenceSet, see:

* [API documentation](https://cloud.google.com/migration-center/docs/reference/rest/v1)
* How-to Guides
    * [Managing Migration Preferences](https://cloud.google.com/migration-center/docs/migration-preferences)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=preference_set_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Preference Set Basic


```hcl
resource "google_migration_center_preference_set" "default" {
  location          = "us-central1"
  preference_set_id = "preference-set-test"
  description       = "Terraform integration test description"
  display_name      = "Terraform integration test display"
  virtual_machine_preferences {
    vmware_engine_preferences {
      cpu_overcommit_ratio = 1.5
    }
    sizing_optimization_strategy = "SIZING_OPTIMIZATION_STRATEGY_SAME_AS_SOURCE"
    target_product = "COMPUTE_MIGRATION_TARGET_PRODUCT_COMPUTE_ENGINE"
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=preference_set_full&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Preference Set Full


```hcl
resource "google_migration_center_preference_set" "default" {
  location          = "us-central1"
  preference_set_id = "preference-set-test"
  description       = "Terraform integration test description"
  display_name      = "Terraform integration test display"
  virtual_machine_preferences {
    vmware_engine_preferences {
      cpu_overcommit_ratio = 1.5
      storage_deduplication_compression_ratio = 1.3
      commitment_plan                         = "ON_DEMAND"
    }
    sizing_optimization_strategy = "SIZING_OPTIMIZATION_STRATEGY_SAME_AS_SOURCE"
    target_product = "COMPUTE_MIGRATION_TARGET_PRODUCT_COMPUTE_ENGINE"
    commitment_plan = "COMMITMENT_PLAN_ONE_YEAR"
    region_preferences {
      preferred_regions = ["us-central1"]
    }
    sole_tenancy_preferences {
      commitment_plan         = "ON_DEMAND"
      cpu_overcommit_ratio    = 1.2
      host_maintenance_policy = "HOST_MAINTENANCE_POLICY_DEFAULT"
      node_types {
        node_name = "tf-test"
      }
    }
    compute_engine_preferences {
      license_type = "LICENSE_TYPE_BRING_YOUR_OWN_LICENSE"
      machine_preferences {
        allowed_machine_series {
          code = "C3"
        }
      }
    }
  }
}
```

## Argument Reference

The following arguments are supported:


* `location` -
  (Required)
  Part of `parent`. See documentation of `projectsId`.

* `preference_set_id` -
  (Required)
  Required. User specified ID for the preference set. It will become the last component of the preference set name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.


* `display_name` -
  (Optional)
  User-friendly display name. Maximum length is 63 characters.

* `description` -
  (Optional)
  A description of the preference set.

* `virtual_machine_preferences` -
  (Optional)
  VirtualMachinePreferences enables you to create sets of assumptions, for example, a geographical location and pricing track, for your migrated virtual machines. The set of preferences influence recommendations for migrating virtual machine assets.
  Structure is [documented below](#nested_virtual_machine_preferences).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_virtual_machine_preferences"></a>The `virtual_machine_preferences` block supports:

* `target_product` -
  (Optional)
  Target product for assets using this preference set. Specify either target product or business goal, but not both. Possible values: `COMPUTE_MIGRATION_TARGET_PRODUCT_UNSPECIFIED`, `COMPUTE_MIGRATION_TARGET_PRODUCT_COMPUTE_ENGINE`, `COMPUTE_MIGRATION_TARGET_PRODUCT_VMWARE_ENGINE`, `COMPUTE_MIGRATION_TARGET_PRODUCT_SOLE_TENANCY`

* `region_preferences` -
  (Optional)
  The user preferences relating to target regions.
  Structure is [documented below](#nested_virtual_machine_preferences_region_preferences).

* `commitment_plan` -
  (Optional)
  Commitment plan to consider when calculating costs for virtual machine insights and recommendations. If you are unsure which value to set, a 3 year commitment plan is often a good value to start with. Possible values: `COMMITMENT_PLAN_UNSPECIFIED`, `COMMITMENT_PLAN_NONE`, `COMMITMENT_PLAN_ONE_YEAR`, `COMMITMENT_PLAN_THREE_YEARS`

* `sizing_optimization_strategy` -
  (Optional)
  Sizing optimization strategy specifies the preferred strategy used when extrapolating usage data to calculate insights and recommendations for a virtual machine. If you are unsure which value to set, a moderate sizing optimization strategy is often a good value to start with. Possible values: `SIZING_OPTIMIZATION_STRATEGY_UNSPECIFIED`, `SIZING_OPTIMIZATION_STRATEGY_SAME_AS_SOURCE`, `SIZING_OPTIMIZATION_STRATEGY_MODERATE`, `SIZING_OPTIMIZATION_STRATEGY_AGGRESSIVE`

* `compute_engine_preferences` -
  (Optional)
  The user preferences relating to Compute Engine target platform.
  Structure is [documented below](#nested_virtual_machine_preferences_compute_engine_preferences).

* `vmware_engine_preferences` -
  (Optional)
  The user preferences relating to Google Cloud VMware Engine target platform.
  Structure is [documented below](#nested_virtual_machine_preferences_vmware_engine_preferences).

* `sole_tenancy_preferences` -
  (Optional)
  Preferences concerning Sole Tenancy nodes and VMs.
  Structure is [documented below](#nested_virtual_machine_preferences_sole_tenancy_preferences).


<a name="nested_virtual_machine_preferences_region_preferences"></a>The `region_preferences` block supports:

* `preferred_regions` -
  (Optional)
  A list of preferred regions, ordered by the most preferred region first. Set only valid Google Cloud region names. See https://cloud.google.com/compute/docs/regions-zones for available regions.

<a name="nested_virtual_machine_preferences_compute_engine_preferences"></a>The `compute_engine_preferences` block supports:

* `machine_preferences` -
  (Optional)
  The type of machines to consider when calculating virtual machine migration insights and recommendations. Not all machine types are available in all zones and regions.
  Structure is [documented below](#nested_virtual_machine_preferences_compute_engine_preferences_machine_preferences).

* `license_type` -
  (Optional)
  License type to consider when calculating costs for virtual machine insights and recommendations. If unspecified, costs are calculated based on the default licensing plan. Possible values: `LICENSE_TYPE_UNSPECIFIED`, `LICENSE_TYPE_DEFAULT`, `LICENSE_TYPE_BRING_YOUR_OWN_LICENSE`


<a name="nested_virtual_machine_preferences_compute_engine_preferences_machine_preferences"></a>The `machine_preferences` block supports:

* `allowed_machine_series` -
  (Optional)
  Compute Engine machine series to consider for insights and recommendations. If empty, no restriction is applied on the machine series.
  Structure is [documented below](#nested_virtual_machine_preferences_compute_engine_preferences_machine_preferences_allowed_machine_series).


<a name="nested_virtual_machine_preferences_compute_engine_preferences_machine_preferences_allowed_machine_series"></a>The `allowed_machine_series` block supports:

* `code` -
  (Optional)
  Code to identify a Compute Engine machine series. Consult https://cloud.google.com/compute/docs/machine-resource#machine_type_comparison for more details on the available series.

<a name="nested_virtual_machine_preferences_vmware_engine_preferences"></a>The `vmware_engine_preferences` block supports:

* `cpu_overcommit_ratio` -
  (Optional)
  CPU overcommit ratio. Acceptable values are between 1.0 and 8.0, with 0.1 increment.

* `memory_overcommit_ratio` -
  (Optional)
  Memory overcommit ratio. Acceptable values are 1.0, 1.25, 1.5, 1.75 and 2.0.

* `storage_deduplication_compression_ratio` -
  (Optional)
  The Deduplication and Compression ratio is based on the logical (Used Before) space required to store data before applying deduplication and compression, in relation to the physical (Used After) space required after applying deduplication and compression. Specifically, the ratio is the Used Before space divided by the Used After space. For example, if the Used Before space is 3 GB, but the physical Used After space is 1 GB, the deduplication and compression ratio is 3x. Acceptable values are between 1.0 and 4.0.

* `commitment_plan` -
  (Optional)
  Commitment plan to consider when calculating costs for virtual machine insights and recommendations. If you are unsure which value to set, a 3 year commitment plan is often a good value to start with. Possible values: `COMMITMENT_PLAN_UNSPECIFIED`, `ON_DEMAND`, `COMMITMENT_1_YEAR_MONTHLY_PAYMENTS`, `COMMITMENT_3_YEAR_MONTHLY_PAYMENTS`, `COMMITMENT_1_YEAR_UPFRONT_PAYMENT`, `COMMITMENT_3_YEAR_UPFRONT_PAYMENT`,

<a name="nested_virtual_machine_preferences_sole_tenancy_preferences"></a>The `sole_tenancy_preferences` block supports:

* `cpu_overcommit_ratio` -
  (Optional)
  CPU overcommit ratio. Acceptable values are between 1.0 and 2.0 inclusive.

* `host_maintenance_policy` -
  (Optional)
  Sole Tenancy nodes maintenance policy. Possible values: `HOST_MAINTENANCE_POLICY_UNSPECIFIED`, `HOST_MAINTENANCE_POLICY_DEFAULT`, `HOST_MAINTENANCE_POLICY_RESTART_IN_PLACE`, `HOST_MAINTENANCE_POLICY_MIGRATE_WITHIN_NODE_GROUP`

* `commitment_plan` -
  (Optional)
  Commitment plan to consider when calculating costs for virtual machine insights and recommendations. If you are unsure which value to set, a 3 year commitment plan is often a good value to start with. Possible values: `COMMITMENT_PLAN_UNSPECIFIED`, `ON_DEMAND`, `COMMITMENT_1_YEAR`, `COMMITMENT_3_YEAR`

* `node_types` -
  (Optional)
  A list of sole tenant node types. An empty list means that all possible node types will be considered.
  Structure is [documented below](#nested_virtual_machine_preferences_sole_tenancy_preferences_node_types).


<a name="nested_virtual_machine_preferences_sole_tenancy_preferences_node_types"></a>The `node_types` block supports:

* `node_name` -
  (Optional)
  Name of the Sole Tenant node. Consult https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/preferenceSets/{{preference_set_id}}`

* `name` -
  Output only. Name of the preference set.

* `create_time` -
  Output only. The timestamp when the preference set was created.

* `update_time` -
  Output only. The timestamp when the preference set was last updated.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


PreferenceSet can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/preferenceSets/{{preference_set_id}}`
* `{{project}}/{{location}}/{{preference_set_id}}`
* `{{location}}/{{preference_set_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import PreferenceSet using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/preferenceSets/{{preference_set_id}}"
  to = google_migration_center_preference_set.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), PreferenceSet can be imported using one of the formats above. For example:

```
$ terraform import google_migration_center_preference_set.default projects/{{project}}/locations/{{location}}/preferenceSets/{{preference_set_id}}
$ terraform import google_migration_center_preference_set.default {{project}}/{{location}}/{{preference_set_id}}
$ terraform import google_migration_center_preference_set.default {{location}}/{{preference_set_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
