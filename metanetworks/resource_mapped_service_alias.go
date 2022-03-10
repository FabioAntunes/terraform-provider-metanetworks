package metanetworks

import (
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMappedServiceAlias() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the mapped service alias.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"mapped_service_id": {
				Description: "The ID of the mapped service.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"alias": {
				Description: "Domain name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
		Create: resourceMappedServiceAliasCreate,
		Read:   resourceMappedServiceAliasRead,
		Delete: resourceMappedServiceAliasDelete,
	}
}

func resourceMappedServiceAliasCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	mappedServiceID := d.Get("mapped_service_id").(string)
	alias := d.Get("alias").(string)

	var networkElement *NetworkElement
	networkElement, err := client.GetNetworkElement(mappedServiceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			return errors.New("That is alias is already present on the Mapped Service")
		}
	}

	_, err = client.SetNetworkElementAlias(mappedServiceID, alias)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s_%s", mappedServiceID, alias))

	return resourceMappedServiceAliasRead(d, m)
}

func resourceMappedServiceAliasRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	mappedServiceID := d.Get("mapped_service_id").(string)
	alias := d.Get("alias").(string)

	var networkElement *NetworkElement
	networkElement, err := client.GetNetworkElement(mappedServiceID)
	if err != nil {
		if apierr, ok := err.(*ApiError); ok && apierr.StatusCode == 404 {
			log.Printf("[WARN] Removing Mapped Service Alias %q because mapped service %q no longer exists", d.Get("id").(string), mappedServiceID)
			d.SetId("")

			return nil
		}
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			return nil
		}
	}
	log.Printf("[WARN] Removing Mapped Service Alias %q because it's gone", d.Get("id").(string))
	d.SetId("")
	return nil
}

func resourceMappedServiceAliasDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	mappedServiceID := d.Get("mapped_service_id").(string)
	alias := d.Get("alias").(string)
	var networkElement *NetworkElement
	networkElement, err := client.GetNetworkElement(mappedServiceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			_, err = client.DeleteNetworkElementAlias(mappedServiceID, alias)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
