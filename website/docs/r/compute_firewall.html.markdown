---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/Firewall.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
description: |-
  Each network has its own firewall controlling access to and from the
  instances.
---

# google_compute_firewall

Each network has its own firewall controlling access to and from the
instances.

All traffic to instances, even from other instances, is blocked by the
firewall unless firewall rules are created to allow it.

The default network has automatically created firewall rules that are
shown in default firewall rules. No manually created network has
automatically created firewall rules except for a default "allow" rule for
outgoing traffic and a default "deny" for incoming traffic. For all
networks except the default network, you must create any firewall rules
you need.


To get more information about Firewall, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/v1/firewalls)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/vpc/docs/firewalls)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=firewall_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Firewall Basic


```hcl
resource "google_compute_firewall" "default" {
  name    = "test-firewall"
  network = google_compute_network.default.name

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["80", "8080", "1000-2000"]
  }

  source_tags = ["web"]
}

resource "google_compute_network" "default" {
  name = "test-network"
}
```
## Example Usage - Firewall With Target Tags


```hcl
resource "google_compute_firewall" "rules" {
  project     = "my-project-name"
  name        = "my-firewall-rule"
  network     = "default"
  description = "Creates firewall rule targeting tagged instances"

  allow {
    protocol  = "tcp"
    ports     = ["80", "8080", "1000-2000"]
  }

  source_tags = ["foo"]
  target_tags = ["web"]
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  Name of the resource. Provided by the client when the resource is
  created. The name must be 1-63 characters long, and comply with
  RFC1035. Specifically, the name must be 1-63 characters long and match
  the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the
  first character must be a lowercase letter, and all following
  characters must be a dash, lowercase letter, or digit, except the last
  character, which cannot be a dash.

* `network` -
  (Required)
  The name or self_link of the network to attach this firewall to.


* `allow` -
  (Optional)
  The list of ALLOW rules specified by this firewall. Each rule
  specifies a protocol and port-range tuple that describes a permitted
  connection.
  Structure is [documented below](#nested_allow).

* `deny` -
  (Optional)
  The list of DENY rules specified by this firewall. Each rule specifies
  a protocol and port-range tuple that describes a denied connection.
  Structure is [documented below](#nested_deny).

* `description` -
  (Optional)
  An optional description of this resource. Provide this property when
  you create the resource.

* `destination_ranges` -
  (Optional)
  If destination ranges are specified, the firewall will apply only to
  traffic that has destination IP address in these ranges. These ranges
  must be expressed in CIDR format. IPv4 or IPv6 ranges are supported.

* `direction` -
  (Optional)
  Direction of traffic to which this firewall applies; default is
  INGRESS. Note: For INGRESS traffic, one of `source_ranges`,
  `source_tags` or `source_service_accounts` is required.
  Possible values are: `INGRESS`, `EGRESS`.

* `disabled` -
  (Optional)
  Denotes whether the firewall rule is disabled, i.e not applied to the
  network it is associated with. When set to true, the firewall rule is
  not enforced and the network behaves as if it did not exist. If this
  is unspecified, the firewall rule will be enabled.

* `log_config` -
  (Optional)
  This field denotes the logging options for a particular firewall rule.
  If defined, logging is enabled, and logs will be exported to Cloud Logging.
  Structure is [documented below](#nested_log_config).

* `priority` -
  (Optional)
  Priority for this rule. This is an integer between 0 and 65535, both
  inclusive. When not specified, the value assumed is 1000. Relative
  priorities determine precedence of conflicting rules. Lower value of
  priority implies higher precedence (eg, a rule with priority 0 has
  higher precedence than a rule with priority 1). DENY rules take
  precedence over ALLOW rules having equal priority.

* `source_ranges` -
  (Optional)
  If source ranges are specified, the firewall will apply only to
  traffic that has source IP address in these ranges. These ranges must
  be expressed in CIDR format. One or both of sourceRanges and
  sourceTags may be set. If both properties are set, the firewall will
  apply to traffic that has source IP address within sourceRanges OR the
  source IP that belongs to a tag listed in the sourceTags property. The
  connection does not need to match both properties for the firewall to
  apply. IPv4 or IPv6 ranges are supported. For INGRESS traffic, one of
  `source_ranges`, `source_tags` or `source_service_accounts` is required.

* `source_service_accounts` -
  (Optional)
  If source service accounts are specified, the firewall will apply only
  to traffic originating from an instance with a service account in this
  list. Source service accounts cannot be used to control traffic to an
  instance's external IP address because service accounts are associated
  with an instance, not an IP address. sourceRanges can be set at the
  same time as sourceServiceAccounts. If both are set, the firewall will
  apply to traffic that has source IP address within sourceRanges OR the
  source IP belongs to an instance with service account listed in
  sourceServiceAccount. The connection does not need to match both
  properties for the firewall to apply. sourceServiceAccounts cannot be
  used at the same time as sourceTags or targetTags. For INGRESS traffic,
  one of `source_ranges`, `source_tags` or `source_service_accounts` is required.

* `source_tags` -
  (Optional)
  If source tags are specified, the firewall will apply only to traffic
  with source IP that belongs to a tag listed in source tags. Source
  tags cannot be used to control traffic to an instance's external IP
  address. Because tags are associated with an instance, not an IP
  address. One or both of sourceRanges and sourceTags may be set. If
  both properties are set, the firewall will apply to traffic that has
  source IP address within sourceRanges OR the source IP that belongs to
  a tag listed in the sourceTags property. The connection does not need
  to match both properties for the firewall to apply. For INGRESS traffic,
  one of `source_ranges`, `source_tags` or `source_service_accounts` is required.

* `target_service_accounts` -
  (Optional)
  A list of service accounts indicating sets of instances located in the
  network that may make network connections as specified in allowed[].
  targetServiceAccounts cannot be used at the same time as targetTags or
  sourceTags. If neither targetServiceAccounts nor targetTags are
  specified, the firewall rule applies to all instances on the specified
  network.

* `target_tags` -
  (Optional)
  A list of instance tags indicating sets of instances located in the
  network that may make network connections as specified in allowed[].
  If no targetTags are specified, the firewall rule applies to all
  instances on the specified network.

* `params` -
  (Optional)
  Additional params passed with the request, but not persisted as part of resource payload
  Structure is [documented below](#nested_params).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.

* `enable_logging` - (Optional, Deprecated) This field denotes whether to enable logging for a particular firewall rule.
If logging is enabled, logs will be exported to Stackdriver. Deprecated in favor of `log_config`


<a name="nested_allow"></a>The `allow` block supports:

* `protocol` -
  (Required)
  The IP protocol to which this rule applies. The protocol type is
  required when creating a firewall rule. This value can either be
  one of the following well known protocol strings (tcp, udp,
  icmp, esp, ah, sctp, ipip, all), or the IP protocol number.

* `ports` -
  (Optional)
  An optional list of ports to which this rule applies. This field
  is only applicable for UDP or TCP protocol. Each entry must be
  either an integer or a range. If not specified, this rule
  applies to connections through any port.
  Example inputs include: [22], [80, 443], and
  ["12345-12349"].

<a name="nested_deny"></a>The `deny` block supports:

* `protocol` -
  (Required)
  The IP protocol to which this rule applies. The protocol type is
  required when creating a firewall rule. This value can either be
  one of the following well known protocol strings (tcp, udp,
  icmp, esp, ah, sctp, ipip, all), or the IP protocol number.

* `ports` -
  (Optional)
  An optional list of ports to which this rule applies. This field
  is only applicable for UDP or TCP protocol. Each entry must be
  either an integer or a range. If not specified, this rule
  applies to connections through any port.
  Example inputs include: [22], [80, 443], and
  ["12345-12349"].

<a name="nested_log_config"></a>The `log_config` block supports:

* `metadata` -
  (Required)
  This field denotes whether to include or exclude metadata for firewall logs.
  Possible values are: `EXCLUDE_ALL_METADATA`, `INCLUDE_ALL_METADATA`.

<a name="nested_params"></a>The `params` block supports:

* `resource_manager_tags` -
  (Optional)
  Resource manager tags to be bound to the firewall. Tag keys and values have the
  same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id},
  and values are in the format tagValues/456. The field is ignored when empty.
  The field is immutable and causes resource replacement when mutated. This field is only
  set at create time and modifying this field after creation will trigger recreation.
  To apply tags to an existing resource, see the google_tags_tag_binding resource.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/global/firewalls/{{name}}`

* `creation_timestamp` -
  Creation timestamp in RFC3339 text format.
* `self_link` - The URI of the created resource.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Firewall can be imported using any of these accepted formats:

* `projects/{{project}}/global/firewalls/{{name}}`
* `{{project}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Firewall using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/global/firewalls/{{name}}"
  to = google_compute_firewall.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Firewall can be imported using one of the formats above. For example:

```
$ terraform import google_compute_firewall.default projects/{{project}}/global/firewalls/{{name}}
$ terraform import google_compute_firewall.default {{project}}/{{name}}
$ terraform import google_compute_firewall.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
