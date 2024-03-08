// Package client specify enrolment client
package client

import (
	"enrolment/client/util"
	"enrolment/client/util/ocr"
	"enrolment/conf"
	"enrolment/logger"
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

type EClient struct {
	*http.Client
	*util.Util
	*conf.Account

	ocr.OcrServer

	store LessonStore // Lesson Storage
	pid   string
	idx   string // Id for Lesson (Map to Lesson)
}

type Lesson struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (l *Lesson) String() string {
	return fmt.Sprintf("{%s:%s}", l.ID, l.Name)
}

type LessonStore map[string]*Lesson

func (e *EClient) String() string {
	c := ""
	if e.idx != "" {
		c += fmt.Sprintf("%s(%s)", e.idx, e.Courses[e.idx])
	} else {
		for k, v := range e.Courses {
			c += fmt.Sprintf("%s(%s),", k, v)
		}
	}
	return fmt.Sprintf("{%s:%s(%s), %s pid:%s}", e.StudentNo, e.Password, e.Comment, c, e.pid)
}

// studentNo, password, comment, courseId, courseComment
func NewEClient(a *conf.Account, oc conf.Ocr, idx string) (*EClient, error) {
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

	o, err := ocr.NewOcrServer(oc)
	if err != nil {
		return nil, err
	}
	if err := o.Setup(); err != nil {
		return nil, err
	}

	e := &EClient{
		Client:    c,
		Util:      u,
		Account:   a,
		OcrServer: o,
		idx:       idx,
	}
	return e, nil
}

func (e *EClient) Notation() string {
	return fmt.Sprintf("%s: %s", e.Comment, e.store[e.idx])
}

type (
	LoginError      struct{ error }
	FetchPidError   struct{ error }
	GetClassesError struct{ error }
)

func (e *EClient) Refresh(le error) error {
	var err error
	if le == nil {
		le = LoginError{}
	}

	switch le.(type) {
	case FetchPidError:
		if err = e.fetchProfileId(); err != nil {
			return FetchPidError{err}
		}

		logger.Tracef("Pid Fetched: %v", e.pid)

		if err = e.getClasses(); err != nil {
			return GetClassesError{err}
		}

	case GetClassesError:
		if err = e.getClasses(); err != nil {
			return GetClassesError{err}
		}

	default:
		if err = e.Login(); err != nil {
			return LoginError{err}
		}
		if err = e.fetchProfileId(); err != nil {
			return FetchPidError{err}
		}
		if err = e.getClasses(); err != nil {
			return GetClassesError{err}
		}
	}
	return nil
}
