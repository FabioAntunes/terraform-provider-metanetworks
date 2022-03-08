resource "metanetworks_peering" "example" {
  name    = "example"
}

resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_peering_attachment" "example" {
  peering_id         = metanetworks_peering.example.id
  network_element_id = metanetworks_mapped_service.example.id
}
