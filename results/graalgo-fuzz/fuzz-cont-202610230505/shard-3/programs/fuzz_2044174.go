package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 uint64 }

func main() {
    v1 := []int{-668, 128, 42}
    var v2 int16 = 1
    v3 := []string{"hello"}
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    v6 := []uint32{242, 695, 4294967295}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int16 = 42
        _ = fi7
        _ = v8
    }
    var v9 float64 = (float64(0.1) / (float64((1.5 / 1.0)) - 0.0))
    var v10 int = 2147483647
    v2 = v2
    if (v10 < (v10 & (v10 << ((18446744073709551615) & 63)))) {
        var v11 uint8 = 0
        _ = len(v3)
        v9 += v9
        _ = v11
    } else {
        v9 = v9
        switch "" {
        case "\n\t":
            v12 := GenS1{F0: v10, F1: false}
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 int16 = -92
            _ = fi13
            _ = v14
        }
    }
    switch "\n\t" {
    case "\n\t":
        if (false && false) {
            for fi15 := 0; fi15 < 15; fi15++ {
                v16 := GenS1{F0: (-(v10)), F1: (!false)}
                _ = fi15
                _ = v16
            }
        }
    default:
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 bool = false
            var v19 uint64 = 18446744073709551615
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    var v20 int64 = 42
    var v21 float64 = 1e10
    v9 = v21
    v22 := GenS2{F0: 188, F1: 0, F2: 315}
    for fi23 := 0; fi23 < 11; fi23++ {
        v9 = v9
        _ = fi23
    }
    v20 += v20
    v2 += v2
    v10 += (v10 ^ v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
