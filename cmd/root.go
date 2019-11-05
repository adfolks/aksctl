/*
Copyright © 2019 Anoop Lekshmanan anoopl@adfolks.com

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

	"github.com/spf13/cobra"
	
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// type configurations struct {
// 	Name   string "name"
// 	Region string "region"
// }

// var configuration configurations

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aksctl",
	Short: "CLI to easily create cluster on Azure AKS",
	Long: `aksctl can create a cluster with single command aksctl create cluster

  It's written in Go and uses azure cli for the Azure resource creation`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aksctl.yaml)")
//	rootCmd.PersistentFlags().StringVar(&configuration, "config", "", "config file (default is $HOME/.aksctl/config/resourcegroup.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigName("resourcegroup")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetConfigType("yaml")
	var configuration c.Configuration
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		
	}
	// if configuration != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := homedir.Dir()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

		// Search config in home directory with name "resourcegroup" (without extension).
		// viper.AddConfigPath(/home/rithvick/Desktop/aksctl-test/aksctl/resourcegroup.yaml)
		// viper.AddConfigPath("./configs")
		// viper.AddConfigPath("$HOME/configs")
		// viper.AddConfigPath("/home/rithvick/Desktop/aksctl-test/aksctl/")
		// viper.SetConfigName("resourcegroup")
	}

	

	// If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
		
	// }

	// err := viper.Unmarshal(&configuration)
	// if err != nil {
	// 	fmt.Println("Unable to decode into struct, %v",err)
	// }

	fmt.Println(viper.GetString("name"))
	fmt.Println("Name\t-", viper.GetString("name"))
	fmt.Println("Region\t\t-", viper.GetString("region"))
	
}
