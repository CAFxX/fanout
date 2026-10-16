package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int = 4294967295
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    v1 += v1
    var v4 int = v1
    var v5 uint64 = 2
    v6 := GenS1{F0: ("" + "\n\t"), F1: false}
    var v7 int = 0
    fmt.Println(v1)
    v8 := []int{(v4 * 4294967295), (v7 << ((255) & 63))}
    for fi9 := 0; fi9 < 3; fi9++ {
        v2[4294967295] = 1
        _ = fi9
    }
    if true {
        v10 := GenS2{F0: 127}
        var v11 int = v7
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int64 = 1
            var v14 float32 = -2.25
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v10
        _ = v11
    }
    v4 += (v4 >> ((0) & 63))
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS2{F0: -60}
        if (-530 < 592) {
            var v17 uint16 = 42
            v5 += v5
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
