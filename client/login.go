package client

import (
	"enrolment/logger"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func (e *EClient) Login() (err error) {
	var (
		req  *http.Request
		resp *http.Response
		body []byte
		lt   string
		ex   string
		rsa  string
		code string
	)

	loginUrl := "https://sso.tju.edu.cn/cas/login"

	if req, err = http.NewRequest(http.MethodGet, loginUrl, nil); err != nil {
		return
	}
	e.SetDefaultHeaders(req)

	if resp, err = e.Do(req); err != nil {
		return
	}

	if resp.StatusCode == 302 {
		logger.Info("Logged In")
		return
	}

	if resp.Request.URL.String() != loginUrl {
		logger.Info("Logged In")
		return
	}

	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	if lt, ex, err = e.ExtractLtExecution(string(body)); err != nil {
		return
	}

	if rsa, err = e.GetRsa(e.StudentNo, e.Password, lt); err != nil {
		return
	}

	if code, err = e.getCode(); err != nil {
		return
	}

	return e.ssoLogin(code, ex, lt, rsa)
}

func (e *EClient) ssoLogin(code, ex, lt, rsa string) (err error) {
	var (
		resp *http.Response
		req  *http.Request
		body []byte
	)

	bd := url.Values{
		"code":      {code},
		"ul":        {strconv.FormatInt(int64(len(e.StudentNo)), 10)},
		"pl":        {strconv.FormatInt(int64(len(e.Password)), 10)},
		"lt":        {lt},
		"rsa":       {rsa},
		"execution": {ex},
		"_eventId":  {"submit"},
	}.Encode()

	req, err = http.NewRequest(http.MethodPost, "https://sso.tju.edu.cn/cas/login", strings.NewReader(bd))
	if err != nil {
		return
	}
	e.SetDefaultHeaders(req)
	resp, err = e.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("sso login Status Code Error: %v", resp.StatusCode)
		return
	}

	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	// Match Name
	r := regexp.MustCompile(`</p> <span class="tit">([^<]+)</span></a>`)
	matches := r.FindStringSubmatch(string(body))
	if len(matches) > 1 {
		uname := matches[1]
		logger.Infof("Hello %v", uname)
	}

	return
}

func (e *EClient) downloadCode(fn string) (err error) {
	err = nil

	var (
		resp      *http.Response
		req       *http.Request
		bodyBytes []byte
	)

	if req, err = http.NewRequest(http.MethodGet, "https://sso.tju.edu.cn/cas/code", nil); err != nil {
		return
	}
	e.SetDefaultHeaders(req)
	if resp, err = e.Do(req); err != nil {
		return
	}

	defer resp.Body.Close()
	if bodyBytes, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("receiving codes error: %v", resp.Status)
		return
	}

	if err = os.WriteFile(fn, bodyBytes, 0666); err != nil {
		return
	}

	return
}

func (e *EClient) recognizeCode(fn string) (code string, err error) {
	code, err = e.OcrFn(fn)
	if err != nil {
		logger.Warnf("Auto recognition failed with: %v", err)
		return
	}
	logger.Infof("Code recognized: %v", code)
	return
}

func (e *EClient) getCode() (code string, err error) {
	fn := filepath.Join(logger.RUNTIME, e.StudentNo+".png")
	if err = e.downloadCode(fn); err != nil {
		return
	}
	logger.Tracef("Captcha saved to: %v", fn)
	return e.recognizeCode(fn)
}
