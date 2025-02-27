package metanetworks

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMappedService() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the mapped service.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the mapped service.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the mapped service.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mapped_service": {
				Description: "Mapped Service IP or Hostname.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"tags": {
				Description: "Tags are key/value attributes that can be used to group elements together.",
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"aliases": {
				Description: "The domain names of the native service.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
			"created_at": {
				Description: "Creation Timestamp.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"dns_name": {
				Description: "`<network_element_id>`.`<org_id>`.nsof",
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
		Create: resourceMappedServiceCreate,
		Read:   resourceMappedServiceRead,
		Update: resourceMappedServiceUpdate,
		Delete: resourceMappedServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceMappedServiceCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mappedService := d.Get("mapped_service").(string)

	networkElement := NetworkElement{
		Name:          name,
		Description:   description,
		MappedService: mappedService,
	}
	var newMappedService *NetworkElement
	newMappedService, err := client.CreateNetworkElement(&networkElement)
	if err != nil {
		return err
	}

	_, err = WaitNetworkElementCreate(client, newMappedService.ID)
	if err != nil {
		return fmt.Errorf("Error waiting for mapped service creation (%s) (%s)", newMappedService.ID, err)
	}

	d.SetId(newMappedService.ID)

	err = mappedServiceToResource(d, newMappedService)
	if err != nil {
		return err
	}
	err = client.SetNetworkElementTags(d)
	if err != nil {
		return err
	}

	return resourceMappedServiceRead(d, m)
}

func resourceMappedServiceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	networkElement, err := client.GetNetworkElement(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = mappedServiceToResource(d, networkElement)
	if err != nil {
		return err
	}

	return nil
}

func resourceMappedServiceUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mappedService := d.Get("mapped_service").(string)

	networkElement := NetworkElement{
		Name:          name,
		Description:   description,
		MappedService: mappedService,
	}
	var updatedMappedService *NetworkElement
	updatedMappedService, err := client.UpdateNetworkElement(d.Id(), &networkElement)
	if err != nil {
		return err
	}

	err = mappedServiceToResource(d, updatedMappedService)
	if err != nil {
		return err
	}
	err = client.SetNetworkElementTags(d)
	if err != nil {
		return err
	}

	return resourceMappedServiceRead(d, m)
}

func resourceMappedServiceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteNetworkElement(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func mappedServiceToResource(d *schema.ResourceData, m *NetworkElement) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("mapped_service", m.MappedService)
	d.Set("aliases", m.Aliases)
	d.Set("created_at", m.CreatedAt)
	d.Set("dns_name", m.DNSName)
	d.Set("expires_at", m.ExpiresAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)

	d.SetId(m.ID)

	return nil
}
