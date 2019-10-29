package coreaksctl

import (
	"fmt"
	"log"
	"os/exec"
)

func CreateDisk(diskName string, diskRegion string, diskLocation string, diskSize string) {

	cmd := exec.Command("az", "disk", "create", "-g", diskRegion, "-n",
		diskName, "-l", diskLocation, "-z", diskSize)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
}
