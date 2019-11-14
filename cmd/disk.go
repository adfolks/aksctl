// /*
// Copyright © 2019 NAME HERE <EMAIL ADDRESS>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
package cmd

import (
	"github.com/adfolks/aksctl/coreaksctl"
	"github.com/spf13/cobra"
)

var diskName, diskResourceGroup, diskLocation, diskSize string
var creatediskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Create and manage Azure Managed Disks.",
	Long: `Create and manage Azure Managed Disks, it will use a random name and default resource group for the disk if not specified.
 	If you need to specify name or other resources use disk.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		diskName, _ = cmd.Flags().GetString("name")
		diskResourceGroup, _ = cmd.Flags().GetString("resourcegroup")
		diskLocation, _ = cmd.Flags().GetString("location")
		diskSize, _ = cmd.Flags().GetString("size")
		coreaksctl.CreateDisk(diskName, diskResourceGroup, diskLocation, diskSize)
	},
}

func init() {
	createCmd.AddCommand(creatediskCmd)
	creatediskCmd.PersistentFlags().StringP("name", "n", "temp", "disk name")
	creatediskCmd.PersistentFlags().StringP("resourcegroup", "g", "opsbrew", "disk resource group")
	creatediskCmd.PersistentFlags().StringP("location", "l", "westus", "disk location")
	creatediskCmd.PersistentFlags().StringP("size", "z", "10", "disk size")

	// Here you will define your flags and configuration settings
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")
	// is called directly, e.g.
	// Cobra supports local flags which will only run when this command
}
