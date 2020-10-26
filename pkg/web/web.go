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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	packr "github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

//
type ServiceConfig struct {
	URL       string
	AccessKey string
	SecretKey string
}

//
type Config struct {
	Port     string
	Storage  ServiceConfig
	Database string
	Certs    string
	Path     string
}

type Web struct {
}

func New() {

}

func (s *Server) ConfigWeb() {

	s.Static = packr.New("app", "../../web/dist")
	s.HTTP = echo.New()
	s.HTTP.HideBanner = true
	staticPackr(s.HTTP, "/", s.Static)
	s.HTTP.Logger.SetOutput(ioutil.Discard)

	s.validate = validator.New()
	//s.HTTP.File("/", "web/dist/index.html")
	//s.HTTP.File("/favicon.ico", "web/dist/favicon.ico")
	s.HTTP.Any("/sse/event", s.eventStreaming)

	s.HTTP.GET("/", s.index)
	s.HTTP.GET("/login", s.index)
	s.HTTP.GET("/sigin", s.index)

	s.HTTP.POST("/login", s.login)
	s.HTTP.POST("/singup", s.singup)
	s.HTTP.GET("/info", s.info)

	api := s.HTTP.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:        &jwtClaims{},
		SigningMethod: "HS512",
		SigningKey:    []byte(s.Opts.SecretKey),
		TokenLookup:   "cookie:JWTCookie",
	}))

	/*
		api.GET("/api/devices", s.listDevices)
		api.GET("/api/devices/:id", s.getDevice)

		api.GET("/api/status", s.getStatus)
		api.GET("/api/status/:id", s.getStatus)

		api.GET("/api/listPlaylist", s.listPlaylist)
		api.GET("/api/listPlaylist/:id", s.getPlaylist)

		api.PUT("/api/playlists", s.updatePlaylist)
		api.PUT("/api/playlists/:id", s.updatePlaylist)

		api.POST("/api/sync", s.actionSync)
		api.POST("/api/sync/:id", s.actionSync)

		api.POST("/api/test_pattern", s.actionTestPatern)
		api.POST("/api/test_pattern/:id", s.actionTestPatern)
	*/

	s.HTTP.HTTPErrorHandler = s.customHTTPErrorHandler
}

func (s *Server) ValidateLogin(u *User) error {
	db := sql.New()
	db.Open("adbus", s.Opts.Database)
	defer db.Close()

	value, err := DB.Get(u.Username)
	if err != nil {
		return err
	}
	user := &User{}
	err = json.Unmarshal([]byte(value), user)
	if err != nil {
		return err
	}
	hashByte, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hash := string(hashByte)
	if hash != user.Hash {
		return errors.New("Incorrect Password")
	}
	return nil
}

func (s *Server) customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		log.Error().Str("type", "web").Err(err).Msg("")
		f, _ := s.Static.Resolve("index.html")
		fi, _ := f.Stat()
		http.ServeContent(c.Response(), c.Request(), "index.html", fi.ModTime(), f)
	}
	log.Error().Str("type", "web").Err(err).Msg("")
}

// Web API
func (s *Server) eventStreaming(c echo.Context) error {
	// Set the headers related to event streaming.
	header := c.Response().Header()
	header.Set("Content-Type", "text/event-stream")
	header.Set("Cache-Control", "no-cache")
	header.Set("Connection", "keep-alive")
	header.Set("Access-Control-Allow-Origin", "*")
	return nil
}

func createJwtToken(secret string) (string, error) {
	// Set custom claims
	claims := &jwtClaims{
		Name:  "AdBus",
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	// Create token with claims
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// Generate encoded token and send it as response.
	token, err := rawToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Server) index(c echo.Context) error {
	html, _ := s.Static.FindString("index.html")
	return c.HTML(http.StatusOK, html)
}

func staticPackr(e *echo.Echo, prefix string, box *packr.Box) *echo.Route {
	h := func(c echo.Context) error {
		//route := c.Request().URL.Path
		route, err := url.PathUnescape(c.Param("*"))
		if err != nil {
			return err
		}
		route = strings.TrimPrefix(route, "/")

		if route == "" {
			return nil
			//route = "index.html"
		}

		f, err := box.Resolve(route)
		if err != nil {
			index, _ := box.Resolve("index.html")
			f = index
			//return nil
		}
		fi, _ := f.Stat()
		http.ServeContent(c.Response(), c.Request(), route, fi.ModTime(), f)
		return nil
	}
	e.GET(prefix, h)
	if prefix == "/" {
		return e.GET(prefix+"*", h)
	}
	return e.GET(prefix+"/*", h)
}

// Use packr with echo for s static files

/*
	s.HTTP.HTTPErrorHandler = func(err error, c echo.Context) {
		f, _ := s.Static.Resolve("index.html")
		fi, _ := f.Stat()
		http.ServeContent(c.Response(), c.Request(), "index.html", fi.ModTime(), f)
	}


	echo.NotFoundHandler = func(c echo.Context) error {
		//f, _ := s.Static.Resolve("index.html")
		//fi, _ := f.Stat()
		// http.ServeContent(c.Response(), c.Request(), "index.html", fi.ModTime(), f)
		return c.String(http.StatusNotFound, "Error404")
	}
*/

//s.HTTP.GET("*", s.all)

//s.HTTP.GET("/", echo.WrapHandler(s.NotFoundHandler()))
//s.HTTP.GET("/*", echo.WrapHandler(http.FileServer(s.Static)))
