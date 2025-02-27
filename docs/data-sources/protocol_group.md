---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "metanetworks_protocol_group Data Source - terraform-provider-metanetworks"
subcategory: ""
description: |-
  Returns a protocol_group of the organization.
---

# metanetworks_protocol_group (Data Source)

Returns a `protocol_group` of the organization.

## Example Usage

```terraform
data "metanetworks_protocol_group" "example" {
  name_regex = "example*"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name_regex** (String) A regex string to apply to the `protocol_group` list returned by Metanetworks. This allows more advanced filtering.

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **created_at** (String) Creation Timestamp.
- **description** (String) Description of the `protocol_group`.
- **modified_at** (String) Modification Timestamp.
- **name** (String) Name of the `protocol_group`.
- **org_id** (String) The ID of the organization.
- **protocols** (List of Object) List of `protocols`. (see [below for nested schema](#nestedatt--protocols))
- **read_only** (Boolean) If `protocol_group` is read only.

<a id="nestedatt--protocols"></a>
### Nested Schema for `protocols`

Read-Only:

- **from_port** (Number)
- **port** (Number)
- **proto** (String)
- **to_port** (Number)


