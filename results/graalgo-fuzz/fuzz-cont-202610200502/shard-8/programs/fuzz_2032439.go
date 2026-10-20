package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint8) uint8 {
    if true {
        fmt.Println(p0)
        var v6 uint8 = 127
        if true {
            var v7 uint64 = 115
            _ = v7
        }
        _ = v6
    }
    p0 = p0
    _ = p0
    return 42
}
func genFn1(p0 int) int16 {
    var v8 float64 = 0.0
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int = (p0 | (p0 % ((32767) | 1)))
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 32767}
    var v12 uint32 = 68
    var v13 float32 = 0.0
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    _ = v8
    return 42
}

func main() {
    var v1 uint8 = 79
    var v2 float32 = (3.14159 - (0.1 / 1e10))
    var v3 float64 = float64(v2)
    var v4 uint = 1
    v5 := []string{"hello", ""}
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint8 = v1
        switch fi14 {
        case 9223372036854775807:
            var v16 int8 = 42
            _ = v16
        case -1:
            var v17 float64 = v3
            _ = v17
        case -9223372036854775808:
            v18 := GenS2{F0: 2, F1: 100}
            _ = v18
        default:
            if (1 > 0) {
                v4 -= (v4 << ((4294967295) & 63))
                v19 := GenS2{F0: 1, F1: 100}
                _ = v19
            } else {
                _ = len(v5)
                switch "\n\t" {
                case "a":
                    var v20 uint32 = 1
                    _ = v20
                case "foo bar":
                    v2 = v2
                case "hello":
                    v2 += -2.25
                default:
                    _ = len(v5)
                }
                fmt.Println(v1, v1, fi14)
            }
        }
        _ = fi14
        _ = v15
    }
    var v21 uint64 = 1
    fmt.Println(v4)
    v2 = v2
    fmt.Println(v2)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        if ((v3 - v3) >= ((float64((float32((((-2.25 * (v3 * (v3 + v3))) + 0.1) * v3)) + v2)) * float64((v2 + v2))) - v3)) {
            v5 = append(v5, "hello")
            var v23 int64 = 366
            _ = v23
        } else {
            var v24 uint16 = 32767
            _ = v24
        }
        v21 -= v21
        _ = fi22
    }
    _ = len(v5)
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 uint8 = v1
        fmt.Println(v21, v21)
        _ = fi25
        _ = v26
    }
    fmt.Println(v21)
    v1 = (^(v1))
    v27 := GenS2{F0: 108, F1: 1}
    v1 -= genFn0((v1 ^ v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
