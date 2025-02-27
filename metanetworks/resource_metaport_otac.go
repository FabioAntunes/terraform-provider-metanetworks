package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMetaportOTAC() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of this resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"metaport_id": {
				Description: "The ID of the metaport.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"secret": {
				Description: "OTAC (Time-limited One-Time Password) for metaport installer download.",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
		},
		Create: resourceMetaportOTACCreate,
		Read:   resourceMetaportOTACRead,
		Delete: resourceMetaportOTACDelete,
	}
}

func resourceMetaportOTACCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	metaportID := d.Get("metaport_id").(string)
	otacSecret, err := client.GenerateMetaPortOTAC(metaportID)
	if err != nil {
		return err
	}

	d.Set("secret", otacSecret)
	d.SetId(otacSecret[0:5])

	return resourceMetaportOTACRead(d, m)
}

func resourceMetaportOTACRead(d *schema.ResourceData, m interface{}) error {
	// fire and forget. The code is valid for a short time, there is no state.
	return nil
}

func resourceMetaportOTACDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
