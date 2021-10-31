package toy

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

//cancelChannelTest 模拟了原始的取消方式
func cancelChannelTest() {
	resp := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		resp <- 1
		close(resp)
	}()

	select {
	case ret := <-resp:
		fmt.Println("resp finished, ret =", ret)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}

//cancelContextTest 模拟了主动取消
func cancelContextTest(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	resp := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		resp <- 110
		close(resp)
	}()
	select {
	case <-ctx.Done():
		fmt.Println("context canceled")
		return errors.New("canceled")
	case r := <-resp:
		fmt.Println("ret:", r)
		return nil
	}
}

//cancelContextTestV2 模拟了超时取消
func cancelContextTestV2(ctx context.Context) {
	subCtx, subCancel := context.WithTimeout(ctx, 2 * time.Second)
	defer subCancel()

	resp := make(chan struct{}, 1)

	go func() {
		time.Sleep(5 * time.Second)
		resp <- struct{}{}
		close(resp)
	}()

	select {
	case <-ctx.Done():
	    fmt.Println("mainCtx timeout")
	case <-subCtx.Done():
		fmt.Println("subCtx timeout")
	case r := <-resp:
		fmt.Println("result=", r)
	}
}
