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
	"time"

	broker "github.com/nats-io/go-nats"

	"github.com/ljgago/adbus/pkg/log"
	"github.com/ljgago/adbus/pkg/pb"
)

// List get the list of devices
func (d *Device) List(msg *broker.Msg) {
	message := &pb.Message{}
	if err := message.Unmarshal(msg.Data); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return
	}

	new := &pb.Message{
		ID:        d.Opts.DeviceUUID,
		Subject:   "adbus.server.list",
		Reply:     "adbus.device.list." + d.Opts.DeviceUUID,
		Sequence:  message.Sequence,
		Timestamp: time.Now().UnixNano(),
	}

	newToSend, err := new.Marshal()
	if err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return
	}
	if err := d.Publish(message.Reply, newToSend); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return
	}
}
