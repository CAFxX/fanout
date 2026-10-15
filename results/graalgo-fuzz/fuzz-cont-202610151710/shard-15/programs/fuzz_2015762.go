package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint }
func genFn0() uint16 {
    var v5 uint = 1
    v5 = (v5 | 462)
    _ = v5
    return 986
}
func genFn1(p0 uint64, p1 string) uint16 {
    if false {
        fmt.Println(p0, p1)
        v6 := GenS1{F0: 72, F1: 414}
        p0 = (p0 + p0)
        _ = v6
    }
    p0 += 175
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []string{("a" + "\n\t"), ""}
    v2 := make(map[int]string)
    v3 := []uint32{}
    var v4 int = -1
    var v7 uint8 = 215
    fmt.Println(v7)
    switch v4 {
    case 0:
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(v7)
            v9 := v2[-531]
            _ = fi8
            _ = v9
        }
    case 1:
        fmt.Println(v4, v7)
    default:
        var v10 uint32 = 2147483648
        _ = v10
    }
    var v11 uint32 = 0
    delete(v2, v4)
    if false {
        if len(v3) > 0 {
            v3[0] = v11
        }
        v11 += (v11 & (v11 * (v11 % ((440) | 1))))
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS1{F0: -128, F1: 589}
            _ = fi12
            _ = v13
        }
    }
    v4 -= (v4 + v4)
    v11 = v11
    var v14 float64 = 1e10
    fmt.Println(v14, v4)
    if (((^(v7)) < uint8(v11)) && ((v4 <= (v4 % ((393) | 1))) || false)) {
        for fi15 := 0; fi15 < 10; fi15++ {
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 float64 = v14
                v18 := GenS1{F0: -128, F1: 0}
                v4 = -9223372036854775808
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi15
        }
        fmt.Println(v7, v4, v14)
    }
    switch (^(v4)) {
    case 127:
        fmt.Println(v4, v7)
    case 635:
        v14 = ((float64(float32((v14 - 0.1))) + v14) - (float64(0.1) * v14))
    default:
        v19 := GenS2{F0: 253}
        _ = v19
    }
    switch "foo bar" {
    case "":
        v1 = append(v1, "")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
