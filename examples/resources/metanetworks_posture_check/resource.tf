resource "metanetworks_metaport" "example" {
  name                  = "CrowdStrike Posture Check"
  description           = "Example Description"
  apply_on_org          = true
  osquery               = "select * from services where name='CSFalconService' and status='RUNNING';"
  platform              = "Windows"
  enabled               = true
  action                = "NONE"
  when                  = ["PRE_CONNECT"]
}
