package oop

import (
    "fmt"
    "testing"
)

func TestNameInstance(t *testing.T) {
    c := China{WithName{Name: "zhangsan"}}
    j := Japan{WithName{Name: "xiba"}}

    PrintName(c)
    PrintName(j)

    PrintName(c)
    PrintName(j)
}

func TestComplete(t *testing.T) {
    s := Square{len: 4}
    var _ Shape = (*Square)(nil)
    fmt.Println(s.Sides())
    fmt.Println(nil)
}

