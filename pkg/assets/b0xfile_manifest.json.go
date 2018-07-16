// Code generaTed by fileb0x at "2018-07-16 13:32:51.501912411 +0200 CEST m=+0.640743318" from config file "b0x.yml" DO NOT EDIT.
// modified(2018-07-16 13:32:38.837376718 +0200 CEST)
// original path: ../../public/dist/manifest.json

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileManifestJSON is "manifest.json"
var FileManifestJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x6a\xc5\x30\x0c\x06\xe0\xdd\xa7\x10\x9a\xcb\x7b\xbc\x6c\xc9\xd6\x1b\x74\x2f\x25\x08\xdb\x38\x02\x47\x36\xb6\x0a\x49\x43\xee\x5e\x6c\x0f\x5d\xba\x68\xf9\xf4\x4b\xff\x65\x00\xb0\x6e\xa9\xe8\x2a\xb4\x7b\x5c\x00\x3f\x28\xb0\x17\x85\xf7\x9c\xf1\xad\xf1\x3f\x10\xd9\x92\x72\x92\xb1\xc0\x36\x49\xc5\x05\x3e\x0d\x00\xc0\xd5\x67\x3b\x5b\x6c\x8b\x3d\x63\x0a\xe9\x91\x25\xf4\xe5\x21\xfc\xe3\x5b\x00\x5f\xf3\x74\xbc\xe6\xe9\x4f\xf4\xcc\xfd\x17\xef\x14\xfc\xb3\x85\xba\xdc\x06\xe0\xab\xff\xaa\x4a\x45\xd7\xef\x12\xfb\x69\x16\xe7\x8f\xc7\xa6\x7b\x1c\x4d\x1c\xd7\x1c\xe9\x6c\x56\x95\xc4\x51\x4c\xe2\x07\xa5\xd2\xca\x8f\xd6\x0b\x60\x24\x71\xd5\x52\xf6\x68\x6e\xf3\x1b\x00\x00\xff\xff\xd5\x88\x8f\xba\x06\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileManifestJSON)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "manifest.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
