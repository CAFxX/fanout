package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 bool }
type GenS3 struct { F0 bool; F1 int; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint64; F1 int8; F2 int64 }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint }
func genFn0(p0 uint8, p1 uint64) uint16 {
    var v4 float32 = (((0.1 * 0.0) - (0.0 / float32((1e10 * ((float64(float32(1.5)) - -2.25) / ((0.1 - (3.14159 * 1e10)) - 3.14159)))))) * 1.5)
    if false {
        v5 := GenS1{F0: (!false), F1: -629}
        v6 := GenS2{F0: (true || ((true || (-9223372036854775808 >= 9223372036854775807)) || true)), F1: true, F2: ((p0 + (p0 + (^(p0)))) <= (p0 - (p0 % ((p0) | 1))))}
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := GenS2{F0: true, F1: true, F2: false}
                v10 := GenS3{F0: (-1 != -2147483648), F1: fi8, F2: 9223372036854775807}
                switch 32768 {
                case -9223372036854775808:
                    var v11 uint8 = 255
                    _ = v11
                case 931:
                    v12 := GenS4{F0: 0, F1: 4294967295, F2: 42}
                    _ = v12
                }
                v4 = 1e10
                _ = fi8
                _ = v10
                _ = v9
            }
            for fi13 := 0; fi13 < 14; fi13++ {
                switch 128 {
                case -9223372036854775808:
                    var v14 bool = (276 < -2147483648)
                    _ = v14
                case 175:
                    var v15 bool = ((-(fi7)) >= (fi7 ^ 9223372036854775807))
                    _ = v15
                default:
                    if (false || false) {
                        fmt.Println(v4, fi13)
                        p1 = uint64(p0)
                    }
                }
                _ = fi13
            }
            _ = fi7
        }
        _ = v5
        _ = v6
    }
    var v16 int64 = 1
    _ = p0
    _ = p1
    _ = v16
    _ = v4
    return 255
}

func main() {
    var v1 int = 128
    var v2 float32 = 0.0
    var v3 uint32 = 2147483648
    v1 += v1
    fmt.Println(v3, v1)
    v1 = 0
    fmt.Println(v2)
    var v17 uint8 = 0
    fmt.Println(v2)
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := GenS5{F0: 439, F1: -1, F2: 720}
        var v20 uint = 65535
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v1, v1)
    switch "a" {
    case "a":
        var v21 int = 127
        _ = v21
    case "foo bar":
        v17 = 255
    case "hello":
        if false {
            fmt.Println(v3)
            var v22 uint32 = (v3 * v3)
            _ = v22
        }
    default:
        v23 := GenS6{F0: v3, F1: v3, F2: 364}
        _ = v23
    }
    v2 += v2
    v24 := []int{(v1 * (v1 >> ((588) & 63)))}
    var v25 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
