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

package proxy // import "github.com/ljgago/adbus/pkg/proxy"

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	nats "github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"

	"github.com/ljgago/adbus/pkg/config"
)

type Proxy struct {
	// Config
	Config config.Proxy

	// Web HTTP API
	Routes *chi.Mux

	// Pub/Sub API
	Conn *nats.Conn
}

func New(cfg *config.Server) *Proxy {
	return &Proxy{
		Config: *cfg,
	}
}

func (p *Proxy) StartPubSub() {
	reconnectDelay := 10 * time.Second

	// Connect Options.
	opts := []nats.Option{nats.Name("SERVER")}

	// User / Password
	opts = append(opts, nats.UserInfo(p.Config.BrokerAccessKey, p.Config.BrokerSecretKey))
	// Token security
	//brokerOpts = append(brokerOpts, broker.Token(opts.Token))

	// Delay for reconnect
	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	// Max reconnections forever
	opts = append(opts, nats.MaxReconnects(-1))
	opts = append(opts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log.Info().Str("type", "proxy-server").Msg("Disconnected: trying to reconnect ...")
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Info().Str("type", "proxy-server").Msgf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatal().Str("type", "proxy-server").Msg("Exiting, no servers available")
	}))

	// Connect to Broker
	showlog := true
	for {
		bc, err := nats.Connect(g.Config.BrokerURL, opts...)
		if err != nil {
			if showlog {
				log.Error().Str("type", "server").Err(err).Msg("Desconected: trying to connect ...")
				showlog = false
			}
			time.Sleep(10 * time.Second)
		} else {
			log.Info().Str("type", "proxy-server").Msgf("Server conected to %s", p.Config.BrokerURL)
			p.Conn = bc
			break
		}
	}

	// if _, err := s.Subscribe("adbus.server.list", s.List); err != nil {
	// 	log.Error().Str("type", "proxy-server").Err(err).Msg("")
	// }
	// if _, err := s.Subscribe("adbus.server.sync", s.Sync); err != nil {
	// 	log.Error().Str("type", "proxy-server").Err(err).Msg("")
	// }
	// if _, err := s.Subscribe("adbus.server.status", s.Status); err != nil {
	// 	log.Error().Str("type", "proxy-server").Err(err).Msg("")
	// }

}

func (p *Proxy) StartAPI() {
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Recoverer,
		middleware.URLFormat,
	)
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/devices", func(r chi.Router) {
			r.Get("/", p.listDevices)   // GET  /api/v1/devices
			r.Post("/", p.createDevice) // POST /api/v1/devices
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", p.getDevice)              // GET    /api/v1/devices/123
				r.Put("/", p.updateDevice)           // PUT    /api/v1/devices/123
				r.Delete("/", p.deleteDevice)        // DELETE /api/v1/devices/123
				r.Get("/playlist", p.getPlaylist)    // GET    /api/v1/devices/123/playlist
				r.Put("/playlist", p.updatePlaylist) // PUT    /api/v1/devices/123/playlist
				r.Post("/sync", p.actionSync)        // POST   /api/v1/devices/123/sync
				r.Post("/test", p.actionTest)        // POST   /api/v1/devices/123/test
			})
			r.Route("/group/{id}", func(r chi.Router) {
				r.Get("/", p.getDevicesByGroup)           // GET  /api/v1/devices/group/1
				r.Put("/", p.updateDevicesByGroup)        // PUT  /api/v1/devices/group/1
				r.Get("/playlist", p.getPlaylistsByGroup) // GET  /api/v1/devices/group/1/playlist
				r.Put("/playlist", p.updatePlaylist)      // PUT  /api/v1/devices/group/1/playlist
				r.Post("/sync", p.actionSyncByGroup)      // POST /api/v1/devices/group/1/sync
				r.Post("/test", p.actionTestByGroup)      // POST /api/v1/devices/group/1/test
			})
			r.Get("/playlist", p.getAllPlaylists) // GET  /api/v1/devices/playlist
			r.Post("/sync", p.actionSyncAll)      // POST /api/v1/devices/sync
			r.Post("/test", p.actionTestAll)      // POST /api/v1/devices/test
		})
		r.Route("/users", func(r chi.Router) {
			r.Get("/", p.getUser)
			r.Post("/", p.createUser)
			r.Put("/", p.updateUser)
			r.Delete("/", p.deleteUser)
		})
	})
	p.Routes = r
}

func (p *Proxy) Run() {
	// Init the

	p.StartPubSub()
	p.StartAPI()

	http.ListenAndServe(p.Config.Port, p.RoutesAPI)
}
