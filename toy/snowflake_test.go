package toy

import "testing"

func TestSnowFlake(t *testing.T) {
    node, err := NewNode(1)
    if err != nil {
        t.Error(err)
        return
    }
    ch := make(chan ID)
    count := 10000
    for i := 0; i < count; i++ {
        go func() {
            ch <- node.Generate()
        }()
    }
    defer close(ch)

    m := make(map[ID]int)
    for i := 0; i < 10000; i++ {
        id := <-ch
        _, ok := m[id]
        if ok {
            t.Logf("ID[%d] is not unique!\n", id)
            return
        }
        m[id]++
    }
    t.Log(m)
}
