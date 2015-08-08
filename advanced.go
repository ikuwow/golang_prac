package main

// import multiple
import (
    "fmt"
    "time"
)

func heavyTask(result chan string) {
    time.Sleep(time.Second * 2)
    fmt.Println("Done! It was hard...")
    // cannot return!
    result <- "Task1 result"
}

func easyTask(result chan string) {
    time.Sleep(time.Second * 1)
    fmt.Println("OK, ... Done!")
    // result <- "Task2 result"
}

func main() {

    result := make(chan string) // chanel of type string

    // go routine ( parallel processing)
    go heavyTask(result) // background
    go easyTask(result) // background

    // if result is empty, wait
    fmt.Println(<-result)

    time.Sleep(time.Second * 4)
}
