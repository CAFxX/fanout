package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint32 = 42
    var v2 uint32 = v1
    v3 := []uint32{4294967295, v1}
    v1 -= (v2 >> ((434) & 63))
    v1 = 636
    for fi4 := 0; fi4 < 18; fi4++ {
        v2 = v1
        _ = fi4
    }
    var v5 string = "hello"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 += 2
    switch -354 {
    case -1:
        for fi6 := 0; fi6 < 12; fi6++ {
            v3 = append(v3, 500)
            v7 := GenS1{F0: 524, F1: 938}
            _ = fi6
            _ = v7
        }
    }
    if ((false && (273 == 42)) && (0.1 == 1e10)) {
        fmt.Println(v1, v1, v5)
    } else {
        v5 = v5
        if len(v3) > 0 {
            v3[0] = (v1 & (^(v2)))
        }
        v2 = (v1 | 251)
    }
    var v8 float32 = (3.14159 / 3.14159)
    var v9 int64 = -1
    var v10 float32 = ((float32(0.0) / v8) + v8)
    fmt.Println(v10)
    _ = len(v3)
    fmt.Println(v5, v1, v9)
    var v11 int32 = -1
    fmt.Println(v9, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v8
    _ = v9
}
