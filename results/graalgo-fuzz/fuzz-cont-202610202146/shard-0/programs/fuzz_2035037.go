package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint64; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 478
    var v3 string = (("hello" + "foo bar") + "hello")
    v4 := []int{128}
    v2 -= 0
    v3 = ("" + v3)
    fmt.Println(v2)
    v5 := GenS1{F0: 0, F1: 128}
    if len(v4) > 0 {
        v4[0] = 306
    }
    v3 = ("a" + ("hello" + "a"))
    v6 := GenS2{F0: true, F1: (v2 & v2), F2: true}
    fmt.Println(v2)
    if (1000 >= -321) {
        _ = len(v4)
        var v7 uint64 = (v2 * 1)
        v7 -= 42
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        if true {
            fmt.Println(v3)
            var v9 bool = ((!((-(fi8)) == fi8)) || (!true))
            _ = v9
        }
        fmt.Println(v2, fi8, v3)
        var v10 int64 = 42
        _ = fi8
        _ = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
