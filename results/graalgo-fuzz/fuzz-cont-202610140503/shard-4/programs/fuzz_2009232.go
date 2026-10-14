package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int8 }
func genFn0(p0 uint, p1 int8, p2 bool) int16 {
    if (255 < 255) {
        p2 = (float32((float64((((1e10 - float32(((0.1 + 1e10) * (1.5 - 1e10)))) + 3.14159) * 3.14159)) - 1.5)) != float32(0.0))
    }
    p2 = (2 < 551)
    v4 := GenS1{F0: -1}
    if (9223372036854775807 != -1) {
        var v5 uint32 = 2147483647
        var v6 int = 0
        _ = v5
        _ = v6
    }
    switch 880 {
    case 127:
        v7 := []int{}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 int16, p1 uint64, p2 int) uint32 {
    p1 = (p1 | 827)
    for fi8 := 0; fi8 < 3; fi8++ {
        switch 32768 {
        case 0:
            switch (fi8 << ((2) & 63)) {
            case -985:
                if true {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        var v10 float32 = 0.0
                        _ = fi9
                        _ = v10
                    }
                    v11 := GenS2{F0: true, F1: 27}
                    _ = v11
                }
            case 32768:
                p0 = (-(p0))
            case -9223372036854775808:
                p2 += fi8
            }
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint16 = 65535
    v3 := make(map[uint32]int)
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint = 159
            v2 = (v2 & (v2 << ((v14) & 63)))
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 int64 = 42
            _ = fi15
            _ = v16
        }
        _ = fi12
    }
    var v17 uint64 = 42
    if (358 < 42) {
        switch 689 {
        case 32767:
            if (!false) {
                v17 -= 42
            }
        case 9223372036854775807:
            v1 -= v1
        default:
            var v18 int8 = 59
            _ = v18
        }
        if true {
            var v19 uint8 = 1
            _ = v19
        } else {
            var v20 float64 = -2.25
            v1 = 1
            _ = v20
        }
    }
    v21 := GenS1{F0: (^(v1))}
    var v22 int16 = 42
    fmt.Println(v1)
    v1 = 829
    v22 = v22
    fmt.Println(v1)
    if ((1.5 * 1.5) == -2.25) {
        v23 := []int{-2147483648, -2147483648}
        v2 = (v2 & 2)
        v24 := GenS2{F0: true, F1: 35}
        _ = v23
        _ = v24
    } else {
        switch -493 {
        case -833:
            var v25 float32 = (0.1 * 0.0)
            _ = v25
        case 32768:
            var v26 string = "foo bar"
            _ = v26
        case 128:
            v3[54] = 562
        default:
            v17 = (v17 ^ v17)
        }
        v22 = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
