package ocr

import (
	localddddocr "enrollment/client/util/ocr/localDdddocr"
	remoteserver "enrollment/client/util/ocr/remoteServer"
	"enrollment/conf"
	"fmt"
)

const (
	LocalDdddocr = iota
	RemoteServer
	LocalServer
)

type OcrServer interface {
	Setup() error
	OcrFn(string) (string, error)
}

func NewOcrServer(c conf.Ocr) (OcrServer, error) {
	switch c.Type {
	case LocalDdddocr:
		return localddddocr.NewDdddocr(c.Payload), nil
	case RemoteServer:
		return remoteserver.NewRemoteServer(c.Payload), nil
	case LocalServer:
		return remoteserver.NewRemoteServer(c.Payload), nil
	default:
		return nil, fmt.Errorf("cannot recognize: %v", c.Type)
	}
}
