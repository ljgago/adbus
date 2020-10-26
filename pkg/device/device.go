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

package device // import "github.com/ljgago/adbus/pkg/device"

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/nats-io/nats.go"

	pb "github.com/ljgago/adbus/pkg/api/v1"
	"github.com/ljgago/adbus/pkg/config"
	"github.com/ljgago/adbus/pkg/log"
)

// Device the device struct
type Device struct {
	sync.Mutex
	conn *nats.Conn

	Config       config.Device
	PlaylistData pb.Playlist

	// list of channes
	playlistCh chan pb.Playlist
	testCh     chan bool

	cmd      *exec.Cmd
	stopPlay context.CancelFunc
}

func NewDevice(cfg *config.Device) *Device {
	return &Device{
		Config: *cfg,
	}
}

func (d *Device) Run() {
	// Init the player
	d.playlistCh = make(chan pb.Playlist)
	d.testCh = make(chan bool)
	go d.Player()

	mediaFolderCheck(d.Config.Path)
	playlist, err := loadPlaylist(d.Config.Path)
	if err == nil {
		d.playlistCh <- *playlist
	}

	reconnectDelay := 10 * time.Second

	// Connect Options.
	natsOpts := []nats.Option{nats.Name("DEVICE: " + d.Config.DeviceID)}
	// User / Password
	natsOpts = append(natsOpts, nats.UserInfo(d.Config.Broker.AccessKey, d.Config.Broker.SecretKey))
	// Delay for reconnect
	natsOpts = append(natsOpts, nats.ReconnectWait(reconnectDelay))
	// Max reconnections forever
	natsOpts = append(natsOpts, nats.MaxReconnects(-1))
	natsOpts = append(natsOpts, nats.DisconnectHandler(func(bc *nats.Conn) {
		log.Info().Str("type", "device").Msg("Disconnected: trying to reconnect ...")
	}))
	natsOpts = append(natsOpts, nats.ReconnectHandler(func(bc *nats.Conn) {
		log.Info().Str("type", "device").Msgf("Reconnected [%s]", bc.ConnectedUrl())
	}))
	natsOpts = append(natsOpts, nats.ClosedHandler(func(bc *nats.Conn) {
		log.Fatal().Str("type", "device").Msg("Exiting, no servers available")
	}))

	// Connect to Broker
	showlog := true
	for {
		bc, err := nats.Connect(d.Config.Broker.URL, natsOpts...)
		if err != nil {
			if showlog {
				log.Error().Str("type", "device").Err(err).Msg("Desconected: trying to connect ...")
				showlog = false
			}
			time.Sleep(10 * time.Second)
		} else {
			log.Info().Str("type", "device").Msgf("Device %s conected to %s", d.Config.DeviceID, d.Config.Broker.URL)
			d.conn = bc
			break
		}
	}

	if _, err := d.Subscribe("get.v1.devices", d.GetDevice); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}
	if _, err := d.Subscribe("get.v1.devices."+d.Config.DeviceID, d.GetDevice); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}

	if _, err := d.Subscribe("put.v1.device", d.UpdateDevice); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}
	if _, err := d.Subscribe("put.v1.device."+d.Config.DeviceID, d.UpdateDevice); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}

	if _, err := d.Subscribe("api.v1.devices.playlist", d.GetPlaylist); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}
	if _, err := d.Subscribe("api.v1.devices."+d.Config.DeviceID+".playlist", d.GetPlaylist); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}

	if _, err := d.Subscribe("api.v1.devices.playlist.update", d.UpdatePlaylist); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}
	if _, err := d.Subscribe("api.v1.devices."+d.Config.DeviceID+".playlist.update", d.UpdatePlaylist); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}

	if _, err := d.Subscribe("adbus.device.status", d.Status); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}
	if _, err := d.Subscribe("adbus.device.status."+d.Config.DeviceID, d.Status); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}

	if _, err := d.Subscribe("adbus.device.test", d.Test); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}
	if _, err := d.Subscribe("adbus.device.test."+d.Config.DeviceID, d.Test); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
	}

	/*
		topics := make(map[string]func(*broker.Msg))

		// List subscribe
		topics["adbus.device.list"] = d.List
		topics["adbus.device.list."+d.opts.DeviceUUID] = d.List

		// Playlist subscribe
		topics["adbus.device.playlist"] = d.Playlist
		topics["adbus.device.playlist"+d.opts.DeviceUUID] = d.Playlist

		// Sync subscribe
		topics["adbus.device.sync"] = d.Sync
		topics["adbus.device.sync"+d.opts.DeviceUUID] = d.Sync

		// Status subscribe
		topics["adbus.device.status"] = d.Status
		topics["adbus.device.status"+d.opts.DeviceUUID] = d.Status

		// Test subscribe
		topics["adbus.device.test"] = d.Test
		topics["adbus.device.test"+d.opts.DeviceUUID] = d.Test

		for topic, run := range topics {
			if _, err := d.Subscribe(topic, run); err != nil {
				log.Error().Str("type", "device").Err(err).Msg("")
			}
		}*/

	<-make(chan struct{})
}

// Subscribe a topic
func (d *Device) Subscribe(subj string, f func(*nats.Msg)) (*nats.Subscription, error) {
	return d.conn.Subscribe(subj, func(msg *nats.Msg) {
		go f(msg)
	})
}

// Publish a message
func (d *Device) Publish(subj string, data []byte) error {
	return d.conn.Publish(subj, data)
}

// Status get the status info
func (d *Device) Status(msg *nats.Msg) {

}

func mediaFolderCheck(path string) {
	path = filepath.Join(path, ".adbus/media")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path does not exist, create it
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Error().Str("type", "device").Err(err).Msg("")
		}
	}
}
