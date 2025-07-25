---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dialogflow/EncryptionSpec.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Dialogflow"
description: |-
  Initializes a location-level encryption key specification.
---

# google_dialogflow_encryption_spec

Initializes a location-level encryption key specification.


To get more information about EncryptionSpec, see:

* [API documentation](https://cloud.google.com/dialogflow/es/docs/reference/rest/v2/projects.locations.encryptionSpec)
* How-to Guides
    * [Official CX Documentation](https://cloud.google.com/dialogflow/cx/docs)
    * [Official ES Documentation](https://cloud.google.com/dialogflow/es/docs)

## Example Usage - Dialogflow Encryption Spec Basic


```hcl
resource "google_project" "project" {
  provider        = google-beta
  project_id      = "my-proj"
  name            = "my-proj"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
  deletion_policy = "DELETE"
}

resource "google_project_service" "cloudkms" {
  provider = google-beta
  project  = google_project.project.project_id
  service  = "cloudkms.googleapis.com"
}

resource "google_project_service" "dialogflow" {
  provider = google-beta
  project  = google_project.project.project_id
  service  = "dialogflow.googleapis.com"
}

resource "time_sleep" "wait_enable_service_api" {
  depends_on = [
    google_project_service.cloudkms,
    google_project_service.dialogflow
  ]
  create_duration = "30s"
}

resource "google_project_service_identity" "gcp_sa" {
  provider   = google-beta
  service    = "dialogflow.googleapis.com"
  project    = google_project.project.project_id
  depends_on = [time_sleep.wait_enable_service_api]
}

resource "time_sleep" "wait_create_sa" {
  depends_on      = [google_project_service_identity.gcp_sa]
  create_duration = "30s"
}

resource "google_kms_key_ring" "keyring" {
  provider   = google-beta
  name       = "my-keyring"
  location   = "us-central1"
  project    = google_project.project.project_id
  depends_on = [time_sleep.wait_enable_service_api]
}

resource "google_kms_crypto_key" "key" {
  provider = google-beta
  name     = "my-key"
  key_ring = google_kms_key_ring.keyring.id
  purpose  = "ENCRYPT_DECRYPT"
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  provider      = google-beta
  crypto_key_id = google_kms_crypto_key.key.id
  member        = "${replace(google_project_service_identity.gcp_sa.member, "@gcp-sa-dialogflow.iam", "@gcp-sa-ccai-cmek.iam")}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  depends_on    = [time_sleep.wait_create_sa]
}

resource "google_dialogflow_encryption_spec" "my-encryption-spec" {
  provider = google-beta
  project  = google_project.project.project_id
  location = "us-central1"
  encryption_spec {
    kms_key = google_kms_crypto_key.key.id
  }
  depends_on = [google_kms_crypto_key_iam_member.crypto_key]
}
```

## Argument Reference

The following arguments are supported:


* `encryption_spec` -
  (Required)
  A nested object resource.
  Structure is [documented below](#nested_encryption_spec).

* `location` -
  (Required)
  The location in which the encryptionSpec is to be initialized.


* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_encryption_spec"></a>The `encryption_spec` block supports:

* `kms_key` -
  (Required)
  The name of customer-managed encryption key that is used to secure a resource and its sub-resources.
  If empty, the resource is secured by the default Google encryption key.
  Only the key in the same location as this resource is allowed to be used for encryption.
  Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{key}

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/encryptionSpec/{{name}}`


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import

This resource does not support import.

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
