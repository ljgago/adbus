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

package config // import "github.com/ljgago/adbus/pkg/config"

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

//
type PlayerConfig struct {
	Video string `json:"video,omitempty"`
	Image string `json:"image,omitempty"`
}

//
type ServiceConfig struct {
	// The URL from service
	URL string `json:"url,omitempty" mapstructure:"url"`

	// The access key credential
	AccessKey string `json:"access_key,omitempty" mapstructure:"access-key"`

	// The secret key credential
	SecretKey string `json:"secret_key,omitempty" mapstructure:"secret-key"`
}

// Device
type Device struct {
	Player  PlayerConfig  `json:"player,omitempty" mapstructure:"player"`
	Broker  ServiceConfig `json:"broker,omitempty" mapstructure:"broker"`
	Storage ServiceConfig `json:"storage,omitempty" mapstructure:"storage"`

	// The unique identification number for the device. This is used to identify it on adbus
	DeviceID string `json:"device_id,omitempty" mapstructure:"device-id"`

	// ID number of the adbus application the device is associated.
	AppID string `json:"app_id,omitempty" mapstructure:"app-id"`

	// The name of the adbus application the device is associated with.
	AppName string `json:"app_name,omitempty" mapstructure:"app-name"`

	// The name of the device on first initialisation.
	DeviceNameAtInit string `json:"device_name_at_init,omitempty" mapstructure:"device-name-at-init"`

	// The type of device the application is running on.
	DeviceType string `json:"device_type,omitempty" mapstructure:"device-type"`

	// The ADBUS=1 variable can be used by your software to detect that it is running on a adbus device.
	Adbus bool `json:"adbus,omitempty" mapstructure:"adbus"`

	// The current version of the supervisor agent running on the device.
	SupervisorVersion string `json:"supervisor_version,omitempty" mapstructure:"supervisor-version"`

	// Authentication key for the supervisor API.
	// This makes sure requests to the supervisor are only coming from containers on the device.
	// See the Supervisor API reference for detailed usage.
	SupervisorAPIKey string `json:"supervisor_api_key,omitempty" mapstructure:"supervisor-api-key"`

	// The network address of the supervisor API. Default: http://127.0.0.1:48484
	SupervisorAddress string `json:"supervisor_address,omitempty" mapstructure:"supervisor-address"`

	// The IP address of the supervisor API. Default: 127.0.0.1
	SupervisorHost string `json:"supervisor_host,omitempty" mapstructure:"supervisor-host"`

	// The network port number for the supervisor API. Default: 48484
	SupervisorPort string `json:"supervisor_port,omitempty" mapstructure:"supervisor-port"`

	// API key which can be used to authenticate requests to the adbus backend.
	// Can be used with the SDKs on the device.
	// WARNING This API key gives the code full user permissions,
	// so can be used to delete and update anything as you would on the Dashboard.
	APIKey string `json:"api_key,omitempty" mapstructure:"api-key"`

	// The version of the host OS.
	HostOsVersion string `json:"host_os_version,omitempty" mapstructure:"host-os-version"`

	Path string `json:"path,omitempty" mapstructure:"path"`
}

// Server
type Server struct {
	// The port used for the HTTP API
	PortHTTP string `json:"port_http" mapstructure:"port-http"`

	// The port used for the gRPC API
	PortGRPC string `json:"port_grpc" mapstructure:"port-grpc"`

	// The broker service config
	Broker ServiceConfig `json:"broker,omitempty" mapstructure:"broker"`

	// The storage service config
	Storage ServiceConfig `json:"storage,omitempty" mapstructure:"storage"`

	// The database credential access
	Database string `json:"database,omitempty" mapstructure:"database"`

	// The certs path
	Certs string `json:"certs,omitempty" mapstructure:"cert"`

	// The config path
	Path string `json:"path,omitempty" mapstructure:"path"`
}

// Init reads in config file and ENV variables if set.
func Init() {
	// Check and/or create config and log path
	config := viper.GetString("adbus.config")
	if config == "" {
		// Find home directory.
		conf, err := homedir.Dir()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		viper.Set("adbus.config", conf)
		config = conf
	}
	adaptConfigDir(config)
	viper.AddConfigPath(config)
	viper.SetConfigName(".adbus/config")
	//configPath := filepath.Join(options.Config, ".adbus")

	viper.SetEnvPrefix("adbus")
	// Environment
	// Replace "." and "-" for "_"
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

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

// GetDevice
func GetDevice() (Device, error) {
	config := &Device{}
	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		viper.Set(key, val)
	}
	if err := viper.UnmarshalKey("adbus.device", config); err != nil {
		return Device{}, err
	}
	config.Path = viper.GetString("adbus.config")
	return *config, nil
}

// GetServer
func GetServer() (Server, error) {
	config := &Server{}
	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		viper.Set(key, val)
	}
	if err := viper.UnmarshalKey("adbus.server", config); err != nil {
		return Server{}, err
	}
	config.Path = viper.GetString("adbus.config")
	return *config, nil
}
