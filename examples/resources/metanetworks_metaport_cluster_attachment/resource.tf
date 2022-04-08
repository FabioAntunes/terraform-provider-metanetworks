resource "metanetworks_metaport_cluster" "example" {
  name = "example"
}


resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}


resource "metanetworks_metaport_cluster_attachment" "example" {
  metaport_id        = metanetworks_metaport_cluster.example.id
  network_element_id = metanetworks_mapped_service.example.id
}
