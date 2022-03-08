package metanetworks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProtocolGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProtocolGroupsRead,
		Schema: map[string]*schema.Schema{
			"protocol_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocols": {
							Type: schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from_port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"proto": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"to_port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"modified_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"org_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"read_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
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
