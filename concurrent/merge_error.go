package concurrent

import "sync"

func Merge(errors ...<-chan error) <-chan error {
    mergedChan := make(chan error)
    var wg sync.WaitGroup
    wg.Add(len(errors))
    go func() {
        wg.Wait()
        close(mergedChan)
    }()
    for i := range errors {
        go func(ch <-chan error) {
            for e := range ch {
                if e != nil {
                    mergedChan <- e
                }
            }
            wg.Done()
        }(errors[i])
    }

    return mergedChan
}
