package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int64; F1 uint8; F2 uint32 }
func genFn0(p0 int8, p1 uint, p2 uint8) int16 {
    var v6 uint = p1
    var v7 uint = 714
    var v8 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -32768
}
func genFn1(p0 uint16, p1 int32) uint64 {
    if true {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 9223372036854775807, F1: 2147483648, F2: 0}
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS2{F0: 89, F1: 18446744073709551615}
    p0 += p0
    if true {
        fmt.Println(p0, p0)
    }
    var v12 float32 = 1e10
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := []int{1}
    var v2 bool = true
    var v3 uint64 = 2
    var v4 uint = 409
    var v5 uint16 = 255
    switch ("" + "a") {
    case "a":
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 int16 = 1000
            v4 += (^(v4))
            _ = fi13
            _ = v14
        }
    case "hello":
        fmt.Println(v4, v3)
    case "\n\t":
        v4 = (v4 >> (((v4 % ((255) | 1))) & 63))
    default:
        var v15 float32 = (1.5 / ((1e10 + 1.5) + (float32(0.0) * (1.5 / -2.25))))
        _ = v15
    }
    v5 += (v5 ^ (^(v5)))
    v16 := GenS3{F0: 9223372036854775807, F1: 126, F2: 2147483648}
    v3 = v3
    if true {
        var v17 int = -9223372036854775808
        if (uint(v5) >= (v4 >> (((v4 | 0)) & 63))) {
            var v18 uint32 = 2
            _ = v18
        } else {
            for fi19 := 0; fi19 < 18; fi19++ {
                v17 += 9223372036854775807
                v20 := GenS1{F0: -1, F1: uint32(v3), F2: -128}
                v5 -= (v5 ^ v5)
                _ = fi19
                _ = v20
            }
        }
        _ = v17
    }
    var v21 uint8 = uint8(v4)
    switch ((("\n\t" + "a") + (("foo bar" + "a") + "")) + "foo bar") {
    case "\n\t":
        switch 0 {
        case 32768:
            var v22 int64 = 965
            _ = v22
        default:
            var v23 float64 = float64(0.0)
            _ = v23
        }
    }
    fmt.Println(v2, v3)
    var v24 uint = v4
    v24 -= (v24 << ((uint(v5)) & 63))
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint16 = v5
        _ = fi25
        _ = v26
    }
    if (406 < 42) {
        fmt.Println(v21)
    }
    v4 = v24
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
