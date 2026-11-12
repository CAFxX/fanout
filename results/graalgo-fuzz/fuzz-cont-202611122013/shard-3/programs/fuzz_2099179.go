package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("" + "\n\t")
    var v3 int32 = 42
    _ = len(v1)
    fmt.Println(v2, v3)
    var v4 int16 = 1000
    v5 := GenS1{F0: 2, F1: 2}
    v6 := GenS2{F0: 2, F1: 2147483647}
    v3 = v3
    v7 := GenS2{F0: 2, F1: v3}
    v8 := []int{}
    if false {
        fmt.Println(v2)
        var v9 uint = 2
        fmt.Println(v9, v3)
        _ = v9
    } else {
        var v10 uint = 1
        switch 839 {
        case 127:
            if (0.1 >= 1.5) {
                var v11 int64 = -798
                fmt.Println(v4, v3, v11)
                _ = v11
            } else {
                var v12 uint = 0
                var v13 int8 = 100
                _ = v12
                _ = v13
            }
        case 44:
            fmt.Println(v3, v3, v4)
        default:
            v2 = ((("foo bar" + v2) + "foo bar") + v2)
        }
        v3 -= v3
        _ = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
