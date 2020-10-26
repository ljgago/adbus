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

package gateway // import "github.com/ljgago/adbus/pkg/gateway

import (
	"time"

	broker "github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"
)

func (s *Server) ConfigClient() {
	// Message client config
	reconnectDelay := 10 * time.Second
	// Connect Options.
	s.BrokerOpts = []broker.Option{broker.Name("SERVER")}
	// User / Password
	s.BrokerOpts = append(s.BrokerOpts, broker.UserInfo(s.Opts.AccessKey, s.Opts.SecretKey))
	// Token security
	//s.BrokerOpts = append(s.BrokerOpts, broker.Token(opts.Token))
	// Delay for reconnect
	s.BrokerOpts = append(s.BrokerOpts, broker.ReconnectWait(reconnectDelay))
	// Max reconnections forever
	s.BrokerOpts = append(s.BrokerOpts, broker.MaxReconnects(-1))
	s.BrokerOpts = append(s.BrokerOpts, broker.DisconnectHandler(func(nc *broker.Conn) {
		log.Info().Str("type", "s").Msg("Disconnected: trying to reconnect ...")
	}))
	s.BrokerOpts = append(s.BrokerOpts, broker.ReconnectHandler(func(nc *broker.Conn) {
		log.Info().Str("type", "s").Msgf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	s.BrokerOpts = append(s.BrokerOpts, broker.ClosedHandler(func(nc *broker.Conn) {
		log.Fatal().Str("type", "s").Msg("Exiting, no ss available")
	}))
}
