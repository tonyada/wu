package zlib_aes_base64

import (
	. "wu"
	"wu/crypto/aes"
	"wu/encoding/base64"
	"wu/zip/zlib"
)

func Encode(s []byte) string {
	// zip
	zippedData, err := zlib.Zip(s)
	if Err(err) {
		Println("zlib error")
		return ""
	}
	// aes encrypt
	enBody, err := aes.Encrypt(zippedData)
	if Err(err) {
		Println("aes error")
		return ""
	}

	// base64 encoding
	base64Body := base64.MyEncode(enBody)
	Println("MyEncode OK len = ", len(base64Body), "\n", string(base64Body))
	return base64Body
}

func Decode(s string) string {
	// base64 decoding
	base64DecodeBody, _ := base64.MyDecode(s)
	// aes decrypt
	deBody, err := aes.Decrypt(base64DecodeBody)

	if Err(err) {
		Println("aes error")
		return ""
	}

	unzippedData, err := zlib.UnZip(deBody)
	if Err(err) {
		Println("unzlib error")
		return ""
	}
	Println("READY TO READ!!! unzip ok body len = ", len(unzippedData), "\n", string(unzippedData))
	return string(unzippedData)
}
