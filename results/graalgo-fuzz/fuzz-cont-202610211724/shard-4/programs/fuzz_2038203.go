package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
func genFn0(p0 bool, p1 int8) uint16 {
    fmt.Println(p1, p0)
    v5 := GenS1{F0: 206}
    if p0 {
        if (1.5 > (-2.25 - 1.5)) {
            p0 = true
            if (!false) {
                v6 := GenS1{F0: 127}
                _ = v6
            }
        } else {
            p1 -= p1
            fmt.Println(p1, p0)
            if p0 {
                var v7 uint32 = 2147483648
                v8 := GenS1{F0: 127}
                _ = v7
                _ = v8
            } else {
                switch 2147483647 {
                case 993:
                    fmt.Println(p1)
                default:
                    v9 := []int{128, 528, 9223372036854775807, -2147483648}
                    _ = v9
                }
            }
        }
    } else {
        p1 -= 102
        fmt.Println(p0, p0)
        if p0 {
            v10 := []int{9223372036854775807}
            _ = v10
        }
    }
    var v11 uint64 = 1
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return 65535
}

func main() {
    var v1 int8 = 0
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    switch "" {
    case "hello":
        var v12 int16 = 32767
        _ = v12
    default:
        var v13 int64 = 9223372036854775807
        _ = v13
    }
    switch ("foo bar" + "hello") {
    case "hello":
        var v14 int16 = 0
        _ = v14
    case "foo bar":
        v3[2147483648] = ((v1 ^ v1) >= (v1 << ((255) & 63)))
    }
    v15 := GenS1{F0: 42}
    if (!false) {
        if false {
            v1 = 127
            if false {
                fmt.Println(v1)
                var v16 float32 = 1.5
                v16 -= float32(((0.1 * float64((float32(-2.25) * v16))) + float64(3.14159)))
                _ = v16
            }
            var v17 uint16 = 965
            _ = v17
        } else {
            switch 32768 {
            case 127:
                var v18 int8 = (v1 << ((392) & 63))
                _ = v18
            case 0:
                v1 -= v1
            case -170:
                switch 127 {
                case -2147483648:
                    var v19 int16 = -1
                    _ = v19
                default:
                    if (0.0 == 1.5) {
                        v1 = 125
                        v20 := GenS1{F0: 127}
                        _ = v20
                    }
                }
            default:
                var v21 uint32 = 65535
                _ = v21
            }
        }
        v1 -= v1
    }
    v22 := GenS2{F0: -94, F1: 0, F2: 2147483647}
    v1 = v1
    v1 -= v1
    fmt.Println(v1)
    v1 = (v1 % ((v1) | 1))
    v23 := GenS2{F0: 1, F1: 44, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
