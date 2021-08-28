/*
Copyright © 2019 Adron Hall <adron@thrashingcode.com>
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

	"github.com/Badmullafo/alphavantage/golang_web/pkg/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string
var pRootFlag bool = true

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Use this command to manage a web golang server",
	Long: `
*** Use this command to manage a web golang server ***

It will eventually be build and used within a container. It's fairly pointless doing it this way as
I could have just used container environment variables, however I am trying to learn cobra`,
}

func Execute() {
	initConfig()
	if err := rootCmd.Execute(); err != nil {
		//  fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	rootCmd.PersistentFlags().BoolVarP(&pRootFlag, "root", "p", false, "The root flag")
	rootCmd.AddCommand(srvCmd)
	configFile = "api-examples.yml"
	viper.AddConfigPath("..")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("VIPER")

	//fmt.Println("Hello")

	for e := range viper.AllSettings() {
		fmt.Println("settings", e)
	}

	helper.HandleError(viper.BindEnv("stock"))
	helper.HandleError(viper.BindEnv("apikey"))
	helper.HandleError(viper.BindEnv("ndays"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found")
		} else {
			// Config file was found but another error was produced
			fmt.Println("Config file not found , error:", err)
		}
	}
}
