package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int64; F1 int32; F2 int }
func genFn0(p0 int8, p1 uint32) string {
    switch 2147483647 {
    case -1:
        p0 = p0
    case 32768:
        v5 := GenS1{F0: 65535, F1: (^(p0))}
        _ = v5
    case 337:
        v6 := []int{-2147483648, 32767, 1, 128}
        _ = v6
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int32 = 0
        var v9 uint16 = 255
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 9223372036854775807 {
    case 1:
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := []int{666, 32768, (fi10 % ((fi10) | 1))}
            var v12 uint32 = p1
            _ = fi10
            _ = v11
            _ = v12
        }
    case -9223372036854775808:
        p1 += 42
    default:
        fmt.Println(p0, p1)
    }
    var v13 float64 = 0.0
    fmt.Println(v13, p0)
    _ = p0
    _ = p1
    _ = v13
    return "\n\t"
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    var v3 uint32 = 968
    v4 := []uint32{v3, v3, v3}
    var v14 uint32 = 1
    var v15 float64 = (1.5 / ((1.5 * (1.5 - -2.25)) / 1.5))
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int16 = -1
        v15 += (((v15 / -2.25) - v15) + (0.1 * 0.0))
        fmt.Println(v15, v15, v3)
        _ = fi16
        _ = v17
    }
    v3 += 797
    var v18 bool = (!false)
    var v19 float64 = 1e10
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 int16 = -32768
            fmt.Println(v3)
            v23 := GenS2{F0: 42, F1: 42, F2: fi20}
            _ = fi21
            _ = v22
            _ = v23
        }
        if (v18 && ((23 > -128) && v18)) {
            for fi24 := 0; fi24 < 19; fi24++ {
                if false {
                    v4 = append(v4, v14)
                } else {
                    var v25 float32 = ((-2.25 - ((1.5 / 1.0) / 1.0)) + -2.25)
                    _ = v25
                }
                _ = fi24
            }
        } else {
            fmt.Println(v19, fi20, v15)
        }
        _ = fi20
    }
    v26 := GenS2{F0: 709, F1: 0, F2: 487}
    v4 = append(v4, v3)
    _ = len(v1)
    var v27 uint32 = v14
    v28 := GenS1{F0: 0, F1: 83}
    if false {
        if len(v4) > 0 {
            v4[0] = v14
        }
    } else {
        v19 = float64(1e10)
    }
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
