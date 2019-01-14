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

var __001_createUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\x4b\xc3\x30\x1c\xc5\xef\x81\x7c\x87\xc7\x4e\x2d\x88\xcc\x83\xa7\xe2\x21\x9b\x99\x06\xd3\x44\x92\xbf\xcc\x9d\xa4\x6b\xc3\x8c\x6e\xb5\x36\x55\xd8\xb7\x97\xd5\x9d\xb6\x9d\x7f\xbf\x07\xef\xbd\x99\x7c\x50\xa6\xe0\x6c\xee\xa4\x20\x09\x12\x33\x2d\xa1\x16\x30\x96\x20\x5f\x95\x27\x8f\xee\x67\xbd\x8d\xf5\xf5\x2e\xa4\x54\x6d\x42\xe2\x2c\xe3\x0c\x40\x6c\x50\xbf\x57\x7d\x55\x0f\xa1\xc7\x6f\xd5\xef\x63\xbb\xc9\x6e\xa6\xd3\x7c\xcc\x9a\x17\xad\xaf\x46\x6f\x52\x6f\x63\x68\x07\xd5\x4c\x2e\xf8\xb7\x67\xfa\xb0\xef\xc2\x65\xf1\x9f\x1f\x6b\xe0\x23\x7d\xb5\xeb\x93\xec\xdc\x1a\x4f\x4e\x28\x43\xe8\x3e\xdf\x62\x83\x67\xa7\x4a\xe1\x56\x78\x92\x2b\x64\xb1\xc9\x39\xcb\x39\x5b\x2a\x7a\xc4\xb8\xc1\xaa\x7b\x8f\x3b\x2c\x84\xf6\x72\x44\x05\x67\x42\x93\x74\xc7\x1f\xce\x96\x03\x76\x69\x0e\xd8\x22\xee\xbe\xd3\xe1\x37\x5b\x96\x8a\x0a\xce\xfe\x02\x00\x00\xff\xff\xcd\x6e\x4d\xbc\x49\x01\x00\x00")

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

	info := bindataFileInfo{name: "001_create.up.sql", size: 329, mode: os.FileMode(438), modTime: time.Unix(1547200509, 0)}
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
