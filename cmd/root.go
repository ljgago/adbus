// Copyright © 2019 Leonardo Javier Gago <ljgago@gmail.com>
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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ljgago/adbus/pkg/config"
)

// Root represents the base command when called without any subcommands
var Root = &cobra.Command{
	Use: "adbus",
	Short: `
An distributed advertising kiosk`,
	Long:              ``,
	SilenceErrors:     true,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(command *cobra.Command, args []string) {
		command.HelpFunc()(command, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.Init)

	pflags := Root.PersistentFlags()
	pflags.String("config", "", "path config file (default is $HOME)")
	pflags.String("mode", "dev", "development or production mode (dev or prod)")

	viper.BindPFlag("adbus.config", pflags.Lookup("config"))
	viper.BindPFlag("adbus.mode", pflags.Lookup("mode"))

	viper.BindEnv("adbus.config", "ADBUS_CONFIG_DIR")
	viper.BindEnv("adbus.mode", "ADBUS_WORK_MODE")
}
