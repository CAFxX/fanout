package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint) int64 {
    p0 = 42
    var v4 int = -2147483648
    v4 += -86
    var v5 uint16 = 42
    _ = p0
    _ = v4
    _ = v5
    return -709
}
func genFn1(p0 uint32, p1 uint32, p2 uint32) bool {
    v6 := []int{-672, 9223372036854775807, 1}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int = fi7
        v8 += (fi7 * (fi7 - (v8 & -312)))
        switch v8 {
        case 833:
            fmt.Println(p2, p0)
        case -2147483648:
            v6 = append(v6, (-(fi7)))
        }
        _ = fi7
        _ = v8
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}

func main() {
    v1 := []int{32767}
    v2 := []string{"\n\t", "\n\t", "\n\t"}
    var v3 uint32 = 2147483648
    _ = len(v1)
    if (32767 < -32768) {
        if ((false || ((!true) || true)) || false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v3)
            fmt.Println(v3)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v3)
        }
        fmt.Println(v3)
        v3 = v3
    } else {
        var v9 uint16 = 142
        v10 := GenS1{F0: -9223372036854775808, F1: genFn0(65535)}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v3 -= v3
        v3 -= v3
        _ = fi11
    }
    var v12 bool = (0.0 == (((float32((1.5 + (1.5 + (0.1 / float64(float32(0.1)))))) / (-2.25 * -2.25)) + 0.0) + (0.0 + 1e10)))
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v15 string = "\n\t"
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v16 := GenS2{F0: 1}
    if v12 {
        fmt.Println(v12, v12)
        v17 := GenS1{F0: genFn0(65535), F1: genFn0(179)}
        _ = v17
    } else {
        v18 := GenS2{F0: 2}
        if ((float64((float32(3.14159) * ((1e10 - 3.14159) / 1.0))) - (-2.25 / 1e10)) == 1.5) {
            var v19 uint32 = (^(v3))
            _ = v19
        } else {
            v3 = (v3 / (((v3 + (v3 << ((2) & 63)))) | 1))
            v3 -= (v3 & 2)
        }
        _ = v18
    }
    v12 = (!v12)
    v12 = (0.1 < 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
}
