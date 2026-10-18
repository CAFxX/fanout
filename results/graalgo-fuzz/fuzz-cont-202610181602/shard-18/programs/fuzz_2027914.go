package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int8; F2 uint16 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]int)
    var v3 uint16 = 0
    var v4 int16 = 42
    v5 := []int{}
    fmt.Println(v1)
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 float32 = 0.1
        fmt.Println(v1, v7)
        var v8 bool = (int32(v3) <= 771)
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v4 += (v4 / (((^(v4))) | 1))
        _ = fi9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int16 = -32768
        var v12 int64 = -1
        fmt.Println(v4, v3)
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: (((42 < -1) || false) && true)}
    fmt.Println(v4)
    var v14 uint8 = 0
    v15 := GenS2{F0: 42, F1: 119, F2: 419}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
