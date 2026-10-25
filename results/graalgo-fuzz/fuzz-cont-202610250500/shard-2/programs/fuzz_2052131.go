package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 uint = 0
    var v2 int = 699
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    v2 = 4294967295
    fmt.Println(v1)
    var v5 int8 = 127
    var v6 uint8 = 145
    fmt.Println(v2, v5, v1)
    var v7 uint32 = 0
    _ = len(v3)
    v8 := []int{v2}
    fmt.Println(v2)
    if true {
        v7 += v7
        if (!((!((1.5 / 1e10) < -2.25)) || (((1.5 - 1.5) > 0.0) && ((uint32(v6) < (^(v7))) || (uint64(v6) <= 230))))) {
            v7 = (v7 / ((564) | 1))
        }
        var v9 uint16 = uint16(v7)
        _ = v9
    }
    switch "hello" {
    case "":
        var v10 uint32 = v7
        _ = v10
    default:
        v2 = (v2 - -2147483648)
    }
    v11 := GenS1{F0: v6, F1: 1}
    if true {
        var v12 string = "hello"
        v13 := GenS1{F0: (^(v6)), F1: (v6 - (v6 ^ v6))}
        _ = v12
        _ = v13
    } else {
        var v14 int = 4294967295
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
