package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float32 = 0.0
    var v2 float32 = 0.1
    v3 := make(map[int]uint32)
    v4 := []uint32{2147483647, 0, 65535, 0}
    var v5 int8 = 1
    var v6 bool = (!false)
    v7 := GenS1{F0: 2, F1: -1, F2: 255}
    if v6 {
        if ((634 > 0) || (v5 != 127)) {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 uint32 = 166
                _ = fi8
                _ = v9
            }
            v6 = (v5 >= v5)
            var v10 uint16 = 1
            _ = v10
        }
    }
    v5 = -127
    var v11 bool = v6
    var v12 float64 = (float64((v1 - float32(-2.25))) * ((0.0 + 0.0) - 1.5))
    v13 := GenS1{F0: 0, F1: 42, F2: 1}
    var v14 int8 = v5
    if (1 <= 255) {
        var v15 float32 = 3.14159
        v16 := GenS2{F0: (173 == 255), F1: 0}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        fmt.Println(v11, v2)
        if v6 {
            v18 := GenS1{F0: 2, F1: fi17, F2: 671}
            var v19 float32 = (float32((v12 - (v12 - float64((0.0 / float32(v12)))))) / ((0.1 / v2) + (1.5 + (((v2 / float32(0.1)) / v2) / float32(float64((float32((v12 / float64(1e10))) / 1.5)))))))
            v20 := GenS3{F0: 2147483647}
            _ = v18
            _ = v19
            _ = v20
        } else {
            v12 = v12
            _ = len(v3)
        }
        var v21 int8 = v5
        _ = fi17
        _ = v21
    }
    v22 := GenS3{F0: -2147483648}
    switch -9223372036854775808 {
    case -2147483648:
        var v23 uint8 = 128
        _ = v23
    case -1:
        var v24 int16 = -104
        _ = v24
    case -9223372036854775808:
        _ = len(v4)
    }
    switch -878 {
    case 436:
        var v25 bool = (2 >= 42)
        _ = v25
    case -89:
        v5 = (v14 - v5)
    default:
        v1 = (float32(v12) * float32((v12 * 1e10)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
