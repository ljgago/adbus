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

package server // import "github.com/ljgago/adbus/cmd/server"

import (
	"context"
	"sync"
	"time"

	pb "github.com/ljgago/adbus/pkg/api/v1"
	"github.com/ljgago/adbus/pkg/log"
	"github.com/nats-io/nats.go"
)

// GetPlaylist get the playlist from device
func (srv *deviceService) GetPlaylist(ctx context.Context, req *pb.GetPlaylistRequest) (*pb.GetPlaylistResponse, error) {
	var (
		wg   sync.WaitGroup
		resp pb.GetPlaylistResponse
	)
	wg.Add(1)
	sub, err := srv.conn.Subscribe("server.get.v1.devices.playlist", func(msg *nats.Msg) {
		go getPlaylist(&resp, msg, &wg)
	})
	if err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
		return &pb.GetPlaylistResponse{}, err
	}
	if err = srv.conn.Publish("get.v1.devices."+req.Id+".playlist", []byte("")); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
	}
	if waitTimeout(&wg, 5000*time.Millisecond) {
		log.Error().Str("type", "server").Msgf("Error timeout, the device %s can be disconected.", req.Id)
	}
	if err = sub.Unsubscribe(); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
		return &pb.GetPlaylistResponse{}, err
	}
	return &resp, nil
}

func getPlaylist(resp *pb.GetPlaylistResponse, msg *nats.Msg, wg *sync.WaitGroup) {
	if err := resp.Unmarshal(msg.Data); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		wg.Done()
		return
	}
	wg.Done()
}
