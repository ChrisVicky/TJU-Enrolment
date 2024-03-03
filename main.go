package main

import (
	"enrollment/logger"
	"enrollment/routine"
)

func main() {
	// Setup Logger
	err := logger.Setup(true)
	if err != nil {
		panic(err)
	}

	r, err := routine.NewRoutine()
	if err != nil {
		logger.Panic(err)
	}

	r.Run()
}
