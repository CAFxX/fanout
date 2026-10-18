package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8; F2 bool }
type GenS3 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{}
    v3 := make(map[string]string)
    var v4 int = 128
    var v5 uint8 = 204
    var v6 uint8 = (v5 / (((v5 | (v5 + v5))) | 1))
    for fi7 := 0; fi7 < 12; fi7++ {
        v5 = v6
        if ((-(fi7)) < fi7) {
            var v8 uint16 = 65535
            _ = v8
        }
        _ = fi7
    }
    if len(v2) > 0 {
        v2[0] = 790
    }
    var v9 int16 = -399
    if true {
        v6 = 215
        var v10 uint = uint(v6)
        v5 += v5
        _ = v10
    } else {
        fmt.Println(v5, v9, v6)
    }
    var v11 uint32 = 0
    v6 = 144
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS1{F0: "\n\t", F1: 420, F2: v11}
            var v15 string = "hello"
            var v16 float32 = 0.0
            var v17 float64 = (((1.5 - 1e10) * 3.14159) / (float64(1.5) * 1.5))
            var v18 float32 = v16
            var v19 float64 = v17
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
            _ = v17
            _ = v18
            _ = v19
        }
        var v20 int32 = -820
        _ = fi12
        _ = v20
    }
    v11 += 2147483647
    fmt.Println(v4)
    v21 := GenS2{F0: -128, F1: 127, F2: true}
    v22 := GenS3{F0: int16(v6), F1: 100}
    fmt.Println(v5, v6, v9)
    var v23 int16 = 226
    v4 -= v4
    v9 = (v9 - 42)
    switch (^(v4)) {
    case 0:
        fmt.Println(v23, v4)
    default:
        fmt.Println(v23)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
