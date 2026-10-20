package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    var v1 int32 = 818
    var v2 bool = false
    var v3 int = 32768
    var v4 uint8 = 2
    v3 = v3
    v4 -= v4
    for fi5 := 0; fi5 < 15; fi5++ {
        v1 -= (-(v1))
        _ = fi5
    }
    if (v1 == v1) {
        for fi6 := 0; fi6 < 2; fi6++ {
            v2 = v2
            v7 := []int{-1, (fi6 + v3), (-(fi6)), 32767}
            _ = fi6
            _ = v7
        }
    }
    v8 := []int{v3, 1}
    v9 := GenS1{F0: v3, F1: -92}
    fmt.Println(v1)
    _ = len(v8)
    v10 := GenS1{F0: 127, F1: 116}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
