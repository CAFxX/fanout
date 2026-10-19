package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int; F1 uint; F2 uint }
type GenS3 struct { F0 uint8; F1 int16; F2 uint32 }
func genFn0(p0 uint8, p1 int8, p2 int8) string {
    var v6 int64 = 855
    v6 = (^(v6))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ("foo bar" + ("\n\t" + "\n\t"))
}
func genFn1() uint32 {
    var v7 bool = false
    v7 = (1.5 < (((0.0 * 0.0) + 3.14159) * (-2.25 / 3.14159)))
    _ = v7
    return 66
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 int8 = 1
    var v4 uint = 144
    var v5 bool = true
    v8 := []int{736, 42, 127, 42}
    if v5 {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 1, F1: false}
            var v11 int8 = (v3 << ((v4) & 63))
            v11 += -1
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint16 = 65535
        _ = v12
    }
    fmt.Println(v3, v3, v5)
    fmt.Println(v3)
    switch "foo bar" {
    case "a":
        v5 = (363 == 18446744073709551615)
    default:
        v13 := GenS2{F0: 9223372036854775807, F1: (v4 >> (((v4 | 42)) & 63)), F2: v4}
        _ = v13
    }
    if ((true || false) || (!((^(v4)) == (v4 - (v4 * v4))))) {
        var v14 float64 = (-2.25 * (0.0 + 1e10))
        var v15 int = 127
        if len(v8) > 0 {
            v8[0] = (^(v15))
        }
        _ = v14
        _ = v15
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        v3 -= (v3 ^ (v3 & (v3 / (((v3 << (((v4 | v4)) & 63))) | 1))))
        _ = fi16
    }
    var v17 int16 = 1000
    if (true && v5) {
        var v18 int32 = 501
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        if v5 {
            var v20 int32 = 0
            var v21 float64 = (1e10 - 1e10)
            _ = v20
            _ = v21
        }
        var v22 int64 = -601
        var v23 uint16 = 255
        _ = fi19
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS3{F0: 129, F1: (-(v17)), F2: 478}
        v3 -= (^(v3))
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
