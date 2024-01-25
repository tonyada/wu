package lzma

import (
	"bytes"
	"io"

	"github.com/itchio/lzma"
)

func Zip(body []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := lzma.NewWriter(&buf)
	_, err := w.Write(body)
	w.Close()
	return buf.Bytes(), err
}

func UnZip(body []byte) ([]byte, error) {
	buf := bytes.NewBuffer(body)
	r := lzma.NewReader(buf)
	_, err := io.Copy(buf, r)
	r.Close()
	return buf.Bytes(), err
}
