package coreaksctl

import (
	"fmt"
	"log"
	"os/exec"
)

func CreateResourceGroup(resourceGroupName string, clusterRegion string) {
	cmd := exec.Command("az", "group", "create", "-l", clusterRegion, "-n",
		resourceGroupName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
}