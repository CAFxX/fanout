package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int16 = 0
    v4 := []uint32{2147483647}
    v5 := GenS1{F0: -1, F1: 65535, F2: (v3 * v3)}
    v3 -= v3
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: (^(v3)), F1: 255, F2: v3}
        _ = fi6
        _ = v7
    }
    v3 = v3
    var v8 int16 = v3
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 int = 32767
        fmt.Println(v3, v8)
        var v11 int64 = -1
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: (-(v3)), F1: 32767, F2: v3}
    v13 := GenS2{F0: 0, F1: 578}
    v3 += v3
    v14 := GenS3{F0: 565, F1: 1}
    var v15 float64 = 1e10
    if (!true) {
        v16 := GenS4{F0: 32767}
        fmt.Println(v3)
        v3 = (v8 % (((v8 * v8)) | 1))
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
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
