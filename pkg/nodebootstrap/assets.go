// Code generated by go-bindata.
// sources:
// assets/10-eksclt.al2.conf
// assets/bootstrap.al2.sh
// assets/bootstrap.ubuntu.sh
// assets/kubelet.yaml
// DO NOT EDIT!

package nodebootstrap

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __10EkscltAl2Conf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4f\x6b\xdb\x4e\x10\xbd\xeb\x53\x2c\x24\x87\xdf\x0f\xbc\x52\xe2\xb8\x39\x04\x74\x50\x63\x25\x18\x54\x27\x44\x0e\x2d\xb4\xc5\x8c\x77\xc7\xce\xe0\xd5\xac\xd8\x5d\xd9\x49\x83\xbf\x7b\x91\x65\xb5\x2e\x09\xa5\x37\xed\xbc\x99\xf7\xde\xfc\xd1\x89\xc0\xb5\x57\xc1\x48\x5f\xa3\xa2\x25\x29\xe1\x5f\x7c\xc0\x4a\x0b\xed\x6c\x2d\x89\x45\xc3\x14\xc4\xd2\x3a\xb1\x6e\x16\x68\x30\x0c\xf6\x8f\xac\x82\x1f\x96\x45\x41\xdc\x3c\x8b\xa1\xf8\x2f\x2b\x86\xff\x47\xd1\xd7\x12\xdd\x86\x14\x7e\x8f\x4e\x44\x61\x15\x18\x51\x61\x00\x0d\x01\x44\x0d\x0e\x2a\x0c\xe8\xfc\x95\x78\xc8\x6f\x27\x77\xd3\x81\xc8\x3e\x97\xf3\x71\x7e\x93\x3d\x16\xb3\x79\x17\x8b\x72\xde\x90\xb3\x5c\x21\x87\x1b\x32\x98\x26\x18\x54\xd2\x59\x4c\x7a\xae\x18\x79\x13\x9d\x88\x5b\x63\x17\x60\x04\xb0\x16\x3e\x40\x20\xf5\x87\xc6\x75\xf1\x58\xce\xf2\x87\xf9\x78\x5a\x0e\xc4\xf4\x6e\x9c\xcf\x8b\xec\x63\x5e\xf4\x8f\x59\x36\x99\xce\xca\xbf\xca\x1d\xfa\x3d\xa8\x75\xed\xb0\x65\xf9\x8e\xd8\x9e\x72\x72\x3f\x10\x93\x69\x39\xcb\xa6\xd7\xf9\x7c\x32\xfe\x27\x6e\xd3\xb2\xee\x15\xa2\xfc\x19\x55\x19\xc0\x85\xf4\xe8\x33\x69\xbc\x4b\x16\xc4\x7d\x81\xf8\x16\x09\x21\x25\x5b\x8d\x92\xea\xf4\xf4\xf5\xa0\xbc\x3b\x06\x0c\x2c\xd0\xf8\x1e\xec\xda\xde\x0d\xc0\xd4\x4f\x10\x77\xfa\x31\xd9\x84\xd8\x07\x60\x85\x92\x74\x7a\xfa\x7a\x64\xbc\xe7\xaa\xe0\x59\xd6\x56\xb7\x44\x9f\xb2\x2f\xf3\xfb\xbb\x71\xd9\x43\x0e\x57\xe4\x03\xba\xbd\x5e\x1a\x5c\x83\xc7\xc1\x2d\x85\x27\x19\x80\x38\xfc\x32\xd1\x8d\xbb\x2f\x07\x63\xec\x56\xd6\x8e\x36\x64\x70\x85\xba\x63\xe8\x30\x65\x6c\xa3\x65\xed\xec\x86\x34\xba\x14\xb6\xbe\x07\x2c\xb7\x9c\xe8\xa4\x6b\x38\x50\x85\xa9\xb6\x6a\x8d\xae\xef\x1c\xc3\xd6\xba\xb5\xac\x4d\xb3\x22\x4e\x15\x53\x5f\xc7\x24\x17\xc4\x52\x93\x4b\x13\x5b\x87\x44\x31\xb5\x23\x3d\x82\x95\xe5\x65\x87\xb7\x2b\x6a\x71\xc6\x10\xeb\x43\x46\x6d\xb5\x24\x5e\x3a\x38\xb2\x40\x15\xac\x30\xbd\x3c\x1b\x8e\xce\xce\xcf\x47\x17\xa3\x0f\xc3\x58\xaf\x5d\x8c\xca\xc5\xa7\xaf\x6f\xcf\x7a\x17\xc3\xfe\x7f\x81\xad\x8f\x95\xad\xda\x2b\x48\x6a\x68\x3c\x4a\xa8\xf4\xe5\xe8\xea\x22\x3e\x3f\x88\xb5\x7b\x6e\xed\xd0\xea\xcd\xbd\x74\xe1\xf8\x05\x2a\xf3\x7b\x24\xef\x25\xb6\x87\xd5\x66\x45\x3f\x03\x00\x00\xff\xff\x5c\x6d\xc6\x40\xdd\x03\x00\x00")

func _10EkscltAl2ConfBytes() ([]byte, error) {
	return bindataRead(
		__10EkscltAl2Conf,
		"10-eksclt.al2.conf",
	)
}

func _10EkscltAl2Conf() (*asset, error) {
	bytes, err := _10EkscltAl2ConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10-eksclt.al2.conf", size: 989, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xd1\x6f\xd3\x3e\x10\xc7\xdf\xef\xaf\xb8\x5f\x7e\xd5\xd4\x6a\x4a\xcd\xd0\x40\x1a\xa5\x48\x13\x2d\x52\x25\xe8\x2a\x6d\x0f\xa0\xaa\x44\x6e\x72\xa1\xd6\x5c\x3b\xb2\x2f\xa3\x23\x0a\x7f\x3b\x72\x9a\x96\x14\xf1\x30\xe0\x29\xb9\xbb\xef\x7d\x7d\xfe\x9c\xff\xff\x4f\xac\x95\x11\x6b\xe9\x37\x00\x9e\x18\x63\x8b\xe4\x1c\xed\x14\x1f\xc2\x42\x15\x94\x4b\xa5\x0f\xb1\xb1\xa5\xf1\xc4\x00\x79\x69\x52\x56\xd6\xe0\x17\xe2\x64\x2b\x77\x49\x61\x33\xdf\x1f\x60\x05\x88\x1f\xae\x3f\x26\x8b\x9b\xc9\x6d\xf2\x6e\xf6\x7e\x3a\x8e\x04\x71\x2a\xe8\xde\xa7\xac\xc5\x41\x99\x6c\x65\x31\xe4\x1d\x47\x80\xf8\x75\xa3\x34\xa1\x23\x99\xa1\x32\x9e\xa5\x49\x29\xe1\xc7\x82\x30\x08\x47\x98\x59\x00\x44\x44\x95\x23\x2e\x97\x18\xf5\xaa\x13\x55\x1d\xe1\x78\x1c\xb2\x17\x75\x84\xab\x15\x9e\x9d\xb5\xaa\xd0\x1d\x8a\xdf\xf1\xf3\xf2\x59\x7c\xb5\x3a\xef\x85\xf2\x08\x79\x43\xa6\x31\x44\xa4\x74\x63\xb1\x55\x8e\xda\x9c\x23\x2e\xdd\x5e\x90\x2b\x1c\x85\xb3\x33\x6b\x08\x5f\x07\xcf\x93\x9b\xd5\x11\xd4\x00\xf3\x9b\xc9\x34\x99\x2d\xc6\xbd\x7e\x5a\x3a\x8d\x71\xec\x95\x26\xc3\xb8\x61\x2e\x5e\x09\x71\xf1\xf2\x6a\xf8\xfc\xc5\xe5\xb0\xfd\x0a\x2d\x99\x3c\x8b\x2d\xb1\x8c\x33\xc9\x52\x68\x9b\x4a\x1d\xab\xe2\xe1\x72\x00\xb3\xf9\xed\xdd\xf5\xfc\xed\x34\x99\x4d\xfe\xda\xef\x00\x27\x56\x59\xc7\xf0\xee\xd3\x62\xfa\xef\x96\x81\xf7\x00\xc0\xdb\xd2\xa5\x84\xdd\xbd\xde\x97\x6b\xd2\xc4\x43\x32\x0f\x00\x2a\x0f\x2b\x88\xbf\x75\x89\x9d\xef\x9a\xf5\x04\xca\x61\x01\x48\xbb\xc2\x3a\x3e\x3e\x95\x71\xaf\xdf\x7d\x48\xd8\xab\x4e\x46\xaf\x07\x23\xc8\x15\x40\xb3\xb0\xe8\x88\xbc\x6a\xff\xea\x08\xdf\xfc\x76\x9c\x06\x6e\x33\xd4\xbe\xf3\x04\x70\xd5\x89\x82\xc3\x1f\x5a\xec\x91\xfe\x32\xe7\x93\x6d\x7e\x5e\xfc\x88\xe8\x29\xcd\xe0\x1f\x3d\xd3\x36\x65\x8d\x99\xa4\xad\x35\xb1\x23\x6d\x65\xd6\xc9\x93\x91\x6b\x4d\xd8\xf6\x76\x0a\x9e\xa5\xe3\x63\xfe\x47\x00\x00\x00\xff\xff\x69\x53\x55\x98\xfd\x03\x00\x00")

func bootstrapAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapAl2Sh,
		"bootstrap.al2.sh",
	)
}

func bootstrapAl2Sh() (*asset, error) {
	bytes, err := bootstrapAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.al2.sh", size: 1021, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapUbuntuSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x6f\x6f\xdb\xb6\x13\x7e\xaf\x4f\x71\x3f\xd7\xe8\x2f\x41\x4b\x29\x49\xd3\x02\x6d\xa6\x61\x5e\xed\x0e\xc6\xb2\xa4\xa8\x5d\x74\x43\x90\x19\x67\xf1\x64\x11\xa6\x48\x81\x3c\xd9\x49\x0d\xef\xb3\x0f\x94\x2c\xd7\x2e\xb0\x3f\x1d\xf2\x22\xe6\xdd\x73\xcf\x1d\x8f\xf7\x9c\x9e\xfc\x2f\x99\x2b\x93\xcc\xd1\x17\x51\xe4\x89\x41\x58\x20\xe7\xe8\x41\x71\x77\xac\x54\x45\x39\x2a\xdd\x9d\x8d\xad\x8d\x27\x8e\xa2\xbc\x36\x19\x2b\x6b\x60\x41\x3c\x2b\xf1\x61\x56\x59\xe9\x4f\x4e\x61\x13\x01\xfc\x32\xf8\x75\xf6\xfe\x76\x38\x99\xbd\x1b\x5f\x8f\xd2\x5e\x42\x9c\x25\xb4\xf4\x19\xeb\xa4\x43\xce\x4a\xac\x62\x7e\xe0\x5e\x04\xb0\x2e\x94\x26\x70\x84\x12\x94\xf1\x8c\x26\xa3\x19\x3f\x56\x04\x01\x78\x05\xd2\x46\x11\x00\x80\xca\x01\xee\xee\xa0\xd7\xdf\x1c\xa1\xb6\x3d\x48\xd3\x60\x3d\xdf\xf6\xe0\xfe\x1e\x9e\x3e\xdd\xa1\x42\x74\x70\xfe\x01\xbf\xdf\x9d\x89\xd7\xf7\xcf\xfa\xc1\x7d\x05\x5c\x90\x69\x08\x01\x28\x2b\x2c\xec\x90\x57\x3b\x9b\x23\xae\x5d\x0b\xc8\x15\x5c\x85\xdc\xd2\x1a\x82\xef\x02\xe7\xd1\xcd\xb6\xbd\x68\x1b\x45\x0d\x47\xef\xe6\x76\x38\x9a\x8d\xdf\xa7\xfd\x93\xac\x76\x1a\x84\xf0\x4a\x93\x61\x28\x98\xab\x37\x49\x72\xfe\xea\x75\x7c\xf1\xf2\x32\xde\xfd\x4f\x34\x32\x79\x4e\x4a\x62\x14\x12\x19\x13\x6d\x33\xd4\x42\x55\xab\xcb\xd3\x1e\x7c\x0f\x87\x2d\x5b\xd6\x73\xd2\xc4\x71\x03\x89\xc9\xac\x76\x29\xc7\x37\x93\xe9\xe0\xe6\xed\x68\x36\x1e\xfe\xe7\xb4\x5d\x27\x85\x92\x21\xef\x37\x26\x9e\xfe\xf6\x7e\xb4\x4f\xed\xbf\x3d\x69\x78\xbe\x7f\x93\x36\xf2\x06\x2b\x40\xad\xd0\xc3\xce\x2b\x68\xe9\xe3\xdd\xef\xce\xf6\x35\x2c\x63\xbd\x87\x65\xac\x3b\x5b\x0b\xf3\x6c\xab\x43\xb2\xc8\x3f\x7a\xa6\x32\xe0\x1c\x79\x62\x11\xa6\x9e\x64\x14\x9d\x44\x00\x4f\x60\x7a\x3b\xbc\x7d\x13\x46\xc7\x13\xf8\xc2\xd6\x5a\xc2\x9c\x40\x5b\xbb\x24\x09\xc8\x40\x2b\x72\x8f\xc0\xaa\xa4\x8e\x14\x3c\xa3\x63\x0f\x75\xf5\xbc\x61\x58\x17\x2a\x2b\x40\x79\x58\x17\xc8\xb0\x26\x90\x16\x94\x81\xc1\xf5\x05\x9c\xec\x7d\x73\xf4\x24\xc1\x1a\xa8\x34\x2a\x03\x6d\x4d\xb2\x25\x40\x23\xa1\x24\x34\x0c\x6c\x43\xf2\xca\x3a\xc6\xb9\xa6\x70\x2c\xad\xe7\x0e\x0d\x52\x79\x76\xd6\x9f\x3e\x87\x79\xcd\xa0\xf8\xff\xbe\x89\x37\x96\x21\xd3\x84\x0e\x0a\xbb\x0e\x41\xda\xa2\xdc\x5d\x29\x77\xb6\xfc\x52\x78\xe8\xcf\x5a\x71\x61\x6b\x86\x02\x57\xca\x2c\x1a\x02\xb6\x90\xd5\x9e\x6d\xa9\x3c\x85\xb8\x16\xa8\xd8\x93\xce\x23\x00\x6f\x6b\x97\xd1\x3f\x3c\xe5\xdf\xc2\xfe\x12\x10\x66\x27\x8c\x4e\x3b\x0d\xcd\x2a\xb8\xbb\x03\xf1\xf9\x50\x92\xcf\x1e\x1a\xfd\xb7\x32\x0e\x1a\x07\x7a\x08\x2d\xda\x6f\xa3\xb4\x7f\x72\xb8\xab\xa0\xbf\x39\x1a\xe4\xed\x69\x08\xcd\x55\xe0\xcf\x35\x2e\x7c\x7a\xd2\x50\xf5\x50\x4a\x47\xde\xa7\x67\x71\xf3\xd7\x6b\xad\xc6\x4a\x12\xaa\x4a\xfb\x9b\x9d\xf4\xb7\x3b\x47\xa6\x6b\xcf\xe4\x84\x34\x3e\xed\x6f\xde\x5e\x7f\x9c\x4c\x47\x1f\x66\xc3\x9b\x49\x07\x28\xf1\x41\x84\x02\xd2\x2f\xb5\x6f\x0f\x49\x35\xce\x49\xfb\x8e\xf8\x7a\xf0\xe3\xe8\x7a\xb2\x7d\x8e\xba\x2a\x30\x6e\xfb\x11\x2b\x7b\xa8\xdc\xf4\xe0\x26\xe3\x61\xc7\x85\x75\x68\x02\xab\x0c\xc3\x9a\x16\x6c\x97\x64\xc4\x9a\xe6\x85\xb5\xcb\x94\x5d\x4d\x07\x38\xeb\xd4\xe7\x16\x56\x5a\x49\xe9\xa7\x16\xd5\x01\xb4\xb6\x6b\x51\x39\xb5\x52\x9a\x16\x24\x0f\x83\x2b\x2b\x85\x32\xb9\x43\x91\x59\xc3\xa8\x0c\x39\xa1\x4a\x5c\x50\xfa\xea\xec\xe2\xf2\xec\xfc\xfc\xf2\xc5\xe5\xcb\x8b\x58\x2e\x5d\x4c\x99\x8b\xfb\x9b\xc1\xa7\xc9\x6c\x38\x7a\x37\xf8\x78\x3d\x9d\x7d\x18\xfd\x34\xbe\xbd\xd9\xc6\x58\xe2\x67\x6b\x70\xed\xe3\xcc\x96\xe1\xc9\x93\x0a\x6b\x4f\x02\x4b\xf9\xea\xf2\xcd\x8b\xf8\x7c\xdf\x59\x5b\x4b\x51\x39\xbb\x52\x92\x5c\x8a\x6b\xff\x75\xcb\x6d\x89\xca\xa4\xbb\x63\x3b\x75\x1d\xc4\x28\x31\x57\x46\x48\xe5\xd2\xc4\x56\x9c\x64\x46\x85\x8f\xdf\x81\x3b\xb3\x26\x6f\xfd\x61\xf2\x82\xdf\x10\xc7\xb2\x43\xec\xef\xe7\x6a\x13\x74\x9e\x4a\x9b\x2d\xc9\x75\x2f\x47\xbc\xb6\x6e\x29\x2a\x5d\x2f\x42\x09\x46\x75\x71\x0b\x67\xeb\x4a\x48\xa7\x56\xe4\xd2\xf6\x94\x77\x85\x3b\x5a\xa8\xa6\xf2\xf0\xf0\x87\x7d\xdd\x3b\x82\x0c\x45\x48\xcc\xfb\x89\x98\x0e\xc6\x37\xd3\xfd\xc8\x34\x4b\xcd\x9a\x5c\x2d\xd2\xaf\x25\xd5\x9a\xe3\x47\x2c\xbb\x2e\xe4\x84\x5c\x3b\x12\x8b\xb0\x92\xd3\x0f\x96\x91\xe9\xe7\x56\x7c\x13\x72\x2b\x72\x6f\xc9\xb1\xca\xc3\xcc\x1c\x95\x83\xc6\x9a\xc7\xd2\xd6\x5e\x84\x69\x49\x73\xd4\x9e\xf6\xbd\x57\x64\x58\x64\x28\x72\xa5\xe9\xa8\x86\x0c\xe3\xcc\x35\x5f\xf8\xd3\x20\xaa\x76\xed\x7e\x59\xd7\x61\xeb\x06\x01\x37\x62\xbb\xfb\xe1\x7e\xdb\x8b\x4e\xa3\x6e\x39\xa3\x3b\xc2\x45\x7f\x06\x00\x00\xff\xff\xab\xe8\x06\xe8\xab\x08\x00\x00")

func bootstrapUbuntuShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapUbuntuSh,
		"bootstrap.ubuntu.sh",
	)
}

func bootstrapUbuntuSh() (*asset, error) {
	bytes, err := bootstrapUbuntuShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.ubuntu.sh", size: 2219, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kubeletYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x4f\x4f\xc3\x30\x0c\xc5\xef\xf9\x14\xfe\x04\x6d\x07\x9a\x04\xb9\x8d\x4d\x70\x60\x27\x36\xe0\xec\xa6\xee\x16\x35\x8d\x27\xc7\x19\x7f\x3e\x3d\x5a\x5a\x90\x26\xa1\x9c\x9e\xde\xb3\xdf\x4f\xce\xe0\x63\x67\xe1\x39\xb7\x14\x48\xd7\x1c\x7b\x7f\xc8\x82\xea\x39\x1a\x3c\xf9\x37\x92\xe4\x39\x5a\x18\xa6\x40\xe5\x4a\xa2\x1a\xee\x52\xe5\xb9\x3e\x2f\x5a\x52\x5c\x18\x83\x5d\x27\x94\x92\x85\xa6\x2a\xcf\xb8\x90\x93\x92\x6c\x78\x44\x1f\x2d\xcc\xb2\x0a\xec\x30\x18\x83\x59\x8f\x14\xd5\xbb\x52\x64\x0d\x00\x46\x8e\x5f\x23\xe7\x74\x11\x00\x14\xb1\x0d\xd4\x59\xe8\x31\x24\x32\x00\x1f\xd4\x1e\x99\x87\xc9\x75\xe8\x8e\xb4\xdf\x6f\x2d\xdc\x8c\x4d\xba\x1e\x50\xc9\x97\xfc\xe7\xb2\xb9\x9f\xc3\xc1\x53\xd4\xf5\xea\xd1\x07\xb2\x50\x93\xba\x9a\x86\xe4\x34\xd4\x0e\x2b\x27\x3a\xd1\xb0\xf8\xef\x3f\x98\x91\x3b\xb2\xf0\x3e\x55\xfe\x5b\xbe\x9a\x47\xa8\x2b\x18\xcb\x5f\x8c\x62\xbe\x46\xbc\xb6\x6f\x9b\x64\x4c\x22\x39\x93\xec\xb7\xbb\x07\x66\x4d\x2a\x78\x9a\x61\x8d\x3b\x08\xe7\xd3\x46\xfc\x99\xc4\xc2\xa4\xfa\x64\x4c\x4f\xa8\x59\xe8\x09\x95\xca\x59\x5e\x58\x51\x69\xfe\xaa\x5d\x59\xb7\x26\x51\xdf\x5f\xee\x48\xf3\xb6\x9f\x00\x00\x00\xff\xff\x1f\x2f\xa9\x0f\xd0\x01\x00\x00")

func kubeletYamlBytes() ([]byte, error) {
	return bindataRead(
		_kubeletYaml,
		"kubelet.yaml",
	)
}

func kubeletYaml() (*asset, error) {
	bytes, err := kubeletYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubelet.yaml", size: 464, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"10-eksclt.al2.conf": _10EkscltAl2Conf,
	"bootstrap.al2.sh": bootstrapAl2Sh,
	"bootstrap.ubuntu.sh": bootstrapUbuntuSh,
	"kubelet.yaml": kubeletYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"10-eksclt.al2.conf": &bintree{_10EkscltAl2Conf, map[string]*bintree{}},
	"bootstrap.al2.sh": &bintree{bootstrapAl2Sh, map[string]*bintree{}},
	"bootstrap.ubuntu.sh": &bintree{bootstrapUbuntuSh, map[string]*bintree{}},
	"kubelet.yaml": &bintree{kubeletYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

