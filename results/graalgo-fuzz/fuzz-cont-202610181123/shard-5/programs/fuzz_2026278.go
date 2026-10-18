package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint8 = 42
    var v2 int8 = 100
    v3 := make(map[string]uint32)
    v4 := []int{32767}
    var v5 int64 = 1
    v5 -= (v5 % ((v5) | 1))
    if (32767 == 2) {
        v6 := GenS1{F0: false}
        _ = v6
    } else {
        v2 = (v2 >> ((267) & 63))
    }
    v2 = (v2 / ((0) | 1))
    switch -220 {
    case -549:
        var v7 uint64 = 18446744073709551615
        _ = v7
    case -497:
        v2 -= v2
    default:
        fmt.Println(v1)
    }
    v8 := GenS1{F0: (!false)}
    var v9 int = 1
    v10 := GenS1{F0: true}
    fmt.Println(v5, v1)
    if false {
        fmt.Println(v9, v5, v2)
        var v11 int16 = 32767
        v2 -= (v2 % (((-(v2))) | 1))
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
