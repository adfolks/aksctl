package coreaksctl

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func CreateCluster(clusterName string, resourceGroupName string, extraflags []string) {
	fmt.Println("Starting to set up your k8s Cluster")
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
	//Create AKS Cluster
	var args = []string{"aks", "create", "--name", clusterName, "--resource-group", resourceGroupName}
	args = append(args, extraflags...)
	cmd := exec.Command("az", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed :" + fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	color.Green("Result: Cluster Created")
	color.Cyan("Fetching Credentials ........")
	GetClusterCredentials(clusterName, resourceGroupName)

}

func GetClusterCredentials(clusterName string, resourceGroupName string) {
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
	//Create AKS Cluster
	var args = []string{"aks", "get-credentials", "--name", clusterName, "--resource-group", resourceGroupName}
	cmd := exec.Command("az", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed :" + fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	color.Cyan(out.String())
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
	fmt.Println("Cluster Deleted")
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
	color.Green("Cluster Updated")
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
}
