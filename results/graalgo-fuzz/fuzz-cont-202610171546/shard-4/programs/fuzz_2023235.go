package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{}
    var v3 float64 = ((1.5 - (3.14159 * float64(-2.25))) / float64(1e10))
    var v4 uint64 = 0
    v5 := []int{4294967295, -412, -9223372036854775808}
    for fi6 := 0; fi6 < 1; fi6++ {
        if false {
            v4 -= v4
        } else {
            var v7 uint32 = 786
            _ = v7
        }
        _ = fi6
    }
    v8 := GenS1{F0: 32768, F1: 1}
    for fi9 := 0; fi9 < 5; fi9++ {
        v4 -= (v4 >> ((0) & 63))
        fmt.Println(v3)
        _ = fi9
    }
    _ = len(v1)
    var v10 uint8 = 1
    v4 = v4
    var v11 float32 = 0.0
    var v12 int8 = 127
    v13 := v1[157]
    v14 := GenS2{F0: 611}
    v11 -= -2.25
    var v15 float32 = v11
    v5 = append(v5, -1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
