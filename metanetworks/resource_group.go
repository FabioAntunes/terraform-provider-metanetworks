package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the group.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "The description of the group",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the group",
				Type:        schema.TypeString,
				Required:    true,
			},
			"expression": {
				Description: "Allows grouping entities by their tags. Filtering by tag value is also supported if provided. Supported operations: AND, OR, XOR, parenthesis.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"provisioned_by": {
				Description: "Groups can be provisioned in the system either by locally creating the groups from the Admin portal or API. Another, more common practice, is to provision groups from an organization directory service, by way of SCIM or LDAP protocols.",
				Type:        schema.TypeString,
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
			"org_id": {
				Description: "The ID of the organization.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"roles": {
				Description: "The group roles.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"users": {
				Description: "The group users.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
		},
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Update: resourceGroupUpdate,
		Delete: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	expression := d.Get("expression").(string)

	group := Group{
		Name:        name,
		Description: description,
		Expression:  expression,
	}

	var newGroup *Group
	newGroup, err := client.CreateGroup(&group)
	if err != nil {
		return err
	}

	d.SetId(newGroup.ID)
	err = groupToResource(d, newGroup)
	if err != nil {
		return err
	}

	err = setGroupRoles(d, client)
	if err != nil {
		return err
	}

	err = setGroupUsers(d, client)
	if err != nil {
		return err
	}

	return resourceGroupRead(d, m)
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	group, err := client.GetGroup(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = groupToResource(d, group)
	if err != nil {
		return err
	}

	return nil
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	expression := d.Get("expression").(string)

	group := Group{
		Name:        name,
		Description: description,
		Expression:  expression,
	}

	var updatedGroup *Group
	updatedGroup, err := client.UpdateGroup(d.Id(), &group)
	if err != nil {
		return err
	}

	d.SetId(updatedGroup.ID)

	err = groupToResource(d, updatedGroup)
	if err != nil {
		return err
	}

	err = setGroupRoles(d, client)
	if err != nil {
		return err
	}

	err = setGroupUsers(d, client)
	if err != nil {
		return err
	}

	return resourceGroupRead(d, m)
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteGroup(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func setGroupRoles(d *schema.ResourceData, client *Client) error {
	if d.HasChange("roles") {
		roles := resourceTypeSetToStringSlice(d.Get("roles").(*schema.Set))
		group, err := client.SetGroupRoles(d.Id(), roles)
		if err != nil {
			return err
		}

		groupToResource(d, group)
	}

	return nil
}

func setGroupUsers(d *schema.ResourceData, client *Client) error {
	if d.HasChange("users") {
		old, new := d.GetChange("users")
		toAddSet := new.(*schema.Set).Difference(old.(*schema.Set))
		toRemoveSet := old.(*schema.Set).Difference(new.(*schema.Set))

		toAdd := resourceTypeSetToStringSlice(toAddSet)
		toRemove := resourceTypeSetToStringSlice(toRemoveSet)

		var group *Group
		var err error
		if len(toAdd) > 0 {
			group, err = client.AddGroupUsers(d.Id(), toAdd)
			if err != nil {
				return err
			}
		}
		if len(toRemove) > 0 {
			group, err = client.RemoveGroupUsers(d.Id(), toAdd)
			if err != nil {
				return err
			}
		}

		groupToResource(d, group)
	}

	return nil
}
