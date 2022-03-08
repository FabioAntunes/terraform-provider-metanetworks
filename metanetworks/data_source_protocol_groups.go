package metanetworks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProtocolGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Returns all `protocol_group` of the organization",
		ReadContext: dataSourceProtocolGroupsRead,
		Schema: map[string]*schema.Schema{
			"protocol_groups": {
				Description: "List of all `protocol_group`.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: ProtocolGroupSchema,
				},
			},
		},
	}
}

func dataSourceProtocolGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	// // Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var protocolGroups []ProtocolGroup
	protocolGroups, err := client.GetProtocolGroups()
	if err != nil {
		return diag.FromErr(err)
	}

	err = ProtocolGroupsToResource(d, &protocolGroups)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
