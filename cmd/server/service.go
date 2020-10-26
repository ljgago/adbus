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
	"sync"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/ljgago/adbus/pkg/config"
)

type deviceService struct {
	cfg  config.Server
	conn *nats.Conn
	user string
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

/*
func (srv *deviceService) Login(ctx context.Context, req *v1.LoginRequest) (*empty.Empty, error) {

	ctxs := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})
	return nil, nil
}

func (srv *deviceService) Logout(ctx context.Context, req *v1.LoginRequest) (*empty.Empty, error) {

}

*/
