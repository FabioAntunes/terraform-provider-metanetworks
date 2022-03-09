data "metanetworks_group" "example" {
  name = "example"
}

resource "metanetworks_routing_group" "example" {
  name    = "example"
  sources = [
    data.metanetworks_group.example.id
  ]
}
