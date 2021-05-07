package toolkit

import (
    "errors"
    "sync"
    "time"
)

const (
    nodeBits  uint8 = 10
    stepBits  uint8 = 12
    nodeMax   int64 = -1 ^ (-1 << nodeBits)
    stepMax   int64 = -1 ^ (-1 << stepBits)
    timeShift uint8 = nodeBits + stepBits
    nodeShift uint8 = stepBits
)

var Epoch int64 = 1288834974657

type ID int64

type Node struct {
    mu        sync.Mutex
    timeStamp int64
    node      int64
    step      int64
}

func NewNode(node int64) (*Node, error) {
    if node <= 0 || node > nodeMax {
        return nil, errors.New("Node number must between 0 and 1023")
    }
    return &Node{
        timeStamp: 0,
        node:      node,
        step:      0,
    }, nil
}

func (n *Node) Generate() ID {
    n.mu.Lock()
    defer n.mu.Unlock()

    now := time.Now().UnixNano() / 1e6

    if n.timeStamp == now {
        n.step++
        if n.step > stepMax {
            for now <= n.timeStamp {
                now = time.Now().UnixNano() / 1e6
            }
        }
    } else {
        n.step = 0
    }

    n.timeStamp = now

    result := (now-Epoch)<<timeShift | (n.node << nodeShift) | (n.step)

    return ID(result)
}
