package coreaksctl

import (
	"fmt"
	"log"
	"os/exec"
)

func CreateCluster(clusterName string, resourceGroupName string) {
	fmt.Println("Starting to set up your k8s Cluster")
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
	//Create AKS Cluster
	cmd := exec.Command("az", "aks", "create", "--name", clusterName,
		"--resource-group", resourceGroupName, "--node-count",
		"6", "--kubernetes-version", "1.11.3")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
}
