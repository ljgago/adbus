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
	"os"
	"testing"
	"time"

	brokerd "github.com/nats-io/gnatsd/server"
	broker "github.com/nats-io/go-nats"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {

	// Run a broker server for test
	opts := &brokerd.Options{}
	opts.Port = 4333
	opts.Authorization = "123456789"

	// Create the broker server with appropriate options.
	s := brokerd.New(opts)

	// Start the server.
	go s.Start()
	if !s.ReadyForConnections(10 * time.Second) {
		os.Exit(1)
	}

	opts1 := []broker.Option{broker.Name("DEVICE: A")}
	opts1 = append(opts1, broker.Token("123456789"))

	conn1, _ := broker.Connect("127.0.0.1:4333", opts1...)

	conn1.Subscribe("adbus.server.list", func(msg *broker.Msg) {
	})
	conn1.Subscribe("adbus.server.sync", func(msg *broker.Msg) {
	})
	conn1.Subscribe("adbus.server.status", func(msg *broker.Msg) {
	})

	opts2 := []broker.Option{broker.Name("DEVICE: Z")}
	opts2 = append(opts1, broker.Token("123456789"))

	conn2, _ := broker.Connect("127.0.0.1:4333", opts2...)

	conn2.Subscribe("adbus.server.list", func(msg *broker.Msg) {
	})
	conn2.Subscribe("adbus.server.sync", func(msg *broker.Msg) {
	})
	conn2.Subscribe("adbus.server.status", func(msg *broker.Msg) {
	})

	assert.Equal(t, 123, 123, "they should be equal")
}
