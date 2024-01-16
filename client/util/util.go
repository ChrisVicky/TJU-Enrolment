package util

import (
	"enrollment/client/util/jsencoder"
	"enrollment/logger"
	"fmt"
	"net/http"
	"regexp"
)

type Util struct {
	*jsencoder.Jsvm
}

func NewUtil() (u *Util, err error) {
	j, err := jsencoder.NewJsvm()
	if err != nil {
		return
	}
	u = &Util{Jsvm: j}
	return
}

func (u *Util) SetDefaultHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/117.0")

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// req.Header.Set("Origin","https://sso.tju.edu.cn")
	// req.Header.Set("Host","sso.tju.edu.cn")
	// req.Header.Set("Referer","https://sso.tju.edu.cn/cas/login")
	// req.Header.Set("Keep-Alive", "timeout=1, max=1000")
	// req.Header.Set("Connection","keep-alive")
	// req.Header.Set("Accept-Language","en-US,en;q=0.5")
	// req.Header.Set("Accept-Charset", "GB2312,utf-8;q=0.7,*;q=0.7")
	// req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
}

func (u *Util) ExtractLtExecution(content string) (string, string, error) {
	ltPattern := `id="lt" name="lt" value="([^"]+)"`
	executionPattern := `name="execution" value="([^"]+)"`

	ltRegex := regexp.MustCompile(ltPattern)
	executionRegex := regexp.MustCompile(executionPattern)

	ltMatches := ltRegex.FindStringSubmatch(content)
	executionMatches := executionRegex.FindStringSubmatch(content)

	if len(ltMatches) <= 1 || len(executionMatches) <= 1 {
		return "", "", fmt.Errorf("cannot find both values")
	}

	lt, ex := ltMatches[1], executionMatches[1]
	logger.Tracef("lt, ex: %s, %s", lt, ex)
	return lt, ex, nil
}

func (u *Util) GetRsa(name, pw, lt string) (rsa string, err error) {
	rsa, err = u.Enc(name + pw + lt)
	logger.Tracef("rsa: %s", rsa)
	return
}
