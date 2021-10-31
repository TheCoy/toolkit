package toy

import (
	"fmt"
	"sync"
	"time"
)

type myLock struct {
	sync.Mutex
}

var status int
var lock myLock

func first() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("first")
	status = 1

}

func second() {
	for {
		lock.Lock()
		t := status
		lock.Unlock()

		if t == 1{
			status = 2
			fmt.Println("second")
			break
		}
	}
}

func third() {
	for {
		lock.Lock()
		t := status
		lock.Unlock()

		if t == 2{
			status = 3
			fmt.Println("third")
			break
		}
	}
}

func SyncRunDemo() {
	//_ = trace.Start(os.Stderr)
	//defer trace.Stop()
	status = 0
	go third()
	go first()
	go second()


	time.Sleep(3*time.Second)
}