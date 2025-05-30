// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/osconfigv2/resource_os_config_v2_policy_orchestrator_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package osconfigv2_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccOSConfigV2PolicyOrchestrator_basic(t *testing.T) {
	t.Parallel()

	acctest.BootstrapIamMembers(t, []acctest.IamMember{
		{
			Member: "serviceAccount:service-{project_number}@gcp-sa-osconfig.iam.gserviceaccount.com",
			Role:   "roles/osconfig.serviceAgent",
		},
		{
			Member: "serviceAccount:service-{project_number}@gcp-sa-osconfig-rollout.iam.gserviceaccount.com",
			Role:   "roles/osconfig.rolloutServiceAgent",
		},
		{
			Member: "serviceAccount:service-{project_number}@gcp-sa-progrollout.iam.gserviceaccount.com",
			Role:   "roles/progressiverollout.serviceAgent",
		},
	})

	context := map[string]interface{}{
		"project":        envvar.GetTestProjectFromEnv(),
		"project_number": envvar.GetTestProjectNumberFromEnv(),
		"random_suffix":  acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOSConfigV2PolicyOrchestrator_basic(context),
			},
			{
				ResourceName:            "google_os_config_v2_policy_orchestrator.policy_orchestrator",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "policy_orchestrator_id", "terraform_labels"},
			},
			{
				Config: testAccOSConfigV2PolicyOrchestrator_update(context),
			},
			{
				ResourceName:            "google_os_config_v2_policy_orchestrator.policy_orchestrator",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "policy_orchestrator_id", "terraform_labels"},
			},
		},
	})
}

func testAccOSConfigV2PolicyOrchestrator_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_os_config_v2_policy_orchestrator" "policy_orchestrator" {
    policy_orchestrator_id = "tf-test-test-po%{random_suffix}"
    
    state = "ACTIVE"
    action = "UPSERT"
    
    orchestrated_resource {
        id = "tf-test-test-orchestrated-resource%{random_suffix}"
        os_policy_assignment_v1_payload {
            os_policies {
                id = "tf-test-test-os-policy%{random_suffix}"
                mode = "VALIDATION"
                resource_groups {
                    resources {
                        id = "resource-tf"
                        file {
                            content = "file-content-tf"
                            path = "file-path-tf-1"
                            state = "PRESENT"
                        }
                    }
                }
            }
            instance_filter {
                inventories {
                    os_short_name = "windows-10"
                }
            }
            rollout {
                disruption_budget {
                    percent = 100
                }
                min_wait_duration = "60s"
            }
        }
    }
    labels = {
        state = "active"
    }
}
`, context)
}

func testAccOSConfigV2PolicyOrchestrator_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_os_config_v2_policy_orchestrator" "policy_orchestrator" {
    policy_orchestrator_id = "tf-test-test-po%{random_suffix}"
    
    state = "STOPPED"
    action = "DELETE"
    description = "Updated description"
    
    orchestrated_resource {
        id = "tf-test-test-orchestrated-resource%{random_suffix}"
        os_policy_assignment_v1_payload {
            os_policies {
                id = "tf-test-test-os-policy%{random_suffix}"
                mode = "VALIDATION"
                resource_groups {
                    resources {
                        id = "resource-tf"
                        file {
                            content = "file-content-tf-2"
                            path = "file-path-tf-2"
                            state = "PRESENT"
                        }
                    }
                }
            }
            instance_filter {
                inventories {
                    os_short_name = "ubuntu"
                }
            }
            rollout {
                disruption_budget {
                    percent = 50
                }
                min_wait_duration = "120s"
            }
        }
    }
    labels = {
        state = "active"
    }
}
`, context)
}
