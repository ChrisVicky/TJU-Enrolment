// File uiLogger.go implements a tview logger,
// pipelining log contents to *tview.TextView
package logger

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rifflock/lfshook"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
)

type logWriter struct {
	textView *tview.TextView
}

var LogUi *logWriter

type customFormatter struct{}

// colorForLevel returns the corresponding Color Indicators
func colorForLevel(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return tcell.ColorRebeccaPurple.TrueColor().String()
	case logrus.InfoLevel:
		return tcell.ColorGreen.TrueColor().String()
	case logrus.WarnLevel:
		return tcell.ColorYellow.TrueColor().String()
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return tcell.ColorRed.TrueColor().String()
	default:
		return tcell.ColorWhite.TrueColor().String()
	}
}

// Format implement logrus.Formatter interfaces
func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := strings.ToUpper(entry.Level.String())
	color := colorForLevel(entry.Level)

	// ANSI is used to formalize output -- Colorized
	_, err := b.WriteString(fmt.Sprintf("[%s]%s (%s) %s [-]\n", color, timestamp, level, entry.Message))
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (lw *logWriter) Write(p []byte) (n int, err error) {
	lw.textView.Write(p) // Write to TextView
	return len(p), nil
}

func SetTview(t *tview.TextView) {
	LogUi = &logWriter{textView: t}
	t.SetChangedFunc(func() {
		t.ScrollToEnd()
	})
	t.SetDynamicColors(true)

	uiHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  LogUi,
		logrus.WarnLevel:  LogUi,
		logrus.ErrorLevel: LogUi,
		logrus.FatalLevel: LogUi,
		logrus.PanicLevel: LogUi,
	}, new(customFormatter))

	logger.AddHook(uiHook)
}
