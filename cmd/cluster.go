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
	"github.com/spf13/viper"
)

// clusterCmd represents the cluster command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create an AKS cluster",
	Long: `Create an AKS cluster, it would use a Random Name for cluster.
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

		/*
			viper default value will be prior than Flag default
			so value selection priority oerder is
				- Flag Value
				- Config File
				- Vipro Default
				- Flag Default
		*/
		// viper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		viper.SetDefault("rgroupName", "opsDefault") // for setting a default value

		rgroupName := viper.GetString("rgroupName") // getting values through viper
		rgroupRegion := viper.GetString("rgroupRegion")
		clusterName := viper.GetString("clusterName")

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
		viper.SetDefault("rgroupName", "opsDefault") // for setting a default value

		viper.SetConfigName("default") // name of config file (without extension)
		viper.AddConfigPath(".")       // optionally look for config in the working directory
		err := viper.ReadInConfig()    // Find and read the config file
		if err != nil {                // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		rgroupName := viper.GetString("rgroupName") // getting values through viper
		clusterName := viper.GetString("clusterName")

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

		// Setting config file with viper
		viper.SetDefault("rgroupName", "opsDefault") // for setting a default value

		viper.SetConfigName("default") // name of config file (without extension)
		viper.AddConfigPath(".")       // optionally look for config in the working directory
		err := viper.ReadInConfig()    // Find and read the config file
		if err != nil {                // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		rgroupName := viper.GetString("rgroupName") // getting values through viper
		clusterName := viper.GetString("clusterName")

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
		viper.SetDefault("rgroupName", "opsDefault") // for setting a default value

		viper.SetConfigName("default") // name of config file (without extension)
		viper.AddConfigPath(".")       // optionally look for config in the working directory
		err := viper.ReadInConfig()    // Find and read the config file
		if err != nil {                // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		rgroupName := viper.GetString("rgroupName") // getting values through viper

		fmt.Println("rgroupName : ", rgroupName)

		coreaksctl.GetCluster(rgroupName)
	},
}

func init() {

	//for create

	createCmd.AddCommand(clusterCmd)
	clusterCmd.PersistentFlags().StringP("clusterName", "n", "opsFlagDefault", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	clusterCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefault", "disk resource group")
	clusterCmd.PersistentFlags().StringP("rgroupRegion", "r", "westus", "disk location")

	viper.BindPFlags(clusterCmd.PersistentFlags())

	//for delete

	deleteCmd.AddCommand(deleteClusterCmd)
	deleteClusterCmd.PersistentFlags().StringP("clusterName", "n", "opsFlagDefault", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	deleteClusterCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefault", "disk resource group")

	viper.BindPFlags(deleteCmd.PersistentFlags())

	//for update

	updateCmd.AddCommand(updateClusterCmd)
	updateClusterCmd.PersistentFlags().StringP("clusterName", "n", "opsFlagDefault", "disk name") //  fullyQualifiedName,shorthand,defalt,description
	updateClusterCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefault", "disk resource group")
	updateClusterCmd.PersistentFlags().StringP("rgroupRegion", "r", "westus", "disk location")

	viper.BindPFlags(updateCmd.PersistentFlags())

	//for get List

	getCmd.AddCommand(getClusterCmd)
	getClusterCmd.PersistentFlags().StringP("rgroupname", "g", "opsFlagDefault", "disk resource group")

	viper.BindPFlags(getClusterCmd.PersistentFlags())
}
