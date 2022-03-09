package metanetworks

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	protocolGroupsEndpoint string = "/v1/protocol_groups"
)

// ProtocolGroup ...
type ProtocolGroup struct {
	Description string     `json:"description"`
	Name        string     `json:"name"`
	Protocols   []Protocol `json:"protocols,omitempty"`
	CreatedAt   string     `json:"created_at,omitempty"`
	ID          string     `json:"id,omitempty"`
	ModifiedAt  string     `json:"modified_at,omitempty"`
	OrgID       string     `json:"org_id,omitempty"`
	ReadOnly    bool       `json:"read_only,omitempty"`
}

// Protocol ...
type Protocol struct {
	FromPort int64  `json:"from_port" type:"integer"`
	Port     int64  `json:"port" type:"integer"`
	Protocol string `json:"proto"`
	ToPort   int64  `json:"to_port" type:"integer"`
}

func flattenProtocolGroup(pg ProtocolGroup) map[string]interface{} {
	out := make(map[string]interface{})
	protocols := make([]map[string]interface{}, len(pg.Protocols), len(pg.Protocols))

	out["description"] = pg.Description
	out["name"] = pg.Name
	for j, pv := range pg.Protocols {
		p := make(map[string]interface{})
		p["from_port"] = pv.FromPort
		p["port"] = pv.Port
		p["proto"] = pv.Protocol
		p["to_port"] = pv.ToPort

		protocols[j] = p
	}
	out["protocols"] = protocols
	out["created_at"] = pg.CreatedAt
	out["modified_at"] = pg.ModifiedAt
	out["org_id"] = pg.OrgID
	out["read_only"] = pg.ReadOnly
	return out
}

func flattenProtocolGroups(pg []ProtocolGroup) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(pg), len(pg))
	for i, v := range pg {
		out[i] = flattenProtocolGroup(v)
	}
	return out
}

func ProtocolGroupToResource(d *schema.ResourceData, m *ProtocolGroup) error {
	flattenedPG := flattenProtocolGroup(*m)

	for key, val := range flattenedPG {
		err := d.Set(key, val)
		if err != nil {
			return err
		}
	}

	d.SetId(m.ID)

	return nil
}

// protocolGroupToResource ...
func ProtocolGroupsToResource(d *schema.ResourceData, m *[]ProtocolGroup) error {
	err := d.Set("protocol_groups", flattenProtocolGroups(*m))

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return err
}

func (c *Client) GetProtocolGroups() ([]ProtocolGroup, error) {
	var protocolGroups []ProtocolGroup
	err := c.Read(protocolGroupsEndpoint, &protocolGroups)
	if err != nil {

		return nil, err
	}

	if len(protocolGroups) == 0 {
		return nil, fmt.Errorf("Protocol Groups Not found")
	}

	return protocolGroups, nil
}

// GetProtocolGroup ...
func (c *Client) GetProtocolGroup(protocolGroupID string) (*ProtocolGroup, error) {
	var protocolGroup ProtocolGroup
	err := c.Read(protocolGroupsEndpoint+"/"+protocolGroupID, &protocolGroup)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning ProtocolGroup from Get: %s", protocolGroup)
	return &protocolGroup, nil
}

// UpdateProtocolGroup ...
func (c *Client) UpdateProtocolGroup(protocolGroupID string, protocolGroup *ProtocolGroup) (*ProtocolGroup, error) {
	resp, err := c.Update(protocolGroupsEndpoint+"/"+protocolGroupID, *protocolGroup)
	if err != nil {
		return nil, err
	}
	updatedProtocolGroup, _ := resp.(*ProtocolGroup)

	log.Printf("Returning ProtocolGroup from Update: %s", updatedProtocolGroup.ID)
	return updatedProtocolGroup, nil
}

// CreateProtocolGroup ...
func (c *Client) CreateProtocolGroup(protocolGroup *ProtocolGroup) (*ProtocolGroup, error) {
	resp, err := c.Create(protocolGroupsEndpoint, *protocolGroup)
	if err != nil {
		return nil, err
	}

	createdProtocolGroup, ok := resp.(*ProtocolGroup)
	if !ok {
		log.Printf("Returned Type is " + reflect.TypeOf(resp).Kind().String())
		return nil, errors.New("Object returned from API was not a ProtocolGroup Pointer")
	}

	log.Printf("Returning ProtocolGroup from Create: %s", createdProtocolGroup.ID)
	return createdProtocolGroup, nil
}

// DeleteProtocolGroup ...
func (c *Client) DeleteProtocolGroup(protocolGroupID string) error {
	err := c.Delete(protocolGroupsEndpoint + "/" + protocolGroupID)
	if err != nil {
		return err
	}

	return nil
}

func StatusProtocolGroupCreate(client *Client, ProtocolGroupID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var metaport *MetaPort
		_, err := client.GetProtocolGroup(ProtocolGroupID)
		if err != nil {
			return metaport, "Pending", nil
		}
		return metaport, "Completed", nil
	}
}

func WaitProtocolGroupCreate(client *Client, ProtocolGroupID string) (*Client, error) {
	createStateConf := &resource.StateChangeConf{
		Pending:    []string{"Pending"},
		Target:     []string{"Completed"},
		Timeout:    30 * time.Second,
		MinTimeout: 5 * time.Second,
		Delay:      2 * time.Second,
		Refresh:    StatusProtocolGroupCreate(client, ProtocolGroupID),
	}

	_, err := createStateConf.WaitForState()
	if err != nil {
		return nil, err
	}

	return client, err
}
