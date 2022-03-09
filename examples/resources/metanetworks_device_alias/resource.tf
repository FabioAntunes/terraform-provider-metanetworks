data "metanetworks_user" "example" {
  email = "user@example.com"
}

resource "metanetworks_device" "example" {
  name     = "example"
  owner_id = data.metanetworks_user.example.id,
  platform = "macOS"
}

resource "metanetworks_device_alias" "example" {
  device_id = metanetworks_device.example.id
  alias     = "example.com"
}
