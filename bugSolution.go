```go
package main

import (

        "fmt"
        "sync"
        "time"
)

func main() {
        var wg sync.WaitGroup
        var m sync.Mutex
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m.Lock()
                        fmt.Printf("Goroutine %d is working...\n", i)
                        // Simulate some work
                        time.Sleep(time.Millisecond * 100) // Added delay
                        m.Unlock()
                }(i)
        }
        wg.Wait()
}
```