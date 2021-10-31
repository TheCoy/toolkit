package publish

import "sync"

type Registrable interface {
    TopicFunc(v interface{}) bool
    ConsumeFunc(ch chan interface{}, wg *sync.WaitGroup)
}

