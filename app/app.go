package app

import (
	"flag"
	"fmt"
)

type appEnv struct {
	logFileName   string
	inputFileName string
	maxWorker     int
	maxQueue      int
}

func (app *appEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet(args[0], flag.ContinueOnError)
	fl.IntVar(&app.maxWorker, "worker", 100, "max Worker number(default 100)")
	fl.IntVar(&app.maxQueue, "queue", 10000, "max Queue number(default 10000)")
	fl.StringVar(&app.logFileName, "logFile", "./auditApp.log", "path to logFile")
	fl.StringVar(&app.inputFileName, "inputFile", "./inputFile.txt", "path to inputFile")

	if err := fl.Parse(args[1:]); err != nil {
		return err
	}

	return nil
}

func (app *appEnv) run() error {
	fmt.Printf("logFileName=%s\n", app.logFileName)
	fmt.Printf("inputFileName=%s\n", app.inputFileName)

	return nil
}

//CLI show the entrance for main program
func CLI(args []string) int {
	var app appEnv
	if err := app.fromArgs(args); err != nil {
		return 2
	}

	if err := app.run(); err != nil {
		return 1
	}

	return 0
}
