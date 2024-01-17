// Package client specify enrollment client
package client

import (
	"enrollment/client/util"
	"enrollment/conf"
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

type EClient struct {
	*http.Client
	*util.Util
	*conf.Account

	pid string
}

func (e *EClient) String() string {
	c := ""
	for idx := range e.CourseComment {
		c += fmt.Sprintf("%s(%s),", e.CourseNo[idx], e.CourseComment[idx])
	}
	return fmt.Sprintf("{%s:%s(%s), %s pid:%s}", e.StudentNo, e.Password, e.Comment, c, e.pid)
}

// studentNo, password, comment, courseId, courseComment
func NewEClient(a *conf.Account) (*EClient, error) {
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
		Client:  c,
		Util:    u,
		Account: a,
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
