package remoteserver

import "testing"

func TestRemote(t *testing.T) {
	r := NewRemoteServer("https://learning.twt.edu.cn/ocr")
	if err := r.Setup(); err != nil {
		t.Errorf("err: %v", err)
	}
}

func TestLocal(t *testing.T) {
	r := NewRemoteServer("http://127.0.0.1:8000/uploadfile/")
	if err := r.Setup(); err != nil {
		t.Errorf("err: %v", err)
	}
}
