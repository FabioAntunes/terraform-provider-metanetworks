package metanetworks

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var ProtocolGroupSchema = map[string]*schema.Schema{
	"name_regex": {
		Description:  "A regex string to apply to the `protocol_group` list returned by Metanetworks. This allows more advanced filtering.",
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringIsValidRegExp,
	},
	"description": {
		Description: "Description of the `protocol_group`.",
		Type:        schema.TypeString,
		Computed:    true,
	},
	"name": {
		Description: "Name of the `protocol_group`.",
		Type:        schema.TypeString,
		Computed:    true,
	},
	"protocols": {
		Description: "List of `protocols`.",
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"from_port": {
					Description: "From port number.",
					Type:        schema.TypeInt,
					Computed:    true,
				},
				"port": {
					Description: "Port number.",
					Type:        schema.TypeInt,
					Computed:    true,
				},
				"proto": {
					Description: "Protocol. Valid values are `tcp`, `udp` and `icmp`.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"to_port": {
					Description: "To port number.",
					Type:        schema.TypeInt,
					Computed:    true,
				},
			},
		},
		Computed: true,
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

	"read_only": {
		Description: "If `protocol_group` is read only.",
		Type:        schema.TypeBool,
		Computed:    true,
	},
}

func dataSourceProtocolGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Returns a `protocol_group` of the organization.",
		ReadContext: dataSourceProtocolGroupRead,
		Schema:      ProtocolGroupSchema,
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
