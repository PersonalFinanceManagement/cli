package main

import (
	"context"
	"fmt"
	"os"

	"github.com/PersonalFinanceManagement/cli/internal"
	"github.com/PersonalFinanceManagement/common/logging"
	cli "github.com/urfave/cli/v3"
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
	(&cli.Command{}).Run(context.Background(), os.Args)
	fmt.Printf("The logger value at the moment is :: %v+\n", logger)
	logger.Info("This is the finance manager!!")
}
