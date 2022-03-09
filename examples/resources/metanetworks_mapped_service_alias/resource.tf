resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_mapped_service_alias" "example" {
  mapped_service_id = metanetworks_mapped_service.example.id
  alias             = "example.com"
}
