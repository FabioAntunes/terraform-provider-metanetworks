package metanetworks

import (
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	metaportClustersEndpoint string = "/v1/metaport_clusters"
)

// Metaport Cluster ...
type MetaportCluster struct {
	CreatedAt      string   `json:"created_at,omitempty" meta_api:"read_only"`
	Description    string   `json:"description"`
	ID             string   `json:"id,omitempty" meta_api:"read_only"`
	MappedElements []string `json:"mapped_elements,omitempty"`
	Metaports      []string `json:"metaports,omitempty"`
	ModifiedAt     string   `json:"modified_at,omitempty" meta_api:"read_only"`
	Name           string   `json:"name"`
}

func (c *Client) GetMetaPortCluster(metaportClusterID string) (*MetaportCluster, error) {
	var metaportCluster MetaportCluster
	err := c.Read(metaportClustersEndpoint+"/"+metaportClusterID+"?expand=true", &metaportCluster)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Metaport Cluster from Get: %s", metaportCluster.ID)
	return &metaportCluster, nil
}

func (c *Client) UpdateMetaPortCluster(metaportClusterID string, metaportCluster *MetaportCluster) (*MetaportCluster, error) {
	resp, err := c.Update(metaportClustersEndpoint+"/"+metaportClusterID, *metaportCluster)
	if err != nil {
		return nil, err
	}
	updatedMetaportCluster, _ := resp.(*MetaportCluster)

	log.Printf("Returning Metaport Cluster from Update: %s", updatedMetaportCluster.ID)
	return updatedMetaportCluster, nil
}

func (c *Client) CreateMetaPortCluster(metaportCluster *MetaportCluster) (*MetaportCluster, error) {
	resp, err := c.Create(metaportClustersEndpoint, *metaportCluster)
	if err != nil {
		return nil, err
	}

	createdMetaportCluster, ok := resp.(*MetaportCluster)
	if !ok {
		log.Printf("Returned Type is " + reflect.TypeOf(resp).Kind().String())
		return nil, errors.New("Object returned from API was not a Metaport Cluster Pointer")
	}

	log.Printf("Returning Metaport Cluster from Create: %s", createdMetaportCluster.ID)
	return createdMetaportCluster, nil
}

func (c *Client) DeleteMetaPortCluster(metaportClusterID string) error {
	err := c.Delete(metaportClustersEndpoint + "/" + metaportClusterID)
	if err != nil {
		return err
	}

	return nil
}

func StatusMetaportClusterAttachmentCreate(client *Client, metaportClusterID string, elementID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var metaportCluster *MetaportCluster
		metaportCluster, err := client.GetMetaPortCluster(metaportClusterID)
		if err != nil {
			return 0, "", err
		}

		for i := 0; i < len(metaportCluster.MappedElements); i++ {
			if metaportCluster.MappedElements[i] == elementID {
				return metaportCluster, "Completed", nil
			}
		}
		return metaportCluster, "Pending", nil
	}
}

func WaitMetaportClusterAttachmentCreate(client *Client, metaportClusterID string, elementID string) (*Client, error) {
	createStateConf := &resource.StateChangeConf{
		Pending:    []string{"Pending"},
		Target:     []string{"Completed"},
		Timeout:    5 * time.Minute,
		MinTimeout: 5 * time.Second,
		Delay:      3 * time.Second,
		Refresh:    StatusMetaportClusterAttachmentCreate(client, metaportClusterID, elementID),
	}

	_, err := createStateConf.WaitForState()
	if err != nil {
		return nil, err
	}

	return client, err
}
