package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 bool }
func genFn0() string {
    var v5 uint = 42
    fmt.Println(v5)
    v6 := []int{-325, 2147483647, -242}
    v5 = (v5 * 1)
    _ = v5
    _ = v6
    return ("a" + ("hello" + "foo bar"))
}
func genFn1(p0 int, p1 uint16) int16 {
    var v7 int32 = 297
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{fi8, fi8, (p0 ^ p0)}
        fmt.Println(v7, fi8, p0)
        v10 := GenS1{F0: -102, F1: 18446744073709551615}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v7, p0, v7)
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    var v4 uint32 = 65535
    if false {
        var v11 int = 32768
        fmt.Println(v4)
        v12 := GenS1{F0: 1, F1: 185}
        _ = v11
        _ = v12
    }
    v4 = (v4 & (^(v4)))
    v4 += (v4 ^ (v4 & v4))
    var v13 uint8 = 233
    var v14 int = 9223372036854775807
    switch "a" {
    case "\n\t":
        v13 += v13
    default:
        v15 := GenS2{F0: 297, F1: v4, F2: 255}
        _ = v15
    }
    v13 -= (^(v13))
    v4 -= v4
    var v16 uint64 = 762
    if false {
        var v17 int16 = -7
        var v18 float64 = (-2.25 - -2.25)
        switch (v14 >> ((0) & 63)) {
        case 32768:
            v18 = -2.25
        case 42:
            if (4294967295 != v4) {
                v19 := GenS3{F0: false}
                v20 := GenS1{F0: -128, F1: 123}
                _ = v19
                _ = v20
            }
        case 467:
            v21 := []int{v14, -788, (^(v14)), 127}
            _ = v21
        default:
            if (42 < 2) {
                v16 += 0
                var v22 bool = true
                var v23 uint16 = 0
                _ = v22
                _ = v23
            }
        }
        _ = v17
        _ = v18
    }
    v24 := GenS1{F0: 69, F1: (v16 * v16)}
    var v25 int16 = 467
    var v26 int = v14
    var v27 int = v26
    v28 := GenS3{F0: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
