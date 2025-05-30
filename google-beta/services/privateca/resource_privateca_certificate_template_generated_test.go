// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package privateca_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccPrivatecaCertificateTemplate_privatecaTemplateBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPrivatecaCertificateTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificateTemplate_privatecaTemplateBasicExample(context),
			},
			{
				ResourceName:            "google_privateca_certificate_template.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccPrivatecaCertificateTemplate_privatecaTemplateBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"
  description = "A sample certificate template"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }

  maximum_lifetime = "86400s"

  passthrough_extensions {
    additional_extensions {
      object_id_path = [1, 6]
    }
    known_extensions = ["EXTENDED_KEY_USAGE"]
  }

  predefined_values {
    additional_extensions {
      object_id {
        object_id_path = [1, 6]
      }
      value    = "c3RyaW5nCg=="
      critical = true
    }
    aia_ocsp_servers = ["string"]
    ca_options {
      is_ca                  = false
      max_issuer_path_length = 6
    }
    key_usage {
      base_key_usage {
        cert_sign          = false
        content_commitment = true
        crl_sign           = false
        data_encipherment  = true
        decipher_only      = true
        digital_signature  = true
        encipher_only      = true
        key_agreement      = true
        key_encipherment   = true
      }
      extended_key_usage {
        client_auth      = true
        code_signing     = true
        email_protection = true
        ocsp_signing     = true
        server_auth      = true
        time_stamping    = true
      }
      unknown_extended_key_usages {
        object_id_path = [1, 6]
      }
    }
    policy_ids {
      object_id_path = [1, 6]
    }
  }

  labels = {
    label-one = "value-one"
  }
}
`, context)
}

func TestAccPrivatecaCertificateTemplate_privatecaTemplateZeroMaxIssuerPathLengthNullCaExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPrivatecaCertificateTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificateTemplate_privatecaTemplateZeroMaxIssuerPathLengthNullCaExample(context),
			},
			{
				ResourceName:            "google_privateca_certificate_template.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccPrivatecaCertificateTemplate_privatecaTemplateZeroMaxIssuerPathLengthNullCaExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"
  description = "A sample certificate template"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }

  maximum_lifetime = "86400s"

  passthrough_extensions {
    additional_extensions {
      object_id_path = [1, 6]
    }
    known_extensions = ["EXTENDED_KEY_USAGE"]
  }

  predefined_values {
    additional_extensions {
      object_id {
        object_id_path = [1, 6]
      }
      value    = "c3RyaW5nCg=="
      critical = true
    }
    aia_ocsp_servers = ["string"]
    ca_options {
      is_ca                       = false
      null_ca                     = true
      zero_max_issuer_path_length = true
      max_issuer_path_length      = 0
    }
    key_usage {
      base_key_usage {
        cert_sign          = false
        content_commitment = true
        crl_sign           = false
        data_encipherment  = true
        decipher_only      = true
        digital_signature  = true
        encipher_only      = true
        key_agreement      = true
        key_encipherment   = true
      }
      extended_key_usage {
        client_auth      = true
        code_signing     = true
        email_protection = true
        ocsp_signing     = true
        server_auth      = true
        time_stamping    = true
      }
      unknown_extended_key_usages {
        object_id_path = [1, 6]
      }
    }
    policy_ids {
      object_id_path = [1, 6]
    }
    name_constraints {
      critical                  = true
      permitted_dns_names       = ["*.example1.com", "*.example2.com"]
      excluded_dns_names        = ["*.deny.example1.com", "*.deny.example2.com"]
      permitted_ip_ranges       = ["10.0.0.0/8", "11.0.0.0/8"]
      excluded_ip_ranges        = ["10.1.1.0/24", "11.1.1.0/24"]
      permitted_email_addresses = [".example1.com", ".example2.com"]
      excluded_email_addresses  = [".deny.example1.com", ".deny.example2.com"]
      permitted_uris            = [".example1.com", ".example2.com"]
      excluded_uris             = [".deny.example1.com", ".deny.example2.com"]
    }
  }

  labels = {
    label-one = "value-one"
  }
}
`, context)
}

func testAccCheckPrivatecaCertificateTemplateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_privateca_certificate_template" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{PrivatecaBasePath}}projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("PrivatecaCertificateTemplate still exists at %s", url)
			}
		}

		return nil
	}
}
