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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createNPViper = viper.New()
var deleteNPViper = viper.New()
var updateNPViper = viper.New()
var scaleNPViper = viper.New()
var getNPViper = viper.New()

var createNodePoolCmd = &cobra.Command{
	Use:   "nodepool",
	Short: "Create and manage Nodepools.",
	Long: `Create and manage Nodepools, it will use a random name, a default resource group and cluster for the nodepool if not specified.
 	If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		createNPViper.SetConfigName("default") // name of config file (without extension)
		createNPViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := createNPViper.ReadInConfig()    // Find and read the config file
		if err != nil {                        // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// createViper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		//createViper.SetDefault("nodePoolName", "nodepoolDefault") // for setting a default value

		nodePoolName := createNPViper.GetString("metadata.nodepool-name") // getting values through viper
		clusterName := createNPViper.GetString("metadata.name")
		rgroupName := createNPViper.GetString("metadata.resource-group")

		fmt.Println("nodePoolName : ", nodePoolName, ", ", "clusterName : ", clusterName, ", ", "rgroupName : ", rgroupName)

		//coreaksctl.CreateNodePool(nodePoolName, clusterName, rgroupName)
	},
}

var deleteNodePoolCmd = &cobra.Command{
	Use:   "nodepool",
	Short: "Delete and manage Nodepools.",
	Long: `Delete and manage Nodepools, it will use a random name, a default resource group for the nodepool if not specified.
 	If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		deleteNPViper.SetConfigName("default") // name of config file (without extension)
		deleteNPViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := deleteNPViper.ReadInConfig()    // Find and read the config file
		if err != nil {                        // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// deleteNPViper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		//deleteNPViper.SetDefault("nodePoolName", "nodepoolDefault") // for setting a default value

		nodePoolName := deleteNPViper.GetString("metadata.nodepool-name") // getting values through viper
		clusterName := deleteNPViper.GetString("metadata.name")
		rgroupName := deleteNPViper.GetString("metadata.resource-group")

		fmt.Println("nodePoolName : ", nodePoolName, ", ", "clusterName : ", clusterName, ", ", "rgroupName : ", rgroupName)

		//coreaksctl.DeleteNodePool(nodePoolName, clusterName, rgroupName)
	},
}

var updateNodePoolCmd = &cobra.Command{
	Use:   "nodepool",
	Short: "Update and manage Nodepools.",
	Long: `Update and manage Nodepools, it will use a random name, a default resource group for the nodepool if not specified.
 	If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		updateNPViper.SetConfigName("default") // name of config file (without extension)
		updateNPViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := updateNPViper.ReadInConfig()    // Find and read the config file
		if err != nil {                        // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// updateNPViper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		//updateNPViper.SetDefault("nodePoolName", "nodepoolDefault") // for setting a default value

		nodePoolName := updateNPViper.GetString("metadata.nodepool-name") // getting values through viper
		clusterName := updateNPViper.GetString("metadata.name")
		rgroupName := updateNPViper.GetString("metadata.resource-group")

		fmt.Println("nodePoolName : ", nodePoolName, ", ", "clusterName : ", clusterName, ", ", "rgroupName : ", rgroupName)

		//coreaksctl.UpdateNodePool(nodePoolName, clusterName, rgroupName)
	},
}

var scaleNodePoolCmd = &cobra.Command{
	Use:   "nodepool",
	Short: "Scale and manage Nodepools.",
	Long: `Scale and manage Nodepools, it will use a random name, a default resource group and cluster for the nodepool if not specified.
 	If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		getNPViper.SetConfigName("default") // name of config file (without extension)
		getNPViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := getNPViper.ReadInConfig()    // Find and read the config file
		if err != nil {                     // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// viper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		//viper.SetDefault("nodePoolName", "nodepoolDefault") // for setting a default value

		nodePoolName := getNPViper.GetString("metadata.nodepool-name") // getting values through viper
		clusterName := getNPViper.GetString("metadata.name")
		rgroupName := getNPViper.GetString("metadata.resource-group")

		fmt.Println("nodePoolName : ", nodePoolName, ", ", "clusterName : ", clusterName, ", ", "rgroupName : ", rgroupName)

		//coreaksctl.ScaleNodePool(nodePoolName, clusterName, rgroupName)
	},
}

var getNodePoolCmd = &cobra.Command{
	Use:   "nodepool",
	Short: "Get and manage Nodepools.",
	Long: `Get and manage Nodepools, it will use a random name, a default resource group and cluster for the nodepool if not specified.
 	If you need to specify name or other resources use yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		// Setting config file with viper

		getNPViper.SetConfigName("default") // name of config file (without extension)
		getNPViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := getNPViper.ReadInConfig()    // Find and read the config file
		if err != nil {                     // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// getNPViper.Set("rgroupName", "opsOverrided")   //setting overide for any value

		//getNPViper.SetDefault("nodePoolName", "nodepoolDefault") // for setting a default value

		// getting values through viper
		clusterName := getNPViper.GetString("metadata.name")
		rgroupName := getNPViper.GetString("metadata.resource-group")

		fmt.Println("clusterName : ", clusterName, ", ", "rgroupName : ", rgroupName)

		//coreaksctl.GetNodePool(clusterName, rgroupName)
	},
}

func init() {

	//for create

	createCmd.AddCommand(createNodePoolCmd)
	createNodePoolCmd.PersistentFlags().StringP("nodepoolname", "n", "nodepoolFlag", "NodePool name")
	createNodePoolCmd.PersistentFlags().StringP("rgroupname", "g", "rgFlag", "NodePool resource group")
	createNodePoolCmd.PersistentFlags().StringP("clustername", "c", "clusterFlag", "NodePool cluster")

	createNPViper.BindPFlag("metadata.name", createNodePoolCmd.PersistentFlags().Lookup("clustername"))
	createNPViper.BindPFlag("metadata.resource-group", createNodePoolCmd.PersistentFlags().Lookup("rgroupname"))
	createNPViper.BindPFlag("metadata.nodepool-name", createNodePoolCmd.PersistentFlags().Lookup("nodepoolname"))

	//viper.BindPFlags(createCmd.PersistentFlags())

	//for delete

	deleteCmd.AddCommand(deleteNodePoolCmd)
	deleteNodePoolCmd.PersistentFlags().StringP("nodepoolname", "n", "nodepoolFlag", "NodePool name")
	deleteNodePoolCmd.PersistentFlags().StringP("rgroupname", "g", "rgFlag", "NodePool resource group")
	deleteNodePoolCmd.PersistentFlags().StringP("clustername", "c", "clusterFlag", "NodePool cluster")

	deleteNPViper.BindPFlag("metadata.name", deleteNodePoolCmd.PersistentFlags().Lookup("clustername"))
	deleteNPViper.BindPFlag("metadata.resource-group", deleteNodePoolCmd.PersistentFlags().Lookup("rgroupname"))
	deleteNPViper.BindPFlag("metadata.nodepool-name", deleteNodePoolCmd.PersistentFlags().Lookup("nodepoolname"))

	//for update

	updateCmd.AddCommand(updateNodePoolCmd)
	updateNodePoolCmd.PersistentFlags().StringP("nodepoolname", "n", "nodepoolFlag", "NodePool name")
	updateNodePoolCmd.PersistentFlags().StringP("rgroupname", "g", "rgFlag", "NodePool resource group")
	updateNodePoolCmd.PersistentFlags().StringP("clustername", "c", "clusterFlag", "NodePool cluster")

	updateNPViper.BindPFlag("metadata.name", updateNodePoolCmd.PersistentFlags().Lookup("clustername"))
	updateNPViper.BindPFlag("metadata.resource-group", updateNodePoolCmd.PersistentFlags().Lookup("rgroupname"))
	updateNPViper.BindPFlag("metadata.nodepool-name", updateNodePoolCmd.PersistentFlags().Lookup("nodepoolname"))

	//for scale

	scaleCmd.AddCommand(scaleNodePoolCmd)
	scaleNodePoolCmd.PersistentFlags().StringP("nodepoolname", "n", "nodepoolFlag", "NodePool name")
	scaleNodePoolCmd.PersistentFlags().StringP("rgroupname", "g", "rgFlag", "NodePool resource group")
	scaleNodePoolCmd.PersistentFlags().StringP("clustername", "c", "clusterFlag", "NodePool cluster")

	scaleNPViper.BindPFlag("metadata.name", scaleNodePoolCmd.PersistentFlags().Lookup("clustername"))
	scaleNPViper.BindPFlag("metadata.resource-group", scaleNodePoolCmd.PersistentFlags().Lookup("rgroupname"))
	scaleNPViper.BindPFlag("metadata.nodepool-name", scaleNodePoolCmd.PersistentFlags().Lookup("nodepoolname"))

	//for get List

	getCmd.AddCommand(getNodePoolCmd)
	getNodePoolCmd.PersistentFlags().StringP("rgroupname", "g", "rgFlag", "NodePool resource group")
	getNodePoolCmd.PersistentFlags().StringP("clustername", "c", "clusterFlag", "NodePool cluster")

	getNPViper.BindPFlag("metadata.name", getNodePoolCmd.PersistentFlags().Lookup("clustername"))
	getNPViper.BindPFlag("metadata.resource-group", getNodePoolCmd.PersistentFlags().Lookup("rgroupname"))

}
