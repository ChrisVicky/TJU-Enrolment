package remoteserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const (
	testImg = "test.png"
	testAns = "3g5t"
)

type RemoteServer struct {
	*http.Client
	api string
}

func NewRemoteServer(api string) *RemoteServer {
	r := &RemoteServer{
		api:    api,
		Client: &http.Client{},
	}
	return r
}

func (r *RemoteServer) Setup() error {
	code, err := r.OcrFn(testImg)
	if err != nil {
		return err
	}

	if code != testAns {
		return fmt.Errorf("wanted: %v, get: %v", testAns, code)
	}
	return nil
}

type learning_return struct {
	Data string `json:"data"`
	Code int    `json:"code"`
}

func (r *RemoteServer) OcrFn(fn string) (code string, err error) {
	var (
		resp      *http.Response
		req       *http.Request
		bodyBytes []byte
		formFile  io.Writer
	)

	bodyBytes, err = os.ReadFile(fn)
	if err != nil {
		return
	}

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	if formFile, err = writer.CreateFormFile("image", "123.png"); err != nil {
		return
	}
	if _, err = io.Copy(formFile, strings.NewReader(string(bodyBytes))); err != nil {
		return
	}
	writer.Close()

	if req, err = http.NewRequest(http.MethodPost, r.api, buf); err != nil {
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/117.0")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	if resp, err = r.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("get code status code error: %v", resp.StatusCode)
		return
	}

	if bodyBytes, err = io.ReadAll(resp.Body); err != nil {
		return
	}
	value := &learning_return{}
	if err = json.Unmarshal(bodyBytes, value); err != nil {
		return
	}
	if value.Code != 200 {
		err = fmt.Errorf("code: %v, data: %v", value.Code, value.Data)
		return
	}
	code = value.Data
	return
}
