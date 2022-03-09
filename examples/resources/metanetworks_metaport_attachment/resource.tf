resource "metanetworks_metaport" "example" {
  name    = "example"
  enabled = false
}


resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}


resource "metanetworks_metaport_attachment" "example" {
  metaport_id        = metanetworks_metaport.example.id
  network_element_id = metanetworks_mapped_service.example.id
}
