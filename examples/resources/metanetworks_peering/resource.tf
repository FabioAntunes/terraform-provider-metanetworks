resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.1.0/28"]
}

resource "metanetworks_peering" "example" {
  name     = " example"
  peers    = [
    metanetworks_mapped_service.example.id,
    metanetworks_mapped_subnets.example.id
  ]
}
