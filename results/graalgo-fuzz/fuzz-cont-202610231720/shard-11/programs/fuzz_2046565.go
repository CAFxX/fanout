package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint64; F1 string; F2 uint32 }
func genFn0() int16 {
    var v5 int8 = -128
    var v6 int8 = (-(v5))
    v5 += (v6 | (v5 >> ((255) & 63)))
    var v7 int32 = 0
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int32 = (v7 + (v7 + v7))
        var v10 bool = (2 >= 42)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return 878
}

func main() {
    var v1 bool = false
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    var v4 int16 = 0
    switch 526 {
    case -2147483648:
        fmt.Println(v4, v4)
    case 128:
        var v11 int16 = (^(v4))
        _ = v11
    case 225:
        _ = len(v3)
    default:
        v1 = v1
    }
    var v12 int = 127
    var v13 int32 = -2147483648
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            _ = len(v3)
            _ = fi15
        }
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = fi14
    }
    if (!false) {
        v4 = v4
        v16 := GenS1{F0: 103, F1: (v12 ^ (-(v12)))}
        var v17 int64 = 1
        _ = v16
        _ = v17
    } else {
        fmt.Println(v4, v12, v13)
        v13 = v13
    }
    if v1 {
        if (("a" < "") && v1) {
            if v1 {
                v1 = (!((0.1 / float32(((float64(1e10) + float64(0.1)) - float64(-2.25)))) >= float32(3.14159)))
                for fi18 := 0; fi18 < 7; fi18++ {
                    v13 -= v13
                    _ = fi18
                }
            }
        }
        var v19 int = 128
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 string = "a"
            _ = fi20
            _ = v21
        }
        _ = v19
    } else {
        fmt.Println(v12, v12)
        var v22 bool = (false && v1)
        v4 -= (v4 + v4)
        _ = v22
    }
    var v23 int64 = -1
    var v24 int16 = (^(v4))
    if false {
        v25 := GenS2{F0: 183, F1: ("" + ("hello" + "hello")), F2: 0}
        _ = v25
    } else {
        var v26 uint8 = 255
        var v27 int32 = 42
        _ = v26
        _ = v27
    }
    v12 += (v12 * (v12 ^ v12))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
