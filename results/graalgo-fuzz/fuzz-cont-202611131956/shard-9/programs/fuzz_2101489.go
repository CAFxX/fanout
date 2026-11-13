package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }

func main() {
    var v1 float32 = (0.0 / (0.1 + 0.0))
    var v2 int64 = -183
    var v3 uint8 = 121
    var v4 int32 = -384
    var v5 uint32 = 0
    v6 := []int{-811, 164, 42, -444}
    var v7 uint32 = 207
    v5 = v5
    var v8 int64 = v2
    v9 := GenS1{F0: 0, F1: 100, F2: (v4 / ((963) | 1))}
    var v10 uint = 4294967295
    switch "foo bar" {
    case "":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        var v11 uint64 = 1
        _ = v11
    }
    fmt.Println(v7, v2, v2)
    fmt.Println(v1, v3, v10)
    if false {
        v3 = 128
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(v3)
    }
    v10 -= (v10 % ((303) | 1))
    var v12 int64 = -1
    var v13 int32 = (^(v4))
    v4 = -1
    var v14 uint = (v10 + uint(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
