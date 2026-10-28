package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    var v1 bool = (-2147483648 > 0)
    var v2 int = 32767
    var v3 int32 = 2147483647
    var v4 uint16 = 32767
    v5 := make(map[uint32]uint32)
    v5[2147483648] = 2147483647
    fmt.Println(v2)
    if (100 != 127) {
        var v6 uint16 = (v4 / (((v4 * (v4 & (v4 & v4)))) | 1))
        v7 := []int{}
        v8 := GenS1{F0: 1000, F1: 350}
        _ = v6
        _ = v7
        _ = v8
    } else {
        v5[42] = 0
        var v9 string = "\n\t"
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 float64 = 0.0
        _ = fi10
        _ = v11
    }
    v3 -= 1
    v12 := GenS1{F0: 42, F1: 860}
    v13 := []int{v2, 128}
    fmt.Println(v2, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
