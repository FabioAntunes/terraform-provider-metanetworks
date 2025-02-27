---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "metanetworks_group Resource - terraform-provider-metanetworks"
subcategory: ""
description: |-
  
---

# metanetworks_group (Resource)



## Example Usage

```terraform
resource "metanetworks_group" "example" {
  name = "example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The name of the group

### Optional

- **description** (String) The description of the group
- **expression** (String) Allows grouping entities by their tags. Filtering by tag value is also supported if provided. Supported operations: AND, OR, XOR, parenthesis.
- **roles** (Set of String) The group roles.
- **users** (Set of String) The group users.

### Read-Only

- **created_at** (String) Creation Timestamp.
- **id** (String) The ID of the group.
- **modified_at** (String) Modification Timestamp.
- **org_id** (String) The ID of the organization.
- **provisioned_by** (String) Groups can be provisioned in the system either by locally creating the groups from the Admin portal or API. Another, more common practice, is to provision groups from an organization directory service, by way of SCIM or LDAP protocols.


