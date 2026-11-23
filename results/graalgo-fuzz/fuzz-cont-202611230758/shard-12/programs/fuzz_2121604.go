package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int }
func genFn0(p0 int8, p1 uint) int16 {
    var v5 int64 = 42
    v6 := GenS1{F0: 1, F1: (-(p0))}
    var v7 uint32 = 0
    fmt.Println(v5, p0, v5)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 1000
}
func genFn1(p0 int8) int16 {
    v8 := GenS2{F0: 1000, F1: 2147483648, F2: 0}
    var v9 string = "hello"
    v10 := []int{-2147483648, -9223372036854775808, 128, 127}
    v11 := GenS1{F0: 1, F1: 1}
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := GenS1{F0: 255, F1: -60}
            if len(v10) > 0 {
                v10[0] = (fi12 & fi12)
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return -32768
}

func main() {
    var v1 float32 = (1e10 / 0.1)
    var v2 uint32 = 65535
    v3 := []uint32{v2}
    v4 := make(map[string]uint32)
    var v15 uint8 = 127
    if len(v3) > 0 {
        v3[0] = 919
    }
    _ = len(v3)
    v2 = v2
    v16 := GenS1{F0: (v15 | 160), F1: 67}
    switch 645 {
    case 533:
        if true {
            v2 = (v2 ^ (v2 - 65535))
            fmt.Println(v15, v15, v2)
        }
    case 588:
        v17 := v4["\n\t"]
        _ = v17
    default:
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    v15 += v15
    v1 = v1
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := GenS2{F0: 513, F1: (v2 * v2), F2: 864}
        if true {
            if (true && (18446744073709551615 == 4294967295)) {
                v20 := GenS3{F0: 803, F1: 32767, F2: 2147483648}
                _ = v20
            }
            if len(v3) > 0 {
                v3[0] = v2
            }
        }
        v1 = (-2.25 * 1e10)
        _ = fi18
        _ = v19
    }
    v4[("hello" + (("hello" + ("a" + "")) + "\n\t"))] = (v2 - (v2 & 2147483648))
    fmt.Println(v15)
    switch 32767 {
    case 4294967295:
        v1 = ((((0.1 + v1) * v1) - (v1 + v1)) - v1)
    case 32768:
        v15 -= (v15 & (v15 & 1))
    }
    var v21 int = -9223372036854775808
    var v22 int32 = -1
    v23 := GenS4{F0: (-(v21))}
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 int32 = -94
        _ = fi24
        _ = v25
    }
    v22 -= (v22 ^ (v22 | v22))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
