package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int8; F1 int16; F2 uint }
func genFn0(p0 int32, p1 uint8, p2 uint64) int8 {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: (fi6 ^ (fi6 | fi6)), F1: 41}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return 127
}
func genFn1(p0 bool, p1 uint, p2 string) uint32 {
    switch 145 {
    case 9223372036854775807:
        v8 := []int{-222, -9223372036854775808, 42, 339}
        _ = v8
    case -209:
        if (((p2 + p2) + p2) >= p2) {
            p2 = (p2 + p2)
        }
    case 4294967295:
        v9 := []int{42, 42}
        _ = v9
    default:
        p0 = (65 == -70)
    }
    if true {
        switch 9223372036854775807 {
        case 32767:
            fmt.Println(p1)
        case -9223372036854775808:
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS2{F0: 2147483648, F1: 542}
                _ = fi10
                _ = v11
            }
        case 87:
            var v12 int32 = 42
            _ = v12
        default:
            if false {
                v13 := GenS2{F0: 476, F1: 32767}
                v14 := GenS3{F0: 658, F1: -9223372036854775808}
                _ = v13
                _ = v14
            }
        }
        var v15 float64 = 1e10
        _ = v15
    }
    v16 := []int{-1}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 2147483648
}

func main() {
    var v1 int8 = -21
    var v2 int64 = 0
    var v3 int32 = 0
    var v4 int8 = 127
    var v5 int8 = v4
    var v17 uint8 = 0
    v18 := []int{-9223372036854775808}
    v1 += 42
    fmt.Println(v5, v4)
    v19 := GenS1{F0: 4294967295, F1: v5}
    if (172 == 1) {
        v20 := GenS2{F0: 1, F1: 660}
        var v21 int = 4294967295
        switch v21 {
        case 971:
            if (false && true) {
                var v22 uint8 = v17
                _ = v22
            } else {
                v2 = int64(v21)
                var v23 uint32 = 9
                v5 -= v4
                _ = v23
            }
        case 32768:
            v21 -= v21
        case 32767:
            v24 := GenS4{F0: v5, F1: -560, F2: 4294967295}
            _ = v24
        }
        _ = v20
        _ = v21
    }
    v25 := GenS1{F0: -314, F1: 1}
    v26 := GenS4{F0: v5, F1: 547, F2: 18446744073709551615}
    fmt.Println(v3, v5, v1)
    if (127 != 0) {
        var v27 uint16 = 255
        var v28 int16 = -947
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        _ = v27
        _ = v28
    }
    v1 = (v4 << ((65535) & 63))
    fmt.Println(v17, v3, v17)
    var v29 float32 = 3.14159
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 float32 = 3.14159
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        v4 += v5
        v29 -= v29
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
