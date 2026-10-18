package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 uint32; F2 int32 }
func genFn0(p0 int32, p1 bool, p2 int32) int64 {
    v4 := GenS1{F0: "foo bar", F1: -9223372036854775808}
    fmt.Println(p0, p1)
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -950
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 42
    var v3 int16 = 1
    if ((false || false) || true) {
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 uint = 1
            v7 := GenS2{F0: 18446744073709551615, F1: (v2 << (((v6 / ((v6) | 1))) & 63)), F2: -2147483648}
            _ = fi5
            _ = v6
            _ = v7
        }
        v3 += (v3 & (v3 >> ((4294967295) & 63)))
        fmt.Println(v3)
    }
    fmt.Println(v2, v2)
    v3 = 42
    fmt.Println(v2, v2)
    v8 := []int{}
    switch 128 {
    case 0:
        v9 := GenS2{F0: 18446744073709551615, F1: v2, F2: 2147483647}
        _ = v9
    case 256:
        v10 := GenS3{F0: 108, F1: -901}
        _ = v10
    case 1:
        v3 += v3
    }
    v11 := GenS2{F0: 42, F1: v2, F2: -1}
    v2 -= (-(v2))
    delete(v1, 2)
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 float32 = ((1.5 - 3.14159) / (3.14159 / -2.25))
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS4{F0: 4294967295, F1: 0, F2: 128}
            var v16 bool = true
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 string = ("" + "a")
        _ = fi12
        _ = v13
        _ = v17
    }
    var v18 uint8 = 128
    v19 := v1[42]
    v2 = (v2 / ((9223372036854775807) | 1))
    switch 695 {
    case 32768:
        v20 := GenS5{F0: (42 == -999)}
        _ = v20
    }
    switch "" {
    case "foo bar":
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 int32 = 539
            _ = fi21
            _ = v22
        }
    }
    switch 128 {
    case -9223372036854775808:
        v19 = v19
    default:
        switch 42 {
        case 286:
            v23 := GenS6{F0: 255, F1: 1, F2: 42}
            _ = v23
        default:
            _ = len(v8)
        }
    }
    var v24 uint8 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v8
}
