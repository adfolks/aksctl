/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// clusterCmd represents the cluster command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create an AKS cluster",
	Long: `Create an AKS cluster, it would use a Random Name for cluster.
	If you need to specify name or other resources use cluster.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	createCmd.AddCommand(clusterCmd)

}

func createCluster(clusterName string, resourceGroupName string) {
	fmt.Println("Starting to set up your k8s Cluster")
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
	//Create AKS Cluster
	cmd := exec.Command("az", "aks", "create", "--name", clusterName,
		"--resource-group", resourceGroupName, "--node-count",
		"6", "--kubernetes-version", "1.11.3")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
}