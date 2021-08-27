// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	_ "errors"
	"fmt"

	"github.com/Badmullafo/alphavantage/golang_web/pkg/helper"
	"github.com/Badmullafo/alphavantage/golang_web/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// byeCmd represents the bye command
var startServer = &cobra.Command{
	Use:   "startserver",
	Short: "This command starts the server",
	Long:  `This command starts the server to serve an api`,
	RunE: func(cmd *cobra.Command, args []string) error {

		apiKey, _ := cmd.Flags().GetString("apiKey")
		stock, _ := cmd.Flags().GetString("stock")
		nDays, _ := cmd.Flags().GetInt("nDays")

		err := server.Startserver(apiKey, stock, nDays)

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startServer)
	startServer.Flags().StringP("apiKey", "k", viper.GetString("apiKey"), "Pass in your api apiKey")
	startServer.Flags().StringP("stock", "s", viper.GetString("stock"), "Pass in the stock you want")
	startServer.Flags().IntP("nDays", "n", viper.GetInt("nDays"), "The number of days you want")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	configFile = "api-examples.yml"
	viper.AddConfigPath("..")
	//viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()
	//viper.SetEnvPrefix("COBRACLISAMPLES")
	helper.HandleError(viper.BindEnv("stock"))
	helper.HandleError(viper.BindEnv("apiKey"))
	helper.HandleError(viper.BindEnv("nDays"))

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
