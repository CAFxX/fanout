package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0(p0 int32, p1 uint64, p2 uint64) uint {
    var v4 uint64 = 1
    v5 := []int{127, 376, 0}
    var v6 bool = true
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "a" {
    case "foo bar":
        for fi7 := 0; fi7 < 15; fi7++ {
            if false {
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 uint64 = (p2 % ((1) | 1))
                    _ = fi8
                    _ = v9
                }
            }
            var v10 int32 = (-(p0))
            _ = fi7
            _ = v10
        }
    case "hello":
        var v11 int8 = 59
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 int = 42
    var v2 uint32 = 0
    var v3 int8 = -85
    fmt.Println(v2, v2)
    var v12 uint32 = v2
    v13 := []int{-9223372036854775808, 128, (v1 | 42), 128}
    v12 += v12
    for fi14 := 0; fi14 < 11; fi14++ {
        v2 = v2
        _ = fi14
    }
    v15 := GenS1{F0: (v1 % ((-804) | 1)), F1: (false || false), F2: v3}
    v16 := GenS1{F0: -9223372036854775808, F1: true, F2: v3}
    var v17 float32 = ((1e10 / (3.14159 * ((1.5 * 1e10) + 0.1))) - (3.14159 / ((0.0 / 1.0) - 1e10)))
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(v2, v12)
        _ = fi18
    }
    v19 := GenS2{F0: v3, F1: v3}
    if false {
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 int32 = -1
            var v22 int = v1
            v23 := GenS3{F0: uint8(v12), F1: v21, F2: 0}
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        if len(v13) > 0 {
            v13[0] = v1
        }
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 int16 = -32768
            v26 := GenS3{F0: 128, F1: 42, F2: 0}
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    var v27 int8 = v3
    v28 := GenS3{F0: 255, F1: -1, F2: 42}
    for fi29 := 0; fi29 < 12; fi29++ {
        fmt.Println(v12, v3)
        _ = fi29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        v12 = v12
        v13 = append(v13, fi30)
        v1 = (v1 / (((-(v1))) | 1))
        _ = fi30
    }
    v27 += -128
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
