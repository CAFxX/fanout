package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint8 }

func main() {
    var v1 float32 = 1.5
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    var v4 int32 = v2
    v5 := make(map[string]uint32)
    v2 = 894
    v6 := GenS1{F0: true, F1: 101}
    if (9223372036854775807 >= 9223372036854775807) {
        var v7 int32 = (v2 / (((v2 >> ((255) & 63))) | 1))
        fmt.Println(v1)
        _ = v7
    } else {
        v8 := GenS2{F0: -398, F1: 56}
        v1 -= v1
        switch -1 {
        case -494:
            if false {
                v2 = (v2 * v2)
                v1 -= v1
                v9 := GenS2{F0: 0, F1: 164}
                _ = v9
            }
        case 127:
            var v10 uint16 = 255
            _ = v10
        case -52:
            if ((1 == 380) && (0 <= 2)) {
                var v11 int16 = -124
                var v12 float64 = 0.1
                fmt.Println(v4, v1, v4)
                _ = v11
                _ = v12
            } else {
                var v13 int64 = 9223372036854775807
                _ = v13
            }
        }
        _ = v8
    }
    v14 := v5["\n\t"]
    v4 = v4
    v15 := v5[(("\n\t" + "hello") + (("a" + "\n\t") + "hello"))]
    var v16 int16 = 0
    if (!true) {
        var v17 uint8 = 255
        fmt.Println(v14, v1)
        _ = v17
    }
    var v18 int64 = 1
    v19 := []int{-9223372036854775808}
    var v20 float64 = (1e10 / ((0.0 - 1.5) * 3.14159))
    _ = len(v3)
    if false {
        v2 = v4
        v4 -= v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
