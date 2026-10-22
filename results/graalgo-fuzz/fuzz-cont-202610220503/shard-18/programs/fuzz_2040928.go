package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int; F1 int16; F2 int8 }
func genFn0() int16 {
    var v6 float32 = -2.25
    switch "hello" {
    case "a":
        fmt.Println(v6)
    case "\n\t":
        var v7 uint32 = 552
        _ = v7
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{(fi8 % (((fi8 << ((1) & 63))) | 1)), fi8, (^(fi8))}
            var v10 float64 = -2.25
            v10 = v10
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    _ = v6
    return 200
}
func genFn1(p0 uint8) int {
    p0 += (p0 / (((^(p0))) | 1))
    switch -9223372036854775808 {
    case 32767:
        var v11 uint16 = 2
        _ = v11
    case 0:
        var v12 uint8 = p0
        _ = v12
    case -583:
        p0 += (p0 & 255)
    default:
        fmt.Println(p0)
    }
    var v13 int8 = -49
    _ = p0
    _ = v13
    return 32768
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (!false)
    v3 := []uint32{4294967295, 2, 402}
    var v4 int32 = 0
    var v5 float32 = (3.14159 * 3.14159)
    v4 -= v4
    switch genFn1(57) {
    case 32768:
        v1[145] = (42 < 199)
    default:
        var v14 int16 = 1000
        _ = v14
    }
    v15 := GenS1{F0: 9223372036854775807, F1: v4}
    v4 = (v4 >> ((1) & 63))
    v5 += (v5 * (((((float32((3.14159 / 1e10)) / ((v5 * 3.14159) * v5)) + v5) / (v5 * v5)) / ((0.0 + (v5 * v5)) / v5)) / 1.5))
    v2 = false
    if false {
        var v16 int32 = -2147483648
        for fi17 := 0; fi17 < 13; fi17++ {
            v5 = 1.5
            for fi18 := 0; fi18 < 18; fi18++ {
                if v2 {
                    v19 := GenS1{F0: 1, F1: v4}
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v16, v4)
                    }
                    _ = v19
                } else {
                    v5 = (0.1 - v5)
                    v16 += (v16 | -1)
                }
                _ = fi18
            }
            switch (fi17 & (fi17 / ((genFn1(0)) | 1))) {
            case 4294967295:
                if len(v3) > 0 {
                    v3[0] = 888
                }
            }
            _ = fi17
        }
        var v20 uint16 = 2
        _ = v16
        _ = v20
    } else {
        v21 := GenS2{F0: 42, F1: 1, F2: 100}
        _ = len(v1)
        if false {
            for fi22 := 0; fi22 < 4; fi22++ {
                v2 = true
                v4 = v4
                for fi23 := 0; fi23 < 11; fi23++ {
                    _ = len(v3)
                    var v24 bool = v2
                    var v25 int = fi22
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi22
            }
            v2 = true
        }
        _ = v21
    }
    v26 := GenS2{F0: 262, F1: genFn0(), F2: 127}
    for fi27 := 0; fi27 < 9; fi27++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v28 int32 = 42
        fmt.Println(v4, v2)
        _ = fi27
        _ = v28
    }
    fmt.Println(v2, v4)
    var v29 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
