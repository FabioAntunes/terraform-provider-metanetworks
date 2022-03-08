resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.1.0/28"]
}

data "metanetworks_protocol_group" "https" {
  name_regex = "HTTPS"
}

data "metanetworks_user" "example" {
  email = "user@example.com"
}

data "metanetworks_group" "example" {
  name = "example"
}

resource "metanetworks_policy" "example" {
  name            = " example"
  destinations    = [
    metanetworks_mapped_service.example.id,
    metanetworks_mapped_subnets.example.id
  ]
  protocol_groups = [
    metanetworks_protocol_group.https.id,
  ]
  sources         = [
    data.metanetworks_user.example.id,
    data.metanetworks_group.example.id
  ]
}
