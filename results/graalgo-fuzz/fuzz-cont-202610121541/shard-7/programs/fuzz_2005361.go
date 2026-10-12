package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 int8 = 42
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 bool = (-1 != 0)
    var v5 float32 = (((float32(0.0) - 1e10) / 1e10) - (3.14159 * (float32(0.1) * (1e10 + -2.25))))
    var v6 uint64 = 1
    v4 = (65535 < 2147483648)
    v7 := GenS1{F0: 359}
    for fi8 := 0; fi8 < 7; fi8++ {
        delete(v3, 2147483647)
        _ = fi8
    }
    var v9 float32 = 0.1
    v10 := []int{0}
    if (v5 < float32(float64(v9))) {
        v5 += v5
        if false {
            v9 += v5
            v9 -= (v5 * (3.14159 - v5))
        } else {
            v5 = v5
            v11 := GenS2{F0: v6, F1: 143, F2: 1}
            _ = v11
        }
        if (127 < 214) {
            var v12 int8 = v1
            v6 = v6
            var v13 bool = (159 == 2)
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 17; fi14++ {
                fmt.Println(v4)
                for fi15 := 0; fi15 < 18; fi15++ {
                    fmt.Println(v6)
                    v16 := GenS3{F0: 0}
                    v4 = false
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        }
    }
    v4 = ((1 < -55) && (v1 != (v1 - (v1 - (v1 / ((v1) | 1))))))
    var v17 float32 = (float32(0.0) + ((float32((((1e10 * (((float64(float32((float64(v9) - (3.14159 - 1.5)))) / 3.14159) + 1.5) * 1e10)) / 1.5) * 1e10)) * v9) * v9))
    v18 := v3[42]
    if (false && (0 != (v6 << ((65535) & 63)))) {
        v19 := GenS2{F0: v6, F1: v6, F2: int64(v6)}
        var v20 int16 = 32767
        _ = v19
        _ = v20
    } else {
        v17 += (v9 / (1e10 / ((((3.14159 - 1e10) / float32(((0.0 * 1.5) * (float64((1e10 - v17)) * (0.0 * 1.5))))) * v9) * (v9 + v17))))
        v6 -= v6
        v5 += float32(3.14159)
    }
    switch 2147483647 {
    case 189:
        var v21 int32 = 42
        _ = v21
    case -236:
        var v22 int64 = -9223372036854775808
        _ = v22
    case -9223372036854775808:
        v1 = (v1 & v1)
    }
    if (v1 <= 125) {
        delete(v3, 128)
        var v23 uint64 = v6
        var v24 int8 = v1
        _ = v23
        _ = v24
    }
    v25 := GenS3{F0: 1}
    v26 := GenS4{F0: -621}
    if len(v10) > 0 {
        v10[0] = int(v6)
    }
    v27 := GenS3{F0: 255}
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
