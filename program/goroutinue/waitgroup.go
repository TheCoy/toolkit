package goroutinue

import (
    "fmt"
    "math/rand"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

func Entrance(n int) {
    wg := &sync.WaitGroup{}
    ch := make(chan struct{})
    for i := 0; i < n; i++ {
        wg.Add(1)
        go worker(i, wg, ch)
    }

    //time.Sleep(3 * time.Second)
    go func(cgroup chan <- struct{}) {
        signals := make(chan os.Signal, 1)
        signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
        tc := time.Tick(3*time.Second)
        for{
            select{
            case <-tc:
                for i := 0; i < n; i++ {
                    cgroup <- struct{}{}
                }
                return
            case <-signals:
                for i := 0; i < n; i++ {
                    cgroup <- struct{}{}
                }
                return
            default:
                //fmt.Println("===running===")
            }
        }

    }(ch)


    wg.Wait()

    fmt.Println("main finished!")
}

func worker(index int, wg *sync.WaitGroup, signal <-chan struct{}) {
    defer func() {
        wg.Done()
    }()

    rand.Seed(time.Now().UnixNano())

    n := time.Duration(rand.Intn(10))
    //time.Sleep(n*time.Second)
    <-signal
    fmt.Printf("worker[%d] @%dfinished!\n", index, n)

    return
}
