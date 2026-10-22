package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 bool; F1 string; F2 uint16 }
func genFn0(p0 uint32) int16 {
    p0 -= p0
    v6 := GenS1{F0: (p0 | (^(p0))), F1: 514, F2: (-32768 > -1)}
    if false {
        p0 += p0
    } else {
        if (!false) {
            switch -2147483648 {
            case -2147483648:
                p0 = p0
            case 9223372036854775807:
                fmt.Println(p0)
            case -916:
                v7 := []int{128, -528}
                _ = v7
            }
            for fi8 := 0; fi8 < 14; fi8++ {
                if false {
                    p0 = (p0 & 2147483647)
                    v9 := []int{2147483647}
                    v10 := GenS1{F0: p0, F1: 65535, F2: true}
                    _ = v10
                    _ = v9
                }
                _ = fi8
            }
            for fi11 := 0; fi11 < 16; fi11++ {
                if (false || false) {
                    v12 := []int{-269, (^(fi11)), (fi11 * (fi11 ^ 281)), fi11}
                    _ = v12
                } else {
                    v13 := GenS1{F0: 5, F1: 504, F2: false}
                    _ = v13
                }
                for fi14 := 0; fi14 < 20; fi14++ {
                    v15 := []int{}
                    p0 += (p0 | 0)
                    v16 := GenS2{F0: (fi11 % ((fi14) | 1))}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                for fi17 := 0; fi17 < 7; fi17++ {
                    var v18 uint16 = 680
                    _ = fi17
                    _ = v18
                }
                _ = fi11
            }
        }
    }
    var v19 int32 = 42
    if (596 > 2147483648) {
        v20 := []int{128, -587, 893}
        var v21 int64 = 9223372036854775807
        _ = v20
        _ = v21
    }
    _ = p0
    _ = v19
    _ = v6
    return 8
}
func genFn1(p0 int8, p1 uint, p2 int16) int16 {
    p0 = p0
    v22 := GenS3{F0: 0, F1: 42}
    fmt.Println(p1)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return genFn0(4294967295)
}

func main() {
    var v1 int64 = 499
    v2 := []uint32{164, 62, 1}
    var v3 uint = 4294967295
    var v4 float32 = 1e10
    var v5 uint32 = 419
    var v23 uint64 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (42 >= -701) {
        if (v1 <= 0) {
            fmt.Println(v5, v23, v3)
            switch -2147483648 {
            case -9223372036854775808:
                v23 -= (^(v23))
            case 32767:
                var v24 int64 = v1
                _ = v24
            default:
                v25 := GenS4{F0: 2147483647, F1: (true && false)}
                _ = v25
            }
            v3 = v3
        }
    } else {
        fmt.Println(v3, v23, v23)
        var v26 int = 1
        fmt.Println(v3)
        _ = v26
    }
    fmt.Println(v3, v1, v5)
    v27 := GenS3{F0: uint8(v23), F1: 702}
    v28 := GenS5{F0: (251 == -1), F1: ("" + "foo bar"), F2: 255}
    var v29 float32 = 3.14159
    v23 -= v23
    v29 += v29
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
