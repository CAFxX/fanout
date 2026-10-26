package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]bool)
    var v3 int16 = 42
    v4 := []int{-1, -785}
    var v5 int8 = 100
    fmt.Println(v1)
    fmt.Println(v1, v3)
    v6 := GenS1{F0: 100}
    v1 = v1
    switch 337 {
    case 127:
        if (152 < 2) {
            v7 := GenS2{F0: 104, F1: 750}
            _ = v7
        } else {
            var v8 int64 = 575
            var v9 uint16 = (v1 | 995)
            fmt.Println(v5, v1)
            _ = v8
            _ = v9
        }
    }
    v10 := GenS1{F0: -1}
    v11 := v2[577]
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v1 = v1
    switch 128 {
    case 32768:
        v12 := v2[65535]
        _ = v12
    }
    if v11 {
        var v13 uint = 0
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
