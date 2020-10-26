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

// go clean -testcache
// go test -count=1 -v

package sync

import (
	"fmt"
	"io/ioutil"
	"testing"

	minio "github.com/minio/minio-go"
	"github.com/stretchr/testify/assert"
)

const endpoint = "127.0.0.1:9000"
const accessKey = "admin"
const secretKey = "adminadmin"
const bucket = "testdata"
const local = "/tmp/testdataFS"

func TestSync(t *testing.T) {
	sync := &Sync{
		endpoint:  endpoint,
		accessKey: accessKey,
		secretKey: secretKey,
	}

	if err := sync.S3toLocal(bucket, local); err != nil {
		t.Log(err)
	}

	s3 := listS3(bucket)
	fs := listFS(local)
	assert.Equal(t, s3, fs, "they should be equal")
}

func listS3(bucket string) []string {
	doneCh := make(chan struct{})

	client, err := minio.New(endpoint, accessKey, secretKey, false)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var list []string
	// Source from minio s3
	for f := range client.ListObjects(bucket, "", false, doneCh) {
		list = append(list, f.Key)
	}
	return list
}

func listFS(path string) []string {
	localFiles, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var list []string
	for _, f := range localFiles {
		list = append(list, f.Name())
	}
	return list
}
