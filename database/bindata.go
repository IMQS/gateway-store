// Code generated by go-bindata.
// sources:
// 001_create.down.sql
// 001_create.up.sql
// DO NOT EDIT!

package database

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

var __001_createDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe5\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x28\x4d\xca\xc9\x4c\xd6\xcb\x4d\x2d\x2e\x4e\x4c\x4f\x2d\xb6\xe6\xe5\x72\xf6\xf7\xf5\xf5\x0c\xb1\x06\x04\x00\x00\xff\xff\xbf\x16\x37\xd6\x2c\x00\x00\x00")

func _001_createDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_createDownSql,
		"001_create.down.sql",
	)
}

func _001_createDownSql() (*asset, error) {
	bytes, err := _001_createDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_create.down.sql", size: 44, mode: os.FileMode(438), modTime: time.Unix(1547203400, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_createUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x31\x4f\xc3\x30\x10\x85\x77\x4b\xfe\x0f\x37\x26\x12\x42\x30\x30\x45\x0c\x6e\xb9\x80\x85\xe3\x80\xe3\xa8\x74\x74\x53\xab\x04\x35\xa1\xc4\x29\x52\xff\x3d\x4a\x9c\x08\x86\x4a\x6d\x50\x57\xdf\xbb\xa7\xef\xbd\xf3\x0c\x1f\xb9\x8c\x28\xa1\x64\xae\x90\x69\x04\xcd\x66\x02\x81\xc7\x20\x53\x0d\xf8\xc6\x33\x9d\x41\xb1\x2d\x6d\xdd\x52\x12\x50\x02\x00\xe5\x1a\x32\x54\x9c\x09\x78\x51\x3c\x61\x6a\x09\xcf\xb8\xbc\xea\x47\x5e\xc8\xd7\x50\xbc\x9b\xc6\x14\xad\x6d\xe0\xdb\x34\x87\xb2\xde\x04\x77\x37\x61\x6f\x29\x73\x21\x20\x97\xfc\x35\x47\xbf\x53\x9b\xca\x9e\xd0\x7b\xa1\x6b\x4d\xbb\x77\x47\xa4\xb7\x7f\xa4\x94\x84\x94\x2c\xb8\x7e\x82\x1e\x36\xe5\x0f\x19\xdc\x43\xcc\x44\x86\xfd\x28\xa2\x84\x09\x8d\x6a\xc8\x39\x26\x03\x48\x17\xb2\x7b\x4d\xa1\xac\xbe\xdc\xa9\x42\xda\xc3\xce\x9e\x51\xc7\x19\xd1\x86\x2a\x26\x62\xef\xf6\xab\x6d\x59\x5c\x7b\x8c\x89\xec\x95\x75\xce\x6c\xac\xbb\xe0\x39\x15\xc6\xa8\x50\xce\x71\xfc\x2a\xc1\xb8\x19\x7a\xa3\x0e\x74\x82\x49\x27\x0f\xba\xf2\x86\xf5\x01\x19\x3e\xdc\x67\xbd\xfa\xef\xa9\x7f\x73\x1f\x2d\x2c\x4d\x12\xae\x23\x4a\x7e\x02\x00\x00\xff\xff\x60\x22\x63\x53\x11\x03\x00\x00")

func _001_createUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_createUpSql,
		"001_create.up.sql",
	)
}

func _001_createUpSql() (*asset, error) {
	bytes, err := _001_createUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_create.up.sql", size: 785, mode: os.FileMode(438), modTime: time.Unix(1552640835, 0)}
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
	"001_create.down.sql": _001_createDownSql,
	"001_create.up.sql":   _001_createUpSql,
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
	"001_create.down.sql": &bintree{_001_createDownSql, map[string]*bintree{}},
	"001_create.up.sql":   &bintree{_001_createUpSql, map[string]*bintree{}},
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
