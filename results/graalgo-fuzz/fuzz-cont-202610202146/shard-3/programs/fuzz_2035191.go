package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint16 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []uint32{1, 1, 908, 2}
    var v2 float32 = 0.0
    var v3 float32 = 3.14159
    v4 := GenS1{F0: 2147483647, F1: 0, F2: 255}
    switch "foo bar" {
    case "hello":
        var v5 uint16 = 0
        _ = v5
    case "":
        v2 = float32(float64(v3))
    case "foo bar":
        _ = len(v1)
    }
    var v6 uint = 2
    v3 -= (float32(1.5) * (0.1 / v2))
    v7 := GenS1{F0: 42, F1: 65535, F2: 238}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: (fi9 * fi8), F1: 218, F2: 2}
            var v11 uint32 = 42
            switch fi8 {
            case -507:
                if (((697 <= 810) || false) && false) {
                    fmt.Println(v11, fi8)
                    v11 += v11
                }
            default:
                v12 := GenS1{F0: fi9, F1: v11, F2: 0}
                _ = v12
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        var v13 float64 = 0.0
        _ = fi8
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            if (-2.25 < ((0.1 + -2.25) / 3.14159)) {
                var v16 int = fi15
                fmt.Println(fi14)
                _ = v16
            }
            v6 = (v6 * (^(v6)))
            v2 = v2
            v17 := GenS2{F0: (!(1 <= 113))}
            _ = fi15
            _ = v17
        }
        _ = fi14
    }
    var v18 int = 42
    fmt.Println(v6)
    fmt.Println(v3)
    var v19 uint32 = uint32(v6)
    for fi20 := 0; fi20 < 2; fi20++ {
        for fi21 := 0; fi21 < 2; fi21++ {
            v19 = (v19 % ((v19) | 1))
            _ = fi21
        }
        if ((255 <= 32767) && (!(0 >= -32768))) {
            v6 = v6
            var v22 int = 128
            v22 += v18
            _ = v22
        } else {
            fmt.Println(v3, v18)
            v18 = v18
        }
        _ = fi20
    }
    switch "\n\t" {
    case "\n\t":
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 int64 = -991
            _ = fi23
            _ = v24
        }
    case "foo bar":
        var v25 float32 = ((1e10 / 1.5) / (v3 / 1.5))
        _ = v25
    default:
        switch 32767 {
        case -1:
            v2 = v2
        case -991:
            var v26 uint = 1
            _ = v26
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
