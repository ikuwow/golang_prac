package main

import "fmt"

func main() {

    // score := 83

    if score := 83; score > 80 { // no round brackets
        fmt.Println("Good")
    } else {
        fmt.Println("Bad")
    }

    // fmt.Println(score) // undefined

    // switch statement

    signal := "asdf" // double quatation, not single quatation

    switch signal {
    case "red":
        fmt.Println("Stop")
    case "yellow":
        fmt.Println("Caution")
    case "green", "blue":
        fmt.Println("Go") // "break;" is not needed
    default:
        fmt.Println("Invalid signal")
    }

    score2 := 79

    switch { // if-else alternate
    case score2 > 80:
        fmt.Println("Good")
    default:
        fmt.Println("Bad")
    }


    /** for statement (no while) **/

    for i:=0; i<10; i++ { // no brackets
        if i == 3 {
            continue
        }
        if i == 8 {
            break
        }
        fmt.Println(i)
    }

    j:=0
    for j < 10 { // like while
        fmt.Println(j)
        j++
    }

    for { // infinite loop
        fmt.Println("test")
        if true {
            break
        }
    }

    // range

    ss := []int{2,3,8}
    for _, v:= range ss { // _ blank
        fmt.Println(v)
    }

    map1 := map[string]int{"ikuwow": 52, "degawa": 33}
    for k, v :=range map1 {
        fmt.Println(k,v)
    }

}
