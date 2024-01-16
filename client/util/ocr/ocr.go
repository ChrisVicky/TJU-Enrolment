package ocr

import (
	"os/exec"
)

// Function OcrFn run ocr.py to recognize via OCR
func OcrFn(fn string) (string, error) {
	cmd := exec.Command("python", "ocr.py", fn)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
