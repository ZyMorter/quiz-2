package main

	

import (
    "fmt"
    "time"
)

	

func gocount(starting string) {
    for i := 0; i < 5; i++ {
        fmt.Println(starting, ":", i)
    }
}

	

func main() {

    go gocount("goroutine")

    go func(message string) {
        fmt.Println(message)
    }("starting count")

    time.Sleep(time.Second*10)
    fmt.Println("done")
}
