package coreaksctl

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func CreateResourceGroup(rgroupName string, rgroupRegion string) {

	fmt.Println("Creating resource group : " + rgroupName)
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")

	cmd := exec.Command("az", "group", "create", "-l", rgroupRegion, "-n", rgroupName)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	color.Green("Resource group Created")
}

func CheckResourceGroup(rgroupName string) bool {

	cmd := exec.Command("az", "group", "exists", "-n", rgroupName)
	var check bool
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	
	if out.String() == "true\n" {
		check=true
		fmt.Println("Using existing resource group")
	} else {
		check=false
		fmt.Println("The resource group", rgroupName, "does not exist.")
	}
	return check
}

func DeleteResourceGroup(rgroupName string) {

	//Delete AKS ResourceGroup

	fmt.Println("Deleting resource group : " + rgroupName)
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")

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
	fmt.Println("Resource group Deleted")
}

func UpdateResourceGroup(rgroupName string) {

	fmt.Println("Updating resource group : " + rgroupName)
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")

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
	color.Green("Resource Group Updated")
}

func GetResourceGroup() {
	fmt.Println("fetching resource group")
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
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
