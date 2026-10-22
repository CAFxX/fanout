package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 2
    var v3 int16 = -1
    v4 := []uint32{}
    v5 := make(map[uint32]string)
    if (-2.25 <= 0.1) {
        var v6 string = ""
        v7 := GenS1{F0: false}
        _ = v6
        _ = v7
    }
    v3 = v3
    switch 32768 {
    case 9223372036854775807:
        var v8 uint64 = 18446744073709551615
        _ = v8
    default:
        v2 += v2
    }
    v2 += v2
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v2, v2)
        _ = fi9
    }
    var v10 int32 = int32(v2)
    var v11 int32 = (v10 % ((2147483647) | 1))
    v12 := GenS1{F0: false}
    v13 := GenS2{F0: 1, F1: 141}
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
