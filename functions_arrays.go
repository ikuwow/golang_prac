package main

import "fmt"

func hi(name string) {
    fmt.Println("Hi, "+name+"!")
}

/*
func getHiMessage(name string) string {
    message := "Hi, "+name+"!"
    return message
}
*/
func getHiMessage(name string) (message string) {
    // "message" is already defined, and automatically returned
    message = "Hi, "+name+"!"
    return
}


func main() {
    fmt.Println("func!")
    hi("ikuwow")
    fmt.Println(getHiMessage("ikuwow"))

    // alternate function defining
    fff := func(a, b int) (int, int) {
        return b, a
    }

    fmt.Println(fff(3,5))

}
