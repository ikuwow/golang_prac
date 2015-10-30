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

    // algorithmic operations

    x := 10 % 3
    x += 3
    x++ // ++x is not allowed
    fmt.Println(x)

    s := "hello "+"world"
    fmt.Println(s)
    a := true
    b := false
    fmt.Println(a&&b)
    fmt.Println(a||b)

    // pointers
    // pointer operation is not allowed

    variable := 5
    var p_variable *int // pointer able to
    p_variable = &variable
    fmt.Println(p_variable)
    fmt.Println(*p_variable)
}

