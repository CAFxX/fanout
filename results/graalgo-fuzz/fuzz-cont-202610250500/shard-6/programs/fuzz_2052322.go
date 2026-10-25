package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint; F1 int8 }
func genFn0(p0 uint32) uint16 {
    p0 = 42
    switch "\n\t" {
    case "\n\t":
        var v4 int32 = 650
        _ = v4
    case "hello":
        v5 := []int{-58, 1, 1, 4294967295}
        _ = v5
    default:
        p0 -= (p0 + p0)
    }
    v6 := GenS1{F0: false, F1: 255, F2: 42}
    switch -2147483648 {
    case 1:
        if (107 <= 42) {
            v7 := GenS2{F0: 2147483647, F1: -751}
            var v8 uint = 384
            _ = v7
            _ = v8
        }
    case -9223372036854775808:
        p0 += p0
    case 128:
        if (!false) {
            if false {
                var v9 int16 = 0
                p0 -= p0
                _ = v9
            }
            fmt.Println(p0)
            fmt.Println(p0)
        } else {
            var v10 uint16 = 897
            _ = v10
        }
    default:
        v11 := []int{}
        _ = v11
    }
    var v12 int64 = -1
    _ = p0
    _ = v12
    _ = v6
    return 65535
}
func genFn1(p0 uint32) int8 {
    var v13 int32 = 400
    fmt.Println(v13)
    p0 = p0
    v14 := GenS2{F0: -2147483648, F1: -1}
    v13 -= (^(v13))
    _ = p0
    _ = v13
    _ = v14
    return -113
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]bool)
    var v3 float64 = 0.0
    var v15 int64 = -9223372036854775808
    v1 -= float32(1e10)
    if false {
        var v16 int8 = genFn1(4294967295)
        v17 := GenS3{F0: 1, F1: v16}
        _ = v16
        _ = v17
    } else {
        var v18 uint = 47
        var v19 int8 = 127
        v20 := GenS4{F0: 5, F1: 691, F2: -298}
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := GenS5{F0: genFn1(130), F1: -732}
    v3 = v3
    v22 := GenS6{F0: 2}
    v3 += v3
    for fi23 := 0; fi23 < 7; fi23++ {
        v1 -= (0.1 * -2.25)
        _ = fi23
    }
    switch -496 {
    case -886:
        fmt.Println(v3, v1, v15)
    default:
        v24 := GenS7{F0: 2, F1: -128}
        _ = v24
    }
    if true {
        v3 -= (float64((v1 / v1)) / ((((float64(v1) / 1.5) * float64(((3.14159 / v1) - 0.1))) / v3) - ((3.14159 + v3) / 1e10)))
    } else {
        v15 = (v15 + -20)
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 uint8 = 2
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 int = fi25
                _ = fi27
                _ = v28
            }
            if (1 != genFn0(42)) {
                v15 += (-(v15))
                v29 := GenS4{F0: (v26 & (v26 / (((v26 ^ (v26 % ((1) | 1)))) | 1))), F1: 706, F2: fi25}
                _ = v29
            }
            _ = fi25
            _ = v26
        }
    }
    v15 = (v15 ^ v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
