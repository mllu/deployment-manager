locals {
  project_id = "cloud-ops-testing"
  region     = "us-central1"
}

provider "google" {
  project = local.project_id
  region  = local.region

  version = "~> 3.36"
}

provider "google-beta" {
  project = local.project_id
  region  = local.region

  version = "~> 3.36"
}
