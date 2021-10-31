package toy

import (
    "testing"
    "time"
)

func TestPingPong(t *testing.T) {
    c := make(chan interface{})
    go pong(c)
    go ping(c)

    time.Sleep(1*time.Second)
}
