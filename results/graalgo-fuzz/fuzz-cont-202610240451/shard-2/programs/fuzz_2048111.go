package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint64; F2 int16 }
func genFn0(p0 string) bool {
    var v5 int16 = 781
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 = p0
        _ = fi6
    }
    p0 = p0
    fmt.Println(p0, p0)
    v7 := GenS1{F0: 65535}
    _ = p0
    _ = v5
    _ = v7
    return true
}
func genFn1(p0 uint64, p1 uint, p2 int8) string {
    fmt.Println(p1)
    v8 := GenS2{F0: 1000, F1: 65535, F2: 0}
    var v9 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return ("" + (("" + ("" + "\n\t")) + "\n\t"))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 int = 4294967295
    v4 := []string{"\n\t", "a", "\n\t"}
    var v10 int = (v3 | 177)
    v11 := GenS2{F0: int16(v10), F1: 4294967295, F2: 128}
    var v12 int = 127
    var v13 int64 = 0
    if true {
        fmt.Println(v12)
    }
    if false {
        v14 := GenS3{F0: 1, F1: 76, F2: 42}
        var v15 float32 = -2.25
        var v16 float32 = ((v15 - v15) * (float32(3.14159) / v15))
        _ = v14
        _ = v15
        _ = v16
    } else {
        v10 = v3
        v12 += (v10 >> ((199) & 63))
    }
    v13 = (v13 + -564)
    if true {
        v13 -= (v13 % (((v13 + (v13 % ((v13) | 1)))) | 1))
        var v17 float64 = 1.5
        _ = v17
    } else {
        var v18 uint32 = 613
        v19 := GenS1{F0: 42}
        for fi20 := 0; fi20 < 10; fi20++ {
            fmt.Println(v10, fi20, v3)
            if (67 >= 1) {
                var v21 string = "\n\t"
                v12 -= (-(v3))
                _ = v21
            }
            var v22 float64 = 1.5
            _ = fi20
            _ = v22
        }
        _ = v18
        _ = v19
    }
    v1[686] = (false || (false || (18446744073709551615 < 1)))
    var v23 uint64 = 4294967295
    var v24 uint8 = 106
    v13 = (v13 + v13)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
