package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint64 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := []int{-839}
    v4 := v2[128]
    v5 := GenS1{F0: 42, F1: (false && false), F2: 711}
    var v6 int64 = 42
    v6 += -607
    var v7 float64 = 1e10
    v8 := GenS2{F0: v6}
    for fi9 := 0; fi9 < 18; fi9++ {
        v3 = append(v3, fi9)
        var v10 uint = 356
        switch (^(fi9)) {
        case 127:
            var v11 uint32 = (v4 / ((uint32(v10)) | 1))
            _ = v11
        case 0:
            v7 = v7
        case -9223372036854775808:
            for fi12 := 0; fi12 < 9; fi12++ {
                v3 = append(v3, fi9)
                var v13 uint64 = 178
                var v14 int8 = 6
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = fi9
        _ = v10
    }
    var v15 int64 = 743
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
