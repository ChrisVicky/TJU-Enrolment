package logger_test

import (
	"enrolment/logger"
	"fmt"
)

func ExampleSetup() {
	if err := logger.Setup(true); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Ok")
	// Output:
	// Ok
}

func ExampleNewLogger() {
	if _, err := logger.NewLogger("test.log", true); err != nil {
		fmt.Printf("Error %+v\n", err)
		return
	}
	fmt.Println(logger.ShowLogfile())
	// Output:
	// [runtime/test.log]
}
