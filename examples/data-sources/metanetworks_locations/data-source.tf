data "metanetworks_locations" "all" {}

locals {
  locations = {
    for location in data.metanetworks_locations.all.locations :
    location.city => location
  }
}
