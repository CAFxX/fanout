package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
func genFn0(p0 uint16) uint64 {
    p0 += (^(p0))
    p0 = p0
    _ = p0
    return 860
}

func main() {
    var v1 bool = true
    var v2 uint8 = 1
    var v3 uint32 = 42
    var v4 int8 = 100
    var v5 float64 = (-2.25 + 1.5)
    v5 = v5
    v6 := []int{127, 1}
    v1 = v1
    var v7 uint32 = (v3 / (((v3 ^ 86)) | 1))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v4 += v4
    v8 := GenS1{F0: 0, F1: -1}
    for fi9 := 0; fi9 < 13; fi9++ {
        v3 -= v7
        _ = fi9
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
