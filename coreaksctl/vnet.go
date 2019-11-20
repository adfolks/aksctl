package coreaksctl

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CreateVNet(vNetName string, rgroupName string) {

	cmd := exec.Command("az", "network", "vnet", "create", "--name",
		vNetName, "--resource-group", rgroupName)
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

func DeleteVNet(vNetName string, rgroupName string) {

	cmd := exec.Command("az", "network", "vnet", "delete", "--name",
		vNetName, "--resource-group", rgroupName)
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
