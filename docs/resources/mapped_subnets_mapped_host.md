---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "metanetworks_mapped_subnets_mapped_host Resource - terraform-provider-metanetworks"
subcategory: ""
description: |-
  
---

# metanetworks_mapped_subnets_mapped_host (Resource)



## Example Usage

```terraform
resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.0.0/22"]
}


resource "metanetworks_mapped_subnets_mapped_host" "example" {
  mapped_subnets_id = metanetworks_mapped_subnets.example.id
  name              = "ec2.internal"
  mapped_host       = "ec2.internal"
  ignore_bounds     = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **mapped_host** (String) Remote hostname or IP.
- **mapped_subnets_id** (String) ID of the mapped subnet network element.
- **name** (String) Mapped hostname.

### Optional

- **ignore_bounds** (Boolean) Allow setting mapped hosts outside of the defined mapped subnets, default=false.

### Read-Only

- **id** (String) The ID of this resource.


