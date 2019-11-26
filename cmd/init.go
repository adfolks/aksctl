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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/adfolks/aksctl/pkg/ctl/utils"
)

// createCmd represents the create command
var initViper = viper.New()
var prodFilef string = "defaultprod"
var devFilef string = "defaultdev"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "To create a yaml file",
	Long:  `You can create a yaml file for production or development using aksctl init`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Do you want to create a Prod file or Dev file ? (prod/dev)")

		okayResponses := []string{"p", "P", "prod", "Prod", "PROD"}
		nokayResponses := []string{"d", "D", "dev", "Dev", "DEV"}
		message := "Please type prod or dev and then press enter:"
		confirmation := utils.AskForConfirmation(okayResponses,nokayResponses,message)
		if confirmation == true {
			_,errc := os.Create(prodFilef+".yaml")
			if errc != nil {
				color.Red("Error creating default prod yaml try creating it mannually");
				os.Exit(0)
			}
			initViper.SetConfigName(prodFilef) // name of config file (without extension)
			initViper.AddConfigPath(".")      // optionally look for config in the working directory 
			initViper.SetConfigType("yaml")
			initViper.Set("metadata.name", "defaulCluster")
			initViper.Set("metadata.resource-group", "defaultRGroup")
			initViper.Set("metadata.location", "eastus")
			errb := initViper.WriteConfig()
			if errb != nil {
				fmt.Print("Error : ",errb);
			}
			color.Green("Default prod yaml generated")
		} else {
			_,errc := os.Create(devFilef+".yaml")
			if errc != nil {
				color.Red("Error creating default dev yaml try creating it mannually");
				os.Exit(0)
			}
			initViper.SetConfigName(devFilef) // name of config file (without extension)
			initViper.AddConfigPath(".")      // optionally look for config in the working directory 
			initViper.SetConfigType("yaml")
			initViper.Set("metadata.name", "defaulCluster")
			initViper.Set("metadata.resource-group", "defaultRGroup")
			initViper.Set("metadata.location", "eastus")
			errb := initViper.WriteConfig()
			if errb != nil {
				fmt.Print("Error : ",errb);
			}
			color.Green("Default dev yaml generated")
			os.Exit(0)
		}
	
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}