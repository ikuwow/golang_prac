package main

import "fmt"

func main() {
    const name = "ikuwow"
    // name = "degawa" // Error
    fmt.Println(name)

    const (
        sun = iota // 0
        mon // 1
        tue // 2
    )

    fmt.Println(sun,mon,tue)
}

