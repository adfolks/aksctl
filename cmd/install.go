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

	"github.com/adfolks/aksctl/pkg/ctl/addon"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var installViper = viper.New()

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "To install add-on",
	Long:  `addon message`,
	Run: func(cmd *cobra.Command, args []string) {
		// Setting config file with viper

		installViper.SetConfigName("default") // name of config file (without extension)
		installViper.AddConfigPath(".")       // optionally look for config in the working directory
		err := installViper.ReadInConfig()    // Find and read the config file
		if err != nil {                        // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		
		//installViper.SetDefault("chartName", "test") // for setting a default value
		chartName, _ := cmd.Flags().GetString("name")
		repoName, _ := cmd.Flags().GetString("reponame")
		//chartName := installViper.GetString("metadata.chartname") // getting values through viper
		//repoName := installViper.GetString("metadata.reponame")

		color.Cyan("chartName : " + chartName + ", repository: " + repoName)

		addon.InstallAddon(chartName, repoName)
	},
}



func init() {
	addOnCmd.AddCommand(installCmd)
	installCmd.PersistentFlags().StringP("name", "n", "test", "chart name")
	installCmd.PersistentFlags().StringP("reponame", "r", "test", "repo name")
}