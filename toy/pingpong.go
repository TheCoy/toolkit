package toy

import "fmt"

func ping(c chan interface{}) {
    for _, v := range []int{1,2,3,4,5} {
        c <- v
        fmt.Println("ping", <-c)
    }
}

func pong(c chan interface{}) {
    for _, v := range []string{"a", "b", "c", "d", "e"} {
        fmt.Println("pong", <-c)
        c <- v
    }
}
