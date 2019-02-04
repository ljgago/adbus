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
	Port    string
	NatsURL string
}

var serverOptions ServerOptions

// serverCmd represents the subcommand
var serverCmd = &cobra.Command{
	Use:   "server [flags]",
	Short: "Start the server",
	Long: `
Start in server mode.

Te server has a web app for administration 
and use a external pub-sub server (nats-streaming) 
for comunicate with the remote devices.`,
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
	flags.StringP("port", "p", "5400", "listen port (format: [port])")
	flags.String("nats-url", "0.0.0.0:4222", "nats streaming server address (format: <ip>[:port])")

	viper.BindPFlag("adbus.server.port", flags.Lookup("port"))
	viper.BindPFlag("adbus.server.nats-url", flags.Lookup("nats-url"))
}

func configServer(cmd *cobra.Command) error {
	if serverOptions.Port = viper.GetString("adbus.server.port"); serverOptions.Port == "" {
		return errors.New("'web-port' the web server need a port")
	}
	if serverOptions.NatsURL = viper.GetString("adbus.server.nats-url"); serverOptions.NatsURL == "" {
		return errors.New("'nats-url' the nats server address")
	}
	return nil
}

func runServer(opts ServerOptions, gopts GlobalOptions, args []string) error {
	log.Debug().Caller().Str("port", opts.Port).Msg("Info")

	return nil
}
