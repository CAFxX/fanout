package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 int32) int32 {
    v7 := GenS1{F0: 42, F1: p0}
    var v8 uint64 = 780
    _ = p0
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 int = 127
    var v2 float32 = (0.1 - 3.14159)
    var v3 int = v1
    v4 := []string{"a"}
    var v5 int64 = 496
    var v6 int = v3
    v3 = (v6 | v3)
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 float32 = (v2 - ((float32((0.0 + 1.5)) * v2) - (-2.25 + (0.0 - 0.0))))
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int64 = (v5 % ((42) | 1))
        _ = fi11
        _ = v12
    }
    var v13 int = -89
    v14 := GenS1{F0: 1, F1: genFn0(2147483647)}
    var v15 uint64 = 774
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        v16 := GenS2{F0: 2, F1: 255}
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 float64 = (float64((1.5 - v2)) + 1e10)
            v19 := GenS3{F0: (v1 == (v3 << ((18446744073709551615) & 63)))}
            v6 -= v1
            _ = fi17
            _ = v18
            _ = v19
        }
        if (1 < genFn0(2147483647)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
