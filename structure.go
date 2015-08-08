package main

import "fmt"

// structure
type user struct {
    name string
    score int
}


// u: receiver
func (u user) show() { // connect to user structure
    // value reference
    fmt.Printf("name:%s, score:%d\n",u.name,u.score)
}

func (u *user) hit() { // pointer reference
    u.score++
}


// interface
type greeter interface {
    greet()
}

type japanese struct {}
type american struct {}

func (j japanese) greet() {
    fmt.Println("こんにちは！")
}

func (a american) greet() {
    fmt.Println("Hello!")
}


// empty interface (all type are compatible with this empty interface)

// it can check all type whether japanese or not japanese
func sayNationality(t interface{}) { // can receive all type
    // type assersion
    _, ok := t.(japanese) // value, is_japanese?
    if ok {
        fmt.Println("I am Japanese")
    } else {
        fmt.Println("I am not Japanese")
    }

    switch t.(type) {
    case japanese:
        fmt.Println("I am Japanese!")
    default:
        fmt.Println("I am not Japanese!")
    }

}

func main() {
    u := new(user) // pointer
    // (*u) := new(user) same meaning
    fmt.Println(u)

    u1 := user{"ikuwow",3}
    u2 := user{ name:"degawa", score:2}
    fmt.Println(u1)
    fmt.Println(u2)

    u1.hit()
    u1.show()

    // can be used as same type with same interface
    greeters := []greeter{japanese{}, american{}}

    for _, greeter := range greeters {
        greeter.greet()
        sayNationality(greeter)
    }


}
