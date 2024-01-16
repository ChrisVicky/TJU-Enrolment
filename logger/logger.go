package logger

import (
	"os"
	"path/filepath"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	logger   *logrus.Logger
	RUNTIME  = "./runtime"
	logFn    = "app.log"
	logfiles = []string{}
)

type nullWriter struct{}

func (*nullWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

// showLog == false remove log information to `stdout`
func Setup(showLog bool) error {
	// Mkdir & Return error if it exists
	os.Mkdir(RUNTIME, os.ModePerm)

	logger = logrus.New()

	if !showLog {
		logger.SetOutput(&nullWriter{})
	}

	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(
		&prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	)

	// Define hooker writes to local file
	fn := filepath.Join(RUNTIME, logFn)
	logFile, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	// logfiles = append(logfiles, fn)

	fileHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  logFile,
		logrus.WarnLevel:  logFile,
		logrus.ErrorLevel: logFile,
		logrus.FatalLevel: logFile,
		logrus.PanicLevel: logFile,
	}, &logrus.JSONFormatter{})

	logger.AddHook(fileHook)
	return nil
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}

func Tracef(str string, args ...interface{}) {
	logger.Tracef(str, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(str string, args ...interface{}) {
	logger.Debugf(str, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(str string, args ...interface{}) {
	logger.Infof(str, args...)
}

func Print(args ...interface{}) {
	logger.Print(args...)
}

func Printf(str string, args ...interface{}) {
	logger.Printf(str, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(str string, args ...interface{}) {
	logger.Warnf(str, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(str string, args ...interface{}) {
	logger.Errorf(str, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(str string, args ...interface{}) {
	logger.Fatalf(str, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(str string, args ...interface{}) {
	logger.Panicf(str, args...)
}

// New a log handler
//
// If fileName != "", Add hook to output log to file
//
// If showLog == true, SetOutput(os.Stdout)
func NewLogger(fileName string, showLog bool) (ret *logrus.Logger, err error) {
	os.Mkdir(RUNTIME, os.ModePerm)

	ret = logrus.New()

	if showLog {
		ret.SetOutput(os.Stdout)
		ret.SetLevel(logrus.TraceLevel)
		ret.SetFormatter(
			&prefixed.TextFormatter{
				DisableColors:   false,
				TimestampFormat: "2006-01-02 15:04:05",
				FullTimestamp:   true,
				ForceFormatting: true,
			},
		)
	} else {
		ret.SetOutput(&nullWriter{})
	}

	if fileName != "" {
		var logFile *os.File
		fn := filepath.Join(RUNTIME, fileName)
		logFile, err = os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return
		}

		logfiles = append(logfiles, fn)

		fileHook := lfshook.NewHook(lfshook.WriterMap{
			logrus.InfoLevel:  logFile,
			logrus.WarnLevel:  logFile,
			logrus.ErrorLevel: logFile,
			logrus.FatalLevel: logFile,
			logrus.PanicLevel: logFile,
		}, &logrus.JSONFormatter{})

		ret.AddHook(fileHook)
	}

	return
}

func ShowLogfile() []string {
	return logfiles
}
