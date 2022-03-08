package metanetworks

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLocations() *schema.Resource {
	return &schema.Resource{
		Description: "Returns all the `locations`.",
		ReadContext: dataSourceLocationsRead,
		Schema: map[string]*schema.Schema{
			"locations": {
				Description: "List of `locations`.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"city": {
							Description: "The city of the location.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"country": {
							Description: "The country of the location.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"latitude": {
							Description: "The latitude of the location.",
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"longitude": {
							Description: "The longitude of the location.",
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"name": {
							Description: "The name of the location.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"state": {
							Description: "The state of the location.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status": {
							Description: "The status of the location. Valid values are `Operational`, `Degraded`, `Outage`, and `Maintenance`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLocationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	// // Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var locations []Location
	locations, err := client.GetLocations()
	if err != nil {
		return diag.FromErr(err)
	}
	err = locationsToResource(d, &locations)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func locationsToResource(d *schema.ResourceData, m *[]Location) error {
	err := d.Set("locations", flattenLocations(*m))
	if err != nil {
		return err
	}

	return nil
}

func flattenLocations(in []Location) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["city"] = v.City
		m["country"] = v.Country
		m["latitude"] = v.Latitude
		m["longitude"] = v.Longitude
		m["name"] = v.Name
		m["state"] = v.State
		m["status"] = v.Status
		out[i] = m
	}
	log.Printf("flattenLocations: %s", out)
	return out
}
