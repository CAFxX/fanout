package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 uint; F1 uint; F2 bool }
func genFn0(p0 uint32) uint8 {
    v5 := []int{}
    v6 := GenS1{F0: 0, F1: -128}
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS2{F0: p0, F1: fi8, F2: -235}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 128, F1: -91}
        var v11 int = fi7
        _ = fi7
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint16 = 4
    var v2 bool = (42 >= 42)
    var v3 int = -240
    var v4 int8 = 0
    v2 = false
    v12 := []int{v3}
    if (0.0 != (1e10 - ((1e10 / (float64((1.5 - 1e10)) * -2.25)) / (float64(-2.25) + float64(-2.25))))) {
        var v13 bool = v2
        _ = v13
    } else {
        var v14 uint8 = 42
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := GenS1{F0: 127, F1: v4}
        var v17 int8 = (v4 >> ((0) & 63))
        _ = len(v12)
        _ = fi15
        _ = v16
        _ = v17
    }
    if v2 {
        fmt.Println(v4, v1)
        var v18 uint32 = 2147483647
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS1{F0: 42, F1: 127}
            fmt.Println(v18)
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    v1 = 0
    if (2 > 42) {
        if (v3 < (v3 >> ((513) & 63))) {
            var v21 float32 = (float32((-2.25 - -2.25)) - 1e10)
            if false {
                v22 := GenS1{F0: 1, F1: v4}
                if (!true) {
                    v3 += (v3 / (((-(v3))) | 1))
                    v23 := GenS1{F0: uint8(v1), F1: (v4 - (v4 % ((0) | 1)))}
                    _ = v23
                }
                _ = v22
            }
            _ = v21
        }
    } else {
        fmt.Println(v2)
        var v24 float64 = 0.0
        if (!v2) {
            var v25 int16 = 0
            var v26 uint64 = 2
            var v27 int = (v3 ^ (v3 >> ((18446744073709551615) & 63)))
            _ = v25
            _ = v26
            _ = v27
        } else {
            var v28 uint8 = genFn0(1)
            v29 := GenS3{F0: 65535, F1: 255, F2: (!(255 == 2))}
            _ = v28
            _ = v29
        }
        _ = v24
    }
    v2 = ((v1 != v1) && false)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
