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

	"github.com/RithvickAR/aksctl/coreaksctl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rgroupName, rgroupRegion string

type configurations struct {
	Name   string "name"
	Region string "region"
}

var configuration configurations

var createresourcegroupCmd = &cobra.Command{
	Use:   "resourcegroup",
	Short: "Create and manage an AKS resource group",
	Long: `Create and manage an AKS resource group, it would use a random name and default location for resource group.
	If you need to specify name or other resources use resourcegroup.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		//	rgroupName, _ = cmd.Flags().GetString("name")
		//	rgroupRegion, _ = cmd.Flags().GetString("region")
		err := viper.Unmarshal(&configuration)
		if err != nil {
			fmt.Printf("Unable to decode into struct, %v", err)
		}
		rgroupName, _ = viper.GetString("name")
		rgroupRegion, _ = configuration.region
		coreaksctl.CreateResourceGroup(rgroupName, rgroupRegion)
		fmt.Println("rgn+", rgroupName)
		fmt.Println("rgr=", rgroupRegion)
		fmt.Println("Resource group created")
	},
}

func init() {
	createCmd.AddCommand(createresourcegroupCmd)
	createresourcegroupCmd.PersistentFlags().StringP("name", "n", viper.GetString("name"), "resource group name")
	createresourcegroupCmd.PersistentFlags().StringP("region", "r", viper.GetString("region"), "resource group location")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
