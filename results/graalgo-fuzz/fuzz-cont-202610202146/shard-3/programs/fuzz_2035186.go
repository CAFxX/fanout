package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0() uint8 {
    var v5 float64 = -2.25
    if false {
        v5 -= 1.5
    } else {
        switch 32768 {
        case 25:
            var v6 float64 = (-2.25 * v5)
            _ = v6
        case 42:
            v5 = float64(float32((v5 * ((1.5 * v5) / v5))))
        case 1:
            v7 := GenS1{F0: 1, F1: 128, F2: -365}
            _ = v7
        }
        v8 := []int{-1, -2147483648, 32767}
        _ = v8
    }
    v9 := []int{32767, 32768, 939}
    switch 122 {
    case 4294967295:
        if len(v9) > 0 {
            v9[0] = 127
        }
    case 9223372036854775807:
        var v10 string = "foo bar"
        _ = v10
    case 2147483647:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    _ = v5
    _ = v9
    return 72
}
func genFn1(p0 uint64, p1 int16, p2 uint8) int {
    for fi11 := 0; fi11 < 10; fi11++ {
        p2 -= (p2 % ((0) | 1))
        _ = fi11
    }
    v12 := []int{-518, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 0
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]bool)
    v3 := []int{}
    v4 := make(map[int]uint32)
    if true {
        for fi13 := 0; fi13 < 12; fi13++ {
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS1{F0: 183, F1: (v1 % ((v1) | 1)), F2: -1}
                v16 := v2[65535]
                var v17 int32 = -90
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            v1 -= 1
            switch (fi13 * fi13) {
            case 1:
                for fi18 := 0; fi18 < 6; fi18++ {
                    switch "a" {
                    case "":
                        _ = len(v3)
                    }
                    fmt.Println(fi13, fi13)
                    v1 = (v1 & (v1 | (^(v1))))
                    _ = fi18
                }
            case -101:
                if true {
                    v4[-9223372036854775808] = 447
                } else {
                    v1 -= v1
                    v19 := GenS2{F0: int32(fi13), F1: 4294967295}
                    var v20 uint = 18446744073709551615
                    _ = v19
                    _ = v20
                }
            case 586:
                v21 := GenS2{F0: -2147483648, F1: 4294967295}
                _ = v21
            }
            _ = fi13
        }
    }
    if true {
        fmt.Println(v1)
    } else {
        v22 := v4[-9223372036854775808]
        _ = v22
    }
    var v23 uint16 = 724
    switch 32768 {
    case 9223372036854775807:
        fmt.Println(v23, v1)
    case 32767:
        fmt.Println(v1, v23)
    case 128:
        v24 := GenS1{F0: 2, F1: (v1 | v1), F2: 2147483647}
        _ = v24
    default:
        v1 -= v1
    }
    v23 += v23
    v1 = (v1 >> ((609) & 63))
    switch 32767 {
    case -54:
        if false {
            v1 -= genFn0()
            delete(v2, 65535)
        }
    case 546:
        v23 = (v23 >> ((18446744073709551615) & 63))
    default:
        var v25 int8 = 42
        _ = v25
    }
    v2[4294967295] = (1 < 1)
    v3 = append(v3, 128)
    var v26 float32 = ((1e10 + (0.0 * -2.25)) / float32(0.1))
    var v27 uint32 = 167
    v26 -= (((0.1 * (v26 + (float32(3.14159) - v26))) * 1.5) + 3.14159)
    v28 := GenS1{F0: 2, F1: (v1 >> ((4294967295) & 63)), F2: 1}
    v1 = genFn0()
    if false {
        v26 = float32(0.0)
        v1 -= uint8(v27)
    } else {
        if (-1 == 9223372036854775807) {
            var v29 uint16 = 684
            var v30 int16 = 1
            v23 -= (^(v29))
            _ = v29
            _ = v30
        } else {
            var v31 int16 = 538
            v31 = (-(v31))
            _ = v31
        }
    }
    fmt.Println(v27, v27, v26)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
