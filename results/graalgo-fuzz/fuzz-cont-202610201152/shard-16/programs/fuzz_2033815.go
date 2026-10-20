package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int8; F2 bool }
func genFn0(p0 uint64) uint16 {
    v4 := []int{4294967295}
    p0 = 121
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 65535
}
func genFn1() int64 {
    switch -393 {
    case -976:
        var v5 int32 = 0
        _ = v5
    case -503:
        var v6 float64 = 0.0
        _ = v6
    }
    var v7 uint8 = 2
    var v8 int32 = -2147483648
    v8 -= (v8 >> ((0) & 63))
    for fi9 := 0; fi9 < 18; fi9++ {
        v7 = (v7 - v7)
        v8 += v8
        _ = fi9
    }
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 uint = 505
    var v2 uint32 = 894
    var v3 int16 = -883
    if (1 != -1) {
        var v10 uint16 = 255
        _ = v10
    }
    fmt.Println(v3, v2, v3)
    for fi11 := 0; fi11 < 12; fi11++ {
        if true {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int = fi11
                var v14 bool = true
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        var v15 uint = 4294967295
        fmt.Println(fi11, v3, fi11)
        _ = fi11
        _ = v15
    }
    if false {
        fmt.Println(v1)
    }
    var v16 int8 = 42
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS1{F0: (v1 / ((255) | 1)), F1: v3, F2: 18446744073709551615}
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: 18446744073709551615}
    v20 := GenS3{F0: 625, F1: 42, F2: -24}
    v1 = 160
    var v21 uint64 = 18446744073709551615
    var v22 uint16 = 836
    v23 := []int{-65}
    if (uint64(v2) != v21) {
        var v24 uint32 = v2
        var v25 uint32 = (v24 % ((824) | 1))
        var v26 uint16 = v22
        _ = v24
        _ = v25
        _ = v26
    }
    switch int(v1) {
    case 944:
        var v27 uint8 = 119
        _ = v27
    case -1:
        if true {
            fmt.Println(v21)
            v22 += (v22 << (((v1 / ((v1) | 1))) & 63))
            fmt.Println(v16, v21)
        } else {
            v21 -= (^(v21))
            _ = len(v23)
            var v28 uint64 = (v21 / ((v21) | 1))
            _ = v28
        }
    default:
        fmt.Println(v21, v21, v1)
    }
    v2 += v2
    for fi29 := 0; fi29 < 7; fi29++ {
        v30 := GenS3{F0: v21, F1: v21, F2: v16}
        if (400 >= v22) {
            _ = len(v23)
            v31 := GenS4{F0: v22}
            _ = v31
        }
        v32 := GenS5{F0: v16, F1: 42, F2: true}
        _ = fi29
        _ = v30
        _ = v32
    }
    v1 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
