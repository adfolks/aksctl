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
	"fmt"

	"github.com/adfolks/aksctl/coreaksctl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createDiskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Create and manage Azure Managed Disks.",
	Long: `Create and manage Azure Managed Disks, it will use a random name and default resource group for the disk if not specified.
 	If you need to specify name or other resources use disk.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		viper.SetConfigName("default") // name of config file (without extension)
		viper.AddConfigPath(".")       // optionally look for config in the working directory
		err := viper.ReadInConfig()    // Find and read the config file
		if err != nil {                // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// viper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		// viper.SetDefault("rgroupName", "opsDefault") // for setting a default value

		diskName := viper.GetString("diskName") // getting values through viper
		diskResourceGroup := viper.GetString("diskResourceGroup")
		diskLocation := viper.GetString("diskLocation")
		diskSize := viper.GetString("diskSize")

		fmt.Println("diskName : ", diskName, ", ", "diskResourceGroup : ", diskResourceGroup, ", ", "diskLocation : ", diskLocation, ", ", "diskSize : ", diskSize)

		// coreaksctl.CreateDisk(diskName, diskResourceGroup, diskLocation, diskSize)
	},
}

// deleteClusterCmd represents the delete operation on cluster command
var deleteDiskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Delete an AKS disk",
	Long: `Delete an AKS disk, it would use a Random Name for disk.
	If you need to specify name or other resources use disk.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		viper.SetConfigName("default") // name of config file (without extension)
		viper.AddConfigPath(".")       // optionally look for config in the working directory
		err := viper.ReadInConfig()    // Find and read the config file
		if err != nil {                // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// viper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		viper.SetDefault("diskName", "opsDiskDefault") // for setting a default value

		diskName := viper.GetString("diskName") // getting values through viper
		diskResourceGroup := viper.GetString("diskResourceGroup")

		fmt.Println("diskName : ", diskName, ", ", "diskResourceGroup : ", diskResourceGroup)

		coreaksctl.DeleteDisk(diskName, diskResourceGroup)
	},
}

// updatediskCmd represents the update operation on disk command
var updatediskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Update an AKS disk",
	Long: `Update an AKS disk, it would use a Random Name for disk.
	If you need to specify name or other resources use cluster.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		viper.SetConfigName("default") // name of config file (without extension)
		viper.AddConfigPath(".")       // optionally look for config in the working directory
		err := viper.ReadInConfig()    // Find and read the config file
		if err != nil {                // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		viper.SetDefault("diskName", "opsDiskDefault") // for setting a default value

		diskName := viper.GetString("diskName") // getting values through viper
		diskResourceGroup := viper.GetString("diskResourceGroup")
		diskSize := viper.GetString("diskSize")

		fmt.Println("diskName : ", diskName, ", ", "diskResourceGroup : ", diskResourceGroup, ", ", "diskLocation : ", diskLocation)

		coreaksctl.UpdateDisk(diskName, diskResourceGroup, diskLocation)
	},
}

// getdiskCmd represents the get list of disk innresource group
var getClusterCmd = &cobra.Command{
	Use:   "disk",
	Short: "Get list of AKS disks",
	Long:  `Get list of AKS disks from a resource group.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		viper.SetConfigName("default") // name of config file (without extension)
		viper.AddConfigPath(".")       // optionally look for config in the working directory
		err := viper.ReadInConfig()    // Find and read the config file
		if err != nil {                // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		viper.SetDefault("diskResourceGroup", "DiskRGDefault") // for setting a default value

		diskResourceGroup := viper.GetString("diskResourceGroup")

		fmt.Println("diskResourceGroup : ", diskResourceGroup)

		coreaksctl.GetDisk(diskResourceGroup)
	},
}

func init() {

	//for create

	createCmd.AddCommand(createDiskCmd)
	createDiskCmd.PersistentFlags().StringP("diskName", "n", "temp", "disk name")
	createDiskCmd.PersistentFlags().StringP("diskResourceGroup", "g", "opsbrew", "disk resource group")
	createDiskCmd.PersistentFlags().StringP("diskLocation", "l", "westus", "disk location")
	createDiskCmd.PersistentFlags().StringP("diskSize", "z", "10", "disk size")
	viper.BindPFlags(createCmd.PersistentFlags())

	//for delete

	deleteCmd.AddCommand(deleteDiskCmd)
	deleteDiskCmd.PersistentFlags().StringP("name", "n", "opsFlagDefault", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	deleteDiskCmd.PersistentFlags().StringP("resourcegroup", "g", "opsFlagDefault", "disk resource group")

	viper.BindPDiskFlags(deleteCmd.PersistentFlags())

	//for update

	updateCmd.AddCommand(updateDiskCmd)
	updateDiskCmd.PersistentFlags().StringP("name", "n", "opsFlagDefault", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	updateDiskCmd.PersistentFlags().StringP("resourcegroup", "g", "opsFlagDefault", "disk resource group")
	updateDiskCmd.PersistentFlags().StringP("location", "r", "westus", "disk location")

	viper.BindPFlags(updateCmd.PersistentFlags())

	//for get List

	getCmd.AddCommand(getDiskCmd)
	getDiskCmd.PersistentFlags().StringP("resourcegroup", "g", "opsFlagDefault", "disk resource group")

	viper.BindPFlags(getDiskCmd.PersistentFlags())

	// Here you will define your flags and configuration settings
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")
	// is called directly, e.g.
	// Cobra supports local flags which will only run when this command