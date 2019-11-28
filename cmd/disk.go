/*Package cmd is used for command line
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
*/
package cmd

import (
	"fmt"

	"github.com/adfolks/aksctl/pkg/ctl/disk"
	"github.com/adfolks/aksctl/pkg/ctl/resourcegroup"
	"github.com/adfolks/aksctl/pkg/ctl/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createDiskViper = viper.New()
var deleteDiskViper = viper.New()
var updateDiskViper = viper.New()
var getDiskViper = viper.New()

var createDiskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Create and manage Azure Managed Disks.",
	Long: `Create and manage Azure Managed Disks, it will use a random name and default resource group for the disk if not specified.
        If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		// Setting config file with viper

		createDiskViper.SetConfigName("default") // name of config file (without extension)
		createDiskViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := createDiskViper.ReadInConfig()    // Find and read the config file
		if err != nil {                          // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// viper.SetDefault("rgroupName", "opsDefault") // for setting a default value

		diskName := createDiskViper.GetString("managedDisk.name") // getting values through viper
		diskResourceGroup := createDiskViper.GetString("managedDisk.resource-group")
		diskLocation := createDiskViper.GetString("managedDisk.location")
		diskSize := createDiskViper.GetString("managedDisk.size-gb")

		color.Cyan("diskName : " + diskName + ", diskResourceGroup : " + diskResourceGroup + ", diskLocation : " + diskLocation + ", diskSize : " + diskSize)
		status := resourcegroup.CheckResourceGroup(diskResourceGroup)

		if !status {
			fmt.Println("Do you want to create a new resource group? (yes/no)")
			okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
			nokayResponses := []string{"n", "N", "no", "No", "NO"}
			message := "Please type yes or no and then press enter:"
			confirmation := utils.AskForConfirmation(okayResponses, nokayResponses, message)
			if confirmation {
				rgroupName := createDiskViper.GetString("managedDisk.resource-group") // getting values through viper
				rgroupRegion := createDiskViper.GetString("managedDisk.location")

				color.Cyan("rgroupName : " + rgroupName + ", rgroupRegion: " + rgroupRegion)

				resourcegroup.CreateResourceGroup(rgroupName, rgroupRegion)
				color.Green("Resource group created")
				disk.CreateDisk(diskName, diskResourceGroup, diskLocation, diskSize)
			} else {
				color.Red("Cannot create disk as the resource group does not exist")
			}
		} else {
			disk.CreateDisk(diskName, diskResourceGroup, diskLocation, diskSize)
		}
	},
}

// deleteClusterCmd represents the delete operation on cluster command
var deleteDiskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Delete an AKS disk",
	Long: `Delete an AKS disk with the specified disk name and resource group.
        If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		// Setting config file with viper

		deleteDiskViper.SetConfigName("default") // name of config file (without extension)
		deleteDiskViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := deleteDiskViper.ReadInConfig()    // Find and read the config file
		if err != nil {                          // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// viper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		// viper.SetDefault("diskName", "opsDiskDefault") // for setting a default value

		diskName := deleteDiskViper.GetString("managedDisk.name") // getting values through viper
		diskResourceGroup := deleteDiskViper.GetString("managedDisk.resource-group")

		color.Cyan("diskName : " + diskName + ", diskResourceGroup : " + diskResourceGroup)

		disk.DeleteDisk(diskName, diskResourceGroup)
	},
}

// updatediskCmd represents the update operation on disk command
var updateDiskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Update an AKS disk",
	Long: `Update an AKS disk with the specified disk name and resource group.
        If you need to specify name or other resources use cluster.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		// Setting config file with viper

		updateDiskViper.SetConfigName("default") // name of config file (without extension)
		updateDiskViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := updateDiskViper.ReadInConfig()    // Find and read the config file
		if err != nil {                          // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		//updateDiskViper.SetDefault("diskName", "opsDiskDefault") // for setting a default value

		diskName := updateDiskViper.GetString("managedDisk.name") // getting values through viper
		diskResourceGroup := updateDiskViper.GetString("managedDisk.resource-group")
		diskSize := updateDiskViper.GetString("managedDisk.size-gb")

		color.Cyan("diskName : " + diskName + ", diskResourceGroup : " + diskResourceGroup + ", diskSize : " + diskSize)

		disk.UpdateDisk(diskName, diskResourceGroup, diskSize)
	},
}

// getdiskCmd represents the get list of disk innresource group
var getDiskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Get list of AKS disks",
	Long:  `Get list of AKS disks from a resource group.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		getDiskViper.SetConfigName("default") // name of config file (without extension)
		getDiskViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := getDiskViper.ReadInConfig()    // Find and read the config file
		if err != nil {                       // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		//getDiskViper.SetDefault("diskResourceGroup", "DiskRGDefault") // for setting a default value

		diskResourceGroup := getDiskViper.GetString("managedDisk.resource-group")
		diskFilter, _ := cmd.Flags().GetString("flag")

		color.Cyan("diskResourceGroup : " + diskResourceGroup)

		disk.GetDisk(diskResourceGroup, diskFilter)
	},
}

func init() {

	//for create

	createCmd.AddCommand(createDiskCmd)
	createDiskCmd.PersistentFlags().StringP("name", "n", "temp", "disk name")
	createDiskCmd.PersistentFlags().StringP("rgroupname", "g", "opsbrew", "disk resource group")
	createDiskCmd.PersistentFlags().StringP("rgroupRegion", "l", "westus", "disk location")
	createDiskCmd.PersistentFlags().StringP("size", "z", "10", "disk size")

	createDiskViper.BindPFlag("managedDisk.name", createDiskCmd.PersistentFlags().Lookup("name"))
	createDiskViper.BindPFlag("managedDisk.resource-group", createDiskCmd.PersistentFlags().Lookup("rgroupname"))
	createDiskViper.BindPFlag("managedDisk.location", createDiskCmd.PersistentFlags().Lookup("rgroupRegion"))
	createDiskViper.BindPFlag("managedDisk.size-gb", createDiskCmd.PersistentFlags().Lookup("size"))

	//for delete

	deleteCmd.AddCommand(deleteDiskCmd)
	deleteDiskCmd.PersistentFlags().StringP("name", "n", "opsFlagDefault", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	deleteDiskCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefault", "disk resource group")

	deleteDiskViper.BindPFlag("managedDisk.name", deleteDiskCmd.PersistentFlags().Lookup("name"))
	deleteDiskViper.BindPFlag("managedDisk.resource-group", deleteDiskCmd.PersistentFlags().Lookup("rgroupname"))

	//for update

	updateCmd.AddCommand(updateDiskCmd)
	updateDiskCmd.PersistentFlags().StringP("name", "n", "opsFlagDefaultU", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	updateDiskCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefaultU", "disk resource group")
	updateDiskCmd.PersistentFlags().StringP("size", "r", "westus", "disk location")

	updateDiskViper.BindPFlag("managedDisk.name", updateDiskCmd.PersistentFlags().Lookup("name"))
	updateDiskViper.BindPFlag("managedDisk.resource-group", updateDiskCmd.PersistentFlags().Lookup("rgroupname"))
	updateDiskViper.BindPFlag("managedDisk.size-gb", updateDiskCmd.PersistentFlags().Lookup("size"))

	//for get List

	getCmd.AddCommand(getDiskCmd)
	getDiskCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefaultG", "disk resource group")
	getDiskCmd.PersistentFlags().StringP("flag", "l", "all", "filtr flag")

	getDiskViper.BindPFlag("managedDisk.resource-group", getDiskCmd.PersistentFlags().Lookup("rgroupname"))

	// Here you will define your flags and configuration settings
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")
	// is called directly, e.g.
	// Cobra supports local flags which will only run when this command
}
