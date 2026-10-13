package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int; F2 uint32 }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{2147483648}
    var v3 int8 = -128
    v4 := make(map[uint32]int)
    var v5 int32 = -1
    v6 := make(map[int]bool)
    fmt.Println(v1, v3, v5)
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v1 = (v1 + (v1 + "hello"))
            var v9 bool = (!false)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if len(v2) > 0 {
        v2[0] = 433
    }
    v10 := v4[2]
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS1{F0: (2 != 38), F1: 786, F2: 776}
            fmt.Println(v10)
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v10 -= v10
    v3 += v3
    v14 := GenS2{F0: -367}
    v1 = v1
    switch 42 {
    case 1:
        v15 := GenS3{F0: 0, F1: v10, F2: 2147483648}
        _ = v15
    default:
        var v16 int = int(v5)
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
