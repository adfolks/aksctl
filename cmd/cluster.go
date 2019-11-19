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
<<<<<<< HEAD
=======

>>>>>>> 8093a6af86a4c465096a34a0332c16dd85b49caf
	"github.com/adfolks/aksctl/coreaksctl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clusterCmd represents the cluster command

var createViper = viper.New()
var deleteViper = viper.New()
var updateViper = viper.New()
var getViper = viper.New()

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create an AKS cluster",
	Long: `Create an AKS cluster, it would use a Random Name for cluster.
	If you need to specify name or other resources use cluster.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
<<<<<<< HEAD
		createCluster("Ops_Brew", "Ops_Brew")
		fmt.Println("cluster called")
=======
		fmt.Println(args)

		// Setting config file with viper

		createViper.SetConfigName("default") // name of config file (without extension)
		createViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := createViper.ReadInConfig()    // Find and read the config file
		if err != nil {                      // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		/*
			viper default value will be prior than Flag default
			so value selection priority oerder is
				- Flag Value
				- Config File
				- Vipro Default
				- Flag Default
		*/
		// viper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		clusterName := createViper.GetString("metadata.name")
		rgroupName := createViper.GetString("metadata.resource-group") // getting values through viper
		rgroupRegion := createViper.GetString("metadata.location")

		fmt.Println("rgroupName : ", rgroupName, ", ", "rgroupRegion : ", rgroupRegion, ", ", "clusterName : ", clusterName)

		coreaksctl.CreateResourceGroup(rgroupName, rgroupRegion)
		coreaksctl.CreateCluster(clusterName, rgroupName)
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

		fmt.Println("rgroupName : ", rgroupName, ", ", "rgroupRegion : ", "clusterName : ", clusterName)

		coreaksctl.DeleteCluster(clusterName, rgroupName)
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

		fmt.Println("rgroupName : ", rgroupName, ", ", "rgroupRegion : ", "clusterName : ", clusterName)

		coreaksctl.UpdateCluster(clusterName, rgroupName)
	},
}

// getClusterCmd represents the get list of cluster innresource group
var getClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Get list of AKS cluster",
	Long:  `Get list of AKS cluster from a resource group.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper
		// getViper.SetDefault("rgroupName", "opsDefault") // for setting a default value

		getViper.SetConfigName("default") // name of config file (without extension)
		getViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := getViper.ReadInConfig()    // Find and read the config file
		if err != nil {                   // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		rgroupName := viper.GetString("metadata.resource-group") // getting values through viper

		fmt.Println("rgroupName : ", rgroupName)

		coreaksctl.GetCluster(rgroupName)
>>>>>>> 8093a6af86a4c465096a34a0332c16dd85b49caf
	},
}

func init() {

	//for create

	createCmd.AddCommand(clusterCmd)
	clusterCmd.PersistentFlags().StringP("name", "n", "opsFlagDefaultCreate", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	clusterCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefaultC", "disk resource group")
	clusterCmd.PersistentFlags().StringP("rgroupRegion", "r", "westus", "disk location")

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
}
