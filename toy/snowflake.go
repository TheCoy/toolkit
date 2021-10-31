package toy

import (
    "sync"
    "time"

    "github.com/pkg/errors"
)

const (
    nodeBits  uint8 = 10
    stepBits  uint8 = 12
    nodeMax   int64 = -1 ^ (-1 << nodeBits)
    stepMax   int64 = -1 ^ (-1 << stepBits)
    timeShift       = nodeBits + stepBits
    nodeShift       = stepBits
)

type Node struct {
    mu        sync.Mutex
    timestamp int64
    node      int64
    step      int64
}

var Epoch int64 = 1288834974657

type ID int64

func NewNode(n int64) (*Node, error) {
    if n < 0 || n > nodeMax {
        return nil, errors.New("invalid node number")
    }

    return &Node{
        timestamp: 0,
        node:      n,
        step:      0,
    }, nil
}

func (n *Node) Generate() ID {
    n.mu.Lock()
    defer n.mu.Unlock()


    now := time.Now().UnixNano() / 1e6
    if n.timestamp < now {
        n.step = 0
        n.timestamp = now
    }

    if n.step > stepMax {
        for {
            now = time.Now().UnixNano() / 1e6
            if now > n.timestamp {
                n.step = 0
                n.timestamp = now
                break
            }
        }
    }else{
        n.step++
    }

    result := ID((n.timestamp-Epoch) << timeShift | n.node << nodeShift | n.step)

    return result
}
