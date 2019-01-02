package main

import (  
    "fmt"
    "time"
)

func thread1(done chan int) {  
    fmt.Println("hello go routine is going to sleep")
    time.Sleep(4 * time.Second)
    fmt.Println("hello go routine awake and going to write to done")
    done <- 1

    // close(done)
}

func thread2(chnl chan int) {  
    for i := 0; i < 10; i++ {
        time.Sleep(1 * time.Second)
        chnl <- i+1000
    }
    // close(chnl)
}

func thread3(chnl chan int) {  
    for {
        time.Sleep(5 * time.Second)
        chnl <- 1
    }
}

// func main() {
//     done := make(chan int)
//     fmt.Println("Main going to call hello go goroutine")
//     go thread1(done)
//     go thread2(done)
//     // go thread3(done)


//     for {
//         v,ok := <- done
//         fmt.Println("Received ",v,ok)
//         if(ok==false) {
//             break
//         }
//     }
//     // check := <- done
//     time.Sleep(1* time.Second)
//     fmt.Println("Main received data")
// }