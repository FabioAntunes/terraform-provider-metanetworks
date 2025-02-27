package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMetaport() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the metaport.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the metaport.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the metaport.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enabled": {
				Description: "default=true.",
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
			},
			"mapped_elements": {
				Description: "Network elements attached to the metaport.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
			"allow_support": {
				Description: "Enable external support to access to this metaport remotely, default=true.",
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
			},
			"created_at": {
				Description: "Creation Timestamp.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"dns_name": {
				Description: "<metaport_id>`.`<org_id>`.nsof",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"expires_at": {
				Description: "Expiration Timestamp.",
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
		Create: resourceMetaportCreate,
		Read:   resourceMetaportRead,
		Update: resourceMetaportUpdate,
		Delete: resourceMetaportDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceMetaportCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	allowSupport := d.Get("allow_support").(bool)

	metaport := MetaPort{
		Name:         name,
		Description:  description,
		Enabled:      enabled,
		AllowSupport: &allowSupport,
	}
	var newMetaport *MetaPort
	newMetaport, err := client.CreateMetaPort(&metaport)
	if err != nil {
		return err
	}

	d.SetId(newMetaport.ID)
	err = metaportToResource(d, newMetaport)
	if err != nil {
		return err
	}
	return resourceMetaportRead(d, m)
}

func resourceMetaportRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	metaport, err := client.GetMetaPort(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = metaportToResource(d, metaport)
	if err != nil {
		return err
	}

	return nil
}

func resourceMetaportUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	allowSupport := d.Get("allow_support").(bool)

	metaport := MetaPort{
		Name:         name,
		Description:  description,
		Enabled:      enabled,
		AllowSupport: &allowSupport,
	}

	var updatedMetaport *MetaPort
	updatedMetaport, err := client.UpdateMetaPort(d.Id(), &metaport)
	if err != nil {
		return err
	}
	err = metaportToResource(d, updatedMetaport)
	if err != nil {
		return err
	}

	return resourceMetaportRead(d, m)
}

func resourceMetaportDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteMetaPort(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func metaportToResource(d *schema.ResourceData, m *MetaPort) error {
	err := d.Set("name", m.Name)
	if err != nil {
		return err
	}

	err = d.Set("description", m.Description)
	if err != nil {
		return err
	}

	err = d.Set("enabled", m.Enabled)
	if err != nil {
		return err
	}

	err = d.Set("mapped_elements", m.MappedElements)
	if err != nil {
		return err
	}

	err = d.Set("allow_support", m.AllowSupport)
	if err != nil {
		return err
	}

	err = d.Set("created_at", m.CreatedAt)
	if err != nil {
		return err
	}

	err = d.Set("dns_name", m.DNSName)
	if err != nil {
		return err
	}

	err = d.Set("expires_at", m.ExpiresAt)
	if err != nil {
		return err
	}

	err = d.Set("modified_at", m.ModifiedAt)
	if err != nil {
		return err
	}

	err = d.Set("org_id", m.OrgID)
	if err != nil {
		return err
	}

	d.SetId(m.ID)

	return nil
}
