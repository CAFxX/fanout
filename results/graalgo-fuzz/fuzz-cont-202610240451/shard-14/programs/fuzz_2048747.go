package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int8 }

func main() {
    var v1 string = "hello"
    var v2 float32 = 0.0
    var v3 int = 2147483647
    var v4 float32 = 3.14159
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(v3)
        v7 := GenS1{F0: 155}
        _ = fi6
        _ = v7
    }
    switch (v1 + ((v1 + ("hello" + "hello")) + ("\n\t" + (v1 + (v1 + (v1 + (v1 + ("\n\t" + "\n\t")))))))) {
    case "":
        var v8 float32 = (((3.14159 + v4) / float32((1.5 - (0.1 / 1.5)))) + v2)
        _ = v8
    default:
        v5[42] = true
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        _ = len(v5)
        v4 += (float32(1e10) - v2)
        var v10 bool = (0.1 >= 1.5)
        _ = fi9
        _ = v10
    }
    v2 -= (float32(1e10) - (((v4 / 1.5) / 1.0) + (((float32(0.0) / -2.25) / (v4 * v2)) / v4)))
    v11 := GenS2{F0: -1, F1: -81}
    v12 := []int{v3, (v3 * v3), 127, (v3 * (v3 - (v3 & (v3 + -265))))}
    for fi13 := 0; fi13 < 12; fi13++ {
        v3 += (fi13 / (((-(v3))) | 1))
        _ = fi13
    }
    v14 := GenS2{F0: -2147483648, F1: -83}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
