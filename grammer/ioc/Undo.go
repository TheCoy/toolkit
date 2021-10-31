package ioc

import "fmt"

type Undo []func()

func (u *Undo) Undo() error {
    functions := *u
    if len(functions) == 0 {
        return fmt.Errorf("without undo functions")
    }

    index := len(functions)-1
    if f := functions[index]; f !=  nil {
        f()
        functions[index] = nil
    }

    *u = functions[:index]

    return nil
}

func (u *Undo) Add(f func()) {
    *u = append(*u, f)
}

