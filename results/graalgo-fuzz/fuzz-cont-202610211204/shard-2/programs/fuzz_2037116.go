package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint16 = 424
    var v2 float32 = 1.5
    var v3 int32 = 0
    v4 := []uint32{1}
    v5 := make(map[int]uint32)
    var v6 uint32 = 1
    var v7 uint8 = 128
    var v8 int64 = int64(v6)
    v9 := GenS1{F0: (v3 + 42)}
    fmt.Println(v7)
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: (^(v3))}
        v12 := GenS1{F0: 0}
        v5[(fi10 % (((fi10 * fi10)) | 1))] = v6
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int8 = 1
    if false {
        var v14 uint8 = v7
        _ = v14
    }
    var v15 float32 = (v2 / v2)
    if false {
        v16 := GenS2{F0: 542, F1: v6, F2: 127}
        _ = v16
    }
    v1 += uint16(v7)
    v17 := GenS1{F0: v3}
    v7 += v7
    _ = len(v5)
    delete(v5, 42)
    fmt.Println(v2)
    v18 := GenS1{F0: 0}
    for fi19 := 0; fi19 < 20; fi19++ {
        switch fi19 {
        case -556:
            if (v1 >= v1) {
                v20 := GenS3{F0: ((false || (true && false)) && ((v1 % ((0) | 1)) != v1))}
                _ = v20
            }
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
