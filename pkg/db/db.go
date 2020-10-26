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

package db // import "github.com/ljgago/adbus/pkg/db"

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username string
	Hash     string
	Avatar   string
	Roles    []string
}

type DB interface {
	Open(string, string) error
	Close() error
	Create(interface{}, interface{}) error
}

type sqlDB struct {
	db *gorm.DB
}

func New() *sqlDB {
	return &DB{}
}

// Open open the database with the name and address
func (d *sqlDB) Open(name, addr string) error {
	// const addr = "postgresql://maxroach@localhost:26257/bank?sslmode=disable"
	var err error
	d.db, err = gorm.Open(name, addr)
	if err != nil {
		return err
	}
	return nil
}

// Close close the database
func (d *DB) Close() error {
	return d.db.Close()
}

// Create create a new table in the database
func (d *DB) Create(schema, value interface{}) {
	// Migrate the schema
	d.db.AutoMigrate(schema)
	d.db.Create(value)
}

// Search search the content of a table
func (d *DB) Search(value interface{}, args ...interface{}) {
	d.db.Where("username = ? AND hash = ?", args).Find(&value)
}

// Update update a table
func (d *DB) Update(value interface{}, attrs ...interface{}) {
	d.db.Model(&value).Update(attrs)
}
