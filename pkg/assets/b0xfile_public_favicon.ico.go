// Code generaTed by fileb0x at "2018-07-05 09:59:02.579552966 +0200 CEST m=+0.041374539" from config file "b0x.yml" DO NOT EDIT.
// modified(2018-07-05 08:51:22.286431294 +0200 CEST)
// original path: ../../public/dist/public/favicon.ico

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FilePublicFaviconIco is "public/favicon.ico"
var FilePublicFaviconIco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\x4d\x68\x13\x51\x10\xc7\x27\x24\x10\xc1\x8f\x46\x11\xc1\x83\x9a\x8b\xb1\x67\x61\xd2\xda\x8f\x24\x83\x87\x74\xb7\xb5\x58\xf0\x20\x94\xbc\xa4\x2a\xa5\x48\xb0\x88\x5f\x54\x4b\x13\x3b\xbd\x89\x8a\xe0\xa1\x20\x08\x22\x82\x28\x2a\x42\xd1\x49\x0f\x7e\x5c\x44\x3c\x78\xf7\x26\x78\x08\x56\x14\x51\x41\x85\x90\x91\x97\xdd\x0d\x8b\x28\x29\x1e\x1c\xf8\xb3\x6f\xe7\xbd\xdf\xcc\xff\x7d\x00\x44\x20\x02\x89\x84\xfd\x26\x61\x2a\x06\xb0\x05\x00\xba\x01\x20\x01\x00\x49\xf0\xf2\xad\x88\xc1\x7f\x09\x16\x0a\xb4\x93\x85\x92\x76\xfc\x0f\xfc\x1a\x16\xba\xc5\x42\x8b\x2c\x14\x5b\x6d\x8d\x50\xef\xfd\x2c\xf4\x8d\x85\x3e\xb1\xd0\xde\x20\xbf\x4a\x7e\x33\x0b\x3d\x67\x21\xf5\xb5\xc4\x42\xeb\x3b\xf1\xa1\xde\xd3\x2c\xd4\xe0\x1a\x69\x4b\x42\x3f\x59\xa8\xd4\xc9\x83\x3f\x9f\x62\xa1\x37\x96\xab\xdc\xcd\x69\xe5\x5e\x2e\xa8\xf1\x9a\x85\xb6\xfd\x8d\xf7\xd9\x28\x0b\x5d\xb6\x9e\xe7\x1f\x91\x96\x4f\xf4\xe9\xb1\x99\x3e\xe5\xc7\xed\x7d\x54\xff\xe4\x21\xe4\x3b\xc3\x42\x1f\x16\x96\x49\x67\x16\x33\x5a\x1a\x4d\xeb\xc4\x58\x5a\xcf\x5e\xcf\x04\x1e\xde\xb1\xd0\xee\xdf\x6b\xf8\xff\x6b\x59\xe8\x81\x5d\x77\xfe\x61\x4e\xa7\x26\x7b\xb5\x90\xc7\x96\x8e\x96\xf7\xe8\xfc\x52\xdb\xc3\x0d\x16\x8a\x07\x7c\xa8\xf7\x38\x0b\x7d\xb7\xfc\xc9\x0b\x03\x6a\x5c\x54\xe3\x60\xd3\xaa\x38\x82\x7a\xfa\xca\x60\xe0\xe1\x0b\x0b\xed\x0b\x71\x56\x5b\x59\xe8\x95\x9d\x9f\xbb\x93\xd3\x23\xe3\x3d\x6a\x86\xd0\xd6\x78\x6a\x5c\x7c\x61\xc7\x93\xa5\x5e\xad\xde\x6f\x9f\xe5\x13\x16\xda\x14\xe2\xcf\xb1\x50\xd3\x9e\xd3\xf4\x6c\xbf\xc7\x3a\xf8\xb1\x38\x8c\xd9\xe2\x30\x8e\x18\x07\xbf\x1a\x07\xf5\x38\xf7\x07\x7b\x68\xb0\x50\xd9\x67\xbb\x59\xe8\x25\xd7\xe8\xfd\xec\xcd\x6c\xfd\xd0\x81\x74\xdd\x0c\xe1\x8a\x71\xf1\xd2\xc4\x58\x3a\x5a\x1a\x4d\xc7\x8d\x8b\xd7\x6c\xee\xf0\xc1\x9e\xfa\xdc\xed\x6c\x9d\x6b\xb4\xe2\x7b\xd8\xce\x42\x1b\xec\x9d\x2f\x2c\x53\xea\xd4\xc5\x81\x54\x21\x8f\x29\xe3\x60\xaa\xe8\x62\x97\x71\x10\x0a\x79\x84\xa2\x8b\x1b\x8d\x83\xbb\x6c\xfe\xcc\xd5\xc1\xd6\x5a\xff\x9d\xac\xeb\xf8\xa0\x3b\x84\x2a\x40\xb3\x02\xd0\xc8\x7a\xfa\xb1\xc3\xd3\xe7\x2e\x80\xb7\x56\x71\x80\x67\x51\x4f\xd5\x88\x27\x1b\x91\xaa\xc7\xfe\x0a\x00\x00\xff\xff\x56\x45\x56\x48\x7e\x04\x00\x00")

func init() {

	rb := bytes.NewReader(FilePublicFaviconIco)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "public/favicon.ico", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
