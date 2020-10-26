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

	pb "github.com/ljgago/adbus/pkg/api/v1"
)

// GetPlaylistByGroup get the playlist of a group
func (srv *deviceService) GetPlaylistByGroup(ctx context.Context, req *pb.GetPlaylistByGroupRequest) (*pb.GetPlaylistByGroupResponse, error) {
	var (
		//wg   sync.WaitGroup
		//m    sync.Mutex
		resp pb.GetPlaylistByGroupResponse
	)

	// TODO get the playlist form database

	return &resp, nil
}
