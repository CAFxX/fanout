package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int16; F1 int64; F2 string }
type GenS3 struct { F0 bool }
func genFn0(p0 int16, p1 uint64) int16 {
    var v4 bool = ((2 <= 385) || (true || true))
    v4 = (v4 || v4)
    switch "foo bar" {
    case "":
        p0 += (p0 + 387)
    case "foo bar":
        switch -2147483648 {
        case -489:
            var v5 float32 = (-2.25 * (0.0 - ((1.5 - ((float32(1e10) + 0.0) - 1e10)) - (((1e10 + float32(3.14159)) * 0.1) - 1.5))))
            _ = v5
        case 42:
            p0 += -45
        case 2147483647:
            v6 := GenS1{F0: 2147483647, F1: 1}
            _ = v6
        }
    default:
        v4 = (255 > 18446744073709551615)
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        v4 = (!v4)
        switch fi7 {
        case 965:
            var v8 int16 = int16(fi7)
            _ = v8
        default:
            fmt.Println(p1, p1, fi7)
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 int, p1 uint) uint {
    fmt.Println(p1)
    for fi9 := 0; fi9 < 1; fi9++ {
        p0 = (p0 * p0)
        _ = fi9
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    var v3 float64 = 0.0
    var v10 int64 = 805
    var v11 int8 = 127
    v1 = (255 == 255)
    var v12 float32 = -2.25
    if (1 > 42) {
        v13 := GenS1{F0: 251, F1: 0}
        v12 += (v12 / 1e10)
        switch -9223372036854775808 {
        case -2147483648:
            v14 := GenS2{F0: 810, F1: (v10 / (((v10 / ((v10) | 1))) | 1)), F2: ("hello" + "a")}
            _ = v14
        case 33:
            v11 -= 127
        case 32767:
            var v15 uint = genFn1(2147483647, 1)
            _ = v15
        }
        _ = v13
    }
    fmt.Println(v11)
    fmt.Println(v3, v3, v3)
    v11 -= -55
    v16 := GenS3{F0: (3.14159 < v3)}
    var v17 float32 = (v12 / float32(v3))
    v12 += v12
    v11 = v11
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v10)
        v11 -= v11
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
