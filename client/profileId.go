package client

import (
	"enrollment/logger"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

func (e *EClient) fetchProfileId() (err error) {
	var (
		content string

		resp      *http.Response
		req       *http.Request
		bodyBytes []byte
		pid       string
	)

	time.Sleep(100 * time.Millisecond)
	if req, err = http.NewRequest(http.MethodGet, "http://classes.tju.edu.cn/eams/stdElectCourse.action", nil); err != nil {
		return
	}
	e.SetDefaultHeaders(req)
	if resp, err = e.Do(req); err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("go to select course: status code error: %v", resp.StatusCode)
		return
	}

	if bodyBytes, err = io.ReadAll(resp.Body); err != nil {
		err = errors.New("error emitted while reading from body with io.readall")
		return
	}
	content = string(bodyBytes)
	if pid, err = extractValue(content); err != nil {
		return
	}
	logger.Tracef("Pid: %v", pid)
	e.pid = pid
	return
}

func extractValue(html string) (string, error) {
	re := regexp.MustCompile(`type='hidden'\s+name='electionProfile\.id'\s+value='(\d+)'`)
	matches := re.FindAllStringSubmatch(html, -1)
	if len(matches) >= 1 {
		return matches[len(matches)-1][1], nil
	}
	return "", errors.New("value not found")
}
