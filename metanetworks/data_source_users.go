package metanetworks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Description: "Description of the `user`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"email": {
				Description: "Email of the `user`.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"enabled": {
				Description: "If `user` is enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"family_name": {
				Description: "Family name of the `user`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"given_name": {
				Description: "Given name of the `user`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"phone": {
				Description: "Phone number of the `user`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"provisioned_by": {
				Description: "Users can be provisioned in the system either by locally creating the users in the Admin portal or API. Another, more common practice, is to provision users from an organization directory service, by way of SCIM or LDAP protocols.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "Creation Timestamp.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"inventory": {
				Description: "Devices used by the `user`",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
			"mfa_enabled": {
				Description: "If mfa is enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"modified_at": {
				Description: "Modification Timestamp.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the `user`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"org_id": {
				Description: "The ID of the organization.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"overlay_mfa_enabled": {
				Description: "If overlay mfa is enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"phone_verified": {
				Description: "If phone is verified.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"roles": {
				Description: "Roles of the `user`.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
			"tags": {
				Description: "Tags.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	// // Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	email := d.Get("email").(string)

	var user []User
	user, err := client.GetUsers(email)
	if err != nil {
		return diag.FromErr(err)
	}
	err = userToResource(d, &user[0])
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
