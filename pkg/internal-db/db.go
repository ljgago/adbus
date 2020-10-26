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

package intdb

import (
	"encoding/json"
	"errors"

	"github.com/dgraph-io/badger"
)

// Dir is the directory path for save the database
var Dir string

// DB is the main database struct
type DB struct {
	db  *badger.DB
	Dir string
}

// New return a new DB struct
func New(dir string) *DB {
	return &DB{
		Dir: dir,
	}
}

// Open open the database
func (d *DB) Open() error {
	// Open the Badger database.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions
	opts.Dir = d.Dir
	opts.ValueDir = d.Dir
	var err error
	d.db, err = badger.Open(opts)
	return err
}

// Get get the value from database
func (d *DB) Get(key string) (string, error) {
	var value []byte
	var err error
	err = d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		value, err = item.Value()
		if err != nil {
			return err
		}
		return nil
	})
	return string(value), err
}

// Set set a new in the database
func (d *DB) Set(key string, v interface{}) error {
	value, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = d.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
	return err
}

func (d *DB) SetSafe(key string, v interface{}) error {
	if _, err := d.Get(key); err == nil {
		return errors.New("The key already exist, you can't overwrite it")
	}
	return d.Set(key, v)
}

func (d *DB) Delete(key string) error {
	err := d.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
	return err
}

// Close close the database
func (d *DB) Close() error {
	return d.db.Close()
}

//hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//if err != nil {
//	return nil, err
//}

/*

func (s *Server) login(c echo.Context) error {
	//username := c.FormValue("username")
	//password := c.FormValue("password")
	db.Dir = s.Opts.Config
	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	if err := s.ValidateLogin(u); err != nil {
		log.Error().Str("type", "web").Err(err).Msg("")
		return echo.ErrUnauthorized
	}
	// TODO: Implement DB with SHA3 password check
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
*/
