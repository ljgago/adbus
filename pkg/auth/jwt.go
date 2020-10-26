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

package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	pb "github.com/ljgago/adbus/pkg/api/v1"
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Auth interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	key string
}

func New(key string) *Auth {
	return &TokenService{
		key: key,
	}
}

// Decode a token string into a token object
func (srv *TokenService) Decode(tokenString, key string) (*CustomClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return srv.key, nil
	})
	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *pb.User, key string) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "go.micro.srv.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(srv.key)
}
