package net

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	. "wu"
)

func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return io.ReadAll(resp.Body)
	}
	if resp.StatusCode == 404 {
		return nil, Errf("404 Not Found")
	}
	return nil, Errf("Bad status code: %v", resp.StatusCode)
}

// http get with referer
func HttpGetWithReferer(url, referer string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if Err(err) {
		return nil, err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", referer)
	res, err := client.Do(req)
	// Log(res.Body)
	if Err(err) {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)

}

// http get with referer (auto use url to referer)
func HttpGetAutoReferer(url string) ([]byte, error) {
	return HttpGetWithReferer(url, url)
}

func HttpGetString(url string) (string, error) {
	body, err := HttpGetWithReferer(url, url)
	return string(body), err
}

// Downlaod file to pathfilename
func DownloadFile(durl string, pathfile string) (size int64, err error) {
	size = 0
	// download
	res, err := http.Get(durl)
	if Err(err) {
		return 0, err
	}
	defer res.Body.Close()

	// not found and try root url
	if res.StatusCode == 404 {
		return 0, ErrNew("404 Not Found")
	}

	if res.StatusCode != 200 {
		return 0, ErrNew(Sprintf("wrong StatusCode %d", res.StatusCode))
	}
	// create
	f, err := os.Create(pathfile)
	if Err(err) {
		return 0, err
	}
	defer f.Close()
	// saving file
	size, err = io.Copy(f, res.Body)
	if Err(err) {
		return 0, err
	}
	// valid download file size
	if size != res.ContentLength && res.ContentLength != -1 {
		Log("warning size != ContentLength", durl)
		Log("contentLength:", res.ContentLength)
		Log("download size:", size)
	}
	// var totalSize int64 = 0
	// totalSize += size
	// Log("saved: [%s] (%dk) total: [%vk]", pathfile, size/1024, totalSize/1024)
	// Log("Total: [%vk]", totalSize/1024)
	return size, nil
}

// HTTP method
func ClientDo(client *http.Client, method, url string, header http.Header, postData io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, url, postData)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", USER_AGENT)
	for k, vs := range header {
		req.Header[k] = vs
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		return resp.Body, nil
	}
	resp.Body.Close()
	if resp.StatusCode == 404 { // 403 can be rate limit error.  || resp.StatusCode == 403 {
		err = Errf("404 not found: %s", url)
	} else {
		err = Errf("%s %s -> %d", method, url, resp.StatusCode)
	}
	return nil, err
}

// HttpGet gets the specified resource.
func ClientGet(client *http.Client, url string, header http.Header) (io.ReadCloser, error) {
	return ClientDo(client, "GET", url, header, nil)
}

func ClientGetBytes(client *http.Client, url string, header http.Header) ([]byte, error) {
	rc, err := ClientGet(client, url, header)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return io.ReadAll(rc)
}

func ClientGetString(client *http.Client, url string, header http.Header) (string, error) {
	body, err := ClientGetBytes(client, url, header)
	return string(body), err
}

// HttpPost posts the specified resource.
func ClientPost(client *http.Client, url string, header http.Header, postData []byte) (io.ReadCloser, error) {
	return ClientDo(client, "POST", url, header, bytes.NewBuffer(postData))
}

// HttpGetJSON gets the specified resource and mapping to struct.
func HttpGetJSON(client *http.Client, url string, v interface{}) error {
	rc, err := ClientGet(client, url, nil)
	if err != nil {
		return err
	}
	defer rc.Close()
	err = json.NewDecoder(rc).Decode(v)
	if _, ok := err.(*json.SyntaxError); ok {
		return Errf("JSON syntax error at %s", url)
	}
	return nil
}

// HttpPostJSON posts the specified resource with struct values,
// and maps results to struct.
func HttpPostJSON(client *http.Client, url string, body, v interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	rc, err := ClientPost(client, url, http.Header{"content-type": []string{"application/json"}}, data)
	if err != nil {
		return err
	}
	defer rc.Close()
	err = json.NewDecoder(rc).Decode(v)
	if _, ok := err.(*json.SyntaxError); ok {
		return Errf("JSON syntax error at %s", url)
	}
	return nil
}

// A RawFile describes a file that can be downloaded.
type RawFile interface {
	Name() string
	RawUrl() string
	Data() []byte
	SetData([]byte)
}

// FetchFiles fetches files specified by the rawURL field in parallel.
func FetchFiles(client *http.Client, files []RawFile, header http.Header) error {
	ch := make(chan error, len(files))
	for i := range files {
		go func(i int) {
			p, err := ClientGetBytes(client, files[i].RawUrl(), nil)
			if err != nil {
				ch <- err
				return
			}
			files[i].SetData(p)
			ch <- nil
		}(i)
	}
	for range files {
		if err := <-ch; err != nil {
			return err
		}
	}
	return nil
}

// FetchFiles uses command `curl` to fetch files specified by the rawURL field in parallel.
// func FetchFilesCurl(files []RawFile, curlOptions ...string) error {
// 	ch := make(chan error, len(files))
// 	for i := range files {
// 		go func(i int) {
// 			stdout, _, err := ExecCmd("curl", append(curlOptions, files[i].RawUrl())...)
// 			if err != nil {
// 				ch <- err
// 				return
// 			}

// 			files[i].SetData([]byte(stdout))
// 			ch <- nil
// 		}(i)
// 	}
// 	for _ = range files {
// 		if err := <-ch; err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
