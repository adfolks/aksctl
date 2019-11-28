/*Package cmd is used for command line
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

	"github.com/adfolks/aksctl/pkg/ctl/resourcegroup"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createRGViper = viper.New()
var deleteRGViper = viper.New()
var updateRGViper = viper.New()
var getRGViper = viper.New()

var createResourceGroupCmd = &cobra.Command{
	Use:   "resourcegroup",
	Short: "Create and manage an AKS resource group",
	Long: `Create and manage an AKS resource group, it would use a random name and default region for resource group.
	If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		// Setting config file with viper

		createRGViper.SetConfigName("default") // name of config file (without extension)
		createRGViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := createRGViper.ReadInConfig()    // Find and read the config file
		if err != nil {                        // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// createRGViper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		//createRGViper.SetDefault("rgroupName", "rgDefault") // for setting a default value

		rgroupName := createRGViper.GetString("metadata.resource-group") // getting values through viper
		rgroupRegion := createRGViper.GetString("metadata.location")

		color.Cyan("rgroupName : " + rgroupName + ", rgroupRegion: " + rgroupRegion)

		resourcegroup.CreateResourceGroup(rgroupName, rgroupRegion)
	},
}

// deleteResourceGroupCmd represents the delete operation on disk command
var deleteResourceGroupCmd = &cobra.Command{
	Use:   "resourcegroup",
	Short: "Delete an AKS resource group",
	Long: `Delete an AKS resource group with the specified resource group name.
	If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		// Setting config file with viper

		deleteRGViper.SetConfigName("default") // name of config file (without extension)
		deleteRGViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := deleteRGViper.ReadInConfig()    // Find and read the config file
		if err != nil {                        // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// deleteViper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		//deleteViper.SetDefault("rgroupName", "rgroupDefault") // for setting a default value

		rgroupName := deleteRGViper.GetString("metadata.resource-group") // getting values through viper

		color.Cyan("rgroupName : " + rgroupName)

		resourcegroup.DeleteResourceGroup(rgroupName)
	},
}

// updateResourceGroupCmd represents the update operation on ResourceGroup command
var updateResourceGroupCmd = &cobra.Command{
	Use:   "resourcegroup",
	Short: "Update an AKS ResourceGroup",
	Long: `Update an AKS ResourceGroup with the specified resource group name.
	If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		// Setting config file with viper

		updateRGViper.SetConfigName("default") // name of config file (without extension)
		updateRGViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := updateRGViper.ReadInConfig()    // Find and read the config file
		if err != nil {                        // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		//updateViper.SetDefault("rgroupName", "rgNameDefault") // for setting a default value

		rgroupName := updateRGViper.GetString("metadata.resource-group") // getting values through viper

		color.Cyan("rgroupName : " + rgroupName)

		resourcegroup.UpdateResourceGroup(rgroupName)
	},
}

// getResourceGroupCmd represents the list operation resource group command
var getResourceGroupCmd = &cobra.Command{
	Use:   "resourcegroup",
	Short: "Get list of AKS resource groups",
	Long:  `Get list of AKS resource groups that are available.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Setting config file with viper

		getRGViper.SetConfigName("default") // name of config file (without extension)
		getRGViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := getRGViper.ReadInConfig()    // Find and read the config file
		if err != nil {                     // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		rGroupFilter, _ := cmd.Flags().GetString("flag")
		resourcegroup.GetResourceGroup(rGroupFilter)
	},
}

func init() {

	//for create

	createCmd.AddCommand(createResourceGroupCmd)
	createResourceGroupCmd.PersistentFlags().StringP("name", "n", "rgnameflagdefault", "resource group name")
	createResourceGroupCmd.PersistentFlags().StringP("rgroupregion", "r", "westus", "resource group region")

	createRGViper.BindPFlag("metadata.resource-group", createResourceGroupCmd.PersistentFlags().Lookup("name"))
	createRGViper.BindPFlag("metadata.location", createResourceGroupCmd.PersistentFlags().Lookup("rgroupregion"))
	//viper.BindPFlags(createCmd.PersistentFlags())

	//for delete

	deleteCmd.AddCommand(deleteResourceGroupCmd)
	deleteResourceGroupCmd.PersistentFlags().StringP("name", "n", "rgnameflagdefault", "resource group name") //  fullyQualifiedName,shorthand,default,description

	deleteRGViper.BindPFlag("metadata.resource-group", deleteResourceGroupCmd.PersistentFlags().Lookup("name"))
	//viper.BindPFlags(deleteCmd.PersistentFlags())

	//for update

	updateCmd.AddCommand(updateResourceGroupCmd)
	updateResourceGroupCmd.PersistentFlags().StringP("name", "n", "rgnameflagdefault", "resource group name") //  fullyQualifiedName,shorthand,default,description

	updateRGViper.BindPFlag("metadata.resource-group", updateResourceGroupCmd.PersistentFlags().Lookup("name"))

	//viper.BindPFlags(updateCmd.PersistentFlags())

	//for get List

	getCmd.AddCommand(getResourceGroupCmd)

	getResourceGroupCmd.PersistentFlags().StringP("flag", "l", "all", "filtr flag")
	getRGViper.BindPFlags(getResourceGroupCmd.PersistentFlags())
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
