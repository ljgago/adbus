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
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	broker "github.com/nats-io/go-nats"

	"github.com/ljgago/adbus/pkg/log"
	"github.com/ljgago/adbus/pkg/pb"
	"github.com/ljgago/adbus/pkg/sync"
)

// Update sync the media files and update the playlist
func (d *Device) Update(msg *broker.Msg) {
	// Get the new playlist
	message := &pb.Message{}
	if err := message.Unmarshal(msg.Data); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return
	}
	newPlaylist := &pb.Playlist{}
	if err := json.Unmarshal(message.Data, &newPlaylist.Items); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return
	}

	// Stop the player
	d.Stop()

	// Update the media files
	d.Lock()
	bucket := "media"
	local := filepath.Join(d.Opts.Config, ".adbus/media")
	update := sync.New(d.Opts.Endpoint, d.Opts.AccessKey, d.Opts.SecretKey)
	if err := update.S3toLocal(bucket, local); err != nil {
		log.Error().Str("type", "update").Err(err).Msg("")
	}

	// Save playlist to file
	path := filepath.Join(d.Opts.Config, ".adbus/playlist.json")
	err := ioutil.WriteFile(path, message.Data, 0644)
	if err != nil {
		log.Error().Str("type", "update").Err(err).Msg("")
	}
	d.Unlock()

	// Play the playlist
	d.playlistCh <- *newPlaylist
}
