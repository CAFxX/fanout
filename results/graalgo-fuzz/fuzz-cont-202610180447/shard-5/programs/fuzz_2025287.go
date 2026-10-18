package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 698
    var v3 uint8 = uint8(v2)
    var v4 int8 = -1
    var v5 int = 1
    v6 := []int{127, (-(v5)), -132, (-(v5))}
    var v7 int8 = v4
    _ = len(v6)
    delete(v1, (v5 << (((^(v2))) & 63)))
    v8 := GenS1{F0: 1000, F1: (42 > 2)}
    switch int(v2) {
    case 9223372036854775807:
        v3 += 127
    case 42:
        if (!true) {
            if (1.5 <= (1e10 - -2.25)) {
                _ = len(v1)
            } else {
                var v9 uint8 = (v3 - (v3 / (((v3 / ((v3) | 1))) | 1)))
                v10 := GenS1{F0: -1, F1: true}
                _ = v10
                _ = v9
            }
        } else {
            v7 = -128
            v11 := GenS2{F0: 32767, F1: 42, F2: 0}
            v5 -= v5
            _ = v11
        }
    default:
        switch (v5 ^ (v5 ^ (v5 * v5))) {
        case 127:
            v5 -= int(v4)
        case -463:
            v12 := v1[v5]
            _ = v12
        default:
            var v13 int32 = 0
            _ = v13
        }
    }
    var v14 int = (^(v5))
    fmt.Println(v4, v14, v3)
    var v15 int32 = -112
    v4 = (-(v4))
    if false {
        v1[42] = 2147483647
    } else {
        if false {
            v4 -= (v4 | (v4 ^ v4))
            v15 = 270
        }
        fmt.Println(v5, v7)
    }
    v16 := GenS3{F0: -32768, F1: (v7 >> (((v2 * 65535)) & 63))}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (32767 >= 32767) {
        var v17 int64 = 9223372036854775807
        var v18 int8 = (v7 / (((v7 >> (((v2 * (v2 >> (((v2 >> ((v2) & 63))) & 63)))) & 63))) | 1))
        var v19 int8 = 42
        _ = v17
        _ = v18
        _ = v19
    } else {
        var v20 int = (v14 | (v14 >> ((v2) & 63)))
        v21 := GenS4{F0: -175}
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v3 += v3
        fmt.Println(v4, v3, v15)
        if len(v6) > 0 {
            v6[0] = (-(v5))
        }
        _ = fi22
    }
    var v23 float64 = -2.25
    delete(v1, (-(v14)))
    var v24 float64 = v23
    var v25 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
