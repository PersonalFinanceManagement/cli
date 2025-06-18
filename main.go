package main

import (
	"fmt"

	"github.com/PersonalFinanceManagement/cli/internal"
	"github.com/PersonalFinanceManagement/common/logging"
)

var logger logging.Logger

func init() {
	var err error
	logger, err = logging.NewLogger(logging.DefaultLoggerConfig())
	if err != nil {
		panic(fmt.Errorf("unable to setup the logger. %w", err))
	}
	logger.Debug("The logger is successfully initialized")
}

func main() {
	internal.PrintHi()
	fmt.Printf("The logger value at the moment is :: %v+\n", logger)
	logger.Info("This is the finance manager!!")
}
