package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 bool; F1 uint8; F2 uint8 }
func genFn0() int16 {
    v6 := []int{-921, -913, 4294967295}
    var v7 float32 = 3.14159
    v6 = append(v6, 513)
    var v8 uint8 = 2
    _ = v6
    _ = v7
    _ = v8
    return -32768
}
func genFn1(p0 int32, p1 string, p2 string) int16 {
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: -9223372036854775808, F1: int32(fi9), F2: 4294967295}
        v11 := []int{(fi9 * 42), (fi9 << ((379) & 63))}
        _ = fi9
        _ = v10
        _ = v11
    }
    p2 = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 bool = true
    var v2 int16 = 150
    var v3 int8 = -27
    v4 := make(map[string]string)
    v5 := []int{-208}
    _ = len(v4)
    if v1 {
        fmt.Println(v3)
    }
    v2 = -32768
    if v1 {
        v12 := GenS2{F0: v2}
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := GenS1{F0: 9223372036854775807, F1: 1, F2: 804}
            var v15 int16 = genFn0()
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    }
    v2 = (v2 & genFn0())
    v1 = true
    v2 -= (^(v2))
    if (!false) {
        var v16 bool = (65535 < 2147483647)
        v5 = append(v5, 42)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v16
    } else {
        if (-2147483648 > 42) {
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 uint64 = 0
                v1 = true
                v19 := GenS1{F0: 229, F1: -2147483648, F2: (^(v18))}
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS3{F0: 4294967295, F1: 42}
            _ = v20
        }
        switch 128 {
        case 2147483647:
            if (-407 != 496) {
                var v21 int32 = 42
                v1 = true
                _ = v21
            }
        case 42:
            fmt.Println(v1, v1)
        }
    }
    var v22 int16 = v2
    var v23 float64 = (3.14159 - (3.14159 - 1e10))
    v1 = (true || false)
    var v24 uint16 = 32767
    v25 := GenS4{F0: v1, F1: 2, F2: 68}
    v2 = 1
    fmt.Println(v24, v3)
    if true {
        var v26 bool = (!(v3 == (v3 / ((84) | 1))))
        v23 -= (v23 - v23)
        var v27 uint = 255
        _ = v26
        _ = v27
    }
    v1 = v1
    v1 = (v1 && v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
