package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 127
    var v3 bool = false
    v4 := GenS1{F0: 127, F1: 4294967295, F2: (v2 | 1)}
    if (0.0 >= 0.1) {
        var v5 int64 = 1
        var v6 uint32 = 235
        _ = v5
        _ = v6
    } else {
        _ = len(v1)
        if v3 {
            v2 = v2
            v7 := GenS2{F0: "", F1: 255, F2: 473}
            _ = v7
        } else {
            var v8 int16 = 1
            if false {
                v2 += (v2 >> ((255) & 63))
                var v9 int8 = 127
                v10 := GenS1{F0: 0, F1: 15, F2: v9}
                _ = v10
                _ = v9
            } else {
                v2 = (v2 >> ((4294967295) & 63))
            }
            switch 32767 {
            case 128:
                v3 = true
            }
            _ = v8
        }
    }
    fmt.Println(v2)
    if v3 {
        var v11 float64 = float64(float32((1e10 + 0.0)))
        v11 = (1e10 * ((float64(0.1) * (float64(1e10) * v11)) + float64(-2.25)))
        v2 += 42
        _ = v11
    }
    var v12 uint = 4294967295
    v2 = (v2 ^ v2)
    v13 := v1[2147483648]
    v13 = v13
    var v14 int32 = int32(v12)
    var v15 float32 = 0.0
    var v16 int16 = 1
    v14 = (v14 + (v14 % (((v14 * 2147483647)) | 1)))
    v17 := GenS2{F0: "\n\t", F1: 233, F2: (v12 ^ 827)}
    var v18 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
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
