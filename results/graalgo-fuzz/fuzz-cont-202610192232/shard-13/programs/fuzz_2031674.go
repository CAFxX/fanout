package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int8 }
func genFn0(p0 int8, p1 int64, p2 int16) uint {
    var v5 int64 = p1
    if (843 <= 2) {
        fmt.Println(v5)
        fmt.Println(p2, p0)
        var v6 uint8 = 176
        _ = v6
    } else {
        var v7 bool = (-501 == 5)
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    v1 := []int{-329}
    var v2 uint = 4294967295
    v3 := make(map[uint32]int)
    var v4 float64 = -2.25
    v2 += genFn0(97, 0, 58)
    v4 -= v4
    v8 := GenS1{F0: 32767, F1: v2, F2: -78}
    for fi9 := 0; fi9 < 5; fi9++ {
        if (false || false) {
            var v10 float32 = 1.5
            _ = v10
        } else {
            v1 = append(v1, (fi9 ^ (fi9 | 1)))
            v2 += (v2 - (v2 << (((v2 * v2)) & 63)))
        }
        if true {
            fmt.Println(v2, fi9, fi9)
        }
        var v11 float32 = (((float32((((v4 - (1e10 - (v4 + v4))) - v4) / -2.25)) * (float32(0.0) + -2.25)) + 3.14159) - 1e10)
        _ = fi9
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: v2, F2: 0}
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(fi14)
            _ = fi14
        }
        v1 = append(v1, (fi13 % (((-(fi13))) | 1)))
        _ = fi13
    }
    v2 = v2
    fmt.Println(v4, v2)
    switch "" {
    case "a":
        switch (("a" + "\n\t") + "a") {
        case "":
            v4 -= v4
        case "\n\t":
            for fi15 := 0; fi15 < 13; fi15++ {
                _ = len(v1)
                _ = fi15
            }
        default:
            var v16 uint8 = 42
            _ = v16
        }
    case "":
        for fi17 := 0; fi17 < 20; fi17++ {
            v3[406] = fi17
            var v18 string = ("\n\t" + "a")
            v4 = v4
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v4, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
