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
	"os"
	"strconv"
	"strings"
	"time"

	brokerd "github.com/nats-io/gnatsd/server"

	"github.com/ljgago/adbus/pkg/log"
)

// BrokerServerInit initilize the broker server
func BrokerServerInit(opts *Options) {
	// Set options
	brokerOpts := &brokerd.Options{}
	url := strings.Split(opts.BrokerURL, ":")
	if len(url) == 2 {
		//opts.Host = url[0]
		port, err := strconv.Atoi(url[1])
		if err == nil {
			brokerOpts.Port = port
		}
	}
	//brokerOpts.Authorization = opts.Token

	brokerOpts.Username = opts.AccessKey
	brokerOpts.Password = opts.SecretKey

	// Create the broker server with appropriate options.
	s := brokerd.New(brokerOpts)

	// Start the server.
	go s.Start()
	if !s.ReadyForConnections(10 * time.Second) {
		log.Fatal().Str("type", "broker-server").Msgf("Unable to start Broker Server on %s", opts.BrokerURL)
		os.Exit(1)
	}
	log.Info().Str("type", "broker-server").Msgf("Listening for devices clients connections on %s", opts.BrokerURL)
}
