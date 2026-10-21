package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }

func main() {
    var v1 int32 = 880
    var v2 bool = false
    v3 := make(map[uint32]string)
    var v4 uint32 = 464
    var v5 uint32 = 263
    switch 1 {
    case 2147483647:
        switch 128 {
        case 42:
            v2 = v2
        case 2147483647:
            var v6 uint = 2
            _ = v6
        }
    default:
        v7 := GenS1{F0: 0, F1: "a", F2: 42}
        _ = v7
    }
    var v8 uint = 469
    v5 -= v4
    v5 = 42
    fmt.Println(v5)
    switch "hello" {
    case "hello":
        fmt.Println(v4)
    case "a":
        var v9 int64 = 42
        _ = v9
    default:
        v10 := GenS1{F0: 42, F1: ("" + ("a" + "\n\t")), F2: v8}
        _ = v10
    }
    fmt.Println(v2)
    var v11 float64 = float64(3.14159)
    var v12 uint32 = v4
    v1 += v1
    var v13 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
