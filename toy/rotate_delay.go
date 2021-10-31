package toy

import (
    "log"
    "time"
)

const (
    up = iota
    down
)

const (
    minDelay  = 35
    maxDelay  = 55
)

var delay int


func RotateDelay() {
    direction := up
    for {
        if direction == up && delay == maxDelay {
            direction = down
        }
        if direction == down && delay == minDelay {
            direction = up
        }

        if direction == up {
            delay += 1
        } else {
            delay -= 1
        }

        time.Sleep(100 * time.Millisecond)
        log.Printf("setting delay to %v", delay)
    }
}
