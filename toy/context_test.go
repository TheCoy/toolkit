package toy

import (
    "context"
    "sync"
    "testing"
    "time"
)

func TestContext(t *testing.T) {
    cancelChannelTest()

    //cancelContextTest
    wg := new(sync.WaitGroup)
    ctx, cancel := context.WithCancel(context.Background())
    wg.Add(1)
    go cancelContextTest(ctx, wg)
    time.Sleep(1 * time.Second)
    cancel()
    wg.Wait()

    //cancelContextTestV2
    ctx2, cancel2 := context.WithTimeout(context.Background(), 1 * time.Second)
    defer cancel2()
    cancelContextTestV2(ctx2)
}
