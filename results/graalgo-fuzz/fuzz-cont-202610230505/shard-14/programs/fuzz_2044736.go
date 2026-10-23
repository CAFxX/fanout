package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 int8; F2 int32 }

func main() {
    v1 := []uint32{180, 0}
    v2 := make(map[uint32]string)
    var v3 int = -9223372036854775808
    var v4 float64 = 0.0
    switch (v3 / ((551) | 1)) {
    case 128:
        var v5 int32 = -1
        _ = v5
    case 9223372036854775807:
        var v6 int32 = -2147483648
        _ = v6
    default:
        var v7 int64 = 958
        _ = v7
    }
    v4 += v4
    if ((42 <= -1) || (((42 < 2) || (!(255 >= 938))) || ((v3 >= (v3 / ((32768) | 1))) || (false && false)))) {
        switch -798 {
        case 128:
            fmt.Println(v3)
        case 0:
            if false {
                fmt.Println(v3, v4)
                v8 := GenS1{F0: 2}
                v4 = float64(((-2.25 * 1e10) / 0.1))
                _ = v8
            } else {
                v2[251] = "a"
            }
        case 886:
            for fi9 := 0; fi9 < 1; fi9++ {
                if false {
                    v4 = v4
                } else {
                    var v10 int8 = -128
                    var v11 float64 = (0.1 / v4)
                    v10 = v10
                    _ = v10
                    _ = v11
                }
                for fi12 := 0; fi12 < 8; fi12++ {
                    v4 = 3.14159
                    v1 = append(v1, 369)
                    if (1000 != 42) {
                        fmt.Println(fi12, fi9, fi12)
                    }
                    _ = fi12
                }
                if (false && false) {
                    if true {
                        for fi13 := 0; fi13 < 4; fi13++ {
                            v3 += 916
                            _ = fi13
                        }
                        v3 -= 127
                    } else {
                        if (128 < 128) {
                            if true {
                                var v14 uint32 = 872
                                _ = v14
                            }
                            v15 := GenS2{F0: 1}
                            v3 = (v3 % ((fi9) | 1))
                            _ = v15
                        } else {
                            v16 := GenS3{F0: 2147483647, F1: -128, F2: -2147483648}
                            _ = v16
                        }
                    }
                } else {
                    fmt.Println(fi9, v3)
                }
                _ = fi9
            }
        default:
            v17 := v2[2147483648]
            _ = v17
        }
        var v18 float64 = v4
        _ = v18
    }
    var v19 int = (v3 << ((255) & 63))
    switch (v19 - -1) {
    case 4294967295:
        var v20 float64 = ((v4 * (1e10 * 1.5)) - v4)
        _ = v20
    case 9223372036854775807:
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 uint16 = 42
            if false {
                v1 = append(v1, 2)
            }
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            _ = fi21
            _ = v22
        }
    case -9223372036854775808:
        v3 -= v19
    }
    var v23 int32 = -2147483648
    switch "a" {
    case "a":
        for fi24 := 0; fi24 < 3; fi24++ {
            v4 = v4
            _ = len(v1)
            var v25 float64 = (v4 / (1e10 + (1e10 / (float64(0.0) / v4))))
            _ = fi24
            _ = v25
        }
    }
    v2[4294967295] = "a"
    v23 = (-(v23))
    var v26 int64 = 412
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
