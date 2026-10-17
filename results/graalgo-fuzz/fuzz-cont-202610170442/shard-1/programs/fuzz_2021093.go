package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint8; F1 uint }

func main() {
    var v1 int32 = 57
    v2 := make(map[uint32]string)
    var v3 uint64 = 214
    var v4 float64 = 3.14159
    if (0 < 2) {
        fmt.Println(v1)
        for fi5 := 0; fi5 < 13; fi5++ {
            v4 = float64(0.0)
            _ = fi5
        }
        v6 := GenS1{F0: 14, F1: -22}
        _ = v6
    }
    switch -772 {
    case 1:
        if ((float64((3.14159 * ((((float32(v4) + 1.5) * 0.1) / float32(v4)) - 3.14159))) >= (((-2.25 / v4) / 1.0) * 3.14159)) && false) {
            v3 += v3
        }
    case 32767:
        var v7 int16 = 1
        _ = v7
    case -588:
        var v8 bool = (1 > 0)
        _ = v8
    default:
        switch ("" + (("" + "a") + "")) {
        case "foo bar":
            v2[65535] = "foo bar"
        case "\n\t":
            var v9 uint32 = 2147483647
            _ = v9
        case "a":
            v4 = 0.1
        }
    }
    var v10 uint64 = 18446744073709551615
    for fi11 := 0; fi11 < 15; fi11++ {
        v10 = (v3 ^ (v3 * 160))
        v12 := GenS2{F0: v1}
        _ = fi11
        _ = v12
    }
    fmt.Println(v4, v1, v10)
    v3 -= v10
    var v13 int16 = -253
    var v14 int8 = 42
    v15 := GenS2{F0: (v1 | (^(v1)))}
    v4 -= (v4 + 0.0)
    v16 := GenS3{F0: 4294967295, F1: v1, F2: 127}
    v17 := []int{32767, int(v14)}
    v13 -= v13
    var v18 int64 = -1
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := GenS4{F0: 127, F1: 1}
            v14 = (v14 * (-(v14)))
            _ = fi20
            _ = v21
        }
        v3 = (^(v10))
        fmt.Println(v13, v1)
        _ = fi19
    }
    v10 = v10
    fmt.Println(v4, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
