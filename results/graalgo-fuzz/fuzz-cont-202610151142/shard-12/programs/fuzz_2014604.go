package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16; F1 uint64; F2 int8 }
func genFn0(p0 string, p1 int8, p2 int64) uint64 {
    v7 := GenS1{F0: 2147483647, F1: 12}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint = 0
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 bool) string {
    var v11 uint32 = 2147483647
    var v12 float32 = 0.1
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := []int{128, fi13, 42, (-(fi13))}
        _ = fi13
        _ = v14
    }
    fmt.Println(v11, v11)
    v12 = v12
    _ = p0
    _ = v11
    _ = v12
    return ((("\n\t" + (("" + "\n\t") + "foo bar")) + "hello") + "a")
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 349
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    var v5 uint64 = 808
    var v6 float32 = 0.0
    fmt.Println(v5, v2, v5)
    _ = len(v1)
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v5, fi15, fi15)
        v16 := GenS1{F0: 220, F1: 878}
        _ = fi15
        _ = v16
    }
    switch "foo bar" {
    case "\n\t":
        v17 := GenS1{F0: 2147483647, F1: 18446744073709551615}
        _ = v17
    case "a":
        v4[-372] = 992
    default:
        var v18 float64 = float64(float32((float64(((v6 * (v6 / (v6 / 0.1))) * -2.25)) * 0.0)))
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 uint16 = 851
        v20 += 1
        _ = fi19
        _ = v20
    }
    var v21 float32 = (v6 - 3.14159)
    var v22 uint64 = v5
    v5 -= (v22 % ((genFn0(("hello" + "hello"), -128, 9223372036854775807)) | 1))
    if (v21 == v6) {
        for fi23 := 0; fi23 < 8; fi23++ {
            if true {
                fmt.Println(v2)
                _ = len(v1)
            }
            _ = fi23
        }
        v24 := v4[127]
        _ = v24
    } else {
        fmt.Println(v21)
        v6 += v21
        _ = len(v3)
    }
    v25 := GenS2{F0: 2, F1: v22, F2: -128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
