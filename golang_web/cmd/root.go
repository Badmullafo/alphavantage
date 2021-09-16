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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//var pRootFlag bool = true

// rootCmd represents the base command when called without any subcommands
var (
	configFile string
	rootCmd    = &cobra.Command{
		Use:   "server",
		Short: "Use this command to manage a web golang server",
		Long: `
*** Use this command to manage a web golang server ***

It will eventually be build and used within a container. It's fairly pointless doing it this way as
I could have just used container environment variables, however I am trying to learn cobra`,
	}
)

func Execute() {
	//initConfig()
	if err := rootCmd.Execute(); err != nil {
		//  fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "api-examples.yml", "config file (default is api-examples.yml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	//viper.BindPFlag("apikey", srvCmd.Flags().Lookup("apikey"))

	rootCmd.AddCommand(srvCmd)
	//rootCmd.AddCommand(initCmd)
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file:%s all keys%v\n", viper.ConfigFileUsed(), viper.AllKeys())
	} else {
		fmt.Println("The config file was not used", err)
	}

}
