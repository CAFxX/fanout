package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int64; F1 int64; F2 int32 }
func genFn0(p0 int64, p1 uint32) int64 {
    p0 = 0
    v5 := []int{1}
    if true {
        var v6 uint = 312
        var v7 float64 = -2.25
        if ((4294967295 > 247) && false) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            p1 = 2147483648
            fmt.Println(v6)
            var v8 int8 = int8(p1)
            _ = v8
        }
        _ = v6
        _ = v7
    } else {
        var v9 uint = 2
        _ = v9
    }
    v5 = append(v5, 1)
    var v10 int16 = -600
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return -56
}
func genFn1(p0 int8, p1 int64) uint64 {
    v11 := []int{}
    fmt.Println(p0)
    p1 = 301
    v12 := GenS1{F0: 541}
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            switch "a" {
            case "hello":
                if len(v11) > 0 {
                    v11[0] = (-(fi13))
                }
            case "\n\t":
                if (65535 >= 18446744073709551615) {
                    p0 -= (^(p0))
                    var v15 bool = (0 < 244)
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = v15
                } else {
                    v16 := GenS2{F0: -967, F1: 83, F2: 1}
                    fmt.Println(p0)
                    _ = v16
                }
            case "a":
                p1 -= -9223372036854775808
            default:
                if true {
                    p1 = p1
                    var v17 float32 = (float32(0.0) + ((1.5 - ((3.14159 / 1.5) - (float32(0.1) - float32(1e10)))) / (float32(1.5) + 1.5)))
                    switch (fi13 / ((fi13) | 1)) {
                    case 0:
                        fmt.Println(fi13, v17, p1)
                    }
                    _ = v17
                }
            }
            p1 = (-(p1))
            _ = len(v11)
            _ = fi14
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]uint32)
    var v3 uint32 = 2147483648
    var v4 int8 = 100
    v18 := GenS3{F0: genFn0(1, v3), F1: 0, F2: (v1 >> ((0) & 63))}
    var v19 uint32 = (^(v3))
    var v20 int32 = 1
    for fi21 := 0; fi21 < 7; fi21++ {
        if true {
            if false {
                v22 := []int{359, fi21, fi21}
                v23 := GenS1{F0: -280}
                v3 = v19
                _ = v22
                _ = v23
            }
            v3 -= (v19 % (((v3 | 1)) | 1))
        }
        _ = fi21
    }
    _ = len(v2)
    fmt.Println(v19, v1, v20)
    if (144 >= 2) {
        v24 := []int{127, 859}
        _ = v24
    }
    switch "foo bar" {
    case "foo bar":
        var v25 int32 = int32(v19)
        _ = v25
    default:
        if (65535 < 255) {
            for fi26 := 0; fi26 < 14; fi26++ {
                _ = len(v2)
                _ = fi26
            }
        } else {
            v27 := []int{-928, 127, -769, 9223372036854775807}
            if len(v27) > 0 {
                v27[0] = 0
            }
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
