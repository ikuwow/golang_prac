package main
// 何らかのパッケージに属している必要がある

import "fmt"
// like "stdout"

func main() {

    // var message string
    // message = "hello world"
    // ↓
    // var message = "hello world"
    // ↓
    message := "hello world" // よく出る

    // when define multiple variables simultaneously
    // var a,b int
    // a, b  = 10, 5
    a, b := 10, 15

    fmt.Println(message)

    fmt.Println(a)
    fmt.Println(b)
    // no semicolons
    // tab indent (not spaces) are prefered

    // 変数や関数名は、
    // 頭が大文字だと他のパッケージからでも参照できる
    // 頭が小文字だとこのパッケージの中から参照する

    /* data types
    string "hello"
    int 53 (int8 int16)
    float64 ( 10.2)
    bool true, false
    nil
    */

    // initial values
    var s string // "" empty string!!
    var f bool // false !
    var i int // 0 ! ¥¥

    fmt.Printf("a: %s, b:%f, c:%d¥n", s, f, i)

}
