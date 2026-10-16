package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint32 = 1
    var v2 bool = true
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]string)
    var v5 uint32 = v1
    var v6 int64 = (v3 ^ 1)
    v7 := []int{127, 0, 42}
    var v8 uint16 = 42
    v2 = (false && v2)
    _ = len(v7)
    v9 := GenS1{F0: -81, F1: (v5 + v1)}
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(v5, v8, v3)
        var v11 uint16 = (v8 % ((v8) | 1))
        _ = fi10
        _ = v11
    }
    switch 1 {
    case 4294967295:
        if (1 != 2147483647) {
            v12 := GenS2{F0: 1, F1: 1}
            _ = v12
        }
    case 42:
        fmt.Println(v8, v8, v6)
    }
    fmt.Println(v1)
    switch 237 {
    case 128:
        _ = len(v7)
    default:
        v2 = (!(1 == 128))
    }
    var v13 int8 = 22
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int16 = 1000
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
