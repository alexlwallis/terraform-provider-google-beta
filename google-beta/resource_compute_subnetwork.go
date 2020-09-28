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
	"context"
	"fmt"
	"log"
	"net"
	"reflect"
	"time"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Whether the IP CIDR change shrinks the block.
func isShrinkageIpCidr(_ context.Context, old, new, _ interface{}) bool {
	_, oldCidr, oldErr := net.ParseCIDR(old.(string))
	_, newCidr, newErr := net.ParseCIDR(new.(string))

	if oldErr != nil || newErr != nil {
		// This should never happen. The ValidateFunc on the field ensures it.
		return false
	}

	oldStart, oldEnd := cidr.AddressRange(oldCidr)

	if newCidr.Contains(oldStart) && newCidr.Contains(oldEnd) {
		// This is a CIDR range expansion, no need to ForceNew, we have an update method for it.
		return false
	}

	return true
}

func resourceComputeSubnetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeSubnetworkCreate,
		Read:   resourceComputeSubnetworkRead,
		Update: resourceComputeSubnetworkUpdate,
		Delete: resourceComputeSubnetworkDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeSubnetworkImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Update: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(6 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			customdiff.ForceNewIfChange("ip_cidr_range", isShrinkageIpCidr),
			resourceComputeSubnetworkSecondaryIpRangeSetStyleDiff,
		),

		Schema: map[string]*schema.Schema{
			"ip_cidr_range": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateIpCidrRange,
				Description: `The range of internal addresses that are owned by this subnetwork.
Provide this property when you create the subnetwork. For example,
10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
non-overlapping within a network. Only IPv4 is supported.`,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateGCPName,
				Description: `The name of the resource, provided by the client when initially
creating the resource. The name must be 1-63 characters long, and
comply with RFC1035. Specifically, the name must be 1-63 characters
long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which
means the first character must be a lowercase letter, and all
following characters must be a dash, lowercase letter, or digit,
except the last character, which cannot be a dash.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The network this subnet belongs to.
Only networks that are in the distributed mode can have subnetworks.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource. This field can be set only at resource
creation time.`,
			},
			"log_config": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Denotes the logging options for the subnetwork flow logs. If logging is enabled
logs will be exported to Stackdriver. This field cannot be set if the 'purpose' of this
subnetwork is 'INTERNAL_HTTPS_LOAD_BALANCER'`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregation_interval": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"INTERVAL_5_SEC", "INTERVAL_30_SEC", "INTERVAL_1_MIN", "INTERVAL_5_MIN", "INTERVAL_10_MIN", "INTERVAL_15_MIN", ""}, false),
							Description: `Can only be specified if VPC flow logging for this subnetwork is enabled.
Toggles the aggregation interval for collecting flow logs. Increasing the
interval time will reduce the amount of generated flow logs for long
lasting connections. Default is an interval of 5 seconds per connection. Default value: "INTERVAL_5_SEC" Possible values: ["INTERVAL_5_SEC", "INTERVAL_30_SEC", "INTERVAL_1_MIN", "INTERVAL_5_MIN", "INTERVAL_10_MIN", "INTERVAL_15_MIN"]`,
							Default:      "INTERVAL_5_SEC",
							AtLeastOneOf: []string{"log_config.0.aggregation_interval", "log_config.0.flow_sampling", "log_config.0.metadata", "log_config.0.filter_expr"},
						},
						"filter_expr": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Export filter used to define which VPC flow logs should be logged, as as CEL expression. See
https://cloud.google.com/vpc/docs/flow-logs#filtering for details on how to format this field.`,
							Default:      "true",
							AtLeastOneOf: []string{"log_config.0.aggregation_interval", "log_config.0.flow_sampling", "log_config.0.metadata", "log_config.0.filter_expr"},
						},
						"flow_sampling": {
							Type:     schema.TypeFloat,
							Optional: true,
							Description: `Can only be specified if VPC flow logging for this subnetwork is enabled.
The value of the field must be in [0, 1]. Set the sampling rate of VPC
flow logs within the subnetwork where 1.0 means all collected logs are
reported and 0.0 means no logs are reported. Default is 0.5 which means
half of all collected logs are reported.`,
							Default:      0.5,
							AtLeastOneOf: []string{"log_config.0.aggregation_interval", "log_config.0.flow_sampling", "log_config.0.metadata", "log_config.0.filter_expr"},
						},
						"metadata": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA", "CUSTOM_METADATA", ""}, false),
							Description: `Can only be specified if VPC flow logging for this subnetwork is enabled.
Configures whether metadata fields should be added to the reported VPC
flow logs. Default value: "INCLUDE_ALL_METADATA" Possible values: ["EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA", "CUSTOM_METADATA"]`,
							Default:      "INCLUDE_ALL_METADATA",
							AtLeastOneOf: []string{"log_config.0.aggregation_interval", "log_config.0.flow_sampling", "log_config.0.metadata", "log_config.0.filter_expr"},
						},
						"metadata_fields": {
							Type:     schema.TypeSet,
							Optional: true,
							Description: `List of metadata fields that should be added to reported logs.
Can only be specified if VPC flow logs for this subnetwork is enabled and "metadata" is set to CUSTOM_METADATA.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
					},
				},
			},
			"private_ip_google_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `When enabled, VMs in this subnetwork without external IP addresses can
access Google APIs and services by using Private Google Access.`,
			},
			"purpose": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"INTERNAL_HTTPS_LOAD_BALANCER", "PRIVATE", ""}, false),
				Description: `The purpose of the resource. This field can be either PRIVATE
or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
reserved for Internal HTTP(S) Load Balancing. If unspecified, the
purpose defaults to PRIVATE.

If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role. Possible values: ["INTERNAL_HTTPS_LOAD_BALANCER", "PRIVATE"]`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The GCP region for this subnetwork.`,
			},
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ACTIVE", "BACKUP", ""}, false),
				Description: `The role of subnetwork. Currently, this field is only used when
purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
or BACKUP. An ACTIVE subnetwork is one that is currently being used
for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
is ready to be promoted to ACTIVE or is currently draining. Possible values: ["ACTIVE", "BACKUP"]`,
			},
			"secondary_ip_range": {
				Type:       schema.TypeList,
				Computed:   true,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Description: `An array of configurations for secondary IP ranges for VM instances
contained in this subnetwork. The primary IP of such VM must belong
to the primary ipCidrRange of the subnetwork. The alias IPs may belong
to either primary or secondary ranges.

**Note**: This field uses [attr-as-block mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html) to avoid
breaking users during the 0.12 upgrade. To explicitly send a list
of zero objects you must use the following syntax:
'example=[]'
For more details about this behavior, see [this section](https://www.terraform.io/docs/configuration/attr-as-blocks.html#defining-a-fixed-object-collection-value).`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_cidr_range": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateIpCidrRange,
							Description: `The range of IP addresses belonging to this subnetwork secondary
range. Provide this property when you create the subnetwork.
Ranges must be unique and non-overlapping with all primary and
secondary IP ranges within a network. Only IPv4 is supported.`,
						},
						"range_name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateGCPName,
							Description: `The name associated with this subnetwork secondary range, used
when adding an alias IP range to a VM instance. The name must
be 1-63 characters long, and comply with RFC1035. The name
must be unique within the subnetwork.`,
						},
					},
				},
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"gateway_address": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The gateway address for default routes to reach destination addresses
outside this subnetwork.`,
			},
			"fingerprint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Fingerprint of this resource. This field is used internally during updates of this resource.",
				Deprecated:  "This field is not useful for users, and has been removed as an output.",
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeSubnetworkSecondaryIpRangeSetStyleDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	keys := diff.GetChangedKeysPrefix("secondary_ip_range")
	if len(keys) == 0 {
		return nil
	}
	oldCount, newCount := diff.GetChange("secondary_ip_range.#")
	var count int
	// There could be duplicates - worth continuing even if the counts are unequal.
	if oldCount.(int) < newCount.(int) {
		count = newCount.(int)
	} else {
		count = oldCount.(int)
	}

	if count < 1 {
		return nil
	}
	old := make([]interface{}, count)
	new := make([]interface{}, count)
	for i := 0; i < count; i++ {
		o, n := diff.GetChange(fmt.Sprintf("secondary_ip_range.%d", i))

		if o != nil {
			old = append(old, o)
		}
		if n != nil {
			new = append(new, n)
		}
	}

	oldSet := schema.NewSet(schema.HashResource(resourceComputeSubnetwork().Schema["secondary_ip_range"].Elem.(*schema.Resource)), old)
	newSet := schema.NewSet(schema.HashResource(resourceComputeSubnetwork().Schema["secondary_ip_range"].Elem.(*schema.Resource)), new)

	if oldSet.Equal(newSet) {
		if err := diff.Clear("secondary_ip_range"); err != nil {
			return err
		}
	}

	return nil
}

func resourceComputeSubnetworkCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeSubnetworkDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	ipCidrRangeProp, err := expandComputeSubnetworkIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !isEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	nameProp, err := expandComputeSubnetworkName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeSubnetworkNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	purposeProp, err := expandComputeSubnetworkPurpose(d.Get("purpose"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("purpose"); !isEmptyValue(reflect.ValueOf(purposeProp)) && (ok || !reflect.DeepEqual(v, purposeProp)) {
		obj["purpose"] = purposeProp
	}
	roleProp, err := expandComputeSubnetworkRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}
	secondaryIpRangesProp, err := expandComputeSubnetworkSecondaryIpRange(d.Get("secondary_ip_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("secondary_ip_range"); ok || !reflect.DeepEqual(v, secondaryIpRangesProp) {
		obj["secondaryIpRanges"] = secondaryIpRangesProp
	}
	privateIpGoogleAccessProp, err := expandComputeSubnetworkPrivateIpGoogleAccess(d.Get("private_ip_google_access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("private_ip_google_access"); !isEmptyValue(reflect.ValueOf(privateIpGoogleAccessProp)) && (ok || !reflect.DeepEqual(v, privateIpGoogleAccessProp)) {
		obj["privateIpGoogleAccess"] = privateIpGoogleAccessProp
	}
	regionProp, err := expandComputeSubnetworkRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	logConfigProp, err := expandComputeSubnetworkLogConfig(d.Get("log_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("log_config"); ok || !reflect.DeepEqual(v, logConfigProp) {
		obj["logConfig"] = logConfigProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Subnetwork: %#v", obj)
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

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Subnetwork: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating Subnetwork", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Subnetwork: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Subnetwork %q: %#v", d.Id(), res)

	return resourceComputeSubnetworkRead(d, meta)
}

func resourceComputeSubnetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
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

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeSubnetwork %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeSubnetworkCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("description", flattenComputeSubnetworkDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("gateway_address", flattenComputeSubnetworkGatewayAddress(res["gatewayAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("ip_cidr_range", flattenComputeSubnetworkIpCidrRange(res["ipCidrRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("name", flattenComputeSubnetworkName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("network", flattenComputeSubnetworkNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("purpose", flattenComputeSubnetworkPurpose(res["purpose"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("role", flattenComputeSubnetworkRole(res["role"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("secondary_ip_range", flattenComputeSubnetworkSecondaryIpRange(res["secondaryIpRanges"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("private_ip_google_access", flattenComputeSubnetworkPrivateIpGoogleAccess(res["privateIpGoogleAccess"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("region", flattenComputeSubnetworkRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("log_config", flattenComputeSubnetworkLogConfig(res["logConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Subnetwork: %s", err)
	}

	return nil
}

func resourceComputeSubnetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("ip_cidr_range") {
		obj := make(map[string]interface{})

		ipCidrRangeProp, err := expandComputeSubnetworkIpCidrRange(d.Get("ip_cidr_range"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("ip_cidr_range"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
			obj["ipCidrRange"] = ipCidrRangeProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}/expandIpCidrRange")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Subnetwork %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Subnetwork %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating Subnetwork", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("private_ip_google_access") {
		obj := make(map[string]interface{})

		privateIpGoogleAccessProp, err := expandComputeSubnetworkPrivateIpGoogleAccess(d.Get("private_ip_google_access"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("private_ip_google_access"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, privateIpGoogleAccessProp)) {
			obj["privateIpGoogleAccess"] = privateIpGoogleAccessProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}/setPrivateIpGoogleAccess")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Subnetwork %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Subnetwork %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating Subnetwork", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("log_config") {
		obj := make(map[string]interface{})

		getUrl, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		getRes, err := sendRequest(config, "GET", billingProject, getUrl, userAgent, nil)
		if err != nil {
			return handleNotFoundError(err, d, fmt.Sprintf("ComputeSubnetwork %q", d.Id()))
		}

		obj["fingerprint"] = getRes["fingerprint"]

		logConfigProp, err := expandComputeSubnetworkLogConfig(d.Get("log_config"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("log_config"); ok || !reflect.DeepEqual(v, logConfigProp) {
			obj["logConfig"] = logConfigProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Subnetwork %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Subnetwork %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating Subnetwork", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("role") {
		obj := make(map[string]interface{})

		getUrl, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		getRes, err := sendRequest(config, "GET", billingProject, getUrl, userAgent, nil)
		if err != nil {
			return handleNotFoundError(err, d, fmt.Sprintf("ComputeSubnetwork %q", d.Id()))
		}

		obj["fingerprint"] = getRes["fingerprint"]

		roleProp, err := expandComputeSubnetworkRole(d.Get("role"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, roleProp)) {
			obj["role"] = roleProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Subnetwork %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Subnetwork %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating Subnetwork", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("secondary_ip_range") {
		obj := make(map[string]interface{})

		getUrl, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		getRes, err := sendRequest(config, "GET", billingProject, getUrl, userAgent, nil)
		if err != nil {
			return handleNotFoundError(err, d, fmt.Sprintf("ComputeSubnetwork %q", d.Id()))
		}

		obj["fingerprint"] = getRes["fingerprint"]

		secondaryIpRangesProp, err := expandComputeSubnetworkSecondaryIpRange(d.Get("secondary_ip_range"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("secondary_ip_range"); ok || !reflect.DeepEqual(v, secondaryIpRangesProp) {
			obj["secondaryIpRanges"] = secondaryIpRangesProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Subnetwork %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Subnetwork %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating Subnetwork", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeSubnetworkRead(d, meta)
}

func resourceComputeSubnetworkDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Subnetwork %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Subnetwork")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting Subnetwork", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Subnetwork %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeSubnetworkImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/subnetworks/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeSubnetworkCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSubnetworkDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSubnetworkGatewayAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSubnetworkIpCidrRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSubnetworkName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSubnetworkNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeSubnetworkPurpose(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSubnetworkRole(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSubnetworkSecondaryIpRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"range_name":    flattenComputeSubnetworkSecondaryIpRangeRangeName(original["rangeName"], d, config),
			"ip_cidr_range": flattenComputeSubnetworkSecondaryIpRangeIpCidrRange(original["ipCidrRange"], d, config),
		})
	}
	return transformed
}
func flattenComputeSubnetworkSecondaryIpRangeRangeName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSubnetworkSecondaryIpRangeIpCidrRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSubnetworkPrivateIpGoogleAccess(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSubnetworkRegion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenComputeSubnetworkLogConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}

	v, ok := original["enable"]
	if ok && !v.(bool) {
		return nil
	}

	transformed := make(map[string]interface{})
	transformed["flow_sampling"] = original["flowSampling"]
	transformed["aggregation_interval"] = original["aggregationInterval"]
	transformed["metadata"] = original["metadata"]
	if original["metadata"].(string) == "CUSTOM_METADATA" {
		transformed["metadata_fields"] = original["metadataFields"]
	} else {
		// MetadataFields can only be set when metadata is CUSTOM_METADATA. However, when updating
		// from custom to include/exclude, the API will return the previous values of the metadata fields,
		// despite not actually having any custom fields at the moment. The API team has confirmed
		// this as WAI (b/162771344), so we work around it by clearing the response if metadata is
		// not custom.
		transformed["metadata_fields"] = nil
	}
	transformed["filter_expr"] = original["filterExpr"]

	return []interface{}{transformed}
}

func expandComputeSubnetworkDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkIpCidrRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSubnetworkPurpose(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkRole(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkSecondaryIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRangeName, err := expandComputeSubnetworkSecondaryIpRangeRangeName(original["range_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRangeName); val.IsValid() && !isEmptyValue(val) {
			transformed["rangeName"] = transformedRangeName
		}

		transformedIpCidrRange, err := expandComputeSubnetworkSecondaryIpRangeIpCidrRange(original["ip_cidr_range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpCidrRange); val.IsValid() && !isEmptyValue(val) {
			transformed["ipCidrRange"] = transformedIpCidrRange
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeSubnetworkSecondaryIpRangeRangeName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkSecondaryIpRangeIpCidrRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkPrivateIpGoogleAccess(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSubnetworkLogConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	transformed := make(map[string]interface{})
	if len(l) == 0 || l[0] == nil {
		purpose, ok := d.GetOkExists("purpose")

		if ok && purpose.(string) == "INTERNAL_HTTPS_LOAD_BALANCER" {
			// Subnetworks for L7ILB do not accept any values for logConfig
			return nil, nil
		}
		// send enable = false to ensure logging is disabled if there is no config
		transformed["enable"] = false
		return transformed, nil
	}

	raw := l[0]
	original := raw.(map[string]interface{})

	// The log_config block is specified, so logging should be enabled
	transformed["enable"] = true
	transformed["aggregationInterval"] = original["aggregation_interval"]
	transformed["flowSampling"] = original["flow_sampling"]
	transformed["metadata"] = original["metadata"]
	transformed["filterExpr"] = original["filter_expr"]

	// make it JSON marshallable
	transformed["metadataFields"] = original["metadata_fields"].(*schema.Set).List()

	return transformed, nil
}
