// Package client specify enrollment client
package client

import (
	"enrollment/client/util"
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

type EClient struct {
	*http.Client
	*util.Util

	studentNo string
	password  string
	comment   string

	courseId      string
	courseComment string

	pid string
}

func (e *EClient) String() string {
	return fmt.Sprintf("{%s:%s(%s), id:%s(%s),  pid:%s}", e.studentNo, e.password, e.comment, e.courseId, e.courseComment, e.pid)
}

// studentNo, password, comment, courseId, courseComment
func NewEClient(s, p, com, id, ccom string) (*EClient, error) {
	cj, err := cookiejar.New(
		&cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		})
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Jar: cj,
	}

	u, err := util.NewUtil()
	if err != nil {
		return nil, err
	}

	e := &EClient{
		Client: c,
		Util:   u,

		studentNo:     s,
		password:      p,
		comment:       com,
		courseId:      id,
		courseComment: ccom,
	}
	return e, nil
}

func (e *EClient) Refresh() (err error) {
	if err = e.Login(); err != nil {
		return
	}

	if err = e.FetchProfileId(); err != nil {
		return
	}
	return
}
