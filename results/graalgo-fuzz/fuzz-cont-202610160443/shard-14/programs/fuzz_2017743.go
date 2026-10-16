package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint }
type GenS2 struct { F0 int32; F1 uint16; F2 int32 }

func main() {
    var v1 int16 = 902
    v2 := []uint32{473, 393, 2, 42}
    var v3 uint64 = 87
    var v4 uint = 1
    switch 128 {
    case -299:
        if true {
            fmt.Println(v4, v4)
            v5 := GenS1{F0: 83, F1: (42 == -1), F2: (^(v4))}
            v6 := GenS2{F0: 42, F1: 255, F2: -995}
            _ = v5
            _ = v6
        }
    }
    var v7 uint = (v4 + v4)
    v3 += 0
    v7 += 1
    var v8 uint32 = 383
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: uint64(v8), F1: (v3 > (v3 + v3)), F2: (^(v4))}
        v7 = (v4 ^ v4)
        if len(v2) > 0 {
            v2[0] = 478
        }
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        _ = len(v2)
        for fi12 := 0; fi12 < 5; fi12++ {
            v3 = 593
            var v13 int64 = 0
            fmt.Println(v3, v4, v3)
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 int8 = 1
    fmt.Println(v1, v3, v8)
    v2 = append(v2, (^(v8)))
    if len(v2) > 0 {
        v2[0] = (v8 ^ v8)
    }
    var v15 uint = 120
    var v16 int32 = -850
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
