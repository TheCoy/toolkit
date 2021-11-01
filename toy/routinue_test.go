package toy

import (
    "testing"
    "time"
)

func TestRoutinue(t *testing.T) {

    var a, b int = 12412, 12424
    ch1 := make(chan int)
    go sum(ch1, a, b)
    t.Logf("%d + %d = %d\n", a, b, <-ch1)

    suck(pump())
    time.Sleep(1e7)
}
