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

package server // import "github.com/ljgago/adbus/pkg/gateway"

import (
	"github.com/labstack/echo"
)

type API struct {
	//opts Options
	echo *echo.Echo
}

func (api *API) listDevices(c echo.Context) error {
	return nil
}

func New(opts *Options) *API {
	api := &API{
		opts: *opts,
	}

	// Add the endpoints
	e := echo.New()
	e.GET("/info", api.Info)

	e.GET("/api/devices", api.listDevices)
	e.GET("/api/devices/:id", api.getDevice)
	e.GET("/api/status", api.getStatus)
	e.GET("/api/status/:id", api.getStatus)
	e.GET("/api/listPlaylist", api.listPlaylist)
	e.GET("/api/listPlaylist/:id", api.getPlaylist)
	e.PUT("/api/playlists", api.updatePlaylist)
	e.PUT("/api/playlists/:id", api.updatePlaylist)
	e.POST("/api/action/sync", api.actionSync)
	e.POST("/api/action/sync/:id", api.actionSync)
	e.POST("/api/action/test", api.actionTestPatern)
	e.POST("/api/action/test/:id", api.actionTestPatern)

	api.echo = e

	return api
}
