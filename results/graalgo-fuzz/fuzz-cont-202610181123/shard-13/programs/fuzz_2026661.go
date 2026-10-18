package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 string) uint32 {
    v7 := []int{-930}
    p0 = ("" + p0)
    switch -2147483648 {
    case -173:
        var v8 uint16 = 65535
        _ = v8
    }
    fmt.Println(p0)
    if len(v7) > 0 {
        v7[0] = 0
    }
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 int64, p1 uint8) uint16 {
    v9 := []int{518}
    if len(v9) > 0 {
        v9[0] = 0
    }
    _ = p0
    _ = p1
    _ = v9
    return 65535
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]bool)
    v3 := []int{32768, -9223372036854775808, 593}
    var v4 uint16 = 111
    var v5 float64 = 1e10
    var v6 uint32 = 2
    if false {
        switch 32767 {
        case 9223372036854775807:
            if false {
                v1 = v1
            } else {
                fmt.Println(v6)
                v10 := GenS1{F0: -2147483648, F1: v1}
                _ = v10
            }
        case 0:
            v11 := GenS1{F0: 0, F1: (v1 % ((v1) | 1))}
            _ = v11
        }
        fmt.Println(v6)
        var v12 float64 = (float64(float32(((float64(1.5) * (3.14159 / v5)) * -2.25))) + 1e10)
        _ = v12
    } else {
        if ("\n\t" != "hello") {
            var v13 uint = 981
            if ("hello" == "") {
                v2[1] = (0 > 18446744073709551615)
                v1 -= v1
                v14 := GenS1{F0: 0, F1: v1}
                _ = v14
            }
            _ = v13
        }
        switch 9223372036854775807 {
        case 18:
            var v15 uint = 1
            _ = v15
        case 865:
            v4 = 42
        default:
            v1 += (v1 | 127)
        }
        v4 = genFn1(-344, v1)
    }
    v1 = uint8(v6)
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 uint16 = 233
        _ = fi16
        _ = v17
    }
    v5 -= (v5 - 3.14159)
    v18 := GenS2{F0: -1}
    var v19 int64 = -770
    fmt.Println(v5, v1)
    if true {
        v20 := GenS3{F0: 65535}
        var v21 int32 = 2147483647
        _ = v20
        _ = v21
    } else {
        v22 := GenS1{F0: -176, F1: v1}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
