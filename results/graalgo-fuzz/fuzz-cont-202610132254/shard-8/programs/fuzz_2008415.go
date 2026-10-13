package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int16 }
type GenS2 struct { F0 int64; F1 int16; F2 uint8 }
func genFn0(p0 uint32, p1 uint64) int64 {
    v6 := []int{-81}
    v7 := GenS1{F0: 2, F1: 0, F2: 0}
    var v8 uint32 = p0
    var v9 uint = 4294967295
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 904
}
func genFn1(p0 int8, p1 uint32) bool {
    p0 += (p0 + (-(p0)))
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(p0, p1, p1)
        var v11 float64 = 0.0
        _ = fi10
        _ = v11
    }
    var v12 uint8 = 1
    var v13 int8 = (p0 | p0)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return ((!true) || (true || ((-32768 < 1) && false)))
}

func main() {
    var v1 int64 = 802
    var v2 int16 = int16(v1)
    v3 := []int{-472, -725}
    v4 := []uint32{2147483648, 201, 0, 2147483648}
    v5 := make(map[uint32]int)
    var v14 uint16 = 65535
    if (1 <= 42) {
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 float64 = -2.25
            for fi17 := 0; fi17 < 13; fi17++ {
                v14 += (v14 / ((v14) | 1))
                _ = fi17
            }
            _ = fi15
            _ = v16
        }
        v18 := v5[244]
        delete(v5, 2147483648)
        _ = v18
    } else {
        v2 += v2
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v1 -= (v1 * v1)
        var v20 int64 = genFn0(42, 330)
        var v21 string = "foo bar"
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v14)
    var v22 int8 = -128
    var v23 bool = true
    if (128 != 0) {
        v24 := GenS1{F0: v14, F1: 0, F2: (v2 >> ((752) & 63))}
        var v25 float64 = 0.0
        for fi26 := 0; fi26 < 15; fi26++ {
            v27 := GenS1{F0: v14, F1: 42, F2: (-(v2))}
            v28 := GenS2{F0: (v1 >> ((617) & 63)), F1: 1, F2: 0}
            fmt.Println(v14)
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v24
        _ = v25
    } else {
        delete(v5, 842)
    }
    switch ("a" + ("hello" + "")) {
    case "hello":
        switch 923 {
        case 2147483647:
            fmt.Println(v23, v23, v1)
        default:
            fmt.Println(v23)
        }
    case "":
        fmt.Println(v23, v2, v1)
    }
    v1 = (-(v1))
    v2 = v2
    fmt.Println(v2, v22, v22)
    v22 = v22
    var v29 uint16 = (v14 * v14)
    if v23 {
        for fi30 := 0; fi30 < 6; fi30++ {
            v5[30] = -1
            _ = fi30
        }
        var v31 int16 = v2
        var v32 int = 0
        _ = v31
        _ = v32
    }
    v14 = v29
    v14 -= v29
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
