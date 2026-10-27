package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 int16) int64 {
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 float64 = 0.0
        if (!true) {
            v8 := GenS1{F0: 245, F1: 889}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    p0 -= p0
    var v9 int32 = 1
    fmt.Println(p0)
    var v10 int8 = 0
    _ = p0
    _ = v10
    _ = v9
    return 1
}
func genFn1(p0 int) int8 {
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 uint64 = 4294967295
        var v13 int16 = -1
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(p0)
    v14 := []int{(p0 << ((255) & 63)), p0, (p0 + (p0 / (((p0 & -2147483648)) | 1))), 1}
    _ = p0
    _ = v14
    return -128
}

func main() {
    var v1 uint16 = 904
    v2 := make(map[int]uint32)
    var v3 bool = false
    v4 := make(map[int]bool)
    var v5 float32 = (1e10 - 0.0)
    v15 := GenS1{F0: 255, F1: 530}
    var v16 float32 = 0.1
    v5 = (v5 + v5)
    var v17 float64 = (3.14159 * -2.25)
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := v4[fi18]
        v20 := GenS1{F0: 242, F1: 1}
        _ = len(v4)
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int8 = 42
    if false {
        if v3 {
            v17 = 0.0
        }
        fmt.Println(v17)
        v5 -= float32(v17)
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        fmt.Println(fi22, v21)
        _ = fi22
    }
    v1 = v1
    v23 := GenS2{F0: -2147483648, F1: 4294967295}
    v1 += v1
    _ = len(v2)
    if false {
        fmt.Println(v1, v16, v1)
        var v24 bool = (42 <= 1)
        _ = v24
    }
    var v25 string = "\n\t"
    v26 := []int{127, 32767, -9223372036854775808, 17}
    fmt.Println(v21, v21)
    var v27 float64 = (((1e10 * 0.1) / 3.14159) / ((0.1 + ((((v17 / (v17 - v17)) - 0.0) * 0.1) - 1.5)) - v17))
    switch 2147483647 {
    case 4294967295:
        var v28 uint16 = (v1 + v1)
        _ = v28
    case -153:
        v17 -= float64((v5 + 3.14159))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
