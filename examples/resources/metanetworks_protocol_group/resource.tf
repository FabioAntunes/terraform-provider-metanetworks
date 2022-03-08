resource "metanetworks_protocol_group" "https" {
  name = "HTTPS"
  protocols {
    from_port = 443
    to_port   = 443
    proto     = "tcp"
  }
}
