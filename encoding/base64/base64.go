package base64

import (
	"encoding/base64"
)

const (
	// my base64 table
	myBase64Table = "KLMNOHIJbctuvwxyijklmn2efgZaYqhFG89-_345rsWXDEz01C67PQRSTUVdABop"
)

// My base64 Encoding (using my base64 table)
func MyEncode(b []byte) string { return base64.NewEncoding(myBase64Table).EncodeToString(b) }

// My base64 Decoding (using my base64 table)
func MyDecode(s string) ([]byte, error) { return base64.NewEncoding(myBase64Table).DecodeString(s) }

// base64 URL Encoding
func UrlEncode(b []byte) string { return base64.URLEncoding.EncodeToString(b) }

// base64 URL Decoding
func UrlDecode(s string) ([]byte, error) { return base64.URLEncoding.DecodeString(s) }

// go std base64 std Encoding
func StdEncode(b []byte) string { return base64.StdEncoding.EncodeToString(b) }

// go std base64 std Decoding
func StdDecode(s string) ([]byte, error) { return base64.StdEncoding.DecodeString(s) }
