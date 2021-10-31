package ioc

type IntSet struct {
    data map[int]bool
    undo Undo
}

func (i *IntSet) Add(n int) {
    if i.data[n] {
        i.undo.Add(nil)
    }else {
        i.data[n] = true
        i.undo.Add(func() {delete(i.data, n)})
    }
}

func (i *IntSet) Del(n int) {
    if i.data[n] {
        delete(i.data, n)
        i.undo.Add(func() {i.data[n] = true})
    }else {
        i.undo.Add(nil)
    }
}