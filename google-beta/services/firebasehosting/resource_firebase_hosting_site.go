// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/firebasehosting/Site.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package firebasehosting

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceFirebaseHostingSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseHostingSiteCreate,
		Read:   resourceFirebaseHostingSiteRead,
		Update: resourceFirebaseHostingSiteUpdate,
		Delete: resourceFirebaseHostingSiteDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseHostingSiteImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Optional. The [ID of a Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)
associated with the Hosting site.`,
			},
			"site_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Required. Immutable. A globally unique identifier for the Hosting site. This identifier is
used to construct the Firebase-provisioned subdomains for the site, so it must also be a valid
domain name label.`,
			},
			"default_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The default URL for the site in the form of https://{name}.web.app`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The fully-qualified resource name of the Hosting site, in
the format: projects/PROJECT_IDENTIFIER/sites/SITE_ID PROJECT_IDENTIFIER: the
Firebase project's
['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its
['ProjectId'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_id).
Learn more about using project identifiers in Google's
[AIP 2510 standard](https://google.aip.dev/cloud/2510).`,
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The type of Hosting site, either 'DEFAULT_SITE' or 'USER_SITE'`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirebaseHostingSiteCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	appIdProp, err := expandFirebaseHostingSiteAppId(d.Get("app_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(appIdProp)) && (ok || !reflect.DeepEqual(v, appIdProp)) {
		obj["appId"] = appIdProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseHostingBasePath}}projects/{{project}}/sites?siteId={{site_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Site: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Site: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	// Check if the Firebase hostng site already exits. Do an update if so.

	getUrl, err := tpgresource.ReplaceVars(d, config, "{{FirebaseHostingBasePath}}projects/{{project}}/sites/{{site_id}}")
	if err != nil {
		return err
	}
	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    getUrl,
		UserAgent: userAgent,
		Headers:   headers,
	})

	if err == nil {
		// Hosting site already exists
		log.Printf("[DEBUG] Firebase hosting site already exists %s", d.Get("site_id"))
		// Replace import id for the resource id
		id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/sites/{{site_id}}")
		if err != nil {
			return fmt.Errorf("Error constructing id: %s", err)
		}
		d.SetId(id)
		return resourceFirebaseHostingSiteUpdate(d, meta)
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Site: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/sites/{{site_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Site %q: %#v", d.Id(), res)

	return resourceFirebaseHostingSiteRead(d, meta)
}

func resourceFirebaseHostingSiteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseHostingBasePath}}projects/{{project}}/sites/{{site_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Site: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirebaseHostingSite %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Site: %s", err)
	}

	if err := d.Set("name", flattenFirebaseHostingSiteName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Site: %s", err)
	}
	if err := d.Set("app_id", flattenFirebaseHostingSiteAppId(res["appId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Site: %s", err)
	}
	if err := d.Set("default_url", flattenFirebaseHostingSiteDefaultUrl(res["defaultUrl"], d, config)); err != nil {
		return fmt.Errorf("Error reading Site: %s", err)
	}
	if err := d.Set("type", flattenFirebaseHostingSiteType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Site: %s", err)
	}

	return nil
}

func resourceFirebaseHostingSiteUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Site: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	appIdProp, err := expandFirebaseHostingSiteAppId(d.Get("app_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, appIdProp)) {
		obj["appId"] = appIdProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseHostingBasePath}}projects/{{project}}/sites/{{site_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Site %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("app_id") {
		updateMask = append(updateMask, "appId")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating Site %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Site %q: %#v", d.Id(), res)
		}

	}

	return resourceFirebaseHostingSiteRead(d, meta)
}

func resourceFirebaseHostingSiteDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Site: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseHostingBasePath}}projects/{{project}}/sites/{{site_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	if siteType := d.Get("type"); siteType == "DEFAULT_SITE" {
		log.Printf("[WARN] Skip deleting default hosting side: %q", d.Get("name").(string))
		return nil
	}

	log.Printf("[DEBUG] Deleting Site %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Site")
	}

	log.Printf("[DEBUG] Finished deleting Site %q: %#v", d.Id(), res)
	return nil
}

func resourceFirebaseHostingSiteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/sites/(?P<site_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<site_id>[^/]+)$",
		"^sites/(?P<site_id>[^/]+)$",
		"^(?P<site_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/sites/{{site_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseHostingSiteName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseHostingSiteAppId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseHostingSiteDefaultUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseHostingSiteType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirebaseHostingSiteAppId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
