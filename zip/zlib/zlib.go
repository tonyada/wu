package zlib

import (
	"bytes"
	"compress/zlib"
	"io"
)

const (
	myZlibDict = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789{}[]:"
)

func Zip(body []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	_, err := w.Write(body)
	w.Close()
	return buf.Bytes(), err
}

func UnZip(body []byte) ([]byte, error) {

	buf := bytes.NewBuffer(body)
	r, err := zlib.NewReader(buf)
	io.Copy(buf, r)
	r.Close()
	return buf.Bytes(), err
}

func ZipDict(body []byte) ([]byte, error) {
	var buf bytes.Buffer
	w, err := zlib.NewWriterLevelDict(&buf, zlib.DefaultCompression, []byte(myZlibDict))
	w.Write(body)
	w.Close()
	return buf.Bytes(), err
}

func UnZipDict(body []byte) ([]byte, error) {
	buf := bytes.NewBuffer(body)
	r, err := zlib.NewReaderDict(buf, []byte(myZlibDict))
	io.Copy(buf, r)
	r.Close()
	return buf.Bytes(), err
}
