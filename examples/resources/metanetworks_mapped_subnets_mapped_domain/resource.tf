resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.0.0/22"]
}


resource "metanetworks_mapped_subnets_mapped_domain" "example" {
  mapped_subnets_id = metanetworks_mapped_subnets.example.id
  name              = "ec2.internal"
  mapped_domain     = "ec2.internal"
  enterprise_dns    = true
}
