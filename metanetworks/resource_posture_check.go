package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePostureCheck() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the posture check.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "The description of the posture check.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the posture check.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"action": {
				Description: "What happens when a posture check is failed. Values: `DISCONNECT`, `NONE`",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"osquery": {
				Description: "OSQuery string to perform posture check.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"platform": {
				Description: "Platform that the posture check is for. Values: `Android`, `macOS`, `iOS`, `Linux`, `Windows`, `ChromeOS`.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"user_message_on_fail": {
				Description: "Failure message to display when posture check fails.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"interval": {
				Description: "Required if `when` contains `PERIODIC`). Time in *minutes* between checks. Values: `5-60`",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"check": {
				Description: "Templated scenario to posture check for.",
				Type:        schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_version": {
							Description: "String of the version.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"type": {
							Description: "Values: `jailbroken_rooted`, `screen_lock_enabled`, `minimum_app_version`, `minimum_os_version`, `malicious_app_detection`, `developer_mode_enabled`.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
				Optional: true,
			},
			"enabled": {
				Description: "default=true",
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
			},
			"apply_to_org": {
				Description: "Required if `sources` is omitted). Applies setting to entire organization.",
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
			},
			"exempt_sources": {
				Description: "Sources to exclude from posture check.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"sources": {
				Description: "Required if `apply_on_org` is omitted). Applies setting to specified sources.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"when": {
				Description: "When the posture check should run. Values: `PRE_CONNECT`, `PERIODIC`.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
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
		},
		Create: resourcePostureCheckCreate,
		Read:   resourcePostureCheckRead,
		Update: resourcePostureCheckUpdate,
		Delete: resourcePostureCheckDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourcePostureCheckCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	action := d.Get("action").(string)
	osQuery := d.Get("osquery").(string)
	platform := d.Get("platform").(string)
	userMessageOnFail := d.Get("user_message_on_fail").(string)
	enabled := d.Get("enabled").(bool)
	applyToOrg := d.Get("apply_to_org").(bool)
	interval := d.Get("interval").(int)
	check := d.Get("check").([]interface{})
	when := resourceTypeSetToStringSlice(d.Get("when").(*schema.Set))
	applyToEntities := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	exemptEntities := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))

	postureCheck := PostureCheck{
		Name:              name,
		Description:       description,
		Action:            action,
		OSQuery:           osQuery,
		Platform:          platform,
		UserMessageOnFail: userMessageOnFail,
		Enabled:           enabled,
		ApplyToOrg:        applyToOrg,
		Interval:          interval,
		Check:             check,
		When:              when,
		ExemptEntities:    exemptEntities,
		ApplyToEntities:   applyToEntities,
	}

	var newPostureCheck *PostureCheck
	newPostureCheck, err := client.CreatePostureCheck(&postureCheck)
	if err != nil {
		return err
	}

	d.SetId(newPostureCheck.ID)

	err = postureCheckToResource(d, newPostureCheck)
	if err != nil {
		return err
	}

	return resourcePostureCheckRead(d, m)
}

func resourcePostureCheckRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	postureCheck, err := client.GetPostureCheck(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = postureCheckToResource(d, postureCheck)
	if err != nil {
		return err
	}

	return nil
}

func resourcePostureCheckUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	action := d.Get("action").(string)
	osQuery := d.Get("osquery").(string)
	platform := d.Get("platform").(string)
	userMessageOnFail := d.Get("user_message_on_fail").(string)
	enabled := d.Get("enabled").(bool)
	applyToOrg := d.Get("apply_to_org").(bool)
	interval := d.Get("interval").(int)
	check := d.Get("check").([]interface{})
	when := resourceTypeSetToStringSlice(d.Get("when").(*schema.Set))
	applyToEntities := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	exemptEntities := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))

	postureCheck := PostureCheck{
		Name:              name,
		Description:       description,
		Action:            action,
		OSQuery:           osQuery,
		Platform:          platform,
		UserMessageOnFail: userMessageOnFail,
		Enabled:           enabled,
		ApplyToOrg:        applyToOrg,
		Interval:          interval,
		Check:             check,
		When:              when,
		ExemptEntities:    exemptEntities,
		ApplyToEntities:   applyToEntities,
	}

	var updatedPostureCheck *PostureCheck
	updatedPostureCheck, err := client.UpdatePostureCheck(d.Id(), &postureCheck)
	if err != nil {
		return err
	}

	d.SetId(updatedPostureCheck.ID)

	err = postureCheckToResource(d, updatedPostureCheck)
	if err != nil {
		return err
	}

	return resourcePostureCheckRead(d, m)
}

func resourcePostureCheckDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeletePostureCheck(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func postureCheckToResource(d *schema.ResourceData, m *PostureCheck) error {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("action", m.Action)
	d.Set("osquery", m.OSQuery)
	d.Set("platform", m.Platform)
	d.Set("enabled", m.Enabled)
	d.Set("apply_to_org", m.ApplyToOrg)
	d.Set("user_message_on_fail", m.UserMessageOnFail)
	d.Set("interval", m.Interval)
	d.Set("check", m.Check)
	d.Set("when", m.When)
	d.Set("exempt_entities", m.ExemptEntities)
	d.Set("apply_to_entities", m.ApplyToEntities)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)

	d.SetId(m.ID)

	return nil
}
