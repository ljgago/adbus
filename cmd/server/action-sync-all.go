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

	types "github.com/gogo/protobuf/types"
	"github.com/ljgago/adbus/pkg/log"
)

// ActionSyncAll sync all devices with the s3 storage
func (srv *deviceService) ActionSyncAll(ctx context.Context, req *types.Empty) (*types.Empty, error) {
	if err := srv.conn.Publish("action.v1.devices.sync", []byte("sync")); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
		return &types.Empty{}, err
	}
	return &types.Empty{}, nil
}
