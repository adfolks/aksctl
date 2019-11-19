package coreaksctl

import (
<<<<<<< HEAD
	"fmt"
	"log"
	"os/exec"
)

func createCluster(clusterName string, resourceGroupName string) {
=======
	"bytes"
	"fmt"
	"os/exec"
)

func CreateCluster(clusterName string, resourceGroupName string) {
>>>>>>> 8093a6af86a4c465096a34a0332c16dd85b49caf
	fmt.Println("Starting to set up your k8s Cluster")
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
	//Create AKS Cluster
	cmd := exec.Command("az", "aks", "create", "--name", clusterName,
		"--resource-group", resourceGroupName, "--node-count",
<<<<<<< HEAD
		"6", "--kubernetes-version", "1.11.3")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
=======
		"2", "--generate-ssh-keys")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}

func DeleteCluster(clusterName string, resourceGroupName string) {
	fmt.Println("Deleting your k8s Cluster " + clusterName)
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
	//Delete AKS Cluster
	cmd := exec.Command("az", "aks", "delete", "--name", clusterName,
		"--resource-group", resourceGroupName, "--yes")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}

func UpdateCluster(clusterName string, resourceGroupName string) {
	fmt.Println("Updating your k8s Cluster")
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
	//Create AKS Cluster
	cmd := exec.Command("az", "aks", "update", "--name", clusterName,
		"--resource-group", resourceGroupName, "--enable-cluster-autoscaler")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}

func GetCluster(resourceGroupName string) {
	fmt.Println("Collection your k8s Cluster informations")
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
	//Create AKS Cluster
	cmd := exec.Command("az", "aks", "list",
		"--resource-group", resourceGroupName)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
>>>>>>> 8093a6af86a4c465096a34a0332c16dd85b49caf
}
