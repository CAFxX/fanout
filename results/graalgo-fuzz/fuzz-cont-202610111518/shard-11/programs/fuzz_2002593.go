package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint64; F2 int64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int8) int64 {
    var v7 float64 = 1.5
    v8 := GenS1{F0: 1}
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 int8 = (-(p0))
    _ = p0
    _ = v12
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 uint = 4294967295
    var v2 uint8 = 128
    v3 := make(map[string]string)
    var v4 int8 = 61
    var v5 int = 9223372036854775807
    v6 := []uint32{2147483647, 2147483647, 65535}
    var v13 float64 = (1.5 + 3.14159)
    var v14 uint64 = 929
    v4 = int8(v5)
    fmt.Println(v4, v4, v5)
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint32 = 2
        var v17 int32 = 0
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS2{F0: v17, F1: (^(v14)), F2: 9223372036854775807}
            if len(v6) > 0 {
                v6[0] = (^(v16))
            }
            _ = fi18
            _ = v19
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        var v20 uint16 = 258
        _ = v20
    } else {
        _ = len(v3)
        var v21 uint16 = 693
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 int32 = 0
            if false {
                v25 := GenS2{F0: (v24 * v24), F1: (v14 / ((v14) | 1)), F2: -664}
                if true {
                    if (false || ((^(v1)) == uint(v14))) {
                        fmt.Println(fi22, v13, fi23)
                        v26 := GenS2{F0: (^(v24)), F1: 0, F2: 42}
                        v14 = (v14 & v14)
                        _ = v26
                    } else {
                        v27 := GenS2{F0: v24, F1: (v14 + (v14 / (((v14 * (v14 & (v14 % ((v14) | 1))))) | 1))), F2: -738}
                        _ = v27
                    }
                }
                _ = v25
            }
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    switch "foo bar" {
    case "\n\t":
        switch v5 {
        case 32768:
            v6 = append(v6, 1)
        case -9223372036854775808:
            v2 += uint8(v14)
        }
    case "foo bar":
        v28 := GenS3{F0: -17}
        _ = v28
    default:
        if true {
            var v29 uint16 = 255
            switch (-(v5)) {
            case 1:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                fmt.Println(v13, v2, v29)
            }
            _ = v29
        } else {
            v13 = float64(float32(v13))
            v13 = ((1e10 / (((-2.25 - (v13 + 0.1)) * v13) - float64((((1.5 + 3.14159) / 0.1) - (3.14159 - 1e10))))) + (((0.1 * (0.1 - (v13 - v13))) + -2.25) * v13))
            v14 += (v14 << ((1) & 63))
        }
    }
    var v30 uint32 = 4294967295
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
