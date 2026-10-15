package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16; F1 int32; F2 string }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 uint64 = 18446744073709551615
    var v4 uint = 255
    fmt.Println(v3, v3)
    fmt.Println(v3, v4)
    if true {
        var v5 uint64 = (v3 << ((v4) & 63))
        switch 127 {
        case 32767:
            fmt.Println(v3, v4, v3)
        case -2147483648:
            v5 = (v5 & (v3 + v3))
        default:
            v4 -= v4
        }
        _ = v5
    } else {
        v6 := GenS1{F0: 0, F1: 106}
        var v7 uint64 = v3
        _ = v6
        _ = v7
    }
    fmt.Println(v3)
    delete(v2, "hello")
    v3 -= (^(v3))
    v3 = 18446744073709551615
    if ((((0.1 > 0.0) && false) || (123 >= 127)) && (((!true) || false) || (!false))) {
        v8 := GenS1{F0: -226, F1: 109}
        fmt.Println(v4)
        v9 := GenS2{F0: 0, F1: 42, F2: ("hello" + (("hello" + "") + ("a" + (("a" + ("hello" + ("hello" + ""))) + "hello"))))}
        _ = v8
        _ = v9
    } else {
        v10 := []int{}
        _ = v10
    }
    fmt.Println(v4, v3)
    var v11 string = "hello"
    v12 := GenS2{F0: 55, F1: -141, F2: v11}
    var v13 int64 = -459
    v14 := GenS2{F0: 0, F1: -2147483648, F2: v11}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
