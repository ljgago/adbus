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
	"context"
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/ljgago/adbus/pkg/log"
	"github.com/ljgago/adbus/pkg/pb"
)

// Player play the files updadted for the playlist
func (d *Device) Player() {
	for {
		select {
		case playlist := <-d.playlistCh:
			ctx, stop := context.WithCancel(context.Background())
			d.stopPlay = stop
			go d.Play(ctx, &playlist)
		case <-d.testCh:
			ctx, stop := context.WithCancel(context.Background())
			d.stopPlay = stop
			go d.PlayTest(ctx)
		}
	}
}

// Stop stop the player
func (d *Device) Stop() {
	d.stopPlay()
}

// Play video or image files
func (d *Device) Play(ctx context.Context, pl *pb.Playlist) {
	d.Lock()
	defer d.Unlock()
	for {
		for _, item := range pl.Items {
			// Check and play video or image files (case insensitive)
			regVid := regexp.MustCompile(`(?mi)(.*?)\.(avi|mov|mkv|mp4|mpeg|m4v)$`)
			regImg := regexp.MustCompile(`(?mi)(.*?)\.(bmp|jpeg|jpg|gif|png|ppm|tif|tiff|webp|xpm|xwd)$`)

			if regVid.MatchString(item.File) {
				d.videoPlay(ctx, item)
			}

			if regImg.MatchString(item.File) {
				d.imageView(ctx, item)
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}
}

// VideoPlay play videos with a value for repeat
func (d *Device) videoPlay(ctx context.Context, item *pb.Item) {
	path := filepath.Join(d.Opts.Config, ".adbus/media")
	path = filepath.Join(path, item.File)
	log.Info().Str("type", "player").Msgf("PLAYING: %s", item.File)
	d.cmd = exec.CommandContext(ctx, d.Opts.VideoPlayer, `"`+path+`"`)
	times := int(item.Time)
	for i := 0; i <= times; i++ {
		d.cmd.Run()
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

// ImageView show pictures with a value for timeout
func (d *Device) imageView(ctx context.Context, item *pb.Item) {
	duration := strconv.Itoa(int(item.Time))
	path := filepath.Join(d.Opts.Config, ".adbus/media")
	path = filepath.Join(path, item.File)
	log.Info().Str("type", "player").Msgf("PLAYING: %s", item.File)
	d.cmd = exec.CommandContext(ctx, d.Opts.ImageViewer, duration, `"`+path+`"`)
	d.cmd.Run()
	select {
	case <-ctx.Done():
		return
	default:
	}
}

// PlayTest play the test patern
func (d *Device) PlayTest(ctx context.Context) {
	path := filepath.Join(d.Opts.Config, ".adbus/bars.mp4")
	log.Info().Str("type", "player").Msg("PLAYING: test bars")
	d.cmd = exec.CommandContext(ctx, d.Opts.VideoPlayer, `"`+path+`"`)
	for {
		d.cmd.Run()
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

// LoadPlaylist open the playlist of file
func loadPlaylist(path string) (*pb.Playlist, error) {
	filename := filepath.Join(path, ".adbus/playlist.json")
	fileJSON, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return nil, err
	}
	playlist := &pb.Playlist{}
	if err := json.Unmarshal(fileJSON, &playlist.Items); err != nil {
		log.Error().Str("type", "device").Err(err).Msg("")
		return nil, err
	}
	return playlist, nil
}

/*
func (d *Device) Play(ctx context.Context, pl *pb.Playlist) {
	d.Lock()
	defer d.Unlock()
	for {
		for _, item := range pl.Items {
			d.PlayFormat(ctx, item)
			select {
			case <-ctx.Done():
				return
			default:
			}
		}

	}
}

// PlayFormat check and play video or image files
func (d *Device) PlayFormat(ctx context.Context, item *pb.Item) {
	regVid := regexp.MustCompile(`(.*?)\.(avi|mov|mkv|mp4|mpeg|m4v)$`)
	regImg := regexp.MustCompile(`(.*?)\.(bmp|jpeg|jpg|gif|png|ppm|tif|tiff|webp|xpm|xwd)$`)

	if regVid.MatchString(item.File) {
		d.VideoPlay(ctx, item)
		return
	}

	if regImg.MatchString(item.File) {
		d.ImageView(ctx, item)
		return
	}
}
*/
