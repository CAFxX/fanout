package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 128
    v3 := []string{("a" + "\n\t")}
    v4 := make(map[int]bool)
    var v5 float32 = -2.25
    var v6 int = 258
    delete(v1, 2147483648)
    v2 += (v2 << ((1) & 63))
    var v7 uint = 42
    var v8 float64 = (float64(v5) + ((float64(0.0) + ((1.5 * (((((3.14159 - (float64(v5) * (0.0 + 3.14159))) - (1e10 - 0.1)) + 0.0) / 1.5) - 0.1)) / (-2.25 + ((float64((v5 - (((0.0 + v5) * 1e10) + (0.1 / 3.14159)))) / ((0.0 / (-2.25 / 3.14159)) - (0.1 + float64((v5 + (float32(0.0) + v5)))))) - ((0.0 * float64(v5)) + -2.25))))) / (1e10 / 1.0)))
    var v9 int16 = 0
    v7 = (^(v7))
    v9 += (v9 + (v9 + v9))
    if false {
        if (v9 == (v9 * (v9 / ((v9) | 1)))) {
            var v10 uint = 255
            switch v6 {
            case 42:
                v11 := GenS1{F0: -583, F1: 1, F2: 1}
                _ = v11
            case -2147483648:
                for fi12 := 0; fi12 < 15; fi12++ {
                    v7 = (v7 - (v10 - (^(v7))))
                    var v13 uint32 = 2
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v8 -= (((0.1 - 0.1) * 0.0) / v8)
                        _ = fi14
                    }
                    _ = fi12
                    _ = v13
                }
            case -9223372036854775808:
                var v15 uint = v7
                _ = v15
            }
            _ = v10
        } else {
            var v16 int = (v6 ^ 1)
            v17 := GenS1{F0: 296, F1: 980, F2: 2147483647}
            _ = v16
            _ = v17
        }
    }
    var v18 bool = true
    v19 := GenS1{F0: 9223372036854775807, F1: 885, F2: 0}
    v20 := GenS1{F0: 9223372036854775807, F1: 602, F2: -397}
    switch "" {
    case "hello":
        var v21 int32 = -87
        _ = v21
    default:
        for fi22 := 0; fi22 < 12; fi22++ {
            v9 -= v9
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 int8 = -6
                v3 = append(v3, "\n\t")
                _ = fi23
                _ = v24
            }
            v25 := GenS1{F0: 245, F1: 144, F2: 2147483647}
            _ = fi22
            _ = v25
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
