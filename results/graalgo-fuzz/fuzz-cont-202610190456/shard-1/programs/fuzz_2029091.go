package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int; F2 int16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int32 = -1
    var v2 uint16 = 113
    v3 := []uint32{}
    switch -421 {
    case 741:
        v2 = v2
    case 1:
        v3 = append(v3, 2147483648)
    case -354:
        v4 := GenS1{F0: 531, F1: 4294967295}
        _ = v4
    default:
        _ = len(v3)
    }
    v2 += v2
    v1 = v1
    v5 := GenS1{F0: 42, F1: 523}
    v1 = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!false) {
        var v6 string = "foo bar"
        if (100 < -128) {
            if false {
                fmt.Println(v6, v6)
                for fi7 := 0; fi7 < 1; fi7++ {
                    var v8 uint32 = 610
                    _ = fi7
                    _ = v8
                }
                if true {
                    switch 2147483647 {
                    case 931:
                        v9 := GenS1{F0: 1, F1: 42}
                        _ = v9
                    case 32768:
                        v2 = v2
                    }
                }
            } else {
                v10 := GenS2{F0: 2, F1: 1, F2: -1}
                v6 = v6
                _ = v10
            }
        }
        if (v2 <= v2) {
            _ = len(v3)
            var v11 uint64 = 495
            _ = v11
        }
        _ = v6
    } else {
        v2 = 505
    }
    v12 := GenS2{F0: 50, F1: 9223372036854775807, F2: 42}
    v13 := GenS2{F0: 81, F1: 2147483647, F2: 42}
    v14 := GenS3{F0: 4294967295}
    v15 := GenS3{F0: 18446744073709551615}
    v3 = append(v3, 65535)
    if true {
        switch 842 {
        case 42:
            v1 -= (v1 | (v1 / (((-(v1))) | 1)))
        case -132:
            if ((1e10 - ((0.0 + -2.25) + -2.25)) >= (1e10 / 1e10)) {
                fmt.Println(v2)
            }
        case -9223372036854775808:
            var v16 int16 = int16(v1)
            _ = v16
        default:
            var v17 string = "a"
            _ = v17
        }
        v3 = append(v3, 2147483648)
    } else {
        for fi18 := 0; fi18 < 6; fi18++ {
            if len(v3) > 0 {
                v3[0] = 1
            }
            var v19 float32 = ((0.0 * (0.0 + 1e10)) / 1.0)
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 4; fi20++ {
            _ = len(v3)
            _ = fi20
        }
        var v21 int32 = v1
        _ = v21
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v5
}
