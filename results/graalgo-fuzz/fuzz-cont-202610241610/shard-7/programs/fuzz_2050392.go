package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
func genFn0(p0 int8) string {
    var v5 int = -1
    if (false || true) {
        v6 := GenS1{F0: 32768, F1: 0}
        var v7 int = 268
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v5
    return "a"
}

func main() {
    var v1 float64 = (0.1 - 1e10)
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 float64 = float64(0.0)
    if false {
        for fi8 := 0; fi8 < 3; fi8++ {
            if ((fi8 & (-(fi8))) <= fi8) {
                fmt.Println(fi8)
                for fi9 := 0; fi9 < 6; fi9++ {
                    var v10 int64 = 9223372036854775807
                    var v11 bool = (0 >= 1)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v1 = v4
            }
            var v12 float32 = 1e10
            var v13 string = "hello"
            _ = fi8
            _ = v12
            _ = v13
        }
    } else {
        v1 -= v4
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 uint = 154
        if (42 < 65535) {
            if (-2147483648 <= -2147483648) {
                fmt.Println(fi14)
                v4 = ((v1 + (float64(1.5) + v4)) + 1e10)
            } else {
                for fi16 := 0; fi16 < 18; fi16++ {
                    v15 -= (v15 * v15)
                    v15 += v15
                    var v17 uint64 = 42
                    _ = fi16
                    _ = v17
                }
            }
        } else {
            var v18 uint = (v15 >> ((v15) & 63))
            v4 += float64(3.14159)
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    v4 = float64(0.0)
    var v19 float32 = -2.25
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 uint16 = 638
        v21 -= (v21 | (v21 % (((v21 + 539)) | 1)))
        var v22 uint = 2
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint = 423
    switch -8 {
    case 32767:
        var v24 int32 = 638
        _ = v24
    case 9223372036854775807:
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint16 = 987
            switch fi25 {
            case 32768:
                v27 := GenS2{F0: 1, F1: 743, F2: 42}
                _ = v27
            default:
                v28 := v3[""]
                _ = v28
            }
            fmt.Println(v1)
            _ = fi25
            _ = v26
        }
    case 128:
        v4 -= 1.5
    default:
        v19 -= v19
    }
    var v29 uint16 = 239
    var v30 int64 = 42
    if false {
        if (!true) {
            for fi31 := 0; fi31 < 8; fi31++ {
                for fi32 := 0; fi32 < 4; fi32++ {
                    v2[fi32] = 805
                    _ = fi32
                }
                _ = fi31
            }
        }
        v2[127] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
