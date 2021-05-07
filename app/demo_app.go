package app

import (
    "flag"
    "fmt"
    "log"
    "os"
    "sync"
    "time"

    "github.com/TheCoy/toolkit/routinepool"
)

type DemoApp struct {
    LogFileName string
    MaxWorker   int
    Times       int64
    *log.Logger
    wg sync.WaitGroup
}

func (app *DemoApp) initConfig(args []string) error {
    fl := flag.NewFlagSet(args[0], flag.ContinueOnError)
    fl.StringVar(&app.LogFileName, "logFile", "./log/demoApp.log", "please input log File name")
    fl.IntVar(&app.MaxWorker, "n", 1, "max  worker number")
    fl.Int64Var(&app.Times, "t", 1000, "times for task cycle")

    if err := fl.Parse(args[1:]); err != nil {
        return err
    }

    logOutput, err := os.OpenFile(app.LogFileName, os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
        return err
    }
    app.Logger = log.New(logOutput, args[0], 1)

    return nil
}

func (app *DemoApp) run() error {
    workPool := routinepool.NewPool(app.MaxWorker)
    go func() {
        defer close(workPool.EntryChannel)
        var i int64
        for i = 0; i < app.Times; i++ {
            app.wg.Add(1)
            workPool.EntryChannel <- app.buildTask(i)
        }
    }()
    workPool.Run()
    app.wg.Wait()

    return nil
}

func (app *DemoApp) buildTask(seq int64) *routinepool.Task {
    task := routinepool.NewTask(func() error {
        time.Sleep(1000 * time.Microsecond)
        fmt.Printf("No[%d] started at %s\n", seq, time.Now().Format("2006-01-02 03:04:05"))
        app.wg.Done()
        return nil
    })

    return task
}
