package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int32 }
type GenS2 struct { F0 bool; F1 int8; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 string; F1 uint64; F2 uint64 }
func genFn0(p0 uint16, p1 int64) int8 {
    v5 := GenS1{F0: 65535, F1: p1, F2: 42}
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint = 42
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 float64 = (((((((1e10 - 1e10) + 0.1) * 0.0) + (((float64(float32(-2.25)) * 1e10) - float64(-2.25)) + 1.5)) / (1e10 / (float64(float32(1.5)) + (0.0 * -2.25)))) - 3.14159) + float64(-2.25))
            var v10 int16 = 32767
            fmt.Println(v7)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v11 := GenS2{F0: true, F1: 42, F2: 0}
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := []int{32768, (fi12 >> ((42) & 63)), fi12, fi12}
        if true {
            var v14 uint64 = 42
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return 126
}

func main() {
    v1 := []uint32{213}
    var v2 int64 = 16
    var v3 float32 = 0.0
    var v4 uint32 = 2147483647
    var v15 int16 = 901
    var v16 int = -798
    v17 := GenS3{F0: true}
    var v18 int8 = -115
    var v19 float32 = (v3 * v3)
    if (!(0 >= 685)) {
        var v20 uint64 = 0
        v4 = v4
        if (true || (640 != 42)) {
            var v21 uint8 = 128
            if true {
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := GenS4{F0: v18, F1: ("a" + ("foo bar" + "a"))}
                    _ = fi22
                    _ = v23
                }
                var v24 int16 = (v15 >> ((0) & 63))
                _ = v24
            }
            _ = v21
        }
        _ = v20
    }
    if ((42 < 1) || (385 > -1)) {
        v25 := GenS5{F0: v15, F1: 18446744073709551615}
        _ = v25
    } else {
        if (!true) {
            if false {
                for fi26 := 0; fi26 < 1; fi26++ {
                    v2 = v2
                    _ = fi26
                }
            }
        } else {
            v27 := GenS3{F0: (!true)}
            v19 = (0.1 * -2.25)
            _ = v27
        }
    }
    v2 -= (v2 % ((42) | 1))
    switch -2147483648 {
    case -406:
        v28 := GenS1{F0: 0, F1: 711, F2: 380}
        _ = v28
    case 195:
        var v29 uint32 = (v4 >> ((0) & 63))
        _ = v29
    case 9223372036854775807:
        for fi30 := 0; fi30 < 2; fi30++ {
            v15 -= v15
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 uint32 = (v4 ^ v4)
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
    default:
        v33 := GenS6{F0: "", F1: 428, F2: 1}
        _ = v33
    }
    v4 += (v4 | v4)
    var v34 int64 = v2
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
