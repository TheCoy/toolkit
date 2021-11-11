package toy

import (
    "context"
    "sync"
    "testing"
    "time"
)

func TestCancelChannelTest(t *testing.T) {
    cancelChannelTest()
}

func TestCancelContextTest(t *testing.T) {
    wg := new(sync.WaitGroup)
    ctx, cancel := context.WithCancel(context.Background())
    wg.Add(1)

    go cancelContextTest(ctx, wg)
    time.Sleep(3 *time.Second)
    cancel()

    wg.Wait()
}

func TestCancelContextTestV2(t *testing.T) {
    ctx2, cancel2 := context.WithTimeout(context.Background(), 3 * time.Second)
    defer cancel2()
    cancelContextTestV2(ctx2)
}