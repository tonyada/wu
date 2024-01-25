package net

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	TEST_URL           = "http://example.com"
	EXAMPLE_WEB_PREFIX = `<!doctype html>
<html>
<head>
    <title>Example Domain</title>
`
)

func TestHttpGet(t *testing.T) {
	s, err := HttpGet(TEST_URL)
	if err != nil {
		t.Errorf("HttpGet:\n Expect => %v\n Got => %s\n", nil, err)
	}
	if !strings.HasPrefix(string(s), EXAMPLE_WEB_PREFIX) {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", EXAMPLE_WEB_PREFIX, s)
	}
}

func TestHttpGetString(t *testing.T) {
	s, err := HttpGetString(TEST_URL)
	if err != nil {
		t.Errorf("HttpGet:\n Expect => %v\n Got => %s\n", nil, err)
	}
	if !strings.HasPrefix(s, EXAMPLE_WEB_PREFIX) {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", EXAMPLE_WEB_PREFIX, s)
	}
}

func TestClientGet(t *testing.T) {
	// 200.
	rc, err := ClientGet(&http.Client{}, TEST_URL, nil)
	if err != nil {
		t.Fatalf("ClientGet:\n Expect => %v\n Got => %s\n", nil, err)
	}
	p, err := ioutil.ReadAll(rc)
	if err != nil {
		t.Errorf("ClientGet:\n Expect => %v\n Got => %s\n", nil, err)
	}
	s := string(p)
	if !strings.HasPrefix(s, EXAMPLE_WEB_PREFIX) {
		t.Errorf("ClientGet:\n Expect => %s\n Got => %s\n", EXAMPLE_WEB_PREFIX, s)
	}
}

func TestClientGetBytes(t *testing.T) {
	p, err := ClientGetBytes(&http.Client{}, TEST_URL, nil)
	if err != nil {
		t.Errorf("ClientGetBytes:\n Expect => %v\n Got => %s\n", nil, err)
	}
	s := string(p)
	if !strings.HasPrefix(s, EXAMPLE_WEB_PREFIX) {
		t.Errorf("ClientHttpGet:\n Expect => %s\n Got => %s\n", EXAMPLE_WEB_PREFIX, s)
	}
}

func TestClientGetString(t *testing.T) {
	p, err := ClientGetString(&http.Client{}, TEST_URL, nil)
	if err != nil {
		t.Errorf("ClientGetBytes:\n Expect => %v\n Got => %s\n", nil, err)
	}
	s := p
	if !strings.HasPrefix(s, EXAMPLE_WEB_PREFIX) {
		t.Errorf("ClientHttpGet:\n Expect => %s\n Got => %s\n", EXAMPLE_WEB_PREFIX, s)
	}
}

// func TestClientHttpGetJSON(t *testing.T) {

// }

// type rawFile struct {
// 	name   string
// 	rawURL string
// 	data   []byte
// }

// func (rf *rawFile) Name() string {
// 	return rf.name
// }

// func (rf *rawFile) RawUrl() string {
// 	return rf.rawURL
// }

// func (rf *rawFile) Data() []byte {
// 	return rf.data
// }

// func (rf *rawFile) SetData(p []byte) {
// 	rf.data = p
// }

// func TestClientFetchFiles(t *testing.T) {
// 	files := []RawFile{
// 		&rawFile{rawURL: TEST_URL},
// 		&rawFile{rawURL: TEST_URL},
// 	}
// 	err := ClientFetchFiles(&http.Client{}, files, nil)
// 	if err != nil {
// 		t.Errorf("ClientFetchFiles:\n Expect => %v\n Got => %s\n", nil, err)
// 	} else if len(files[0].Data()) != 1270 {
// 		t.Errorf("ClientFetchFiles:\n Expect => %d\n Got => %d\n", 1270, len(files[0].Data()))
// 	} else if len(files[1].Data()) != 1270 {
// 		t.Errorf("ClientFetchFiles:\n Expect => %d\n Got => %d\n", 1270, len(files[1].Data()))
// 	}
// }

// func TestClientFetchFilesCurl(t *testing.T) {
// 	files := []RawFile{
// 		&rawFile{rawURL: TEST_URL},
// 		&rawFile{rawURL: TEST_URL},
// 	}
// 	err := ClientFetchFilesCurl(files)
// 	if err != nil {
// 		t.Errorf("ClientFetchFilesCurl:\n Expect => %v\n Got => %s\n", nil, err)
// 	} else if len(files[0].Data()) != 1270 {
// 		t.Errorf("ClientFetchFilesCurl:\n Expect => %d\n Got => %d\n", 1270, len(files[0].Data()))
// 	} else if len(files[1].Data()) != 1270 {
// 		t.Errorf("ClientFetchFilesCurl:\n Expect => %d\n Got => %d\n", 1270, len(files[1].Data()))
// 	}
// }
