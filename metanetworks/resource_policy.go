package metanetworks

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the policy.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "The description of the policy",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the policy.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"destinations": {
				Description: "Set of users and/or groups/devices/mapped subnets/mapped services to attach to the policy.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"enabled": {
				Description: "default=true.",
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
			},
			"protocol_groups": {
				Description: "Set of protocol groups.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"exempt_sources": {
				Description: "Set of users and/or groups/devices/mapped subnets/mapped services to exempt from the policy.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"sources": {
				Description: "Set of users and/or groups/devices/mapped subnets/mapped services to attach to the policy.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"created_at": {
				Description: "Creation Timestamp.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"modified_at": {
				Description: "Modification Timestamp.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"org_id": {
				Description: "The ID of the organization.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
		Create: resourcePolicyCreate,
		Read:   resourcePolicyRead,
		Update: resourcePolicyUpdate,
		Delete: resourcePolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourcePolicyCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	destinations := resourceTypeSetToStringSlice(d.Get("destinations").(*schema.Set))
	protocolGroups := resourceTypeSetToStringSlice(d.Get("protocol_groups").(*schema.Set))

	policy := Policy{
		Name:           name,
		Description:    description,
		Enabled:        enabled,
		Destinations:   destinations,
		ExemptSources:  exemptSources,
		Sources:        sources,
		ProtocolGroups: protocolGroups,
	}

	var newPolicy *Policy
	newPolicy, err := client.CreatePolicy(&policy)
	if err != nil {
		return err
	}

	_, err = WaitPolicyCreate(client, newPolicy.ID)
	if err != nil {
		return fmt.Errorf("Error waiting for policy creation (%s) (%s)", newPolicy.ID, err)
	}

	d.SetId(newPolicy.ID)

	err = policyToResource(d, newPolicy)
	if err != nil {
		return err
	}

	return resourcePolicyRead(d, m)
}

func resourcePolicyRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	policy, err := client.GetPolicy(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = policyToResource(d, policy)
	if err != nil {
		return err
	}

	return nil
}

func resourcePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	destinations := resourceTypeSetToStringSlice(d.Get("destinations").(*schema.Set))
	protocolGroups := resourceTypeSetToStringSlice(d.Get("protocol_groups").(*schema.Set))

	policy := Policy{
		Name:           name,
		Description:    description,
		Enabled:        enabled,
		Destinations:   destinations,
		ExemptSources:  exemptSources,
		Sources:        sources,
		ProtocolGroups: protocolGroups,
	}

	var updatedPolicy *Policy
	updatedPolicy, err := client.UpdatePolicy(d.Id(), &policy)
	if err != nil {
		return err
	}

	d.SetId(updatedPolicy.ID)

	err = policyToResource(d, updatedPolicy)
	if err != nil {
		return err
	}

	return resourcePolicyRead(d, m)
}

func resourcePolicyDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeletePolicy(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func policyToResource(d *schema.ResourceData, m *Policy) error {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("destinations", m.Destinations)
	d.Set("enabled", m.Enabled)
	d.Set("protocol_groups", m.ProtocolGroups)
	d.Set("exempt_sources", m.ExemptSources)
	d.Set("sources", m.Sources)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)

	d.SetId(m.ID)

	return nil
}
