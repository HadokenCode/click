// Code generated by go-bindata.
// sources:
// static/migrations/1_initial.sql
// DO NOT EDIT!

package static

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

var _staticMigrations1_initialSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x5d\x6f\xe2\x48\x10\x7c\x9f\x5f\xd1\xb2\x22\x01\x39\x12\x5d\xa2\x9c\x74\x3a\xeb\x1e\x0c\x6e\xc8\xdc\x99\x31\x3b\x1e\x6f\x92\x27\xe4\xc5\x23\xd6\x5a\x63\x58\x7b\x9c\xec\xcf\x5f\x8d\xbf\x30\x60\x48\x76\x13\xbf\x79\x52\x5d\xe9\xaa\xea\x1e\x73\x75\x05\x7f\xac\xa3\x55\x1a\x28\x09\xfe\x96\x90\x31\x47\x4b\x20\xe0\xa3\x40\xe6\x51\x97\x01\x9d\x00\x73\x05\xe0\x23\xf5\x84\x07\x46\x9e\x47\xe1\xd5\x26\xcb\xb6\x86\x49\x48\xbb\xd8\x53\x81\x92\x6b\x99\xa8\x91\x5c\x45\x49\xcd\x33\xf1\xd9\x58\x68\x9a\x7c\x1b\x06\x4a\x2e\x54\xb4\x96\x99\x0a\xd6\xdb\xfe\x80\x00\x70\x14\x3e\x67\x1e\x08\x4e\xa7\x53\xe4\x60\x79\x70\x71\x08\xbc\x20\x23\x9c\x52\x46\xa0\xe8\x04\x1f\xae\x2f\x81\x7a\x60\x53\x4f\x50\x36\x16\x30\xe1\xee\x0c\x5c\xc7\xbe\xbe\x24\x00\xe2\x1e\x35\x10\x0a\x5c\x49\x14\x2e\x02\x05\xff\xfc\x0b\xcb\x3c\x4d\x65\xa2\x76\xbc\x66\x01\x2c\x3b\xd0\x78\xfd\x8e\x8e\x87\xed\x63\xd7\xb1\x8b\x63\x66\x03\x9d\x98\x04\x99\x6d\x92\x8e\x06\x1d\x8b\x4d\x7d\x6b\x8a\xb0\x8d\xb7\xab\xec\x7b\x6c\x76\x1b\x83\x49\xd8\xf8\x2b\x9e\xe6\x08\x9e\xb0\x84\xef\x69\xd5\xc8\xfc\x19\xf4\x7b\xc1\x52\x45\xcf\xb2\x37\x84\xde\xd7\x28\x0c\x65\xd2\x1b\x98\xbb\x0a\x6b\xe4\x20\x18\x71\x94\x7c\x33\xa0\x4f\x00\x8c\x28\x34\xa0\x7e\x7c\x9f\xda\xcd\x8b\xce\x8b\xf9\x8e\x03\x73\x4e\x67\x16\x7f\x82\xff\xf1\x09\x6c\x9c\x58\xbe\x23\x40\x07\xb8\x58\xc9\x44\xea\xe6\x16\xcf\x77\xfd\xc1\x50\x93\xe5\x99\x4c\x8d\x73\x64\x05\x2a\x09\xd6\xb2\x46\x7d\xb6\xf8\xf8\xde\xe2\xfd\xbf\x6e\x6e\x07\xfb\xa8\x4c\x05\x2a\xcf\x4a\x5c\xa5\xf1\xa0\xb1\xf6\x53\x37\xd6\xa8\xd7\x14\xcb\x54\x56\xe1\x19\x20\xe8\x0c\x3d\x61\xcd\xe6\xaf\x52\x24\x9b\x97\x5a\x4e\x13\xfe\x61\xfd\x61\x6d\xbb\x5e\xff\x91\xb4\x3d\xaf\xe6\xb2\x70\x7d\x51\x71\x1a\x04\x60\x84\x13\x97\x23\xf8\x73\xdb\x12\x7a\x62\x5c\x56\x25\x43\x00\x26\x2e\x07\xb4\xc6\xf7\xc0\xdd\x07\xc0\x47\x1c\xfb\x02\x61\xce\xdd\x31\xda\x3e\xc7\x8e\x45\x38\xca\x38\x88\xa3\x20\xeb\x08\xd9\x43\x4e\xad\xbd\x54\x0b\xad\x45\x73\x15\xee\x7c\x74\xd9\x36\x58\xea\xfc\xea\xe8\x6e\x6e\xff\x1e\xbc\x12\xca\x2a\xde\x7c\x09\xe2\x32\x94\x3c\x4d\x9a\x66\x4e\xa7\xff\xde\xe8\x42\x19\xcb\xdf\x8b\x4e\x97\xfb\x8c\x7e\xf2\x11\xfa\x2d\xbd\xc3\xb2\xf3\x01\x39\x76\x5a\x05\xe9\x4a\xaa\x0f\xb0\x7a\xdf\x82\x3c\x8d\x8e\x8c\xba\xf9\xf3\xf6\xee\xc0\xa9\x34\x8f\x9b\x6d\xfa\xcf\x73\xd9\x68\xc7\x56\x4b\x3d\x94\x77\xda\xdd\x9d\xbd\xbf\xb0\x0d\xa7\xff\x51\xe3\x63\x2d\x79\x58\xca\x1a\x74\xee\x47\xe9\xe3\x6b\x1b\x52\xb9\xfd\x21\x3b\x12\x6f\x56\x1d\xb1\x8d\xe8\xf4\x7d\xc9\x15\xab\x57\xe1\x28\x13\xa8\xd5\x75\xc0\x2a\xbd\x05\xee\x0c\xec\x8d\x73\xb0\xdc\x84\xcd\x1c\x9c\x61\x5b\x6e\x12\x25\x7f\x28\xa3\x63\x5c\xde\xb4\x7f\x27\x26\xa4\x08\xb4\xfd\xc1\xb2\x37\x2f\x09\x21\x36\x77\xe7\x6d\xb3\xcd\xfa\xe8\x54\xe6\xad\x84\xcd\xfd\xf2\xee\xd3\xf2\x96\x3b\xa2\xdd\xbf\x6a\x9b\x8b\xf5\xa0\x78\xff\x6c\xf7\x31\xad\x8f\xce\xfc\xee\x30\xc9\xcf\x00\x00\x00\xff\xff\xb4\xb2\x39\x83\xf4\x08\x00\x00")

func staticMigrations1_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_staticMigrations1_initialSql,
		"static/migrations/1_initial.sql",
	)
}

func staticMigrations1_initialSql() (*asset, error) {
	bytes, err := staticMigrations1_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/migrations/1_initial.sql", size: 2292, mode: os.FileMode(420), modTime: time.Unix(1527161229, 0)}
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
	"static/migrations/1_initial.sql": staticMigrations1_initialSql,
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
	"static": &bintree{nil, map[string]*bintree{
		"migrations": &bintree{nil, map[string]*bintree{
			"1_initial.sql": &bintree{staticMigrations1_initialSql, map[string]*bintree{}},
		}},
	}},
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

