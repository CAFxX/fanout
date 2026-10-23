package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 bool = (((1e10 / -2.25) + (0.1 + 1.5)) != 0.1)
    var v2 string = "\n\t"
    var v3 float64 = ((((1.5 + (float64(3.14159) / -2.25)) - ((0.1 - -2.25) / float64((float32((-2.25 - 1e10)) / 1.0)))) / 1.5) - (((float64((0.1 + 1e10)) / ((3.14159 * float64(0.1)) / (1.5 + 3.14159))) + 3.14159) + 1.5))
    v4 := make(map[int]uint32)
    var v5 float64 = -2.25
    var v6 uint = 1
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int16 = 360
        v9 := GenS1{F0: 1, F1: 127}
        v5 = (-2.25 + (float64(1.5) * float64(((0.1 - -2.25) / -2.25))))
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-2147483648, 32767, 2147483647}
    v11 := GenS1{F0: 1, F1: uint8(v6)}
    fmt.Println(v2)
    fmt.Println(v5, v1, v2)
    fmt.Println(v6, v5, v3)
    v2 = v2
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 int8 = -58
        for fi14 := 0; fi14 < 9; fi14++ {
            v4[(-(fi14))] = 42
            v15 := v4[880]
            if (false && v1) {
                v10 = append(v10, 0)
                v6 -= (^(v6))
                var v16 int = (fi12 + fi12)
                _ = v16
            } else {
                if (!false) {
                    var v17 uint8 = 94
                    var v18 int64 = 216
                    _ = v17
                    _ = v18
                } else {
                    v19 := GenS2{F0: (v13 / (((v13 % ((int8(v6)) | 1))) | 1)), F1: 42}
                    _ = v19
                }
                var v20 string = "a"
                _ = len(v10)
                _ = v20
            }
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v6, v2, v2)
    var v21 int8 = 42
    var v22 int = -938
    v5 += v5
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(v1, fi23, v2)
        _ = fi23
    }
    delete(v4, (v22 ^ (^(v22))))
    switch v2 {
    case "a":
        v10 = append(v10, (v22 ^ v22))
    default:
        switch "" {
        case "":
            v24 := GenS2{F0: (v21 * v21), F1: 0}
            _ = v24
        case "hello":
            v25 := GenS3{F0: 2}
            _ = v25
        default:
            v5 = 3.14159
        }
    }
    v2 = v2
    v4[-185] = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
