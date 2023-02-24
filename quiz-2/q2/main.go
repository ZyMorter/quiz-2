package main
  
import "fmt"
  
// checking channel if its open 
func channel(mychnl chan string) {
  
    for a := 0; a < 2; a++ {
        mychnl <- "sharing data "
    }
    close(mychnl)
}
  

func main() {
  
    // Creating a channel
    c := make(chan string)

    go channel(c)
  
    for {
        res, ok := <-c
        if ok == false {
            fmt.Println("Channel Close ", ok)
            break
        }
        fmt.Println("Channel Open ", res, ok)
    }
}