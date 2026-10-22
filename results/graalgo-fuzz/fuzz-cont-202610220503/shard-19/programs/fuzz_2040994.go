package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 bool; F2 int16 }
func genFn0(p0 string, p1 int8, p2 int) int16 {
    var v5 uint32 = 425
    var v6 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -7
}

func main() {
    var v1 int8 = 42
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    for fi7 := 0; fi7 < 10; fi7++ {
        _ = len(v3)
        v8 := GenS1{F0: 2, F1: 65535, F2: true}
        v9 := GenS2{F0: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS3{F0: 720, F1: true, F2: 1}
    v1 = v1
    v11 := GenS1{F0: 209, F1: 65535, F2: false}
    for fi12 := 0; fi12 < 2; fi12++ {
        v1 += (v1 + v1)
        _ = fi12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint16 = 997
        _ = fi13
        _ = v14
    }
    _ = len(v2)
    var v15 float64 = 1.5
    fmt.Println(v15, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
