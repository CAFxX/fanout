package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []uint32{1}
    var v2 int = -9223372036854775808
    var v3 uint = 1
    v4 := make(map[int]string)
    var v5 int8 = 117
    v2 -= v2
    var v6 uint16 = 255
    fmt.Println(v3, v3)
    v7 := v4[int(v5)]
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: v7, F1: 594, F2: 135}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: (v3 ^ (v3 % ((uint(v6)) | 1)))}
    fmt.Println(v7)
    var v11 int32 = 2147483647
    v12 := GenS3{F0: (v6 - v6)}
    fmt.Println(v11, v7, v5)
    v7 = "hello"
    v13 := GenS2{F0: 1}
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 int16 = 778
        var v16 int32 = v11
        _ = fi14
        _ = v15
        _ = v16
    }
    switch "a" {
    case "foo bar":
        v6 = (v6 ^ v6)
    case "hello":
        if len(v1) > 0 {
            v1[0] = 848
        }
    default:
        v6 += v6
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := GenS3{F0: (v6 << ((v3) & 63))}
        _ = fi17
        _ = v18
    }
    v3 += v3
    for fi19 := 0; fi19 < 14; fi19++ {
        switch -382 {
        case -668:
            fmt.Println(v2, v2)
        case -934:
            fmt.Println(v11)
        default:
            if (-2.25 != ((1.5 / (-2.25 / 0.1)) + 0.0)) {
                _ = len(v1)
                v3 += v3
                fmt.Println(v3, v2)
            }
        }
        var v20 uint8 = 121
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: (v3 << (((v3 + (v3 % ((v3) | 1)))) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
