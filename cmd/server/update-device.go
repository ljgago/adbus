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

// UpdateDevice update the device
func (srv *deviceService) UpdateDevice(ctx context.Context, req *pb.UpdateDeviceRequest) (*pb.UpdateDeviceResponse, error) {
	var (
		wg   sync.WaitGroup
		resp pb.UpdateDeviceResponse
	)

	wg.Add(1)
	sub, err := srv.conn.Subscribe("server.put.v1.devices", func(msg *nats.Msg) {
		go updateDevice(&resp, msg, &wg)
	})

	// send the message in protobuf
	data, err := req.Marshal()
	if err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
		return &pb.UpdateDeviceResponse{}, err
	}

	if err = srv.conn.Publish("put.v1.devices."+req.Id, data); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
	}

	if waitTimeout(&wg, 5000*time.Millisecond) {
		log.Error().Str("type", "server").Msgf("Error timeout, the device %s can be disconected.", req.Id)
	}
	if err = sub.Unsubscribe(); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
		return &pb.UpdateDeviceResponse{}, err
	}
	return &resp, nil
}

func updateDevice(resp *pb.UpdateDeviceResponse, msg *nats.Msg, wg *sync.WaitGroup) {
	if err := resp.Unmarshal(msg.Data); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return
	}
	wg.Done()
}
