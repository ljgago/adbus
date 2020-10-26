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

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type authInterceptor struct {
}

type jwtInterceptor struct {
	token string
}

// Authorization unary interceptor function to handle authorize per RPC call
func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// start := time.Now()
	// Skip authorize when Login or Logout is requested
	if info.FullMethod != "/pb.DeviceService/Login" && info.FullMethod != "/pb.DeviceService/Logout" {
		if err := authorize(ctx); err != nil {
			return nil, err
		}
	}
	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}

// authorize function authorizes the token received from Metadata
func authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["token"]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	token := authHeader[0]
	// validateToken function validates the token
	err := validateToken(token)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}
	return nil
}

func validateToken(token string) error {
	return nil
}
