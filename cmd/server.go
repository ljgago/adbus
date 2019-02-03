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

package cmd // import "github.com/ljgago/adbus/cmd"

import (
	"errors"

	"github.com/ljgago/adbus/cmd/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServerOptions options running server
type ServerOptions struct {
	WebPort  string
	NatsAddr string
}

var serverOptions ServerOptions

// serverCmd represents the subcommand
var serverCmd = &cobra.Command{
	Use:   "server [flags]",
	Short: "Start the server",
	Long: `
Start in server mode`,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := configServer(cmd); err != nil {
			return err
		}
		return runServer(serverOptions, globalOptions, args)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	flags := serverCmd.Flags()
	flags.String("web-port", "5400", "listen web port (format: [port])")
	flags.String("nats-addr", ":4222", "nats streaming server address (format: <ip>[:port])")

	viper.BindPFlag("adbus.server.web-port", flags.Lookup("web-port"))
	viper.BindPFlag("adbus.server.nats-addr", flags.Lookup("nats-addr"))
}

func configServer(cmd *cobra.Command) error {
	if serverOptions.WebPort = viper.GetString("adbus.server.web-port"); serverOptions.WebPort == "" {
		return errors.New("'web-port' the web server need a port")
	}
	if serverOptions.NatsAddr = viper.GetString("adbus.server.nats-addr"); serverOptions.NatsAddr == "" {
		return errors.New("'nats-addr' the nats server address")
	}
	return nil
}

func runServer(opts ServerOptions, gopts GlobalOptions, args []string) error {
	log.Info().Str("web-port", opts.WebPort).Msg("Info")

	return nil
}
