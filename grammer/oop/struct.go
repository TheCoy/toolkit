package oop

import "fmt"

type Engine interface {
    start()
    stop()
}

type Car struct {
    Engine
    name       string
    wheelCount int
}

func (c *Car) start() {
    fmt.Println(c.name, "started")
}

func (c *Car) stop() {
    fmt.Println(c.name, "stoped")
}

func (c Car) numsOfWheel() int {
    return c.wheelCount
}

type Mercedes struct {
    Car
}

// func (m *Mercedes) start() {
// 	fmt.Println("mercedes start")
// }

func (m *Mercedes) sayHiToMerkel() {
    fmt.Println("Hi, Merkel")
}

