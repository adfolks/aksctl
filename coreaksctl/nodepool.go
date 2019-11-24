package coreaksctl

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CreateNodePool(clusterName string, nodePoolName string, rgroupName string, npNodeCount string) {

	cmd := exec.Command("az", "aks", "nodepool", "add", "--cluster-name", clusterName, "--name",
		nodePoolName, "--resource-group", rgroupName, "--node-count", npNodeCount)
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

func DeleteNodePool(clusterName string, nodePoolName string, rgroupName string) {

	cmd := exec.Command("az", "aks", "nodepool", "delete", "--cluster-name", clusterName, "--name",
		nodePoolName, "--resource-group", rgroupName)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Nodepool Deleted")
}

func ScaleNodePool(clusterName string, nodePoolName string, rgroupName string) {

	cmd := exec.Command("az", "aks", "nodepool", "scale", "--cluster-name", clusterName, "--name",
		nodePoolName, "--resource-group", rgroupName)
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

func UpdateNodePool(clusterName string, nodePoolName string, rgroupName string) {

	cmd := exec.Command("az", "aks", "nodepool", "update", "--cluster-name", clusterName, "--name",
		nodePoolName, "--resource-group", rgroupName)
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

func GetNodePool(clusterName string, rgroupName string) {

	cmd := exec.Command("az", "aks", "nodepool", "list", "--cluster-name", clusterName, "--resource-group", rgroupName)
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
