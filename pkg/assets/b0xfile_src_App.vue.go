// Code generaTed by fileb0x at "2018-07-05 09:59:02.571685404 +0200 CEST m=+0.033506970" from config file "b0x.yml" DO NOT EDIT.
// modified(2018-07-05 08:51:22.286431294 +0200 CEST)
// original path: ../../public/dist/src/App.vue

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileSrcAppVue is "src/App.vue"
var FileSrcAppVue = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\xc1\x8e\xdb\x30\x0c\x44\xef\xfa\x0a\xc2\x7b\xf0\x25\xb6\x17\x2d\xda\x83\xe3\x04\xc8\xad\x5f\xd0\x45\x8f\xac\x4c\xab\x6c\x25\x51\x10\x19\x37\x69\xb1\xff\x5e\xc4\x8b\x45\x02\xf4\xca\x79\x33\xc3\x99\x8c\x52\x89\x68\x74\x74\x00\xd3\xcc\x2b\xf0\x7c\x68\xb0\x94\xe6\x76\x00\x98\x38\x05\xd0\xea\x0f\x4d\x3f\xa0\x2a\x99\x0e\x51\x82\xf4\x25\x87\x77\xe2\x0b\xc5\x28\x2f\x52\xe3\x0c\x49\xc3\xa1\x79\xa1\xe8\x25\x11\x98\xc0\x37\x39\x57\xf8\x7a\xa6\xfe\xa7\xc2\xa9\x94\x66\xd8\x5a\x86\x99\xd7\xa3\x9b\x86\x7b\xb5\x9b\xd4\x57\x2e\x76\x74\x9c\x8a\x54\x83\x87\xcc\xa5\x4a\x82\xb6\x1f\xbc\xa4\x22\x99\xb2\xe9\x70\x57\xfb\xf5\x4c\xad\x73\x74\xd9\x5c\x33\x2d\x78\x8e\x06\x7f\x1d\x40\xc6\x44\x23\xb4\x58\x4a\xbb\x73\x00\x77\xf7\xb8\xc9\xf0\x50\xe1\x00\x5e\xdd\xab\x9b\x86\xf7\x27\xdc\xa4\x76\x8d\x74\x74\x4f\x58\xca\x86\x2f\x92\xad\x5b\x30\x71\xbc\x8e\xd0\x9e\x56\xca\x5c\xdb\xdd\x2d\x63\x25\x63\x8f\x3b\x38\x55\xc6\xb8\x03\xc5\xac\x9d\x52\xe5\x65\xef\x00\xba\xdf\xf4\xfd\x17\x5b\xb7\xd9\x35\x89\xd8\x0f\xce\x61\x04\xcc\xc6\x18\x19\x95\xe6\x0d\x4b\xf2\xa7\x13\xbd\xfc\xc7\x85\x8a\x57\xf5\x18\xe9\x46\x19\x5d\xac\xc3\xc8\x21\x8f\xe0\x29\x1b\xd5\xfd\x36\x2c\x4a\x1d\xe1\xe9\x83\xff\x48\x9f\x9e\x6f\x97\x84\x35\x70\xee\x4c\xca\x08\x9f\x9f\xcb\x65\xff\xb6\xed\x6d\xd1\xbf\x00\x00\x00\xff\xff\xff\xe8\x7c\x47\xf0\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileSrcAppVue)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "src/App.vue", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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