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

package device // import "github.com/ljgago/adbus/cmd/device"

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ljgago/adbus/cmd"
	"github.com/ljgago/adbus/pkg/log"
)

// Options for device client
type Options struct {
	// The Broker server URL
	BrokerURL string

	// Endpoint URL
	Endpoint string

	// User access
	AccessKey string

	// Secret password
	SecretKey string

	// The command line app used for play video
	VideoPlayer string

	// The command line app used for play pictures
	ImageViewer string

	// Config path
	Config string

	// ADBUS ENVIROMENT

	// The unique identification number for the device. This is used to identify it on adbus
	DeviceUUID string

	// ID number of the adbus application the device is associated.
	AppID string

	// The name of the adbus application the device is associated with.
	AppName string

	// The name of the device on first initialisation.
	DeviceNameAtInit string

	// The type of device the application is running on.
	DeviceType string

	// The ADBUS=1 variable can be used by your software to detect that it is running on a adbus device.
	Adbus bool

	// The current version of the supervisor agent running on the device.
	SupervisorVersion string

	// Authentication key for the supervisor API.
	// This makes sure requests to the supervisor are only coming from containers on the device.
	// See the Supervisor API reference for detailed usage.
	SupervisorAPIKey string

	// The network address of the supervisor API. Default: http://127.0.0.1:48484
	SupervisorAddress string

	// The IP address of the supervisor API. Default: 127.0.0.1
	SupervisorHost string

	// The network port number for the supervisor API. Default: 48484
	SupervisorPort string

	// API key which can be used to authenticate requests to the adbus backend.
	// Can be used with the SDKs on the device.
	// WARNING This API key gives the code full user permissions,
	// so can be used to delete and update anything as you would on the Dashboard.
	APIKey string

	// The version of the host OS.
	HostOsVersion string
}

const logName string = ".adbus/adbus-device.log"

// commandDefine represents the subcommand
var commandDefine = &cobra.Command{
	Use:   "device [flags]",
	Short: "Start the device client",
	Long: `
Start in device mode.

The device plays videos
and images from a playlist.`,
	DisableAutoGenTag: true,
	Run: func(command *cobra.Command, args []string) {
		// options, err := configDevice()
		// if err != nil {
		// 	log.Fatal().Str("type", "device").Err(err).Msg("")
		// 	os.Exit(1)
		// }
		runDevice(args)
	},
}

func init() {
	cmd.Root.AddCommand(commandDefine)

	flags := commandDefine.Flags()
	flags.String("player-video", "omxplayer", "external video player")
	flags.String("player-image", "fbi", "external image viewer")
	flags.StringP("broker-url", "b", "127.0.0.1:4222", "broker server address (format: <ip>[:port])")
	flags.String("broker-access-key", "", "broker access key")
	flags.String("broker-secret-key", "", "broker secret key")
	flags.StringP("storage-url", "s", "127.0.0.1:9000", "s3 storage url")
	flags.String("storage-access-key", "admin", "storage access key")
	flags.String("storage-secret-key", "123456789", "storage secret key")

	//flags.StringP("token", "t", "123456789", "token for connect")

	viper.BindPFlag("adbus.device.player.video", flags.Lookup("player-video"))
	viper.BindPFlag("adbus.device.player.image", flags.Lookup("player-image"))
	viper.BindPFlag("adbus.device.broker.url", flags.Lookup("broker-url"))
	viper.BindPFlag("adbus.device.broker.access-key", flags.Lookup("broker-access-key"))
	viper.BindPFlag("adbus.device.broker.secret-key", flags.Lookup("broker-secret-key"))
	viper.BindPFlag("adbus.device.storage.url", flags.Lookup("storage-url"))
	viper.BindPFlag("adbus.device.storage.access-key", flags.Lookup("storage-access-key"))
	viper.BindPFlag("adbus.device.storage.secret-key", flags.Lookup("storage-secret-key"))
	//viper.BindPFlag("adbus.device.token", flags.Lookup("token"))

	viper.BindEnv("adbus.device.player.video", "ADBUS_DEVICE_PLAYER_VIDEO")
	viper.BindEnv("adbus.device.player.image", "ADBUS_DEVICE_PLAYER_IMAGE")
	viper.BindEnv("adbus.device.broker.url", "ADBUS_DEVICE_BROKER_URL")
	viper.BindEnv("adbus.device.broker.access-key", "ADBUS_DEVICE_BROKER_ACCESS_KEY")
	viper.BindEnv("adbus.device.broker.secret-key", "ADBUS_DEVICE_BROKER_SECRET_KEY")
	viper.BindEnv("adbus.device.storage.url", "ADBUS_DEVICE_STORAGE_URL")
	viper.BindEnv("adbus.device.storage.access-key", "ADBUS_DEVICE_STORAGE_ACCESS_KEY")
	viper.BindEnv("adbus.device.storage.secret-key", "ADBUS_DEVICE_STORAGE_SECRET_KEY")

	viper.BindEnv("adbus.device.id", "ADBUS_DEVICE_ID")
	viper.BindEnv("adbus.device.appid", "ADBUS_APP_ID")
	viper.BindEnv("adbus.device.appname", "ADBUS_APP_NAME")
	viper.BindEnv("adbus.device.name-at-init", "ADBUS_DEVICE_NAME_AT_INIT")
	viper.BindEnv("adbus.device.type", "ADBUS_DEVICE_TYPE")
	viper.BindEnv("adbus", "ADBUS")
	viper.BindEnv("adbus.device.supervisor-version", "ADBUS_SUPERVISOR_VERSION")
	viper.BindEnv("adbus.device.supervisor-api-key", "ADBUS_SUPERVISOR_API_KEY")
	viper.BindEnv("adbus.device.supervisor-address", "ADBUS_SUPERVISOR_ADDRESS")
	viper.BindEnv("adbus.device.supervisor-host", "ADBUS_SUPERVISOR_HOST")
	viper.BindEnv("adbus.device.supervisor-port", "ADBUS_SUPERVISOR_PORT")
	viper.BindEnv("adbus.device.api-key", "ADBUS_API_KEY")
	viper.BindEnv("adbus.device.host-os-version", "ADBUS_HOST_OS_VERSION")

}

/*
func configDevice() (*Options, error) {
	options := &Options{}
	if options.BrokerURL = viper.GetString("adbus.device.broker-url"); options.BrokerURL == "" {
		return nil, errors.New("'broker-url' the broker server address")
	}
	if options.Endpoint = viper.GetString("adbus.device.endpoint"); options.Endpoint == "" {
		return nil, errors.New("'endpoint' need an endpoint url")
	}
	if options.AccessKey = viper.GetString("adbus.device.access-key"); options.AccessKey == "" {
		return nil, errors.New("'access-key' need access key")
	}
	if options.SecretKey = viper.GetString("adbus.device.secret-key"); options.SecretKey == "" {
		return nil, errors.New("'secret-key' need secret key")
	}
	if options.VideoPlayer = viper.GetString("adbus.device.video-player"); options.VideoPlayer == "" {
		return nil, errors.New("'video-player' need an external video player")
	}
	if options.ImageViewer = viper.GetString("adbus.device.image-viewer"); options.ImageViewer == "" {
		return nil, errors.New("'image-player' need an external image viewer ")
	}

	options.Config = viper.GetString("adbus.config")

	// Balena enviroment
	options.DeviceUUID = viper.GetString("adbus-device-uuid")
	options.AppID = viper.GetString("adbus-app-id")
	options.AppName = viper.GetString("adbus-app-name")
	options.DeviceNameAtInit = viper.GetString("adbus-device-name-at-init")
	options.DeviceType = viper.GetString("adbus-device-type")
	options.Adbus = viper.GetBool("adbus")
	options.SupervisorVersion = viper.GetString("adbus-supervisor-version")
	options.SupervisorAPIKey = viper.GetString("adbus-supervisor-api-key")
	options.SupervisorAddress = viper.GetString("adbus-supervisor-address")
	options.SupervisorHost = viper.GetString("adbus-supervisor-host")
	options.SupervisorPort = viper.GetString("adbus-supervisor-port")
	options.APIKey = viper.GetString("adbus-api-key")
	options.HostOsVersion = viper.GetString("adbus-host-os-version")

	log.Init(options.Config, logName)

	return options, nil
}
*/

func runDevice(args []string) {
	path := viper.GetString("adbus.config")
	mode := viper.GetString("adbus.mode")
	log.Init(path, mode, logName)
	// Start the device client
	// device := NewDevice(opts)
	// device.Run()

	//go Init(opts)

	// Waiting
	<-make(chan struct{})
}

/*
var lastProcessed uint64
func handleMsg(msg *stan.Msg) {
	// Only process messages greater than the last one processed.
	// If it has been seen, skip to acknowledge it to the server.
	if msg.Sequence > lastProcessed {
		if err := process(msg); err != nil {
			// Log error and/or close subscription.
			return
		}

		// Processing successful, set the `lastProcessed` value.
		atomic.SwapUnint64(&lastProcessed, msg.Sequence)
	}

	// ACK with the server.
	msg.Ack()

}
*/
