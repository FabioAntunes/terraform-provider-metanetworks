resource "metanetworks_swg_threat_categories" "example" {
  name              = "example"
  description       = "example threat category"
  types             = ["Bot","Brute Forcer"]
  confidence_level  = "HIGH"
  risk_level        = "HIGH"
}
