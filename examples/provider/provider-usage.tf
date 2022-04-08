terraform {
  required_providers {
    metanetworks = {
      source  = "FabioAntunes/metanetworks"
      version = "1.0.0-pre-2.4"
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
