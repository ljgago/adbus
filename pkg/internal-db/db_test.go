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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestExample struct {
	Username string `json:"username"`
	Hash     string `json:"hash"`
	Avatar   string `json:"avatar"`
}

var mydb *intdb.DB
var toWrite TestExample
var toRead TestExample
var key = "key"

func TestMain(m *testing.M) {
	mydb = intdb.New("/tmp/badger")
	mydb.Open()

	defer mydb.Close()

	// Remove the temporal folder
	os.RemoveAll("/tmp/badger")

	os.Exit(m.Run())
}

func TestSet(t *testing.T) {
	toWrite = TestExample{
		Username: "username",
		Hash:     "hash",
		Avatar:   "avatar",
	}
	if err := mydb.Set(key, toWrite); err != nil {
		t.Error("Set operation failed: ", err)
	}
}

func TestGet(t *testing.T) {
	// Read from database
	data, err := mydb.Get(key)
	if err != nil {
		t.Error("Get operation failed: ", err)
	}
	_, err = mydb.Get("key2")
	if err != nil {
		t.Log("OK, the key is not exist")
	}
	json.Unmarshal([]byte(data), &toRead)

}

func TestCompare(t *testing.T) {
	assert.Equal(t, toWrite, toRead, "they should be equal")
}

func TestSetSafe(t *testing.T) {
	// Tries to write to same key
	if err := mydb.SetSafe(key, toWrite); err != nil {
		t.Log("OK, SetSafe operation does not write by to be the same key")
	}
	// Write with diferent key
	if err := mydb.SetSafe("key2", toWrite); err != nil {
		t.Error("SetSafe operation failed: ", err)
	}
}

func TestDelete(t *testing.T) {
	// Delete the key
	if err := mydb.Delete(key); err != nil {
		t.Error("Delete operation failed: ", err)
	}
}
