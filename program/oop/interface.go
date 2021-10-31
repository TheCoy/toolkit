package oop

import "fmt"

type WithName struct {
    Name string
}

type Printable interface {
    PrintName()
}

func (w WithName) PrintName() {
    fmt.Println("Name:", w.Name)
}

type Stringable interface {
    ToString() string
}

type China struct {
    WithName
}

func (c China) ToString() string {
    return "Name@" + c.Name
}

type Japan struct {
    WithName
}

func (j Japan) ToString() string {
    return "Name@" + j.Name
}

func PrintName(s Stringable) {
    fmt.Println(s.ToString())
}


