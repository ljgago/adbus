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
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ljgago/adbus/internal/log"
)

// GlobalOptions are global options for all the application
type GlobalOptions struct {
	// The config path folder
	Config string
}

// GlobalOpts variables
var options GlobalOptions

// Root represents the base command when called without any subcommands
var Root = &cobra.Command{
	Use: "adbus",
	Short: `
An advertising kiosk for buses with led displays`,
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
	cobra.OnInitialize(initConfig)

	pflags := Root.PersistentFlags()
	pflags.String("config", "", "path config file (default is $HOME/.adbus)")

	viper.BindPFlag("adbus.config", pflags.Lookup("config"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Check and/or create config and log path
	config := viper.GetString("adbus.config")
	if config == "" {
		// Find home directory.
		config, err := homedir.Dir()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		viper.Set("adbus.config", config)
	}

	adaptConfigDir(config)
	viper.AddConfigPath(config)
	viper.SetConfigName(".adbus/config")
	//configPath := filepath.Join(options.Config, ".adbus")

	//viper.SetEnvPrefix("adbus")
	// Environment
	// Replace "." and "-" for "_"
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	log.Init(config)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error:", err)
	}

}

func adaptConfigDir(path string) {
	subpath := filepath.Join(path, ".adbus")
	if _, err := os.Stat(subpath); os.IsNotExist(err) {
		// subPath no exists, create it
		if err := os.MkdirAll(subpath, 0755); err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		// Create config file
		configpath := filepath.Join(subpath, "config.yaml")
		file, err := os.Create(configpath)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		defer file.Close()
		fmt.Printf("Config path and file created on:\n%s\n%s\n", subpath, configpath)
		os.Exit(1)
	}
}
