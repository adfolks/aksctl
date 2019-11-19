package coreaksctl

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CreateResourceGroup(rgroupName string, rgroupRegion string) {

	cmd := exec.Command("az", "group", "create", "-l", rgroupRegion, "-n", rgroupName)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	fmt.Println("Result: " + out.String())
}

func DeleteResourceGroup(rgroupName string) {

	//Delete AKS ResourceGroup
	cmd := exec.Command("az", "group", "delete", "--name", rgroupName, "--yes")
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

func UpdateResourceGroup(rgroupName string) {

	//Update AKS ResourceGroup
	cmd := exec.Command("az", "group", "update", "--name", rgroupName)
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

func GetResourceGroup() {

	//List AKS ResourceGroup
	cmd := exec.Command("az", "group", "list")
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
