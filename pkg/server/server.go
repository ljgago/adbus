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
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	packr "github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo"
	broker "github.com/nats-io/go-nats"
	"gopkg.in/bluesuncorp/validator.v5"

	"github.com/ljgago/adbus/pkg/log"
)

// Options for API, broker server and UI dashboard
type Options struct {
	// API Server
	// The port used for the API server
	Port string

	// Enable o disable the UI dashboard
	UI bool

	// The port used for enter to UI dashboard
	PortUI string

	// Broker server
	// The Broker server URL
	BrokerURL string

	// Use external broker server (default: false) or internal nats server (true)
	BrokerExternal bool

	// Broker access key
	BrokerAccessKey string

	// Broker secret key
	BrokerSecretKey string

	// Storage
	// Endpoint S3 URL
	StorageURL string

	// Storage access key
	StorageAccessKey string

	// Storage secret key
	StorageSecretKey string

	// Database
	// Database URL
	DatabaseURL string

	// Database access key
	DatabaseAccessKey string

	// Database secret key
	DatabaseSecretKey string

	// Config path
	Config string
}

type jwtClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type Server struct {
	Port             string
	StorageAccessKey string
	StorageSecretKey string

	Opts Options

	HTTP       *echo.Echo
	Conn       *broker.Conn
	BrokerOpts []broker.Option

	// Static assets
	Static *packr.Box

	// Validate request
	validate *validator.Validate
}

func New(port, storageAccessKey, storageSecretKey string) *Server {
	server := &Server{}
	server.Opts = *opts

	server.ConfigClient()
	server.ConfigWeb()

	return server
}

func (s *Server) NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, _ := s.Static.Resolve("index.html")
		fi, _ := f.Stat()
		http.ServeContent(w, r, "index.html", fi.ModTime(), f)
	})
}

func (s *Server) Run() {
	// Connect to Broker
	showlog := true
	for {
		bc, err := broker.Connect(s.Opts.BrokerURL, s.BrokerOpts...)
		if err != nil {
			if showlog {
				log.Error().Str("type", "server").Err(err).Msg("Desconected: trying to connect ...")
				showlog = false
			}
		} else {
			log.Info().Str("type", "server").Msgf("Server conected to %s", s.Opts.BrokerURL)
			s.Conn = bc
			break
		}
	}

	// Set the topics on the broker

	if _, err := s.Subscribe("adbus.server.list", s.List); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
	}
	if _, err := s.Subscribe("adbus.server.sync", s.Sync); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
	}
	if _, err := s.Subscribe("adbus.server.status", s.Status); err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
	}

	// Start the web server
	log.Info().Str("type", "web").Msgf("Listening for web client connections on :%s", s.Opts.Port)
	if err := s.HTTP.Start(":" + s.Opts.Port); err != nil {
		log.Fatal().Str("type", "web").Err(err).Msg("")
		//Msg("Unable to start the web server on port :" + port)
		os.Exit(1)
	}

	<-make(chan struct{})
}

// Subscribe a topic
func (s *Server) Subscribe(subj string, f func(*broker.Msg)) (*broker.Subscription, error) {
	return s.Conn.Subscribe(subj, func(msg *broker.Msg) {
		go f(msg)
	})
}

// Publish a message
func (s *Server) Publish(subj string, data []byte) error {
	return s.Conn.Publish(subj, data)
}

// List get the list of devices
func (s *Server) List(msg *broker.Msg) {

}

// Sync sync the media files
func (s *Server) Sync(msg *broker.Msg) {

}

// Status get the status info
func (s *Server) Status(msg *broker.Msg) {

}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hash     string `json:"hash"`
	Avatar   string `json:"avatar"`
}

func (s *Server) login(c echo.Context) error {
	//username := c.FormValue("username")
	//password := c.FormValue("password")
	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	if err := s.ValidateLogin(u); err != nil {
		log.Error().Str("type", "web").Err(err).Msg("")
		return echo.ErrUnauthorized
	}
	// TODO: Implement DB with SHA3 password check
	ValidateLogin()
	fmt.Println("Datos:", u.Username, u.Password)

	// Throws unauthorized error
	if u.Username != "admin" || u.Password != "adminadmin" {
		fmt.Println("Error")
		return echo.ErrUnauthorized
	}
	fmt.Println("continue")

	token, err := createJwtToken(s.Opts.SecretKey)
	if err != nil {
		log.Error().Str("type", "web").Err(err).Msg("")
		return err
	}

	cookie := &http.Cookie{}
	cookie.Name = "adbus_dashboard"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	log.Info().Str("type", "web").Msg("token: " + token)
	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (s *Server) singup(c echo.Context) error {
	return nil
}

func (s *Server) info(c echo.Context) error {
	//user := c.Get("user").(*jwt.Token)
	//claims := user.Claims.(*jwtClaims)
	//name := claims.Name
	return nil
}
