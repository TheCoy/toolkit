package toy

import (
    "fmt"
)


func sum(ch chan int, a, b int) {
    res := a + b
    ch <- res
}

func pump() chan int {
    ch := make(chan int)
    go func() {
        for i := 0; ; i++ {
            ch <- i
        }
    }()

    return ch
}

func suck(ch chan int) {
    go func() {
        for v := range ch {
            fmt.Println("====", v)
        }
    }()
}

