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
	"github.com/Badmullafo/alphavantage/golang_web/pkg/server"
	"github.com/spf13/cobra"
	_ "github.com/spf13/viper"
)

// byeCmd represents the bye command
var startServer = &cobra.Command{
	Use:   "startserver",
	Short: "This command starts the server",
	Long:  `This command starts the server to serve an api`,
	RunE: func(cmd *cobra.Command, args []string) error {

		key, _ := cmd.Flags().GetString("key")

		server.Startserver(key, "IBM", 3)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startServer)
	startServer.Flags().StringP("key", "k", "RABZYXWVHB8MX5GO", "Pass in your api key")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
