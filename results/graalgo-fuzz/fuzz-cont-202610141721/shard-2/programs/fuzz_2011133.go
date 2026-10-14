package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 bool }

func main() {
    var v1 int64 = -33
    var v2 uint32 = 65535
    var v3 int32 = 0
    var v4 int16 = -1
    var v5 string = ("foo bar" + ("a" + "a"))
    v2 += (^(v2))
    v2 = v2
    if true {
        fmt.Println(v4, v2)
    } else {
        switch 1 {
        case -1:
            var v6 uint = 18446744073709551615
            _ = v6
        case -832:
            fmt.Println(v3, v1, v4)
        default:
            v7 := GenS1{F0: (v4 ^ 1), F1: v1, F2: (false && true)}
            _ = v7
        }
        fmt.Println(v1, v1, v3)
    }
    if ((!false) || true) {
        v8 := []int{138, 47}
        v9 := GenS1{F0: v4, F1: v1, F2: true}
        _ = v8
        _ = v9
    }
    var v10 int = 9223372036854775807
    var v11 int64 = v1
    var v12 string = v5
    v12 = "\n\t"
    var v13 uint32 = (v2 | (v2 * 0))
    var v14 int64 = v11
    fmt.Println(v3, v4, v1)
    v1 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
