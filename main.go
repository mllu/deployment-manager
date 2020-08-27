package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	dm "google.golang.org/api/deploymentmanager/v2"
)

func main() {
	dmService, err := dm.NewService(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	ds := dm.NewDeploymentsService(dmService)
	dmListRsp, err := ds.List("cloud-ops-testing").Do()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, deploy := range dmListRsp.Deployments {
		log.Printf("Deployment %s found", deploy.Name)
	}

	config, _ := ioutil.ReadFile("vm.yaml")
	deployment := &dm.Deployment{
		Name: "vm",
		Target: &dm.TargetConfiguration{
			Config: &dm.ConfigFile{
				Content: string(config),
			},
		},
	}
	dmInsertRsp, err := ds.Insert("cloud-ops-testing", deployment).Do()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(dmInsertRsp, err)
}
