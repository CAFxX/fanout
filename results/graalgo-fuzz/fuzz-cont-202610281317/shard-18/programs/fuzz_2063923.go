package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint64 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]uint32)
    var v3 uint32 = 263
    var v4 int64 = 42
    v5 := GenS1{F0: 229, F1: -1}
    var v6 int32 = -2147483648
    for fi7 := 0; fi7 < 18; fi7++ {
        if (439 <= 42) {
            v8 := v2[(fi7 & fi7)]
            _ = v8
        }
        _ = fi7
    }
    var v9 int = 128
    if (319 != 632) {
        v3 += (v3 & (v3 + (v3 << ((597) & 63))))
    }
    if (488 < 4294967295) {
        var v10 bool = ((^(v3)) != 65535)
        fmt.Println(v4)
        var v11 int8 = -128
        _ = v10
        _ = v11
    }
    if false {
        var v12 bool = (true && (!true))
        var v13 uint64 = 42
        _ = v12
        _ = v13
    }
    var v14 uint32 = (v3 ^ (v3 / (((v3 << ((255) & 63))) | 1)))
    var v15 int16 = -32768
    fmt.Println(v14, v3)
    var v16 bool = (v9 >= -821)
    v17 := GenS2{F0: 42, F1: 4294967295}
    v18 := GenS2{F0: 128, F1: 255}
    if v16 {
        fmt.Println(v4)
        v6 = 0
        v19 := GenS2{F0: 255, F1: 42}
        _ = v19
    } else {
        var v20 int16 = int16(v6)
        var v21 float64 = (-2.25 / 3.14159)
        if (false || v16) {
            for fi22 := 0; fi22 < 7; fi22++ {
                switch v9 {
                case -558:
                    for fi23 := 0; fi23 < 16; fi23++ {
                        var v24 uint32 = v14
                        _ = fi23
                        _ = v24
                    }
                case 331:
                    var v25 int16 = v20
                    _ = v25
                }
                if (v6 >= v6) {
                    for fi26 := 0; fi26 < 14; fi26++ {
                        v27 := v2[fi26]
                        v28 := []int{fi26, (v9 * -9223372036854775808)}
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(v3)
                } else {
                    v29 := GenS1{F0: 127, F1: 1000}
                    v3 += v3
                    _ = v29
                }
                _ = fi22
            }
        }
        _ = v20
        _ = v21
    }
    v4 = v4
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 int = (v9 * (^(fi30)))
        v32 := GenS3{F0: 0, F1: 210, F2: 956}
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
