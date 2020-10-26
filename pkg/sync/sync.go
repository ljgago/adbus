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

package sync // import "github.com/ljgago/adbus/pkg/sync"

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	minio "github.com/minio/minio-go"
)

// FileURI struct is the information data from files
type FileURI struct {
	Bucket string
	Name   string
	Path   string
	Scheme string
	Size   int64
}

// Action struct are the data to send to workers
type Action struct {
	Src  *FileURI // source data
	Dst  *FileURI // destination data
	Type int      // action (COPY, DELETE, CHECKSUM)
}

// Sync struct is the
type Sync struct {
	// Minio / S3 credentials:
	endpoint  string
	accessKey string
	secretKey string
}

// New create a new instance Sync
func New(endpoint, accessKey, secretKey string) *Sync {
	return &Sync{
		endpoint:  endpoint,
		accessKey: accessKey,
		secretKey: secretKey,
	}
}

// S3toLocal sync all files in a S3 bucket to local disk
func (s *Sync) S3toLocal(bucket, local string) error {
	const (
		ActCOPY     = iota
		ActREMOVE   = iota
		NumCOPY     = 4
		NumCHECKSUM = 1
		QueueSIZE   = 1000000 // 1M
	)

	var (
		wg      sync.WaitGroup
		srcFile = make(map[string](*FileURI))
		dstFile = make(map[string](*FileURI))
	)

	copyCh := make(chan Action, QueueSIZE)
	removeCh := make(chan Action, QueueSIZE)
	doneCh := make(chan struct{})

	// Ckeck if the folder exist
	if _, err := os.Stat(local); os.IsNotExist(err) {
		return err
	}

	wg.Add(1)
	go s.workerRemove(&wg, removeCh)

	wg.Add(NumCOPY)
	for i := 0; i < NumCOPY; i++ {
		go s.workerCopy(&wg, copyCh)
	}

	allFiles := make(map[string]bool)

	client, err := minio.New(s.endpoint, s.accessKey, s.secretKey, false)
	if err != nil {
		return err
	}
	// Source from minio s3
	for f := range client.ListObjects(bucket, "", false, doneCh) {
		srcFile[f.Key] = &FileURI{
			Scheme: "s3",
			Bucket: bucket,
			Name:   f.Key,
			Size:   f.Size,
		}
		allFiles[f.Key] = true
	}

	// Local files from disk
	localFiles, err := ioutil.ReadDir(local)
	if err != nil {
		return err
	}
	// Get all files names
	for _, f := range localFiles {
		dstFile[f.Name()] = &FileURI{
			Scheme: "file",
			Name:   f.Name(),
			Path:   local,
			Size:   f.Size(),
		}
		allFiles[f.Name()] = true
	}

	for name := range allFiles {
		// Copy src file to dst file if dst is empy
		if srcFile[name] != nil && dstFile[name] == nil {
			copyCh <- Action{
				Src: srcFile[name],
				Dst: &FileURI{
					Scheme: "file",
					Name:   name,
					Path:   local,
				},
				Type: ActCOPY,
			}
			continue
		}
		// Copy src file to dst file if the src file was modified
		if srcFile[name] != nil && dstFile[name] != nil {
			if (srcFile[name]).Size != (dstFile[name]).Size {
				copyCh <- Action{
					Src:  srcFile[name],
					Dst:  dstFile[name],
					Type: ActCOPY,
				}
				continue
			}
		}
		// Remove the dst file if src file does not exists
		if srcFile[name] == nil && dstFile[name] != nil {
			removeCh <- Action{
				Dst:  dstFile[name],
				Type: ActREMOVE,
			}
			continue
		}
	}
	close(copyCh)
	close(removeCh)
	wg.Wait()

	return nil
}

func (s *Sync) workerCopy(wg *sync.WaitGroup, jobs <-chan Action) {
	defer wg.Done()
	for item := range jobs {
		client, err := minio.New(s.endpoint, s.accessKey, s.secretKey, false)
		if err != nil {
			//log.Error().Str("type", "sync").Err(err).Msg("")
			return
		}
		file := filepath.Join(item.Dst.Path, item.Dst.Name)
		err = client.FGetObject(item.Src.Bucket, item.Src.Name, file, minio.GetObjectOptions{})
		if err != nil {
			//log.Error().Str("type", "sync").Err(err).Msg("")
			return
		}
	}
}

func (s *Sync) workerRemove(wg *sync.WaitGroup, jobs <-chan Action) {
	defer wg.Done()
	for item := range jobs {
		file := filepath.Join(item.Dst.Path, item.Dst.Name)
		if err := os.Remove(file); err != nil {
			//log.Error().Str("type", "sync").Err(err).Msg("")
		}
	}
}
