// Code generaTed by fileb0x at "2018-07-05 09:59:02.580145547 +0200 CEST m=+0.041967111" from config file "b0x.yml" DO NOT EDIT.
// modified(2018-07-05 08:51:22.286431294 +0200 CEST)
// original path: ../../public/dist/index.html

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileIndexHTML is "index.html"
var FileIndexHTML = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4f\x4b\x03\x31\x10\xc5\xbf\x4a\xec\xd9\xdd\xb8\x85\x6d\x2b\x24\x81\x52\x7b\xf0\xa4\x07\x15\x3d\xa6\xc9\xd4\x4c\x9b\xcd\xc6\x64\x76\x4b\xbf\xbd\xd8\x5d\x54\x14\xff\x5c\x02\x2f\xbc\x1f\xf3\xde\x30\xe2\xec\xea\x66\x75\xf7\x74\xbb\x66\x8e\x1a\xaf\xc4\xf8\x82\xb6\x4a\x34\x40\x9a\x19\xa7\x53\x06\x92\x1d\x6d\x8b\xc5\xf8\xe7\x88\x62\x01\x2f\x1d\xf6\xf2\xb1\xb8\x5f\x16\xab\xb6\x89\x9a\x70\xe3\x81\x99\x36\x10\x04\x92\x93\xeb\xb5\x04\xfb\x0c\x93\x11\x09\xba\x01\xd9\x23\x1c\x62\x9b\xe8\xc3\x75\x40\x4b\x4e\x5a\xe8\xd1\x40\x71\x12\xe7\x18\x90\x50\xfb\x22\x1b\xed\x41\x56\x13\x25\x08\xc9\x83\x7a\xe8\x80\x2d\x63\x14\x7c\x90\xc2\x63\xd8\x33\x9d\x65\xa6\xa3\x07\xe6\x12\x6c\x25\x37\x39\x73\x1d\x63\x69\xab\xd9\xdc\x9a\x59\x5d\x9a\x9c\x59\x02\x2f\x63\x02\xdf\xbe\x75\x7a\xa7\x4c\xc2\x48\x23\xb6\x1b\xa8\x6a\x5e\x5f\x2c\xec\xf4\xb2\xdc\xfd\x13\x32\xae\x0b\xfb\xa2\x87\x60\xdb\x94\xcb\xba\x9a\xc2\xac\x9e\xff\x80\xff\x11\xf0\xd4\x22\x3b\x00\x52\x82\x0f\xeb\xdf\xb4\xf6\xa8\x84\xc5\x9e\xa1\x95\x3a\x46\x25\xb8\xc5\x5e\x89\x31\x45\x4e\xe6\xf7\x10\x4a\xf0\xc1\xfa\x0d\xf9\x52\xf6\x93\x91\x0f\x43\xf9\xe9\x0c\x5e\x03\x00\x00\xff\xff\xa4\x3c\xee\x6e\x1c\x02\x00\x00")

func init() {

	rb := bytes.NewReader(FileIndexHTML)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
