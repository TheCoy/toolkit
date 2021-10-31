package mem

import (
    "fmt"
    "runtime"
    "time"
)

func ConcurrentAccessData() {
    runtime.GOMAXPROCS(2)
    var x int
    go func() {
        x = 0
        inTc := time.Tick(200 * time.Millisecond)
        endTc := time.Tick(10 * time.Second)
        for {
            select {
            case <-inTc:
                x = 0
            case <-endTc:
                fmt.Println("gorotinue  exited")
                return
            default:
                x = 0
            }
        }
    }()
    for i := 0; i < 100; i++ {
        time.Sleep(200*time.Millisecond)
        x = 1
        fmt.Println(x)
    }

}
