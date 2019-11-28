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

	"github.com/adfolks/aksctl/pkg/ctl/resourcegroup"
	"github.com/adfolks/aksctl/pkg/ctl/utils"
	"github.com/adfolks/aksctl/pkg/ctl/vnet"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createVNetViper = viper.New()
var deleteVNetViper = viper.New()

// var updateNPViper = viper.New()
// var scaleNPViper = viper.New()
// var getNPViper = viper.New()

var createVNetCmd = &cobra.Command{
	Use:   "vnet",
	Short: "Create a Virtual Network.",
	Long: `Create and manage Virtual Networks, it will use a random name, a default resource group for the Virtual Network if not specified.
        If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		// Setting config file with viper

		createVNetViper.SetConfigName("default") // name of config file (without extension)
		createVNetViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := createVNetViper.ReadInConfig()    // Find and read the config file
		if err != nil {                          // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		//createViper.SetDefault("vNetName", "vNetDefault") // for setting a default value

		vNetName := createVNetViper.GetString("vnet.name") // getting values through viper
		rgroupName := createVNetViper.GetString("vnet.resource-group")

		fmt.Println("vNetName : ", vNetName, ", ", "rgroupName : ", rgroupName)

		status := resourcegroup.CheckResourceGroup(rgroupName)
		fmt.Println("status =", status)
		if status {
			fmt.Println("Do you want to create a new resource group? (yes/no)")
			okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
			nokayResponses := []string{"n", "N", "no", "No", "NO"}
			message := "Please type yes or no and then press enter:"
			confirmation := utils.AskForConfirmation(okayResponses, nokayResponses, message)
			if confirmation {
				rgroupName := createVNetViper.GetString("vnet.resource-group") // getting values through viper
				rgroupRegion := createVNetViper.GetString("metadata.location")

				fmt.Println("rgroupName : ", rgroupName, ", ", "rgroupRegion: ", rgroupRegion)

				resourcegroup.CreateResourceGroup(rgroupName, rgroupRegion)
				fmt.Println("Resource group created")
				fmt.Println("vNetName : ", vNetName, ", ", "rgroupName : ", rgroupName)
				vnet.CreateVNet(vNetName, rgroupName)
			} else {
				fmt.Println("The resource group does not exist")
			}
		} else {
			fmt.Println("vNetName : ", vNetName, ", ", "rgroupName : ", rgroupName)
			vnet.CreateVNet(vNetName, rgroupName)
		}
	},
}

var deleteVNetCmd = &cobra.Command{
	Use:   "vnet",
	Short: "Delete a Virtual Networks.",
	Long: `Delete Virtual Network with the specified name and resource group.
        If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		deleteVNetViper.SetConfigName("default") // name of config file (without extension)
		deleteVNetViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := deleteVNetViper.ReadInConfig()    // Find and read the config file
		if err != nil {                          // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		//deleteVNetViper.SetDefault("vNetName", "vNetDefault") // for setting a default value

		vNetName := deleteVNetViper.GetString("vnet.name") // getting values through viper
		rgroupName := deleteVNetViper.GetString("vnet.resource-group")

		fmt.Println("vNetName : ", vNetName, ", ", "rgroupName : ", rgroupName)

		vnet.DeleteVNet(vNetName, rgroupName)

	},
}

func init() {

	//for create

	createCmd.AddCommand(createVNetCmd)
	createVNetCmd.PersistentFlags().StringP("vnetname", "n", "vNetFlag", "VNet name")
	createVNetCmd.PersistentFlags().StringP("rgroupname", "g", "rgFlag", "VNet resource group")

	createNPViper.BindPFlag("metadata.resource-group", createVNetCmd.PersistentFlags().Lookup("rgroupname"))
	createNPViper.BindPFlag("vnet.name", createVNetCmd.PersistentFlags().Lookup("vnetname"))

	//for delete

	deleteCmd.AddCommand(deleteVNetCmd)
	deleteVNetCmd.PersistentFlags().StringP("vnetname", "n", "vNetFlag", "VNet name")
	deleteVNetCmd.PersistentFlags().StringP("rgroupname", "g", "rgFlag", "VNet resource group")

	deleteNPViper.BindPFlag("metadata.resource-group", deleteVNetCmd.PersistentFlags().Lookup("rgroupname"))
	deleteNPViper.BindPFlag("vnet.name", deleteVNetCmd.PersistentFlags().Lookup("vnetname"))

}
