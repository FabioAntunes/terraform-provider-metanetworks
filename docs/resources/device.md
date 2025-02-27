---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "metanetworks_device Resource - terraform-provider-metanetworks"
subcategory: ""
description: |-
  
---

# metanetworks_device (Resource)



## Example Usage

```terraform
data "metanetworks_user" "example" {
  email = "user@example.com"
}


resource "metanetworks_device" "example" {
  name     = "example"
  owner_id = data.metanetworks_user.example.id,
  platform = "macOS"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The name of the device
- **owner_id** (String) The ID of owner of the device.
- **platform** (String) The platform of the device. Valid values are `Android`, `macOS`, `iOS`, `Linux`, `Windows` and `ChromeOS`.

### Optional

- **description** (String) The description of the device
- **enabled** (Boolean) default=true
- **tags** (Map of String) Tags are key/value attributes that can be used to group elements together.

### Read-Only

- **aliases** (Set of String) The domain names of the device.
- **created_at** (String) Creation Timestamp.
- **dns_name** (String) `<network_element_id>`.`<org_id>`.nsof
- **expires_at** (String) Expiration Timestamp.
- **id** (String) The ID of the device.
- **modified_at** (String) Modification Timestamp.
- **org_id** (String) The ID of the organization.


