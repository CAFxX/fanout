package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 int; F1 int32; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint16; F1 int16; F2 int }

func main() {
    var v1 float64 = float64(-2.25)
    var v2 int32 = 2147483647
    var v3 uint8 = 128
    var v4 float32 = 1.5
    v5 := GenS1{F0: 0, F1: v2}
    v6 := []int{-57}
    v7 := GenS2{F0: (^(v3)), F1: 318, F2: ""}
    if ((712 == 42) || ((0.0 != (0.1 / v4)) || true)) {
        if true {
            if false {
                var v8 float64 = (0.0 - ((0.1 / ((1.5 * v1) + v1)) / v1))
                _ = v8
            }
            var v9 int32 = (v2 >> ((865) & 63))
            _ = v9
        }
        var v10 int8 = 0
        fmt.Println(v1, v3, v10)
        _ = v10
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 10; fi12++ {
                v1 -= v1
                fmt.Println(fi11, v2)
                v3 = (v3 ^ (v3 << ((0) & 63)))
                _ = fi12
            }
            v13 := GenS3{F0: 9223372036854775807, F1: v2, F2: 105}
            _ = fi11
            _ = v13
        }
    }
    fmt.Println(v3)
    if (1 != 18446744073709551615) {
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS4{F0: 42}
            _ = fi14
            _ = v15
        }
    }
    switch 4294967295 {
    case 128:
        if true {
            var v16 float32 = v4
            v2 += v2
            _ = v16
        }
    case 32768:
        v17 := GenS2{F0: v3, F1: 593, F2: ("" + "\n\t")}
        _ = v17
    default:
        v18 := GenS2{F0: (v3 / ((255) | 1)), F1: 65535, F2: "foo bar"}
        _ = v18
    }
    if true {
        switch 9223372036854775807 {
        case 192:
            v3 = (v3 * (v3 % ((v3) | 1)))
        case -9223372036854775808:
            var v19 int64 = 969
            _ = v19
        default:
            for fi20 := 0; fi20 < 13; fi20++ {
                if ((true && false) && false) {
                    switch (fi20 | (-(fi20))) {
                    case -1:
                        v21 := GenS5{F0: 65535, F1: 1}
                        _ = v21
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = fi20
            }
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(v1, v3)
    } else {
        v1 = v1
        v22 := GenS6{F0: 255, F1: 0, F2: -9223372036854775808}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
