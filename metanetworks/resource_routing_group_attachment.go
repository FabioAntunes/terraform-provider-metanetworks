package metanetworks

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRoutingGroupAttachment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of this resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"routing_group_id": {
				Description: "The ID of the routing group.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"network_element_id": {
				Description: "The ID of the network element to attach to the routing group.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
		Create: resourceRoutingGroupAttachmentCreate,
		Read:   resourceRoutingGroupAttachmentRead,
		Delete: resourceRoutingGroupAttachmentDelete,
	}
}

func resourceRoutingGroupAttachmentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	routingGroupID := d.Get("routing_group_id").(string)

	metanetworksMutexKV.Lock(routingGroupID)
	defer metanetworksMutexKV.Unlock(routingGroupID)

	var routingGroup *RoutingGroup
	routingGroup, err := client.GetRoutingGroup(routingGroupID)
	if err != nil {
		return err
	}

	for i := 0; i < len(routingGroup.MappedElements); i++ {
		if routingGroup.MappedElements[i] == elementID {
			return errors.New("That network element is already mapped to this RoutingGroup")
		}

	}

	routingGroup.MappedElements = append(routingGroup.MappedElements, elementID)
	_, err = client.UpdateRoutingGroup(routingGroupID, routingGroup)
	if err != nil {
		return err
	}

	_, err = WaitRoutingGroupAttachmentCreate(client, routingGroupID, elementID)

	if err != nil {
		return fmt.Errorf("Error waiting for routing group attachment creation (%s) (%s)", routingGroupID, err)
	}

	d.SetId(fmt.Sprintf("%s_%s", routingGroupID, elementID))

	return resourceRoutingGroupAttachmentRead(d, m)
}

func resourceRoutingGroupAttachmentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	routingGroupID := d.Get("routing_group_id").(string)

	var routingGroup *RoutingGroup
	routingGroup, err := client.GetRoutingGroup(routingGroupID)
	if err != nil {
		return err
	}

	found := false
	for i := 0; i < len(routingGroup.MappedElements); i++ {
		if routingGroup.MappedElements[i] == elementID {
			found = true
			break
		}

	}

	// If not present we need to destroy the terraform resource so that it is recreated.
	if !found {
		d.SetId("")
	}

	return nil
}

func resourceRoutingGroupAttachmentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	routingGroupID := d.Get("routing_group_id").(string)

	metanetworksMutexKV.Lock(routingGroupID)
	defer metanetworksMutexKV.Unlock(routingGroupID)

	var routingGroup *RoutingGroup
	routingGroup, err := client.GetRoutingGroup(routingGroupID)
	if err != nil {
		return err
	}

	// Note that if the entry has already been deleted this won't fail.
	for i := 0; i < len(routingGroup.MappedElements); i++ {
		if routingGroup.MappedElements[i] == elementID {
			routingGroup.MappedElements = append(routingGroup.MappedElements[:i], routingGroup.MappedElements[i+1:]...)
			break
		}
	}

	_, err = client.UpdateRoutingGroup(routingGroupID, routingGroup)
	return err
}
