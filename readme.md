

https://cloud.google.com/deployment-manager/docs/quickstart
API: https://cloud.google.com/deployment-manager/docs/reference/latest/deployments/insert

https://www.terraform.io/docs/providers/google/r/deployment_manager_deployment.html


Instructions

### To invoke deployment manager directly
```
$ gcloud deployment-manager deployments create vm --config vm.yaml

# delete deployment
# $ gcloud deployment-manager deployments delete vm
```

### To invoke deployment manager via API
```
$ go run .
```

### To invoke deployment manager via Terraform
```
$ terraform init
$ terraform plan -out=plan.tfstate
$ terraform apply plan.tfstate

# destroy resource
# $ terraform destroy
```
