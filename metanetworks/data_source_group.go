package metanetworks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Returns a group of the organization.",
		ReadContext: dataSourceGroupRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Description: "The description of the group.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"expression": {
				Description: "Allows grouping entities by their tags. Filtering by tag value is also supported if provided. Supported operations: AND, OR, XOR, parenthesis.",
				Type:        schema.TypeString,
				Computed:    true,
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
				Computed:    true,
			},
			"users": {
				Description: "The group users.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
		},
	}
}

func dataSourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	// // Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("name").(string)

	var group []Group
	group, err := client.GetGroups(name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = groupToResource(d, &group[0])
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
