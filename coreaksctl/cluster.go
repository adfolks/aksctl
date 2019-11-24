package coreaksctl

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

func CreateCluster(clusterName string, resourceGroupName string, extraflags []string) {
	fmt.Println("Starting to set up your k8s Cluster")
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
	//Create AKS Cluster
	var args = []string{"aks", "create", "--name", clusterName, "--resource-group", resourceGroupName}
	//handle the SSH keys generation in case of basic usage or the key value is not present on the config file
	fmt.Println(extraflags)
	_, found := Find(extraflags, "ssh-key-value")
	if !found {
		extraflags = append(extraflags, "--generate-ssh-keys")
	}
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
	color.Green("Cluster Created")
	color.Cyan("Fetching Credentials ........")
	GetClusterCredentials(clusterName, resourceGroupName)
	color.Green("Updated Kubeconfig")
    emoji.Println(":beer: Cheers!!!")
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

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
