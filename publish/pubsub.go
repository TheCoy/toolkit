package publish

import (
    "sync"
    "time"
)

type (
    subscriber chan interface{}
    topicFunc func(v interface{}) bool
    consumeFunc func(sub chan interface{}, wg *sync.WaitGroup)
)

//发布者
type Publisher struct {
    m           sync.RWMutex
    buffer      int
    timeout     time.Duration
    subscribers map[subscriber]topicFunc
}

func NewPublisher(publishBuffer int, publishTimeout time.Duration) *Publisher {
    return &Publisher{
        buffer:      publishBuffer,
        timeout:     publishTimeout,
        subscribers: make(map[subscriber]topicFunc),
    }
}

func (p *Publisher) AddSubscriber(topic topicFunc) chan interface{} {
    ch := make(chan interface{}, p.buffer)

    p.m.Lock()
    p.subscribers[ch] = topic
    p.m.Unlock()

    return ch
}

func (p *Publisher) AddSubscriberWithConsume(topic topicFunc, consume consumeFunc, consumeCount int) *sync.WaitGroup {
    ch := make(chan interface{}, p.buffer)

    p.m.Lock()
    p.subscribers[ch] = topic
    p.m.Unlock()

    var wg sync.WaitGroup
    for i := 0; i < consumeCount; i++ {
        wg.Add(1)
        go consume(ch, &wg)
    }

    return &wg
}

func (p *Publisher) AutoRegister(r Registrable, count int) *sync.WaitGroup {
    ch := make(chan interface{}, p.buffer)

    p.m.Lock()
    p.subscribers[ch] = r.TopicFunc
    p.m.Unlock()

    var wg sync.WaitGroup
    for i := 0; i < count; i++ {
        wg.Add(1)
        go r.ConsumeFunc(ch, &wg)
    }

    return &wg
}

func (p *Publisher) EvictSubscriber(sub subscriber) {
    p.m.Lock()
    defer p.m.Unlock()

    delete(p.subscribers, sub)
    close(sub)
}

func (p *Publisher) Publish(v interface{}) {
    p.m.RLock()
    defer p.m.RUnlock()

    var wg sync.WaitGroup
    for sub, topic := range p.subscribers {
        wg.Add(1)
        go p.sendTopic(sub, topic, v, &wg)
    }

    wg.Wait()
}

func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
    defer wg.Done()

    if topic != nil && !topic(v) {
        return
    }

    select {
    case sub <- v:
    case <-time.After(p.timeout):
    }
}

func (p *Publisher) Close() {
    p.m.Lock()
    defer p.m.Unlock()

    for sub := range p.subscribers {
        delete(p.subscribers, sub)
        close(sub)
    }
}
