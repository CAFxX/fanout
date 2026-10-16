package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 int; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int16 }
func genFn0(p0 uint16, p1 uint8) int32 {
    switch 0 {
    case 32768:
        v6 := GenS1{F0: 1, F1: 39}
        _ = v6
    case 128:
        v7 := []int{-9223372036854775808}
        _ = v7
    default:
        var v8 int8 = 100
        _ = v8
    }
    var v9 uint32 = 0
    if (!false) {
        v9 = v9
    } else {
        var v10 int64 = 778
        p1 = p1
        var v11 uint32 = (^(v9))
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 1, F1: (p1 << ((0) & 63))}
    _ = p0
    _ = p1
    _ = v12
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    v4 := []uint32{}
    v5 := make(map[uint32]int)
    v13 := GenS1{F0: -393, F1: 73}
    var v14 bool = (false && false)
    v14 = v14
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int16 = -1
        for fi17 := 0; fi17 < 17; fi17++ {
            if (0 > 2) {
                v18 := GenS2{F0: (v16 & (-(v16))), F1: (fi15 - 634), F2: genFn0(83, 255)}
                _ = v18
            }
            var v19 uint64 = 4294967295
            _ = fi17
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    v14 = v14
    v20 := v2[2]
    var v21 float64 = -2.25
    for fi22 := 0; fi22 < 7; fi22++ {
        fmt.Println(v21, v21)
        fmt.Println(v21, v21, v21)
        v14 = v14
        _ = fi22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        if (32767 >= 255) {
            v14 = v14
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            v25 := GenS3{F0: 42}
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v26 := GenS4{F0: v20, F1: 1}
    fmt.Println(v14)
    v21 += (v21 + float64(0.0))
    var v27 uint16 = 255
    v2[2147483647] = "\n\t"
    if (42 != 2147483648) {
        v20 = "\n\t"
        var v28 uint32 = 361
        _ = v28
    }
    var v29 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
