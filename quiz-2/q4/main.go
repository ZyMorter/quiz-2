package main

	

import (
    "fmt"
    "sync"
    "time"
)
	
func employee(id int) {
    fmt.Printf("Worker %d started working\n", id)

    time.Sleep(time.Second*10)
    fmt.Printf("Worker %d finish working\n", id)
}

func main() {
//create waitgroup 
    var wg sync.WaitGroup


    for i := 1; i <= 5; i++ {
        wg.Add(1)

        i := i

        go func() {
            defer wg.Done()
            employee(i)
        }()
    }
	
    wg.Wait()


	

}


