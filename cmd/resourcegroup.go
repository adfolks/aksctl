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

	"github.com/adfolks/aksctl/coreaksctl"
	"github.com/spf13/cobra"
)

var rgroupName, rgroupRegion string

var resourcegroupCmd = &cobra.Command{
	Use:   "resourcegroup",
	Short: "Create an AKS resource group",
	Long: `Create an AKS resource group, it would use a Random Name for resource group.
	If you need to specify name or other resources use resourcegroup.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		rgroupName, _ = cmd.Flags().GetString("name")
		rgroupRegion, _ = cmd.Flags().GetString("region")
		coreaksctl.CreateResourceGroup(rgroupName, rgroupRegion)
		fmt.Println("Resource group created")
	},
}

func init() {
	createCmd.AddCommand(resourcegroupCmd)
	resourcegroupCmd.PersistentFlags().StringP("name", "n", "temp", "resource group name")
	resourcegroupCmd.PersistentFlags().StringP("region", "r", "westus", "resource group location")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
