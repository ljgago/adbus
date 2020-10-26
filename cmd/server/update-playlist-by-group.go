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

// UpdatePlaylistByGroup update all devices of one group
func (srv *deviceService) UpdatePlaylistByGroup(ctx context.Context, req *pb.UpdatePlaylistByGroupRequest) (*pb.UpdatePlaylistByGroupResponse, error) {
	var (
		wg   sync.WaitGroup
		m    sync.Mutex
		resp pb.UpdatePlaylistByGroupResponse
	)

	sub, err := srv.conn.Subscribe("server.put.v1.devices.groups.playlist", func(msg *nats.Msg) {
		wg.Add(1)
		go updatePlaylistByGroup(&resp, msg, &wg, &m)
	})

	if err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
		return &pb.UpdatePlaylistByGroupResponse{}, err
	}
	// send the message in protobuf
	data, err := req.Marshal()
	if err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
		return &pb.UpdatePlaylistByGroupResponse{}, err
	}
	if err = srv.conn.Publish("put.v1.devices.groups."+req.GroupId+".playlist", data); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
	}
	if waitTimeout(&wg, 5000*time.Millisecond) {
		log.Error().Str("type", "server").Msgf("Error timeout, the device %s can be disconected.", req.GroupId)
	}
	if err = sub.Unsubscribe(); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
		return &pb.UpdatePlaylistByGroupResponse{}, err
	}
	return &resp, nil
}

func updatePlaylistByGroup(resp *pb.UpdatePlaylistByGroupResponse, msg *nats.Msg, wg *sync.WaitGroup, m *sync.Mutex) {
	data := &pb.Device{}
	if err := data.Unmarshal(msg.Data); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return
	}
	wg.Done()
}
