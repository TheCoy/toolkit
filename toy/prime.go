package toy

func generate() chan int {
    ch := make(chan int)
    go func(){
        for i := 2;;i++ {
            ch <- i
        }
    }()

    return ch
}

func filter(ch chan int, prime int) chan int {
    res := make(chan int)
    go func() {
        for {
            i := <-ch
            if i % prime != 0 {
                res  <- i
            }
        }
    }()

    return res
}

func sieve() chan int {
    output := make(chan int)

    go func() {
        ch := generate()
        for {
            prime := <-ch
            ch = filter(ch, prime)
            output <- prime
        }

    }()

    return output
}
