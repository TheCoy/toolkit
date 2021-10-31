package toy

import (
    "fmt"
    "runtime"
    "time"
)

const (
    FIRST = "What the"
    SECOND = "F*ck"
)
func ConcurrentWriteString(){
    var s string
    runtime.GOMAXPROCS(2)
    go func(){
        var i int
        for {
            i = 1 - i
            if i == 0 {
                s = FIRST
            }else {
                s = SECOND
            }
            time.Sleep(10)
        }
    }()

    for {
        if s == FIRST || s == SECOND {

        }else {
            fmt.Println(s)
        }
        time.Sleep(10)
    }
}
