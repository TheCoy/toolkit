package designpattern

import (
	"sync"
	"sync/atomic"
)


var initialized uint32
var mu sync.Mutex

type SingleTon struct{}

var instance *SingleTon

func GetInstance() *SingleTon {
	if atomic.LoadUint32(&initialized) == 1{
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &SingleTon{}
		atomic.AddUint32(&initialized, 1)
	}

	return instance
}


var once sync.Once

func GetInstanceV2() *SingleTon {
	once.Do(func() {
		instance = &SingleTon{}
		atomic.AddUint32(&initialized, 1)
	})

	return instance
}