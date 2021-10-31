package oop

type Shape interface {
    Sides() int
    Area() int
}

type Square struct {
    len int
}

func (s *Square) Sides() int{
    return s.len
}

func (s *Square) Area() int {
    return s.len * s.len
}