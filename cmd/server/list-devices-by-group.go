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

// ListDevicesByGroup list all devices of one group
func (srv *deviceService) ListDevicesByGroup(ctx context.Context, req *pb.ListDevicesByGroupRequest) (*pb.ListDevicesByGroupResponse, error) {
	var (
		wg   sync.WaitGroup
		m    sync.Mutex
		resp pb.ListDevicesByGroupResponse
	)

	sub, err := srv.conn.Subscribe("server.list.v1.devices.groups", func(msg *nats.Msg) {
		wg.Add(1)
		go listDevicesByGroup(&resp, msg, &wg, &m)
	})
	if err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
		return &pb.ListDevicesByGroupResponse{}, err
	}
	if err = srv.conn.Publish("get.v1.devices.groups."+req.GroupId, []byte("")); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
	}
	// Wait 100 ms for get all device information
	time.Sleep(100 * time.Millisecond)
	if waitTimeout(&wg, 5000*time.Millisecond) {
		log.Error().Str("type", "server").Msgf("Error timeout")
	}

	if err = sub.Unsubscribe(); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
		return &pb.ListDevicesByGroupResponse{}, err
	}

	return &resp, nil
}

func listDevicesByGroup(resp *pb.ListDevicesByGroupResponse, msg *nats.Msg, wg *sync.WaitGroup, m *sync.Mutex) {
	data := &pb.Device{}
	if err := data.Unmarshal(msg.Data); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return
	}
	m.Lock()
	resp.Devices = append(resp.Devices, data)
	m.Unlock()
	wg.Done()
}
