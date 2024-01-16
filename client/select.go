package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrorSelected = errors.New("失败:你已经选过")
	ErrorMultiple = errors.New("失败:多课程类别")
	ErrorMax2     = errors.New("当前学期限制门数上限2")
	ErrorSuccess  = errors.New("成功")
)

var keywords = map[string]error{
	"失败:你已经选过":    ErrorSelected,
	"失败:多课程类别":    ErrorMultiple,
	"当前学期限制门数上限2": ErrorMax2,
	"成功":          ErrorSuccess,
}

// Select Class
func (e *EClient) Select() (err error) {
	var (
		resp *http.Response
		req  *http.Request
		bd   []byte
	)

	u := fmt.Sprintf("http://classes.tju.edu.cn/eams/stdElectCourse!batchOperator.action?profileId=%v", e.pid)

	data := url.Values{
		"optype":    {"true"},
		"operator0": {fmt.Sprintf("%v:true:0", e.courseId)},
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
			return v
		}
	}
	return
}
