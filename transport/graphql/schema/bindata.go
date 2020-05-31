package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _schema_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xb1\x8a\xc3\x30\x0c\x86\x77\x3d\x85\xbc\xdd\xc1\x3d\x81\xb7\x3b\x6e\xc9\x70\xc3\xd1\x74\x2a\x1d\x4c\x24\x5a\x43\x62\xb9\x89\x3d\x84\x92\x77\x2f\x76\x9c\xb4\x69\xa1\x93\xad\xff\xd7\xf7\x4b\x68\x68\xce\xdc\x19\xbc\x02\x22\xe2\x25\x72\x3f\x6a\xfc\x4f\x4f\x16\xba\x18\x4c\xb0\xe2\x34\xfe\x95\x1f\x4c\x00\x10\x46\xcf\x73\x57\x01\x83\x90\x0c\x1a\x0f\xb5\x90\xa8\xa3\x5a\xb5\x0f\x4b\x1a\xab\x5f\xf5\xa9\x31\x59\x49\x4e\x01\x99\x5f\x12\x4b\x84\x21\xaa\x33\xe1\x7c\x0c\x1a\xbf\xe7\xb2\x4a\xd5\x82\xe7\xbe\xe8\xc9\x04\xae\x1f\xc2\xbf\xb0\x30\xfb\xd5\x7a\xc5\x88\x5b\xde\x62\x8b\x3b\x01\x64\x7e\x33\xb2\x2c\x45\x3c\x34\xbd\xf5\xf3\x09\x76\xa1\xb7\xee\xa4\xee\xc0\xd3\xbc\x77\x4c\x32\x1a\xe9\x7c\x5a\x82\x34\xfe\x88\xb4\x6c\x9c\x82\xe9\x16\x00\x00\xff\xff\x44\xdd\x23\x03\x80\x01\x00\x00")

func schema_graphql() ([]byte, error) {
	return bindata_read(
		_schema_graphql,
		"schema.graphql",
	)
}

var _type_scalar_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\x4e\xcc\x49\x2c\x52\x08\xc9\xcc\x4d\x05\x04\x00\x00\xff\xff\x86\x36\xc3\xbe\x0b\x00\x00\x00")

func type_scalar_graphql() ([]byte, error) {
	return bindata_read(
		_type_scalar_graphql,
		"type/scalar.graphql",
	)
}

var _type_todo_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\xc9\x4f\xc9\x57\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xb3\x53\x52\x8b\x93\x8b\x32\x0b\x4a\x32\xf3\xf3\xac\x14\x82\x4b\x8a\x32\xf3\xd2\x21\x12\xc9\xf9\xb9\x05\x39\xa9\x25\xa9\x29\x56\x0a\x4e\xf9\xf9\x39\xa9\x89\x79\x8a\x5c\xb5\x5c\x5c\x99\x79\x05\xa5\x25\x0a\x8e\x29\x29\x48\xe6\x61\x35\x03\xae\x36\xb4\x20\x25\xb1\x24\x95\x90\x72\x9c\x56\x02\x02\x00\x00\xff\xff\x03\x38\x7a\x5b\xbd\x00\x00\x00")

func type_todo_graphql() ([]byte, error) {
	return bindata_read(
		_type_todo_graphql,
		"type/todo.graphql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"schema.graphql": schema_graphql,
	"type/scalar.graphql": type_scalar_graphql,
	"type/todo.graphql": type_todo_graphql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"schema.graphql": &_bintree_t{schema_graphql, map[string]*_bintree_t{
	}},
	"type": &_bintree_t{nil, map[string]*_bintree_t{
		"scalar.graphql": &_bintree_t{type_scalar_graphql, map[string]*_bintree_t{
		}},
		"todo.graphql": &_bintree_t{type_todo_graphql, map[string]*_bintree_t{
		}},
	}},
}}
