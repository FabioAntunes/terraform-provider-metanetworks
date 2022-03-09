resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.0.0/22"]
}


resource "metanetworks_mapped_subnets_mapped_host" "example" {
  mapped_subnets_id = metanetworks_mapped_subnets.example.id
  name              = "ec2.internal"
  mapped_host       = "ec2.internal"
  ignore_bounds     = true
}
