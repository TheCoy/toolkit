package toy

import (
    "testing"
    "time"
)

func TestSieve(t *testing.T) {
    ch := sieve()
    go func() {
        for {
            t.Logf("prime=%d", <- ch)
        }
    }()
    time.Sleep(2 * time.Second)
}