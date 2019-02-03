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
	"github.com/ljgago/adbus/cmd/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServerOptions options running server
type ClientOptions struct {
	ServerAddr  string
	VideoPlayer string
	ImagePlayer string
}

var clientOptions ClientOptions

// serverCmd represents the subcommand
var clientCmd = &cobra.Command{
	Use:   "client [flags]",
	Short: "Start the client",
	Long: `
Start in server mode`,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := configClient(); err != nil {
			return err
		}
		return runClient(clientOptions, globalOptions, args)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	flags := clientCmd.Flags()
	flags.StringVar(&clientOptions.ServerAddr, "server-addr", "", "server address (format: <ip>[:port])")
	flags.StringVar(&clientOptions.VideoPlayer, "video-player", "", "external video player")
	flags.StringVar(&clientOptions.ImagePlayer, "image-player", ":", ")")

	viper.BindPFlag("adbus.client.server-address", flags.Lookup("address"))
}

func configClient() error {

	//if clientOptions.ServerAddr = viper.GetString("adbus.client.servr-address"); clientOptions.ServerAddr == "" {
	//	log.Error().Msg("server-addr" + clientOptions.ServerAddr)
	//	return errors.New("'server-addr' need an server address")
	//}
	return nil
}

func runClient(opts ClientOptions, gopts GlobalOptions, args []string) error {
	log.Info().Msg("server-address" + opts.ServerAddr)
	return nil
}
