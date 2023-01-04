# ---------------------------------
# locals
# ---------------------------------
locals {
  resource_prefix = "${var.team}-${var.env}-${var.app}"
  default_tags = {
    provisioner = "terraform"
    env         = var.env
    team        = var.team
    app         = var.app
  }
}

resource "random_string" "suffix" {
  length  = 8
  lower   = true
  upper   = false
  special = false
}

# ---------------------------------
# modules
# ---------------------------------
module "public_vpc" {
  source = "../../modules/vpc"

  vpc_name = "${local.resource_prefix}-${random_string.suffix.id}"
  tags     = local.default_tags
}