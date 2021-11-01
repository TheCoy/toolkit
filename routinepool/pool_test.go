package routinepool

import (
    "fmt"
    "testing"
    "time"
)

func TestRoutinePool(t *testing.T) {
    task := NewTask(func() error {
        fmt.Println(time.Now())
        return nil
    })

    pool := NewPool(2)
    go func() {
        for {
            pool.EntryChannel <- task
        }
    }()

    pool.Run()
}
