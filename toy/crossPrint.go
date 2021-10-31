package toy

import (
	"fmt"
	"time"
)

func RunCrossPrint() {
	c := make(chan interface{})
	arrA := []interface{}{1,2,3,4,5}
	arrB := []interface{}{"a", "b", "c", "d", "e","f"}

	go func() {
		for _, v := range arrA {
			<- c
			fmt.Print(v)
			c <- struct {

			}{}
		}
	}()

	go func() {
		for _, v := range arrB {
			c <- struct {

			}{}
			fmt.Print(v)
			<- c
		}
	}()

	time.Sleep(1*time.Second)
}