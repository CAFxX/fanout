package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 float32 = float32(1.5)
    v4 := make(map[uint32]uint32)
    v5 := []int{128}
    var v6 uint = 65535
    for fi7 := 0; fi7 < 12; fi7++ {
        if true {
            var v8 uint64 = 0
            var v9 uint64 = (v8 >> ((v6) & 63))
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 88}
            _ = v10
        }
        v3 -= ((((((v3 * v3) - (1e10 + v3)) / (((v3 / 1.0) - v3) - ((float32(float64((v3 + (float32((3.14159 + 0.1)) - (v3 - (v3 - -2.25)))))) - (float32(((1e10 + ((0.1 * ((((1e10 * 1.5) - (0.0 - 1e10)) + -2.25) / (-2.25 + -2.25))) - 1e10)) * 3.14159)) * v3)) + ((v3 - v3) - 0.1)))) / (1.5 * 0.1)) * (((v3 / v3) - v3) / v3)) * v3)
        v11 := GenS2{F0: 4294967295, F1: 4294967295, F2: 9223372036854775807}
        _ = fi7
        _ = v11
    }
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    v12 := GenS2{F0: (^(v6)), F1: 248, F2: 9223372036854775807}
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS3{F0: 119, F1: 2, F2: -728}
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 160
    v15 -= v15
    v1[-2147483648] = "a"
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 bool = (885 == 42)
        switch "\n\t" {
        case "\n\t":
            v18 := GenS4{F0: 703}
            _ = v18
        default:
            switch "" {
            case "\n\t":
                v19 := GenS5{F0: v15}
                _ = v19
            case "hello":
                var v20 int32 = -905
                _ = v20
            case "a":
                for fi21 := 0; fi21 < 18; fi21++ {
                    v22 := GenS3{F0: 127, F1: 0, F2: 32767}
                    v17 = true
                    _ = fi21
                    _ = v22
                }
            default:
                var v23 int64 = 643
                _ = v23
            }
        }
        _ = fi16
        _ = v17
    }
    v24 := GenS3{F0: v15, F1: 4294967295, F2: 107}
    switch 2147483647 {
    case -35:
        v1[128] = "\n\t"
    default:
        if (2147483648 > 740) {
            fmt.Println(v15, v6, v6)
            v6 -= v6
        }
    }
    v6 -= v6
    fmt.Println(v6, v15)
    var v25 float32 = (3.14159 - (1.5 + (v3 / v3)))
    var v26 uint16 = 42
    delete(v4, 2147483648)
    var v27 uint = v6
    v28 := GenS1{F0: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
