package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int8 }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint16) uint64 {
    v7 := GenS1{F0: 100, F1: 2147483648, F2: 56}
    switch 128 {
    case 128:
        p0 += (p0 / ((p0) | 1))
    case -1:
        var v8 float64 = (3.14159 + (1.5 / (1.5 + 0.0)))
        _ = v8
    default:
        var v9 int64 = -1
        _ = v9
    }
    p0 = p0
    var v10 string = ("" + "foo bar")
    _ = p0
    _ = v10
    _ = v7
    return 132
}
func genFn1(p0 uint64, p1 int) uint16 {
    var v11 float32 = -2.25
    var v12 float64 = 1e10
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := []int{}
            v11 += (0.1 + (v11 * 1e10))
            _ = fi14
            _ = v15
        }
        var v16 uint16 = 546
        v17 := GenS1{F0: -128, F1: 689, F2: 127}
        _ = fi13
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 uint64 = 2
    var v2 uint32 = 65535
    v3 := []int{-1, -9223372036854775808}
    v4 := make(map[string]bool)
    var v5 int8 = 0
    v6 := make(map[int]uint32)
    v2 = (v2 % (((v2 - (v2 | 4294967295))) | 1))
    var v18 float64 = 0.0
    switch ("a" + "hello") {
    case "\n\t":
        v19 := GenS1{F0: v5, F1: (v2 >> ((2) & 63)), F2: -128}
        _ = v19
    }
    var v20 uint8 = 35
    var v21 int = 0
    v22 := GenS2{F0: false, F1: v2}
    var v23 int64 = 0
    fmt.Println(v23, v23)
    var v24 uint32 = 996
    v24 += (^(v24))
    var v25 uint64 = 311
    if (-571 != int64(v21)) {
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 float64 = float64((((float32(v18) / 3.14159) - -2.25) - 0.1))
            var v28 int32 = 832
            v23 += (v23 / (((v23 & 664)) | 1))
            _ = fi26
            _ = v27
            _ = v28
        }
        v25 = 1
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
