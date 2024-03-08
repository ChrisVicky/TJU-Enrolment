package client

import (
	"enrolment/logger"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrorSelected        = errors.New("失败:你已经选过")
	ErrorMultiple        = errors.New("失败:多课程类别")
	ErrorMax2            = errors.New("当前学期限制门数上限2")
	ErrorSuccess         = errors.New("成功")
	ErrorFast            = errors.New("点击过快")
	ErrorService         = errors.New("服务器内部错误")
	ErrorNotFinish error = nil
)

var keywords = map[string]error{
	"失败:你已经选过":    ErrorSelected,
	"失败:多课程类别":    ErrorMultiple,
	"当前学期限制门数上限2": ErrorMax2,
	"点击过快":        ErrorFast,
	"服务器内部错误":     ErrorService,
	"选课开放时间":      ErrorNotFinish,
	"成功":          ErrorSuccess,
}

// Select Class
//
// NOTE: sub goroutines share the same `Client` cause issues?
func (e *EClient) Select() (err error) {
	var (
		resp *http.Response
		req  *http.Request
		bd   []byte
	)
	id := e.store[e.idx].ID
	name := e.store[e.idx].Name
	logger.Infof("[%s] selecting for %s [%s]", e.Comment, id, name)

	u := fmt.Sprintf("http://classes.tju.edu.cn/eams/stdElectCourse!batchOperator.action?profileId=%v", e.pid)

	data := url.Values{
		"optype":    {"true"},
		"operator0": {fmt.Sprintf("%v:true:0", id)},
	}.Encode()
	if req, err = http.NewRequest(http.MethodPost, u, strings.NewReader(data)); err != nil {
		return
	}
	e.SetDefaultHeaders(req)

	if resp, err = e.Do(req); err != nil {
		return
	}

	if bd, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	for k, v := range keywords {
		if ok := strings.Contains(string(bd), k); ok {
			if v == ErrorSuccess {
				logger.Debug(string(bd))
			}
			return v
		}
	}

	logger.Trace(string(bd))

	return
}
