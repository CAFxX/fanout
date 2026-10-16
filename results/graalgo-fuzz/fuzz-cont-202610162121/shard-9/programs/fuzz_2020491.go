package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 bool; F2 int8 }
func genFn0() int32 {
    var v7 int = 885
    v7 += (v7 & (v7 | (v7 / ((v7) | 1))))
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float32 = ((0.0 / 1.5) - (float32(1e10) - (float32((3.14159 / ((1e10 * (-2.25 + -2.25)) / 3.14159))) * 1.5)))
    var v3 int64 = -9223372036854775808
    var v4 int = 0
    v5 := make(map[uint32]uint32)
    var v6 float32 = float32(float64(v2))
    v2 = v2
    var v8 uint = 860
    if ((499 > -2147483648) && ((v8 - (v8 >> ((v8) & 63))) != v8)) {
        v9 := []int{(v4 + v4), (v4 << (((v8 - v8)) & 63)), v4}
        v4 = v4
        _ = v9
    } else {
        v10 := GenS1{F0: 0, F1: false, F2: v3}
        var v11 int8 = 33
        if ((^(v4)) < (^(v4))) {
            v12 := GenS2{F0: 65535, F1: 751}
            _ = v12
        } else {
            v13 := v5[2]
            var v14 float32 = -2.25
            _ = v13
            _ = v14
        }
        _ = v10
        _ = v11
    }
    v15 := []int{}
    var v16 int = v4
    for fi17 := 0; fi17 < 12; fi17++ {
        if true {
            if ((127 == 2) && true) {
                var v18 uint8 = 116
                _ = v18
            } else {
                var v19 int64 = 0
                _ = v19
            }
            var v20 int32 = genFn0()
            _ = v20
        } else {
            if (true || true) {
                fmt.Println(v4)
                v2 = ((float32((0.0 - (float64(v6) / (0.1 + 0.1)))) * v6) * v2)
                var v21 uint32 = 2
                _ = v21
            }
            var v22 float64 = 3.14159
            if (!false) {
                v23 := GenS1{F0: 2147483647, F1: false, F2: v3}
                var v24 int = fi17
                _ = v23
                _ = v24
            }
            _ = v22
        }
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi17
    }
    fmt.Println(v6, v16)
    if ((v3 + v3) == v3) {
        v25 := GenS3{F0: 784}
        _ = v25
    }
    switch "\n\t" {
    case "":
        for fi26 := 0; fi26 < 2; fi26++ {
            v4 = v4
            v5[455] = uint32(v1)
            _ = fi26
        }
    default:
        for fi27 := 0; fi27 < 16; fi27++ {
            _ = len(v5)
            switch (v16 | (v4 & v4)) {
            case -9223372036854775808:
                v16 += fi27
            case 306:
                fmt.Println(v3, v8)
            case 0:
                v28 := GenS4{F0: false, F1: (18446744073709551615 < 115), F2: -128}
                _ = v28
            }
            v2 = v6
            _ = fi27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
