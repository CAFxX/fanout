package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int }

func main() {
    v1 := []uint32{18, 2147483647, 2147483648}
    v2 := []int{}
    v3 := make(map[uint32]bool)
    v4 := GenS1{F0: 42, F1: 685, F2: 2147483647}
    v5 := GenS1{F0: -1, F1: 673, F2: 131}
    var v6 int16 = -210
    delete(v3, 2147483648)
    delete(v3, 261)
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint8 = 2
        for fi9 := 0; fi9 < 12; fi9++ {
            v6 += (-(v6))
            fmt.Println(v6, v8)
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    var v10 int16 = v6
    v3[1] = (!(!false))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
