package toy

import (
    "testing"
    "time"
)

func TestRotateDelay(t *testing.T) {
    go RotateDelay()
    time.Sleep(30 * time.Second)
}