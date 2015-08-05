package main

import "fmt"

func main() {

    // score := 83

    if score := 83; score > 80 { // no brackets
        fmt.Println("Good")
    } else {
        fmt.Println("Bad")
    }

    fmt.Println(score) // undefined
}
