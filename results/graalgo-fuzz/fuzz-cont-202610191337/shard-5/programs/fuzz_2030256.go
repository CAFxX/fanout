package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = -2.25
    var v3 int = -55
    v4 := []int{406, 0, 1, (v3 ^ v3)}
    for fi5 := 0; fi5 < 18; fi5++ {
        v2 -= (v2 / (v2 / (v2 + v2)))
        _ = fi5
    }
    v2 -= 0.1
    v1[v3] = 248
    v4 = append(v4, v3)
    v6 := GenS1{F0: 42}
    var v7 int = v3
    v8 := GenS2{F0: 639, F1: 857, F2: 1}
    for fi9 := 0; fi9 < 12; fi9++ {
        v3 += (v7 ^ v3)
        _ = fi9
    }
    fmt.Println(v7, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
