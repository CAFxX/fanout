package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint32 }
func genFn0() bool {
    var v5 int = 641
    v6 := GenS1{F0: true, F1: 100}
    _ = v5
    _ = v6
    return (17 > -1)
}

func main() {
    var v1 int8 = -125
    var v2 string = ("hello" + "foo bar")
    var v3 uint = 571
    var v4 int = 152
    v4 = v4
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    switch v2 {
    case "hello":
        v3 -= 18446744073709551615
    case "foo bar":
        for fi9 := 0; fi9 < 1; fi9++ {
            v3 = v3
            v4 = v4
            _ = fi9
        }
    default:
        fmt.Println(v2, v3, v1)
    }
    var v10 uint32 = 0
    v11 := []int{(v4 / (((v4 / ((4294967295) | 1))) | 1)), (v4 % ((v4) | 1)), (-(v4))}
    var v12 string = (((("\n\t" + ((v2 + v2) + ("hello" + (v2 + v2)))) + "hello") + v2) + v2)
    v3 = (^(v3))
    v2 = v12
    switch v12 {
    case "":
        v12 = v2
    default:
        v3 = 0
    }
    v13 := GenS2{F0: 42}
    fmt.Println(v3, v2, v10)
    v14 := GenS1{F0: (false && (!(v1 < 42))), F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
