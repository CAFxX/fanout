package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint; F1 int8; F2 int64 }
type GenS5 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint16) int32 {
    v7 := []int{666, 127, 2147483647, 0}
    if false {
        switch "foo bar" {
        case "a":
            var v8 uint32 = 996
            _ = v8
        case "foo bar":
            p0 += p0
        case "hello":
            p0 += (p0 | (p0 & p0))
        default:
            var v9 uint16 = p0
            _ = v9
        }
        switch "" {
        case "":
            p0 = (^(p0))
        case "\n\t":
            v10 := GenS1{F0: 0, F1: p0, F2: 211}
            _ = v10
        case "foo bar":
            switch -77 {
            case -2147483648:
                if len(v7) > 0 {
                    v7[0] = -403
                }
            case 32767:
                var v11 float32 = (-2.25 - 1e10)
                _ = v11
            case 128:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
        }
        switch 671 {
        case -570:
            var v12 uint16 = p0
            _ = v12
        default:
            p0 = 952
        }
    }
    v13 := GenS1{F0: 0, F1: (p0 / ((p0) | 1)), F2: 0}
    for fi14 := 0; fi14 < 11; fi14++ {
        p0 = 255
        _ = len(v7)
        _ = fi14
    }
    _ = p0
    _ = v13
    _ = v7
    return -2147483648
}
func genFn1(p0 uint32) uint16 {
    if true {
        switch "foo bar" {
        case "":
            v15 := []int{32767, 2147483647, -1}
            _ = v15
        case "\n\t":
            var v16 uint16 = 171
            _ = v16
        }
        v17 := GenS2{F0: 0}
        p0 -= 2147483648
        _ = v17
    }
    v18 := GenS3{F0: 0, F1: (588 != 255)}
    if false {
        var v19 uint = 227
        _ = v19
    }
    fmt.Println(p0)
    v20 := GenS1{F0: -9223372036854775808, F1: 358, F2: 96}
    _ = p0
    _ = v18
    _ = v20
    return 1
}

func main() {
    var v1 float64 = 1.5
    var v2 bool = (!(2 >= 65535))
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 127
    var v5 float64 = (v1 + v1)
    var v6 uint32 = 4294967295
    var v21 uint64 = 42
    v22 := GenS2{F0: 439}
    v23 := GenS4{F0: 42, F1: 108, F2: -1}
    var v24 int64 = -9223372036854775808
    var v25 uint32 = (v6 & v6)
    if v2 {
        v4 -= (v4 - (v4 * 175))
        _ = len(v3)
    } else {
        v24 -= v24
        v26 := GenS3{F0: -25, F1: v2}
        _ = v26
    }
    v27 := GenS5{F0: (v21 + v21), F1: -1}
    v28 := GenS1{F0: v24, F1: 1, F2: v4}
    if (v4 <= (v4 * v4)) {
        v29 := []int{2147483647, 309, -852}
        _ = v29
    }
    v30 := GenS4{F0: 765, F1: -128, F2: (v24 << ((2) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
