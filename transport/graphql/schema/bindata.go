// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema.graphql
// type/scalar.graphql
// type/todo.graphql
// type/user.graphql
package schema

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x4e\xbd\x8a\xc3\x30\x0c\xde\xfd\x14\x5f\xb6\x3b\xb8\x27\xf0\x7c\x4b\x86\x0e\x85\x64\x2a\x1d\x4c\x24\x68\xa0\x89\x53\x5b\x1e\x42\xe9\xbb\x17\xd9\x4e\x48\x97\x4e\x12\xdf\x7f\x1c\x6e\x3c\x39\x3c\x0d\x00\x3c\x12\x87\xd5\xe2\xac\x27\x03\x53\x12\x27\xa3\x9f\x2d\x4e\xf5\x33\x2f\x63\x8c\xac\x0b\x17\x55\x35\x8a\x27\x1f\x2d\x2e\x9d\x27\xdf\x5c\x9b\x1d\xfb\x19\xc9\xa2\xfd\x6f\x7e\x2d\x94\x02\x32\x93\x22\x07\x55\xf7\x91\xc3\xa6\x56\xec\xa0\x56\x4a\xab\x72\xd3\xd6\x5d\xcb\x1c\x51\x97\xb3\xe7\x25\x49\x09\x6e\xf5\xdd\x5a\x4a\xde\x42\x4e\xb8\x3b\x6c\xf8\xc3\x57\x03\xf1\x9d\x3f\x0d\x47\x76\x08\xec\x84\xfb\x3c\xb2\xa4\xe8\xbf\xa7\xd4\xb9\xef\x00\x00\x00\xff\xff\xcc\x74\x00\x2d\x4d\x01\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 333, mode: os.FileMode(436), modTime: time.Unix(1600126092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeScalarGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\x4e\xcc\x49\x2c\x52\x08\xc9\xcc\x4d\x05\x04\x00\x00\xff\xff\x86\x36\xc3\xbe\x0b\x00\x00\x00")

func typeScalarGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeScalarGraphql,
		"type/scalar.graphql",
	)
}

func typeScalarGraphql() (*asset, error) {
	bytes, err := typeScalarGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/scalar.graphql", size: 11, mode: os.FileMode(436), modTime: time.Unix(1590875830, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeTodoGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8e\x31\xaa\xc3\x40\x0c\x44\xfb\x3d\xc5\xfc\x6b\x6c\xf7\x83\x9b\xad\xe3\x0b\x18\x6b\x08\x02\x7b\x25\xd6\x72\x11\x8c\xef\x1e\xd8\x18\x42\xea\x74\x4f\x4f\x30\x33\xf1\x74\x62\x34\x31\x1c\x09\x00\x54\x32\xca\xf0\xd7\x79\xdf\xd8\xca\xf0\xb9\x85\xdb\xdc\xd4\x43\xad\x66\xdc\xa3\x69\x7d\xbc\x1f\xb3\xad\xbe\x30\x28\x19\x37\xb3\x85\x53\xbd\x7c\xe3\x14\x94\xff\xc8\x18\x75\xe5\x95\xea\xf2\x2d\xcf\x94\xb4\xfa\x1e\x7d\x46\xe9\x74\xfc\xde\x7f\xbe\x02\x00\x00\xff\xff\xc6\x71\x13\x04\xda\x00\x00\x00")

func typeTodoGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeTodoGraphql,
		"type/todo.graphql",
	)
}

func typeTodoGraphql() (*asset, error) {
	bytes, err := typeTodoGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/todo.graphql", size: 218, mode: os.FileMode(436), modTime: time.Unix(1600177957, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2d\x4e\x2d\x52\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xb3\xd3\x32\x8b\x8a\x4b\xfc\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\x21\xc2\x39\x89\xd8\x44\x53\x73\x13\x33\x73\x50\x85\x94\x15\x4a\xf2\x53\xf2\x8b\xad\x14\xa2\x43\xf2\x53\xf2\x15\x63\xb9\x6a\xb9\xb8\x32\xf3\x0a\x4a\x4b\xc0\x16\x7a\x82\x59\xd5\x14\xdb\x54\x0b\x08\x00\x00\xff\xff\x12\x52\xa0\x71\xc3\x00\x00\x00")

func typeUserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeUserGraphql,
		"type/user.graphql",
	)
}

func typeUserGraphql() (*asset, error) {
	bytes, err := typeUserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/user.graphql", size: 195, mode: os.FileMode(436), modTime: time.Unix(1600177968, 0)}
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
	"schema.graphql":      schemaGraphql,
	"type/scalar.graphql": typeScalarGraphql,
	"type/todo.graphql":   typeTodoGraphql,
	"type/user.graphql":   typeUserGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
	"type": &bintree{nil, map[string]*bintree{
		"scalar.graphql": &bintree{typeScalarGraphql, map[string]*bintree{}},
		"todo.graphql":   &bintree{typeTodoGraphql, map[string]*bintree{}},
		"user.graphql":   &bintree{typeUserGraphql, map[string]*bintree{}},
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
