resource "metanetworks_swg_url_filtering_rules" "example" {
  name        = "example"
  description = "example url filtering rule"
  action      = "ISOLATION"
  priority    = 1
  enabled     = true
  sources     = ["grp-exampleid"]
}
