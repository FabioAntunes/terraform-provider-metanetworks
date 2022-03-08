data "metanetworks_group" "example" {
  name = "example"
}

resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.0.0/22"]
}

resource "metanetworks_egress_route" "example" {
  name = "example"
  destinations = [
    "example.com",
  ]
  sources = [
    data.metanetworks_group.example.id
  ]
  via = metanetworks_mapped_subnets.example.id
}

data "metanetworks_group" "example" {
  name = "example"
}

data "metanetworks_locations" "all" {}

locals {
  locations = {
    for location in data.metanetworks_locations.all.locations :
    location.city => location
  }
}

resource "metanetworks_egress_route" "example" {
  name = "example"
  destinations = [
    "example.com",
  ]
  sources = [
    data.metanetworks_group.example.id
  ]
  via = local.locations["New York"].name
}
