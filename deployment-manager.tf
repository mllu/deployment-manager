resource "google_deployment_manager_deployment" "vm" {
  name = "vm"

  target {
    config {
      content = file("vm.yaml")
    }
  }
}
