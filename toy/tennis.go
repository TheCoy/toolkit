package toy

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func MockTennisBallGame() {

    rand.Seed(time.Now().UnixNano())
    court := make(chan int)

    wg := &sync.WaitGroup{}

    wg.Add(2)

    go player(wg, "bob", court)
    go player(wg, "mike", court)

    court <- 1
    wg.Wait()

}

func player(wg *sync.WaitGroup, name string, ch chan int) {
    defer wg.Done()

    for  {
        ball, ok := <- ch
        if !ok {
            fmt.Printf("%s winned!\n", name)
            return
        }
        n := rand.Intn(1000)
        if n % 11 == 0 {
            fmt.Printf("%s lossed\n", name)
            close(ch)

            return
        }

        fmt.Printf("player[%s] hit ball[%d]\n", name, ball)
        ball++

        ch <- ball
    }
}
