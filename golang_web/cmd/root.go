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
	"os"

	"github.com/spf13/cobra"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "startserver",
	Short: "Use this command to manage a web golang server",
	Long: `
*** Use this command to manage a web golang server ***

It will eventually be build and used within a container. It's fairly pointless doing it this way as
I could have just used container environment variables, however I am trying to learn cobra`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//  fmt.Println(err)
		os.Exit(1)
	}
}
