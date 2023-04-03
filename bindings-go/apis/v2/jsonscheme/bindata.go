// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../../language-independent/component-descriptor-v2-schema.yaml

package jsonscheme


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataLanguageindependentComponentdescriptorv2schemaYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1a\x6d\x6f\xdb\xb8\xf9\xbb\x7f\xc5\x83\x4b\x00\x26\x4d\x64\xa7\xde\x3a\xa0\xfe\x12\x64\x77\xd8\x70\xd8\x76\x19\xda\xdb\x3e\x2c\xf5\x0e\xb4\xf4\xd8\x66\x47\x91\x1e\x49\xb9\xd1\xb5\xfd\xef\x03\x49\x51\x6f\x96\x64\x3b\xee\xcb\x0e\x68\xbe\x44\xa4\x9e\xf7\x77\x52\x3e\x67\xc9\x0c\xc8\xda\x98\x8d\x9e\x4d\x26\x2b\xaa\x12\x14\xa8\xc6\x31\x97\x59\x32\xd1\xf1\x1a\x53\xaa\x27\xb1\x4c\x37\x52\xa0\x30\x51\x82\x3a\x56\x6c\x63\xa4\x8a\xb6\x53\x32\x3a\xf7\x10\x35\x0a\x6f\xb5\x14\x91\xdf\x1d\x4b\xb5\x9a\x24\x8a\x2e\xcd\x64\x7a\x33\xbd\x89\x9e\x4f\x0b\x82\x64\x14\xc8\x30\x29\x66\x40\xfe\x5c\x70\x85\xef\x03\x1f\xf8\xa1\xe4\x03\xdb\x29\x54\x68\x4b\x26\x98\xc5\xd2\xb3\x11\x40\x8a\x86\xda\xff\x00\x26\xdf\xe0\x0c\x88\x5c\xbc\xc5\xd8\x10\xb7\xd5\x64\x51\x6a\x00\x95\x06\x0e\x3f\xa1\x86\x7a\x04\x85\xff\xcd\x98\xc2\xc4\x53\x04\x88\x80\x78\xbe\xff\x44\xa5\x99\x14\x1e\x6a\xa3\xe4\x06\x95\x61\xa8\x03\x5c\x03\x28\x6c\x96\x22\x69\xa3\x98\x58\x91\xd1\x08\x80\xd3\x05\xf2\x5e\x79\x3b\xd8\x0b\x9a\x22\xa9\x96\x5b\xca\x33\x74\x94\x4a\x6d\x7e\xa2\x29\x36\x28\x06\x76\x76\x2b\xa5\x8f\x7f\x45\xb1\x32\xeb\x19\x4c\x5f\xbc\xf0\xd2\x53\x63\x50\x59\x83\xfc\xfb\x81\x46\xbf\xde\x44\x2f\xc7\x6f\xa2\xf9\xd5\xc3\x78\x6e\x97\xf3\xf7\xd3\xeb\xdf\x7f\x9c\x3c\x44\xfe\xd5\xe4\x97\xf1\xfc\xd9\xb9\x63\xc8\x12\x14\x86\x99\xfc\xce\x18\xc5\x16\x99\xc1\xbf\x60\xee\xf9\xa6\x4c\x94\x4c\x7a\x58\xcc\x2f\x1e\xa2\x5f\xae\x8a\xe7\x67\x61\xf3\xf2\xd6\x93\x56\xc8\xe9\x23\x26\xaf\x31\xdd\xa2\xf2\x34\xcf\xc0\xd0\xff\xa0\x80\xa5\x92\x29\x68\xf7\xc2\xc6\x12\x50\x91\x00\x4d\xde\x66\xda\x60\x02\x46\x02\xe5\x5c\xbe\x03\x2a\x40\x3a\x37\x53\x0e\x1c\x69\xc2\xc4\x0a\xc8\x96\x5c\x43\x4a\xdf\x4a\x15\x49\xc1\xf3\x6b\x87\xea\xd6\xe3\x94\x89\x62\x37\xf0\x5a\x33\x0d\x29\x52\xa1\xc1\xac\x11\x96\xd2\x52\xb5\x44\xbc\x2d\x35\x50\x85\x96\x15\x6c\x29\x67\x49\x53\x5e\x1d\x04\x7e\x3e\x9e\x8e\x7f\x57\x7f\x8e\x96\x52\x5e\x2d\xa8\x2a\xf6\xb6\x75\x80\x6d\x17\xc4\xf3\xf1\x34\x3c\x95\x60\x35\xf8\xf2\xb1\x81\x56\x37\xf6\x76\x7e\x7b\x71\xf3\xe1\xe1\x79\xf4\x72\xfe\x26\x79\x76\x79\x71\x3b\x7b\x33\xae\x6f\x5c\xde\x76\x6f\x45\x17\x17\xb7\xb3\x6a\xf3\xc3\x9b\xc4\xf9\xe8\x2e\xfa\x57\x34\x7f\xb8\x89\x5e\x86\xe7\x40\xf2\x40\xe0\xcb\xc0\xf1\xea\xa2\xfe\xe2\xca\x11\x69\xec\x38\xc8\x73\xd2\x15\xc6\x5d\xa1\xd7\x9b\x41\x45\x6a\xe6\x36\x29\xf4\x0c\xde\xc3\xb9\xc2\xe5\x0c\xc8\xd9\xa4\x56\x37\x26\x5d\xa1\x4c\xe0\xa3\x0f\xc5\x8d\xd4\xcc\x48\x95\x7f\x2f\x85\xc1\x47\x73\x4c\xb2\x5a\xa8\xbe\x12\xe1\x28\x0c\x54\x06\x19\xb3\x57\xdd\xbc\x29\xe7\xf7\xcb\x8a\x4b\xa7\x46\x3b\x62\x57\x35\xa3\x2d\xa7\x93\x74\x41\x35\xfe\x43\x71\x52\xee\xed\x0a\x6c\xff\x0a\xb0\xfa\x56\x67\x99\xf1\x7f\x8d\x92\xf4\x37\xba\xd9\x30\xb1\x3a\x10\x15\x00\x45\x96\xce\xe0\x81\x64\x8a\xff\x9d\x9a\x35\xb9\x06\xa2\xd7\x74\xfa\xe2\x0f\x51\xc2\x56\xa8\x0d\x99\x8f\x5a\x74\x8e\xa5\xec\x2c\xbc\x62\xda\xa8\x9c\xcc\xad\xc9\x69\x1c\xa3\xd6\x07\x76\x0f\x6b\x0a\x07\x05\x4b\xa9\x0a\x54\xd4\x70\x61\x57\xf8\x68\x50\xd8\xd2\xaf\x2f\xf7\xc4\xc6\x08\x60\xc5\xcc\x3a\x5b\xdc\x0d\xf3\x1e\x0c\x2e\xb7\xb4\x1e\xaf\x79\xd0\xed\x2c\x9f\x14\x7c\x6d\x3b\x79\x01\x4b\x7b\x17\x8c\xf6\xa0\xdb\xa0\x1c\x86\x88\x65\x9a\x32\x33\x94\x02\x42\x0a\x3c\xc5\x2e\x27\xea\xfd\x93\x14\xe8\x03\x43\xcb\x4c\xc5\xf8\x43\x99\x5f\x47\x88\x63\xfb\x75\xb9\xd8\xfa\x81\xa0\x5c\x5b\x0a\xe5\xc2\x87\x50\x8f\xe0\xa2\x6c\xea\x03\x82\x1f\x5e\xdb\x0a\x14\x7c\x34\x8a\xfe\x58\x00\xcc\x8e\xa4\x13\x88\x6c\xdb\x53\x4e\x4f\x41\xaa\xb5\x48\xf2\xc4\x30\x2c\x63\xd0\x8d\x4d\x7a\x07\x95\x2a\x45\xf3\x0a\x93\x19\x4c\x1b\xe5\xab\x53\x32\x47\x2b\x20\xd5\x4b\x80\x5b\x8b\xfc\x7e\x59\x27\xd1\x53\x6f\x3d\x1e\xd9\x0f\x58\xcf\xf6\x03\xc0\xed\x08\x1d\x80\x47\x00\xbe\xf4\xbd\xde\x60\x7c\x44\x08\xae\xa9\x5e\xdf\xf1\x95\x54\xcc\xac\xd3\x2a\x30\xa5\x4a\x29\x67\x9a\x5a\x46\xbb\xaf\xdd\x60\xd9\x13\x8c\x0d\x82\x6d\x27\x78\xf7\x85\xb0\xed\x64\x32\x88\xe2\x18\xf7\x40\xd8\x54\x64\x2b\x41\x4d\xa6\xf0\x48\x23\xd0\x01\x0d\xed\x2a\xc5\x84\xd1\x9f\x43\x3e\xee\xea\x4c\x4f\x16\xde\x6f\x95\x7c\x2a\xa8\x66\x5f\xf9\x79\x8d\x1e\xc8\x37\x17\xb9\x74\x13\x68\xa9\x36\x14\x13\xff\x5e\xfb\x3c\xb5\x46\xf9\x10\x2b\x97\x25\xbd\x23\x0a\x53\x43\x61\x4f\x6f\x4f\x75\xa8\xe2\x3a\x68\xd6\xd2\xa3\x17\xb3\x11\x0f\x2e\x47\xb4\x8a\x5f\x85\xe6\xb3\xb7\x8b\x53\xdb\xa8\x50\xa1\x88\xd1\x9d\x1e\xe0\xa2\x3a\xd7\x72\x19\x53\x7e\x59\x14\xff\xbe\x8e\x12\xca\xe2\x6b\xe4\x18\x1b\xa9\x9e\x5a\x45\x3f\x43\x45\xab\x1f\x0a\x5f\x05\x2d\x9f\x6a\x97\x92\xd2\xa1\x27\xd3\xc6\xf0\x57\x3f\xb1\x0e\x9f\x9c\x3b\x8e\xb1\xbd\x7a\x76\xb2\x18\xea\x94\x70\x06\x34\x36\x19\xe5\x3c\x9f\x55\x9c\x22\x97\x68\xef\x26\xa0\x37\x18\x33\xca\x41\xa1\x85\x8f\x1d\x93\xdf\x6e\x73\xfd\x6c\x3d\xb2\x9d\xd1\x52\x60\xbd\x47\x46\x81\x93\xc8\x78\xed\x40\xd1\xd3\xe0\xea\x99\xef\x8e\x5b\x3e\xdd\xaa\x0a\x79\xe4\x20\x1e\x08\xe8\x83\x6f\x50\x8a\x78\x84\x33\x87\xef\x92\xbe\xa2\x72\x5d\x5c\x05\x64\xda\x40\x4a\x4d\xbc\xae\x25\x82\xde\x99\xe7\x76\x67\x72\xee\x3a\x5f\x6d\xab\x3e\x28\x7c\x1b\xf3\x4a\xad\x7c\xd1\xfe\x44\xd1\xea\x89\x55\x27\x11\xef\x84\x83\x07\x4d\x17\x02\xf6\xbc\x69\x8f\x71\x4a\x50\xfe\xd5\xc7\xce\x03\x87\xce\x1e\x30\x19\xb3\x3f\x72\xb9\x33\x73\xf6\x40\x3b\xed\xff\xc4\x38\xea\x5c\x1b\x4c\x8f\xc5\xbc\xef\x62\xf6\x39\x2b\x86\x8c\xd9\x8f\x29\x5d\x9d\x74\x50\x74\x4b\x66\xa9\x94\x7d\xf2\x93\x9c\x20\x9b\x37\x0c\x45\x74\x34\xd8\xec\xb9\x01\xaa\x4c\x79\xa0\x62\x0d\xb5\x22\x20\x9c\xe6\x21\x0f\x4f\xd3\x05\x48\x21\x0e\x81\xea\x22\x60\xd9\x37\xc4\xde\x59\xe1\x9b\x23\x84\x9d\x62\x53\x2a\xd8\x12\xb5\x69\x8f\xaf\x2d\xa6\x4f\x9c\x91\xbd\x55\x7c\xc1\xf6\xa9\xe1\x25\xd0\x60\xe4\x1e\x8e\xed\x00\xdd\x65\xe7\x21\x02\x2b\x43\xd5\x0a\x0d\x26\x10\x4b\x61\xca\xa1\xa8\x97\xbc\x66\xbf\x0e\xea\x62\xdf\x03\x13\xb0\xc8\x0d\xea\xc0\x63\x61\x8d\xdd\xa6\x2b\xb2\x74\x61\x1d\x3a\x02\xe8\x4d\xd4\x13\x72\x60\xc9\x38\x56\xfd\xf1\xd4\x88\xe9\x90\xb0\x8a\x9e\xc0\xaa\xcf\x2e\xe1\x7d\xdd\x1c\x60\xd6\xd4\x00\xd3\x4e\x77\x6b\x7e\x26\xdc\xbb\xef\xec\x4b\xfd\x1d\x24\x4c\xb9\x21\x3c\xef\xf5\x47\xb0\xdb\xfd\x13\x72\xeb\x0b\x19\xec\xbe\x9d\x67\xc3\xc1\xd9\x0c\x4c\x97\xef\xf0\x8e\x99\x75\x61\x9a\x38\x53\x0a\x85\x81\xae\x8f\x5d\x43\x56\x0a\x65\xf5\x55\x31\x09\x9d\xf2\x8d\xaa\x3e\xf1\x77\x19\xf1\xdb\x4c\xb4\xbf\x8f\x38\x67\x7c\xf9\x41\xa4\x6f\xa0\xa8\xb5\xdc\x2f\xd1\xe4\xab\x6b\xb0\x13\x72\x35\x0b\xb7\xe3\x27\x76\x75\x2b\x4c\xe9\x89\x6c\xe0\x26\x7c\x04\xb0\x42\x81\x8a\xc5\x5f\xf1\x16\xbb\x90\xc0\x5f\x64\x17\x8b\x6f\x49\xfd\x7f\x90\xd4\x95\x63\xfc\xfe\xd7\xcd\xe9\x46\xa0\x7e\x89\x94\x2e\x1b\xd2\xc1\x37\x52\x47\x5f\x41\xed\xc6\xe8\xce\x67\x51\x5d\x7b\xb9\x51\x72\xcb\x92\xca\x9b\x11\x90\xc6\x5d\x42\xf3\x5a\xab\x1c\xe1\x75\x83\x7e\x03\x63\x5f\xdc\x1f\x7e\xab\x75\x42\x50\xee\xea\x7c\x74\x8c\xed\x7c\x05\x19\x3a\x6b\xee\x7c\xb5\x26\x70\x16\xc6\x10\x9e\x5f\xc3\x3b\x04\x29\x78\x5e\xfc\x52\xc3\x4d\xeb\x52\x84\xfb\xe7\xe0\x83\x3d\x59\xf4\xd9\x72\xa5\x70\xdf\x27\xba\x87\x68\x7d\x36\x0c\xf8\x1d\x31\xf4\x69\x18\xee\x12\xae\x82\xe0\xa9\x9a\x1d\xee\xfb\xfa\xdd\x1d\x39\x30\x58\x1a\x33\xe6\x41\x48\xad\x16\xe6\x6a\x49\xb7\x49\xe1\xfd\xc7\xd1\x68\xd4\x2a\x2c\xf5\xaa\x11\x01\x49\xd1\xff\xd4\xab\x9e\xd9\x64\xd4\xcc\xdb\xea\x27\x65\x9d\x02\x05\x12\xad\x82\x36\xec\x20\x52\xff\x54\xd3\x1c\x0c\x6a\x0e\x69\x38\x63\xf8\xf3\x07\x19\xfd\x2f\x00\x00\xff\xff\x7e\x13\x49\xc0\xb6\x27\x00\x00")

func bindataLanguageindependentComponentdescriptorv2schemaYamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataLanguageindependentComponentdescriptorv2schemaYaml,
		"../../../../language-independent/component-descriptor-v2-schema.yaml",
	)
}



func bindataLanguageindependentComponentdescriptorv2schemaYaml() (*asset, error) {
	bytes, err := bindataLanguageindependentComponentdescriptorv2schemaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "../../../../language-independent/component-descriptor-v2-schema.yaml",
		size: 10166,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1678808714, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"../../../../language-independent/component-descriptor-v2-schema.yaml": bindataLanguageindependentComponentdescriptorv2schemaYaml,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"..": {Func: nil, Children: map[string]*bintree{
		"..": {Func: nil, Children: map[string]*bintree{
			"..": {Func: nil, Children: map[string]*bintree{
				"..": {Func: nil, Children: map[string]*bintree{
					"language-independent": {Func: nil, Children: map[string]*bintree{
						"component-descriptor-v2-schema.yaml": {Func: bindataLanguageindependentComponentdescriptorv2schemaYaml, Children: map[string]*bintree{}},
					}},
				}},
			}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
