package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    var v1 int8 = -128
    var v2 uint8 = 42
    var v3 uint16 = 255
    var v4 uint64 = 1
    if ("a" < "foo bar") {
        if (true || (v2 == (v2 | v2))) {
            var v5 float32 = (-2.25 / 0.1)
            v6 := GenS1{F0: (v3 / (((v3 >> ((2) & 63))) | 1)), F1: -2147483648}
            _ = v5
            _ = v6
        }
        v2 = 214
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 11; fi8++ {
                v9 := GenS2{F0: 42, F1: false}
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: (v3 % ((uint16(v2)) | 1)), F1: (-(fi7))}
            if (-785 >= 42) {
                v2 = (v2 - v2)
                var v11 float64 = 0.1
                var v12 uint = 18446744073709551615
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := []int{fi7, fi13, fi13, 42}
                    fmt.Println(v2, v3)
                    v1 -= (v1 - 1)
                    if (-2.25 != (float32(float64((1e10 * 1e10))) + 0.0)) {
                        v4 += 2
                        if (0 >= -125) {
                            if false {
                                v3 -= (^(v3))
                                v4 = v4
                            }
                            v14 = append(v14, 1)
                            fmt.Println(v1, v1)
                        }
                    }
                    var v15 string = "hello"
                    var v16 uint16 = v3
                    v2 -= v2
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
            _ = fi7
            _ = v10
        }
    }
    var v17 uint32 = 100
    v18 := []int{-9223372036854775808, -9223372036854775808}
    if false {
        var v19 uint8 = (v2 + v2)
        v3 += (v3 << ((42) & 63))
        var v20 float64 = 3.14159
        _ = v19
        _ = v20
    }
    v3 -= v3
    v2 -= (v2 / ((128) | 1))
    for fi21 := 0; fi21 < 9; fi21++ {
        for fi22 := 0; fi22 < 18; fi22++ {
            if true {
                v4 = (^(v4))
            } else {
                var v23 int8 = 127
                v4 = 442
                _ = v23
            }
            var v24 float32 = float32(0.1)
            var v25 int32 = 1
            _ = fi22
            _ = v24
            _ = v25
        }
        _ = fi21
    }
    if (2147483647 > -1) {
        switch 32767 {
        case 1:
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 int16 = -975
                _ = fi26
                _ = v27
            }
        default:
            if false {
                fmt.Println(v4)
            }
        }
    }
    var v28 uint64 = 18446744073709551615
    v1 += (v1 ^ (v1 % (((v1 % ((v1) | 1))) | 1)))
    v18 = append(v18, int(v17))
    _ = len(v18)
    fmt.Println(v28)
    var v29 uint16 = v3
    var v30 float64 = ((float64(0.1) + 3.14159) - 0.1)
    var v31 uint = 18446744073709551615
    v32 := GenS1{F0: (v3 / ((v29) | 1)), F1: 128}
    v29 = (v3 * 360)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
