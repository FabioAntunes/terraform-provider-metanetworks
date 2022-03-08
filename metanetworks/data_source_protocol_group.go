package metanetworks

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceProtocolGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProtocolGroupRead,
		Schema: map[string]*schema.Schema{
			"name_regex": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsValidRegExp,
			},
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
	}
}

func dataSourceProtocolGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	// // Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var protocolGroups []ProtocolGroup
	var filteredProtocolGroups []ProtocolGroup

	protocolGroups, err := client.GetProtocolGroups()
	if err != nil {
		return diag.FromErr(err)
	}

	if nameRegex, ok := d.GetOk("name_regex"); ok {
		r := regexp.MustCompile(nameRegex.(string))
		for _, protocolGroup := range protocolGroups {
			if r.MatchString(protocolGroup.Name) {
				filteredProtocolGroups = append(filteredProtocolGroups, protocolGroup)
			}
		}
	} else {
		filteredProtocolGroups = protocolGroups[:]
	}

	if len(filteredProtocolGroups) < 1 {
		return diag.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredProtocolGroups) > 1 {
		return diag.Errorf("Your query returned more than one result. Please try a more specific search criteria")
	}

	err = ProtocolGroupToResource(d, &filteredProtocolGroups[0])
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
