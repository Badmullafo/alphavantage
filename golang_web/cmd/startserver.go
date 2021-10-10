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
	"context"
	_ "errors"
	"fmt"
	"time"

	"github.com/Badmullafo/alphavantage/golang_web/pkg/request"
	"github.com/Badmullafo/alphavantage/golang_web/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const rs = 10

// byeCmd represents the bye command
var srvCmd = &cobra.Command{
	Use:   "startserver",
	Short: "This command starts the server",
	Args:  cobra.ExactValidArgs(1),
	Long:  `This command starts the server to serve an api`,
	RunE: func(cmd *cobra.Command, args []string) error {

		//If its not defined at command line look it up from viper
		apikey, _ := cmd.Flags().GetString("apikey")
		if apikey == "" {
			apikey = viper.GetViper().GetString("apikey")
		}

		stock, _ := cmd.Flags().GetString("stock")
		if stock == "" {
			stock = viper.GetViper().GetString("stock")
		}

		ndays, _ := cmd.Flags().GetInt("ndays")
		if ndays == 0 {
			ndays = viper.GetViper().GetInt("ndays")
		}

		for {
			r, err := request.GetJson(apikey, stock, ndays)

			if err != nil {
				return err
			}

			switch action := args[0]; action {
			case "total":
				r.Getot()
			case "average":
				r.Getavg()
			default:
				return fmt.Errorf("you must choose total")
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*rs))
			server.Startserver(ctx, r)
			cancel()
		}
		//return nil
	},
}

func init() {

	//cobra.OnInitialize(initConfig)

	srvCmd.Flags().StringP("apikey", "k", "", "Pass in your api apikey")
	srvCmd.Flags().StringP("stock", "s", "", "Pass in the stock you want")
	srvCmd.Flags().IntP("ndays", "n", 0, "The number of days you want")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
