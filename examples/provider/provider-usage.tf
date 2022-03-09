terraform {
  required_providers {
    metanetworks = {
      source  = "FabioAntunes/metanetworks"
      version = "1.0.0-alpha"
    }
  }
}

provider "metanetworks" {
  org = "example_organization"
}

# Example resource configuration
resource "metanetworks_resource" "example" {
  # ...
}
