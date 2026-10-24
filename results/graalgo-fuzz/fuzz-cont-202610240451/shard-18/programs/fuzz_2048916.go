package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0() uint32 {
    var v6 int16 = 32767
    if (1.5 != (3.14159 / 3.14159)) {
        if false {
            var v7 uint16 = 32767
            v7 -= v7
            var v8 int64 = 956
            _ = v7
            _ = v8
        }
        v6 += 32767
        v9 := []int{-401, 1, 0}
        _ = v9
    }
    v10 := []int{-9223372036854775808, 127, 541}
    _ = v10
    _ = v6
    return 2147483647
}
func genFn1(p0 int32) int32 {
    p0 -= p0
    var v11 uint32 = 4294967295
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int64 = -717
        var v14 uint = 0
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS1{F0: uint32(v14), F1: 42, F2: 100}
            _ = fi15
            _ = v16
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v17 bool = false
    p0 = p0
    _ = p0
    _ = v11
    _ = v17
    return p0
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 int64 = -9223372036854775808
    var v5 int8 = 42
    v5 -= (v5 % (((^(v5))) | 1))
    v4 -= 1
    fmt.Println(v5, v5)
    v18 := GenS1{F0: 4294967295, F1: 2, F2: 100}
    var v19 uint8 = 127
    v5 += (v5 >> ((255) & 63))
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 uint = 255
        var v22 float32 = float32((1.5 + 1.5))
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: v4, F1: 2147483647}
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 float32 = 1e10
        for fi26 := 0; fi26 < 16; fi26++ {
            if (2 <= genFn0()) {
                v27 := []int{473, (fi26 % ((42) | 1))}
                _ = v27
            }
            fmt.Println(fi26, fi26)
            var v28 float64 = 0.0
            _ = fi26
            _ = v28
        }
        _ = fi24
        _ = v25
    }
    v1 += (v1 / v1)
    for fi29 := 0; fi29 < 1; fi29++ {
        if (156 < 260) {
            v5 += v5
        } else {
            var v30 int32 = 0
            v31 := []int{-1, (fi29 & fi29)}
            _ = v30
            _ = v31
        }
        var v32 int = (fi29 & (fi29 % (((^(fi29))) | 1)))
        _ = fi29
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
