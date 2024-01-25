package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Zip(body []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(body)
	w.Close()
	return buf.Bytes(), err
}

func UnZip(body []byte) ([]byte, error) {
	buf := bytes.NewBuffer(body)
	r, err := gzip.NewReader(buf)
	io.Copy(buf, r)
	r.Close()
	return buf.Bytes(), err
}
