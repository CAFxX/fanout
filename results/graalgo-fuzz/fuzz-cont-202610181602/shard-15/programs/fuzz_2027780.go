package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 uint64 = 33
    var v4 uint8 = 127
    var v5 float32 = 0.1
    v6 := GenS1{F0: -9223372036854775808, F1: "foo bar"}
    if false {
        fmt.Println(v5, v4)
        v7 := []int{}
        _ = v7
    }
    v8 := GenS2{F0: 1}
    var v9 float64 = (-2.25 / 1e10)
    v5 -= v5
    switch 42 {
    case 127:
        v10 := GenS2{F0: 983}
        _ = v10
    case 32768:
        var v11 uint = uint(v3)
        _ = v11
    case 2147483647:
        delete(v1, 668)
    default:
        switch ("foo bar" + "foo bar") {
        case "a":
            v3 = 465
        default:
            var v12 uint64 = v3
            _ = v12
        }
    }
    v13 := GenS2{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
