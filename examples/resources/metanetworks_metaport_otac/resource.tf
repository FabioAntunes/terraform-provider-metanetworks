resource "metanetworks_metaport" "example" {
  name    = "example"
  enabled = false
}


resource "metanetworks_metaport_otac" "example" {
  metaport_id = metanetworks_metaport.example.id
}
