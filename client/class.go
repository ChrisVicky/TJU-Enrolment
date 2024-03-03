package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func (e *EClient) getClasses() (err error) {
	var (
		resp      *http.Response
		req       *http.Request
		bodyBytes []byte
	)

	if resp, err = e.PostForm("http://classes.tju.edu.cn/eams/stdElectCourse!defaultPage.action", url.Values{
		"electionProfile.id": {e.pid},
	}); err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("post for default page of select returns not 200, but %v", resp.StatusCode)
		return
	}

	u := fmt.Sprintf("http://classes.tju.edu.cn/eams/stdElectCourse!data.action?profileId=%v", e.pid)
	if req, err = http.NewRequest(http.MethodGet, u, nil); err != nil {
		return
	}

	if resp, err = e.Do(req); err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("go to select course: status code Error: %v", resp.StatusCode)
		return
	}

	if bodyBytes, err = io.ReadAll(resp.Body); err != nil {
		err = errors.New("error emitted while reading from body with io.ReadAll")
		return
	}
	e.store = newLessonStore(string(bodyBytes))
	return
}

func newLessonStore(data string) LessonStore {
	re := regexp.MustCompile(`\{id:(\d+),no:'(\d+)',name:'([^']+)`)
	matches := re.FindAllStringSubmatch(data, -1)

	r := make(LessonStore)
	for _, match := range matches {
		r[match[2]] = &Lesson{ID: match[1], Name: match[3]}
	}

	return r
}
