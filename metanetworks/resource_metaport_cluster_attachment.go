package metanetworks

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMetaportClusterAttachment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of this resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"metaport_cluster_id": {
				Description: "The ID of the Metaport Cluster.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"network_element_id": {
				Description: "The ID of the network element to attach to the Metaport Cluster.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
		Create: resourceMetaportClusterAttachmentCreate,
		Read:   resourceMetaportClusterAttachmentRead,
		Delete: resourceMetaportClusterAttachmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceMetaportClusterAttachmentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	metaporClustertID := d.Get("metaport_cluster_id").(string)

	metanetworksMutexKV.Lock(metaporClustertID)
	defer metanetworksMutexKV.Unlock(metaporClustertID)

	var metaportCluster *MetaportCluster
	metaportCluster, err := client.GetMetaPortCluster(metaporClustertID)
	if err != nil {
		return err
	}

	for i := 0; i < len(metaportCluster.MappedElements); i++ {
		if metaportCluster.MappedElements[i] == elementID {
			return errors.New("That network element is already mapped to this Metaport Cluster")
		}

	}

	metaportCluster.MappedElements = append(metaportCluster.MappedElements, elementID)
	_, err = client.UpdateMetaPortCluster(metaporClustertID, metaportCluster)
	if err != nil {
		return err
	}

	_, err = WaitMetaportClusterAttachmentCreate(client, metaporClustertID, elementID)

	if err != nil {
		return fmt.Errorf("Error waiting for metaport attachment creation (%s) (%s)", metaporClustertID, err)
	}

	d.SetId(fmt.Sprintf("%s_%s", metaporClustertID, elementID))

	return resourceMetaportClusterAttachmentRead(d, m)
}

func resourceMetaportClusterAttachmentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	id := d.Get("id").(string)
	ids := strings.Split(id, "_")
	if len(ids) != 2 {
		return fmt.Errorf("Error missing id for metaport cluster attachment got (%s)", id)
	}
	metaportClusterID := ids[0]
	elementID := ids[1]

	var metaportCluster *MetaportCluster
	metaportCluster, err := client.GetMetaPortCluster(metaportClusterID)
	if err != nil {
		return err
	}

	found := false
	for i := 0; i < len(metaportCluster.MappedElements); i++ {
		if metaportCluster.MappedElements[i] == elementID {
			found = true
			break
		}
	}

	// If not present we need to destroy the terraform resource so that it is recreated.
	if !found {
		d.SetId("")
	} else {
		d.Set("network_element_id", elementID)
		d.Set("metaport_cluster_id", metaportClusterID)
	}

	return nil
}

func resourceMetaportClusterAttachmentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	metaportClusterID := d.Get("metaport_cluster_id").(string)

	metanetworksMutexKV.Lock(metaportClusterID)
	defer metanetworksMutexKV.Unlock(metaportClusterID)

	var metaportCluster *MetaportCluster
	metaportCluster, err := client.GetMetaPortCluster(metaportClusterID)
	if err != nil {
		return err
	}

	// Note that if the entry has already been deleted this won't fail.
	for i := 0; i < len(metaportCluster.MappedElements); i++ {
		if metaportCluster.MappedElements[i] == elementID {
			metaportCluster.MappedElements = append(metaportCluster.MappedElements[:i], metaportCluster.MappedElements[i+1:]...)
			break
		}
	}

	err = resource.Retry(5*time.Second, func() *resource.RetryError {
		if _, err := client.UpdateMetaPortCluster(metaportClusterID, metaportCluster); err != nil {
			if !strings.Contains(err.Error(), "is busy. Try again later.") {
				return resource.NonRetryableError(err)
			}
			return resource.RetryableError(err)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("Error in Metaport Cluster attachment deletion (%s) (%s)", metaportClusterID, err)
	}

	return nil
}
