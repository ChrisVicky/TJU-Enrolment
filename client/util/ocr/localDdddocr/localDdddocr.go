package localddddocr

import (
	"fmt"
	"os/exec"
)

var (
	script  = "ocr.py"
	testImg = "test.png"
	testAns = "3g5t"
)

type ddddocr struct {
	script string
}

func NewDdddocr(fn string) *ddddocr {
	d := &ddddocr{
		script: fn,
	}
	return d
}

func (d *ddddocr) OcrFn(fn string) (code string, err error) {
	cmd := exec.Command("python", d.script, fn)
	cmd_str := cmd.String()
	fmt.Println("cmd:", cmd_str)
	out, err := cmd.Output()
	if err != nil {
		return
	}
	code = string(out)
	return
}

func (d *ddddocr) Setup() error {
	code, err := d.OcrFn(testImg)
	if err != nil {
		return err
	}
	if code != testAns {
		return fmt.Errorf("expected: %v, get: %v", testAns, code)
	}

	return nil
}
