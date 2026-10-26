package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint }
func genFn0(p0 uint) int32 {
    var v6 int = -887
    fmt.Println(v6, p0)
    _ = p0
    _ = v6
    return 390
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]uint32)
    var v3 int16 = 443
    v4 := []int{}
    v5 := make(map[string]uint32)
    if (-998 > genFn0(65535)) {
        v3 += 155
        for fi7 := 0; fi7 < 17; fi7++ {
            switch (^(fi7)) {
            case 707:
                v3 -= (v3 | v3)
            }
            v4 = append(v4, (^(fi7)))
            _ = fi7
        }
        v5[""] = 4294967295
    }
    fmt.Println(v3, v3)
    v3 += (v3 ^ v3)
    v3 -= 163
    if (false && false) {
        for fi8 := 0; fi8 < 6; fi8++ {
            v1 = (v1 + v1)
            _ = fi8
        }
        var v9 int64 = -9223372036854775808
        var v10 float64 = -2.25
        _ = v10
        _ = v9
    } else {
        var v11 int32 = -2147483648
        _ = v11
    }
    if false {
        var v12 string = "\n\t"
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint16 = v1
            var v15 bool = ((v14 * v14) <= (v14 ^ (v14 / ((v1) | 1))))
            switch v12 {
            case "foo bar":
                v16 := GenS1{F0: 1}
                _ = v16
            default:
                v14 = 32767
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        v17 := GenS2{F0: 818, F1: -2147483648, F2: 1}
        _ = v12
        _ = v17
    }
    v3 = (v3 & v3)
    var v18 int = -225
    var v19 uint32 = 42
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := GenS1{F0: 35}
        var v22 uint32 = 2147483648
        v23 := GenS2{F0: v22, F1: (v18 & 364), F2: 255}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    if (((221 >= 0) || false) && (36 < -1)) {
        v5["a"] = v19
        fmt.Println(v1, v19)
        v3 = (v3 % ((v3) | 1))
    } else {
        var v24 bool = false
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
