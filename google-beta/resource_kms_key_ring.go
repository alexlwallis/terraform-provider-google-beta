// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceKMSKeyRing() *schema.Resource {
	return &schema.Resource{
		Create: resourceKMSKeyRingCreate,
		Read:   resourceKMSKeyRingRead,
		Delete: resourceKMSKeyRingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceKMSKeyRingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The location for the KeyRing.
A full list of valid locations can be found by running 'gcloud kms locations list'.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource name for the KeyRing.`,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKMSKeyRingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandKMSKeyRingName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	locationProp, err := expandKMSKeyRingLocation(d.Get("location"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location"); !isEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}

	obj, err = resourceKMSKeyRingEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{KMSBasePath}}projects/{{project}}/locations/{{location}}/keyRings?keyRingId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new KeyRing: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating KeyRing: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/keyRings/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating KeyRing %q: %#v", d.Id(), res)

	return resourceKMSKeyRingRead(d, meta)
}

func resourceKMSKeyRingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{KMSBasePath}}projects/{{project}}/locations/{{location}}/keyRings/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("KMSKeyRing %q", d.Id()))
	}

	res, err = resourceKMSKeyRingDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing KMSKeyRing because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading KeyRing: %s", err)
	}

	if err := d.Set("name", flattenKMSKeyRingName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeyRing: %s", err)
	}

	return nil
}

func resourceKMSKeyRingDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] KMS KeyRing resources"+
		" cannot be deleted from GCP. The resource %s will be removed from Terraform"+
		" state, but will still be present on the server.", d.Id())
	d.SetId("")

	return nil
}

func resourceKMSKeyRingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/keyRings/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/keyRings/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenKMSKeyRingName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandKMSKeyRingName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandKMSKeyRingLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceKMSKeyRingEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func resourceKMSKeyRingDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Take the returned long form of the name and use it as `self_link`.
	// Then modify the name to be the user specified form.
	// We can't just ignore_read on `name` as the linter will
	// complain that the returned `res` is never used afterwards.
	// Some field needs to be actually set, and we chose `name`.
	d.Set("self_link", res["name"].(string))
	res["name"] = d.Get("name").(string)
	return res, nil
}
