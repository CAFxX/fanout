package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 string; F1 int; F2 bool }
type GenS6 struct { F0 uint32; F1 uint }
func genFn0(p0 int16, p1 string, p2 uint16) uint32 {
    var v5 int = 4294967295
    v6 := GenS1{F0: (-(v5)), F1: 1, F2: (-(p0))}
    var v7 int8 = 1
    var v8 uint8 = 169
    v9 := GenS1{F0: v5, F1: 2, F2: -171}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 195
}
func genFn1() uint {
    var v10 float32 = 1e10
    var v11 uint32 = 2
    var v12 int16 = 153
    v13 := GenS2{F0: 42, F1: (65535 <= 255)}
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 int64 = -577
    var v2 int = 9223372036854775807
    v3 := make(map[int]int)
    var v4 float64 = -2.25
    for fi14 := 0; fi14 < 13; fi14++ {
        if ((1 != 84) && (32767 > 255)) {
            for fi15 := 0; fi15 < 15; fi15++ {
                if ((0.1 >= ((float32(v4) - 1.5) + (-2.25 - (1.5 / -2.25)))) && (2 != 95)) {
                    if true {
                        v4 = 0.1
                        var v16 float64 = v4
                        v17 := GenS3{F0: 1}
                        _ = v16
                        _ = v17
                    }
                    var v18 int32 = -1
                    _ = v18
                }
                var v19 int8 = -128
                _ = fi15
                _ = v19
            }
            fmt.Println(fi14)
            v20 := []int{fi14}
            _ = v20
        }
        fmt.Println(v2)
        for fi21 := 0; fi21 < 3; fi21++ {
            v22 := GenS1{F0: (^(fi14)), F1: 42, F2: -32768}
            _ = fi21
            _ = v22
        }
        _ = fi14
    }
    v3[v2] = -730
    if true {
        v2 += v2
    } else {
        var v23 float32 = -2.25
        v23 = v23
        fmt.Println(v4, v1, v4)
        _ = v23
    }
    v24 := GenS4{F0: 65535, F1: 0}
    var v25 int64 = v1
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := GenS2{F0: 789, F1: ((-(fi26)) >= (-(fi26)))}
        _ = fi26
        _ = v27
    }
    var v28 int32 = 1
    var v29 int64 = v25
    v30 := GenS2{F0: 32767, F1: true}
    v31 := v3[v2]
    var v32 uint = 42
    v33 := []int{v2, v2}
    var v34 float32 = 1.5
    fmt.Println(v28)
    if (false && (35 < v32)) {
        v35 := GenS5{F0: "", F1: (v31 * (v2 ^ (-(v31)))), F2: true}
        v36 := GenS6{F0: 742, F1: (v32 + 355)}
        _ = v35
        _ = v36
    } else {
        var v37 float64 = 0.1
        v37 -= float64(((v34 - ((v34 - v34) + ((v34 + -2.25) + v34))) - v34))
        _ = v37
    }
    v38 := GenS1{F0: v31, F1: 372, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v4
}
