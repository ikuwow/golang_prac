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

    // Arrays

    var a [5]int // a[0]-a[4]
    a[2] = 4
    a[4] = 3
    fmt.Println(a, a[2])

    b := [3]int{1,2,4}
    fmt.Println(b, len(b))

    // a is not pointer like c
    // 値渡し

    // slice: reference type
    // 配列への参照

    s := b[1:]
    fmt.Println(s)
    fmt.Println(len(s))
    fmt.Println(cap(s))

    // sss: = make([]int, 3) // [0,0,0] slice
    // sss: = []int{1, 3, 5} // [0,0,0] slice
    // sss: =

    // fmt.Println(sss)

    // map (hash)

    // m := make(map[string]int) // string => int
    // m["degawa"] = 100
    // m["ikuwow"] = 80

    m := map[string]int{"degawa": 100, "ikuwow": 80}
    fmt.Println(m)
    fmt.Println(len(m))
    delete(m,"degawa") // if nothing, no action
    fmt.Println(m)

    value, exists := m["degawa"]
    fmt.Println(value,exists)
    value2, exists2 := m["ikuwow"]
    fmt.Println(value2,exists2)

}

