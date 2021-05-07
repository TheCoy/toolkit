package app

import "fmt"

type BaseApp interface {
    run() error
    initConfig(strings []string) error
}

//CLI show the entrance for main program
func CLI(args []string, app BaseApp) int {
    if err := app.initConfig(args); err != nil {
        fmt.Println(err)
        return 2
    }

    if err := app.run(); err != nil {
        fmt.Println(err)
        return 3
    }

    return 0
}