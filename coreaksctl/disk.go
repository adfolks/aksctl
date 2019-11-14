package coreaksctl

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CreateDisk(diskName string, diskResourcegroup string, diskLocation string, diskSize string) {
	cmd := exec.Command("az", "disk", "create", "-g", diskResourcegroup, "-n",
		diskName, "-l", diskLocation, "-z", diskSize)
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
