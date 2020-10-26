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

package device

import (
	"encoding/json"

	"github.com/ljgago/adbus/pkg/log"
	"github.com/ljgago/adbus/pkg/pb"
	broker "github.com/nats-io/go-nats"
)

// Test enable o disable test pattern
func (d *Device) Test(msg *broker.Msg) {
	// Get test data
	message := &pb.Message{}
	if err := message.Unmarshal(msg.Data); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return
	}
	test := &pb.Test{}
	if err := json.Unmarshal(message.Data, &test); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return
	}

	// Stop the playlist
	d.Stop()

	if test.Enabled {
		d.testCh <- true
	} else {
		playlist, err := loadPlaylist(d.Opts.Config)
		if err == nil {
			d.playlistCh <- *playlist
		}
	}
}
