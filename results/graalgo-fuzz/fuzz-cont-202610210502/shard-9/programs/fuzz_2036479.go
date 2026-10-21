package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint }
func genFn0() uint64 {
    var v7 int64 = 1
    fmt.Println(v7)
    v8 := GenS1{F0: true, F1: 128, F2: 1}
    var v9 string = "a"
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint8) int16 {
    fmt.Println(p0)
    p0 += 255
    switch 2147483647 {
    case -1:
        p0 -= (^(p0))
    case -372:
        v10 := GenS2{F0: p0}
        _ = v10
    case -9223372036854775808:
        p0 = (p0 + (p0 * 128))
    default:
        v11 := GenS1{F0: true, F1: -225, F2: 0}
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 int8 = 100
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: false, F1: 30, F2: 288}
    _ = p0
    _ = v14
    return 1
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int32 = -2147483648
    var v3 bool = (0.0 >= (float32(0.1) + (1e10 - 1e10)))
    v4 := make(map[string]uint32)
    var v5 int32 = v2
    v6 := []uint32{}
    var v15 int32 = (v5 - v5)
    var v16 string = ("" + ("foo bar" + "\n\t"))
    if len(v6) > 0 {
        v6[0] = 790
    }
    if v3 {
        var v17 float64 = 0.1
        _ = v17
    }
    fmt.Println(v15, v5)
    v18 := GenS1{F0: ((58 >= 42) && (170 != 2)), F1: 32768, F2: 65535}
    v2 += -2147483648
    v1 -= v1
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := GenS3{F0: (v15 & (v2 >> ((18446744073709551615) & 63))), F1: (v1 ^ v1), F2: 636}
        _ = fi19
        _ = v20
    }
    if true {
        var v21 int = 2147483647
        _ = v21
    }
    if (((1e10 / 1.0) <= float64(float32(1e10))) && true) {
        fmt.Println(v3, v2, v5)
        fmt.Println(v5, v5)
    } else {
        _ = len(v6)
        switch 127 {
        case 621:
            v3 = (32767 != 644)
        }
    }
    fmt.Println(v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
