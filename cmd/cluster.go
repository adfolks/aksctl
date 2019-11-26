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
	"os"

	"github.com/adfolks/aksctl/pkg/ctl/cluster"
	"github.com/adfolks/aksctl/pkg/ctl/resourcegroup"
	"github.com/adfolks/aksctl/pkg/ctl/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clusterCmd represents the cluster command

var createViper = viper.New()
var deleteViper = viper.New()
var updateViper = viper.New()
var getViper = viper.New()
var credViper = viper.New()
var cfgFilef string = "default"

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create an AKS cluster",
	Long: `Create an AKS cluster, it would use a Random Name for cluster.
	If you need to specify name or other resources use cluster.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		createViper.SetConfigName(cfgFilef) // name of config file (without extension)
		createViper.AddConfigPath(".")      // optionally look for config in the working directory
		err := createViper.ReadInConfig()   // Find and read the config file
		if err != nil {                     // Handle errors reading the config file
			
			color.Red("Config yaml file doesn't exist")
			fmt.Println("Do you want to create a default yaml file? (yes/no)")

			okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
			nokayResponses := []string{"n", "N", "no", "No", "NO"}
			message := "Please type yes or no and then press enter:"
			confirmation := utils.AskForConfirmation(okayResponses,nokayResponses,message)
			if confirmation == true {
					_,errc := os.Create(cfgFilef+".yaml")
					if errc != nil {
						color.Red("Error creating default yaml try creating it mannually");
						os.Exit(0)
					}
					createViper.SetConfigType("yaml")
					createViper.Set("metadata.name", "defaulCluster")
					createViper.Set("metadata.resource-group", "defaultRGroup")
					createViper.Set("metadata.location", "eastus")
					errb := createViper.WriteConfig()
					if errb != nil {
						fmt.Print("Error : ",errb);
					}
					color.Green("Default config yaml generated")
				} else {
					color.Red("Can't continue without yaml file")
					os.Exit(0)
				}
			
			
		}

		/*
			viper default value will be prior than Flag default
			so value selection priority order is
				- Flag Value
				- Config File
				- Vipro Default
				- Flag Default
		*/
		// viper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		clusterName := createViper.GetString("metadata.name")
		rgroupName := createViper.GetString("metadata.resource-group") // getting values through viper
		rgroupRegion := createViper.GetString("metadata.location")
		aadServerAppId := createViper.GetStringMap("metadata")

		color.Cyan("rgroupName : " + rgroupName + ", rgroupRegion : " + rgroupRegion + ", clusterName : " + clusterName)
		var extraflags []string
		for k, v := range aadServerAppId {
			if k != "name" && k != "resource-group" && k != "location" {
				extraflags = append(extraflags, "--"+k)
				if fmt.Sprintf("%v", v) != "nil" {
					extraflags = append(extraflags, fmt.Sprintf("%v", v))
				}
			}
		}
		status := resourcegroup.CheckResourceGroup(rgroupName)
		if status == false {
			color.Red("Resource group doesn't exist")
			fmt.Println("Do you want to create a new resource group? (yes/no)")
			okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
			nokayResponses := []string{"n", "N", "no", "No", "NO"}
			message := "Please type yes or no and then press enter:"
			confirmation := utils.AskForConfirmation(okayResponses,nokayResponses,message)
			if confirmation == true {

				rgroupName := createViper.GetString("metadata.resource-group") // getting values through viper
				rgroupRegion := createViper.GetString("metadata.location")

				color.Cyan("rgroupName : " + rgroupName + ", rgroupRegion: " + rgroupRegion)

				resourcegroup.CreateResourceGroup(rgroupName, rgroupRegion)

				cluster.CreateCluster(clusterName, rgroupName, extraflags)
			} else {
				color.Red("Cannot create cluster as the resource group does not exist")
			}
		} else {
			cluster.CreateCluster(clusterName, rgroupName, extraflags)
		}
	},
}

// deleteClusterCmd represents the delete operation on cluster command
var deleteClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Delete an AKS cluster",
	Long: `Delete an AKS cluster, it would use a Random Name for cluster.
	If you need to specify name or other resources use cluster.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper
		// deleteViper.SetDefault("rgroupName", "opsDefault") // for setting a default value

		deleteViper.SetConfigName("default") // name of config file (without extension)
		deleteViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := deleteViper.ReadInConfig()    // Find and read the config file
		if err != nil {                      // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		clusterName := deleteViper.GetString("metadata.name")
		rgroupName := deleteViper.GetString("metadata.resource-group") // getting values through viper

		color.Cyan("rgroupName : " + rgroupName + ", clusterName : " + clusterName)

		cluster.DeleteCluster(clusterName, rgroupName)
	},
}

// updateClusterCmd represents the update operation on cluster command
var updateClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Update an AKS cluster",
	Long: `Update an AKS cluster, it would use a Random Name for cluster.
	If you need to specify name or other resources use cluster.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		updateViper.SetConfigName("default") // name of config file (without extension)
		updateViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := updateViper.ReadInConfig()    // Find and read the config file
		if err != nil {                      // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		clusterName := updateViper.GetString("metadata.name")
		rgroupName := updateViper.GetString("metadata.resource-group") // getting values through viper

		color.Cyan("rgroupName : " + rgroupName + ", clusterName : " + clusterName)

		cluster.UpdateCluster(clusterName, rgroupName)
	},
}

// getClusterCmd represents the get list of cluster innresource group
var getClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Get list of AKS cluster",
	Long:  `Get list of AKS cluster from a resource group.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Setting config file with viper
		// getViper.SetDefault("rgroupName", "opsDefault") // for setting a default value

		getViper.SetConfigName("default") // name of config file (without extension)
		getViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := getViper.ReadInConfig()    // Find and read the config file
		if err != nil {                   // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		rgroupName := getViper.GetString("metadata.resource-group") // getting values through viper

		color.Cyan("rgroupName : " + rgroupName)

		cluster.GetCluster(rgroupName)
	},
}


var getCredentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "Get list of AKS cluster",
	Long:  `Get list of AKS cluster from a resource group.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Setting config file with viper
		// getViper.SetDefault("rgroupName", "opsDefault") // for setting a default value

		credViper.SetConfigName("default") // name of config file (without extension)
		credViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := credViper.ReadInConfig()    // Find and read the config file
		if err != nil {                   // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		name := credViper.GetString("metadata.name") // getting values through viper
		rgroupName := credViper.GetString("metadata.resource-group") // getting values through viper

		color.Cyan("cluster name : " + name)

		cluster.GetClusterCredentials(name, rgroupName)
	},
}

func init() {

	//for create

	createCmd.AddCommand(clusterCmd)
	clusterCmd.PersistentFlags().StringP("name", "n", "opsFlagDefaultCreate", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	clusterCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefaultC", "disk resource group")
	clusterCmd.PersistentFlags().StringP("rgroupRegion", "r", "westus", "disk location")
	clusterCmd.PersistentFlags().StringVar(&cfgFilef, "file", "default", "default")

	createViper.BindPFlag("metadata.name", clusterCmd.PersistentFlags().Lookup("name"))
	createViper.BindPFlag("metadata.resource-group", clusterCmd.PersistentFlags().Lookup("rgroupname"))
	createViper.BindPFlag("metadata.location", clusterCmd.PersistentFlags().Lookup("rgroupRegion"))

	//for delete

	deleteCmd.AddCommand(deleteClusterCmd)
	deleteClusterCmd.PersistentFlags().StringP("name", "n", "opsFlagDefaultDelete", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	deleteClusterCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefault", "disk resource group")

	deleteViper.BindPFlag("metadata.name", deleteClusterCmd.PersistentFlags().Lookup("name"))
	deleteViper.BindPFlag("metadata.resource-group", deleteClusterCmd.PersistentFlags().Lookup("rgroupname"))

	//for update

	updateCmd.AddCommand(updateClusterCmd)
	updateClusterCmd.PersistentFlags().StringP("name", "n", "opsFlagDefaultUpdate", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	updateClusterCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefault", "disk resource group")
	updateClusterCmd.PersistentFlags().StringP("rgroupRegion", "r", "westus", "disk location")

	updateViper.BindPFlag("metadata.name", updateClusterCmd.PersistentFlags().Lookup("name"))
	updateViper.BindPFlag("metadata.resource-group", updateClusterCmd.PersistentFlags().Lookup("rgroupname"))
	updateViper.BindPFlag("metadata.location", updateClusterCmd.PersistentFlags().Lookup("rgroupRegion"))

	//for get List

	getCmd.AddCommand(getClusterCmd)
	getClusterCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefault", "disk resource group")

	getViper.BindPFlag("metadata.resource-group", getClusterCmd.PersistentFlags().Lookup("rgroupname"))


	getCmd.AddCommand(getCredentialCmd)
	getCredentialCmd.PersistentFlags().StringP("name", "n", "opsFlagDefault", "cluster name")
	getCredentialCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefault", "resource group name")
	credViper.BindPFlag("metadata.name", getCredentialCmd.PersistentFlags().Lookup("name"))
	credViper.BindPFlag("metadata.resource-group", getCredentialCmd.PersistentFlags().Lookup("rgroupname"))
}
