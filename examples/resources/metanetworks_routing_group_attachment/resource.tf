data "metanetworks_group" "example" {
  name = "example"
}

resource "metanetworks_routing_group" "example" {
  name    = "example"
  sources = [
    data.metanetworks_group.example.id
  ]
}

resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_routing_group_attachment" "example" {
  routing_group_id   = metanetworks_routing_group.example.id
  network_element_id = metanetworks_mapped_service.example.id
}
