package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMetaportCluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the Metaport Cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the Metaport Cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the Metaport Cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mapped_elements": {
				Description: "Network elements attached to the Metaport Cluster.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
			"metaports": {
				Description: "Metaport elements attached to the Metaport Cluster.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
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
		},
		Create: resourceMetaportClusterCreate,
		Read:   resourceMetaportClusterRead,
		Update: resourceMetaportClusterUpdate,
		Delete: resourceMetaportClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceMetaportClusterCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	metaportCluster := MetaportCluster{
		Name:        name,
		Description: description,
	}
	var newMetaportCluster *MetaportCluster
	newMetaportCluster, err := client.CreateMetaPortCluster(&metaportCluster)
	if err != nil {
		return err
	}

	d.SetId(newMetaportCluster.ID)
	err = metaportClusterToResource(d, newMetaportCluster)
	if err != nil {
		return err
	}
	return resourceMetaportClusterRead(d, m)
}

func resourceMetaportClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	metaportCluster, err := client.GetMetaPortCluster(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = metaportClusterToResource(d, metaportCluster)
	if err != nil {
		return err
	}

	return nil
}

func resourceMetaportClusterUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	metaportCluster := MetaportCluster{
		Name:        name,
		Description: description,
	}

	var updatedMetaportCluster *MetaportCluster
	updatedMetaportCluster, err := client.UpdateMetaPortCluster(d.Id(), &metaportCluster)
	if err != nil {
		return err
	}
	err = metaportClusterToResource(d, updatedMetaportCluster)
	if err != nil {
		return err
	}

	return resourceMetaportClusterRead(d, m)
}

func resourceMetaportClusterDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteMetaPortCluster(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func metaportClusterToResource(d *schema.ResourceData, m *MetaportCluster) error {
	err := d.Set("name", m.Name)
	if err != nil {
		return err
	}

	err = d.Set("description", m.Description)
	if err != nil {
		return err
	}

	err = d.Set("mapped_elements", m.MappedElements)
	if err != nil {
		return err
	}

	err = d.Set("metaports", m.Metaports)
	if err != nil {
		return err
	}

	err = d.Set("created_at", m.CreatedAt)
	if err != nil {
		return err
	}

	err = d.Set("modified_at", m.ModifiedAt)
	if err != nil {
		return err
	}

	d.SetId(m.ID)

	return nil
}
