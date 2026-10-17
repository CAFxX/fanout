package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint = 0
    v2 := make(map[int]int)
    v3 := []int{-190, 903, 32767, -2147483648}
    var v4 float32 = 0.1
    v1 = v1
    v5 := GenS1{F0: 1, F1: 2}
    v4 = float32(0.1)
    var v6 int64 = 935
    for fi7 := 0; fi7 < 16; fi7++ {
        v2[(fi7 & (fi7 & fi7))] = -755
        var v8 uint32 = 2147483648
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint32 = 261
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: v1, F1: 196}
    v4 = v4
    v12 := GenS3{F0: -607, F1: 506}
    if true {
        var v13 uint32 = 1
        var v14 int16 = 0
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 int8 = 0
            _ = fi15
            _ = v16
        }
        _ = v13
        _ = v14
    } else {
        v1 = v1
    }
    var v17 int32 = -2147483648
    v4 += 0.0
    var v18 int = 12
    v19 := GenS3{F0: -561, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
