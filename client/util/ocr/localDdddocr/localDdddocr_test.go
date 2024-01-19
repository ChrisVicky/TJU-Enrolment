package localddddocr

import "testing"

func TestLocal(t *testing.T) {
	d := NewDdddocr("./ocr.py")
	if err := d.Setup(); err != nil {
		t.Errorf("err: %v", err)
	}
}
