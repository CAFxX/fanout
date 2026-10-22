package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint }
type GenS6 struct { F0 int8; F1 uint; F2 bool }
type GenS7 struct { F0 int32; F1 int64; F2 int32 }
func genFn0() uint16 {
    if (128 >= -2147483648) {
        switch -1 {
        case 120:
            var v6 int64 = 42
            _ = v6
        case 32768:
            var v7 uint32 = 42
            _ = v7
        default:
            var v8 uint = 65535
            _ = v8
        }
    }
    v9 := GenS1{F0: 4294967295}
    var v10 float32 = float32(0.1)
    var v11 float32 = -2.25
    _ = v10
    _ = v11
    _ = v9
    return 1
}
func genFn1(p0 int8, p1 uint, p2 int64) uint16 {
    v12 := GenS1{F0: 18446744073709551615}
    v13 := GenS2{F0: 1, F1: 255, F2: p1}
    p2 -= (p2 + p2)
    var v14 uint8 = 157
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return genFn0()
}

func main() {
    var v1 uint8 = 0
    var v2 uint32 = 65535
    v3 := []int{128, 677, 42, 2147483647}
    var v4 int16 = 42
    var v5 int32 = -2147483648
    var v15 uint32 = v2
    for fi16 := 0; fi16 < 10; fi16++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v2 = v15
        v17 := GenS3{F0: 1, F1: (v4 >> ((42) & 63))}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        if false {
            v15 = v15
            fmt.Println(v5, v4)
        } else {
            v3 = append(v3, 128)
        }
        if true {
            v19 := GenS3{F0: 4294967295, F1: (v4 << ((2) & 63))}
            v5 -= v5
            var v20 int32 = (v5 ^ (v5 / ((v5) | 1)))
            _ = v19
            _ = v20
        } else {
            v5 = v5
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 = (v15 >> ((319) & 63))
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v5 = v5
        _ = fi21
    }
    v5 = v5
    fmt.Println(v5)
    v22 := GenS4{F0: 42, F1: (v4 / (((v4 * -613)) | 1))}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v23 := GenS5{F0: genFn1(0, 2, -1), F1: 101, F2: 580}
    v24 := GenS6{F0: 90, F1: 255, F2: false}
    switch "hello" {
    case "":
        if (genFn0() == 2) {
            v25 := GenS5{F0: 42, F1: -128, F2: 1}
            switch "a" {
            case "a":
                v26 := GenS7{F0: (v5 / (((v5 << ((65535) & 63))) | 1)), F1: -9223372036854775808, F2: v5}
                _ = v26
            default:
                v27 := GenS5{F0: 578, F1: -1, F2: 65535}
                _ = v27
            }
            _ = v25
        }
    }
    v28 := GenS7{F0: 0, F1: 42, F2: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
