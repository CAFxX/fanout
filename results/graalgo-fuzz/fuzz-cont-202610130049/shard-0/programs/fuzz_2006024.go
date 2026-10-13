package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 string }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int8; F1 bool; F2 uint8 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]bool)
    var v3 float32 = (((0.1 * ((float32(0.1) / 0.1) * 0.1)) * 1.5) + -2.25)
    var v4 float32 = v3
    v5 := []uint32{}
    var v6 bool = ((1e10 >= float32(3.14159)) || true)
    switch "" {
    case "\n\t":
        v4 += (v4 * 0.0)
    default:
        if v6 {
            fmt.Println(v6)
        } else {
            v1 += v1
            if (199 != 93) {
                for fi7 := 0; fi7 < 17; fi7++ {
                    v8 := GenS1{F0: 42, F1: 542, F2: "\n\t"}
                    if v6 {
                        v6 = v6
                        var v9 int16 = 0
                        v4 = ((3.14159 + ((1e10 - v4) * v4)) + (1e10 * float32(-2.25)))
                        _ = v9
                    }
                    _ = fi7
                    _ = v8
                }
            } else {
                switch 32767 {
                case 128:
                    v2[4294967295] = (0 != 4294967295)
                default:
                    v2[1] = v6
                }
                v10 := GenS2{F0: v1, F1: ("" + "\n\t")}
                v11 := GenS2{F0: 0, F1: "a"}
                _ = v10
                _ = v11
            }
            var v12 float64 = 1e10
            _ = v12
        }
    }
    if (4294967295 == 553) {
        switch 0 {
        case 2147483647:
            var v13 int8 = 42
            _ = v13
        case 127:
            _ = len(v5)
        case -9223372036854775808:
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 uint16 = 255
                fmt.Println(v4, v4)
                _ = fi14
                _ = v15
            }
        }
    }
    fmt.Println(v6, v3, v6)
    _ = len(v5)
    v16 := GenS2{F0: v1, F1: "a"}
    fmt.Println(v1)
    var v17 uint32 = 288
    v18 := GenS3{F0: int8(v1), F1: (!v6), F2: 42}
    for fi19 := 0; fi19 < 6; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 uint16 = 0
            _ = fi20
            _ = v21
        }
        v1 += (^(v1))
        var v22 uint16 = 42
        _ = fi19
        _ = v22
    }
    if (!v6) {
        fmt.Println(v1, v1)
        v1 -= v1
        v23 := GenS2{F0: (v1 >> ((6) & 63)), F1: "foo bar"}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
