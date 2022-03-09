package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSwgThreatCategories() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the threat category.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the threat category.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the threat category.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"countries": {
				Description: "Access restricted countries. Enum by Alpha-2 code (ISO-3166). EG 'AU' -> Australia, 'US' -> United States",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"confidence_level": {
				Description: "Confidence of the classification when the classification engine classifies a URL",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"risk_level": {
				Description: "Risk threshold that will not be tolerated while browsing URL categories under selected threat types. Enum: 'LOW', 'MEDIUM', 'HIGH'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"types": {
				Description: "Predefined threat types to protect against. Enum: 'Abused TLD', 'Bitcoin Related', 'Blackhole', 'Bot', 'Brute Forcer', 'Chat Server', 'CnC', 'Compromised', 'DDoS Target', 'Drive By Src', 'Drop', 'DynDNS', 'EXE Source', 'Fake AV', 'IP Check', 'Mobile CnC', 'Mobile Spyware CnC', 'Online Gaming', 'P2P CnC', 'P2P', 'Parking', 'Phishing', 'Proxy', 'Remote Access Service', 'Scanner', 'Self Signed SSL', 'Spam', 'Spyware CnC', 'Tor', 'Undesirable', 'Utility', 'VPN'",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
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
		Create: resourceSwgThreatCategoriesCreate,
		Read:   resourceSwgThreatCategoriesRead,
		Update: resourceSwgThreatCategoriesUpdate,
		Delete: resourceSwgThreatCategoriesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceSwgThreatCategoriesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	countries := resourceTypeSetToStringSlice(d.Get("countries").(*schema.Set))
	types := resourceTypeSetToStringSlice(d.Get("types").(*schema.Set))
	confidenceLevel := d.Get("confidence_level").(string)
	riskLevel := d.Get("risk_level").(string)

	swgThreatCategories := SwgThreatCategories{
		Name:            name,
		Description:     description,
		Countries:       countries,
		Types:           types,
		ConfidenceLevel: confidenceLevel,
		RiskLevel:       riskLevel,
	}

	var newSwgThreatCategories *SwgThreatCategories
	newSwgThreatCategories, err := client.CreateSwgThreatCategories(&swgThreatCategories)
	if err != nil {
		return err
	}

	d.SetId(newSwgThreatCategories.ID)

	err = swgThreatCategoriesToResource(d, newSwgThreatCategories)
	if err != nil {
		return err
	}

	return resourceSwgThreatCategoriesRead(d, m)
}

func resourceSwgThreatCategoriesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	var swgThreatCategories *SwgThreatCategories
	swgThreatCategories, err := client.GetSwgThreatCategories(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = swgThreatCategoriesToResource(d, swgThreatCategories)
	if err != nil {
		return err
	}

	return nil
}

func resourceSwgThreatCategoriesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	countries := resourceTypeSetToStringSlice(d.Get("countries").(*schema.Set))
	types := resourceTypeSetToStringSlice(d.Get("types").(*schema.Set))
	confidenceLevel := d.Get("confidence_level").(string)
	riskLevel := d.Get("risk_level").(string)

	swgThreatCategories := SwgThreatCategories{
		Name:            name,
		Description:     description,
		Countries:       countries,
		Types:           types,
		ConfidenceLevel: confidenceLevel,
		RiskLevel:       riskLevel,
	}

	var updatedSwgThreatCategories *SwgThreatCategories
	updatedSwgThreatCategories, err := client.UpdateSwgThreatCategories(d.Id(), &swgThreatCategories)
	if err != nil {
		return err
	}

	d.SetId(updatedSwgThreatCategories.ID)

	err = swgThreatCategoriesToResource(d, updatedSwgThreatCategories)
	if err != nil {
		return err
	}

	return resourceSwgThreatCategoriesRead(d, m)
}

func resourceSwgThreatCategoriesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteSwgThreatCategories(d.Id())
	return err
}

func swgThreatCategoriesToResource(d *schema.ResourceData, m *SwgThreatCategories) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("countries", m.Countries)
	d.Set("types", m.Types)
	d.Set("confidence_level", m.ConfidenceLevel)
	d.Set("risk_level", m.RiskLevel)

	d.SetId(m.ID)

	return nil
}
