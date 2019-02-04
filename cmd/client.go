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

// ClientOptions options running server
type ClientOptions struct {
	ServerURL   string
	VideoPlayer string
	ImageViewer string
}

var clientOptions ClientOptions

// clientCmd represents the subcommand
var clientCmd = &cobra.Command{
	Use:   "client [flags]",
	Short: "Start the client",
	Long: `
Start the client mode.

The client plays videos
and images from a playlist.`,
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
	flags.StringVar(&clientOptions.ServerURL, "server-url", "", "server url address (format: <ip>[:port])")
	flags.StringVar(&clientOptions.VideoPlayer, "video-player", "omxplayer", "external video player")
	flags.StringVar(&clientOptions.ImageViewer, "image-viewer", "fbi", "external image viewer")

	viper.BindPFlag("adbus.client.server-url", flags.Lookup("server-url"))
	viper.BindPFlag("adbus.client.video-player", flags.Lookup("video-player"))
	viper.BindPFlag("adbus.client.image-viewer", flags.Lookup("image-viewer"))
}

func configClient() error {
	if clientOptions.ServerURL = viper.GetString("adbus.client.server-url"); clientOptions.ServerURL == "" {
		return errors.New("'server-url' need an server address")
	}
	if clientOptions.VideoPlayer = viper.GetString("adbus.client.video-player"); clientOptions.VideoPlayer == "" {
		return errors.New("'video-player' need an external video player")
	}
	if clientOptions.ImageViewer = viper.GetString("adbus.client.image-player"); clientOptions.ImageViewer == "" {
		return errors.New("'server-url' need an external image viewer ")
	}
	return nil
}

func runClient(opts ClientOptions, gopts GlobalOptions, args []string) error {
	log.Debug().Caller().Msg("server-url" + opts.ServerURL)
	return nil
}
