package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int32; F2 int16 }
type GenS5 struct { F0 int8; F1 int64 }
func genFn0(p0 int) int32 {
    if (!false) {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := []int{p0, (p0 & fi6), 477}
            var v8 int16 = 32767
            p0 += p0
            _ = fi6
            _ = v7
            _ = v8
        }
        p0 -= -29
        var v9 float64 = ((1.5 + 1.5) + (1.5 * 3.14159))
        _ = v9
    }
    if false {
        v10 := GenS1{F0: (true && (!((148 <= 128) || true)))}
        _ = v10
    }
    var v11 int16 = 42
    v12 := []int{(p0 % (((-(p0))) | 1)), 0}
    var v13 float32 = ((1.5 - 3.14159) * (0.1 / (3.14159 - 0.1)))
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -180
    var v3 int16 = 42
    var v4 int64 = 1
    var v5 bool = (false && false)
    if v5 {
        var v14 uint64 = 683
        var v15 bool = true
        fmt.Println(v5)
        _ = v14
        _ = v15
    }
    v5 = v5
    v4 = v4
    fmt.Println(v5, v2, v4)
    v5 = (v4 <= (v4 + (v4 | 287)))
    for fi16 := 0; fi16 < 18; fi16++ {
        if ((v2 >> ((255) & 63)) <= v2) {
            v17 := []int{(-(fi16)), fi16, (^(fi16))}
            for fi18 := 0; fi18 < 18; fi18++ {
                v4 -= (v4 | -485)
                _ = fi18
            }
            _ = v17
        }
        v3 = (v3 ^ (-(v3)))
        if (1e10 >= -2.25) {
            v2 -= 42
            v19 := GenS2{F0: "", F1: 874, F2: 409}
            _ = v19
        }
        _ = fi16
    }
    v5 = true
    v20 := GenS3{F0: 132, F1: 32767}
    for fi21 := 0; fi21 < 13; fi21++ {
        v1[1] = 65535
        if false {
            v3 = (v3 / ((v3) | 1))
            v5 = (286 == 42)
        }
        _ = fi21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS4{F0: 1, F1: v2, F2: (v3 << ((970) & 63))}
        var v24 int16 = v3
        v4 -= v4
        _ = fi22
        _ = v23
        _ = v24
    }
    if ((!v5) || (2147483648 != 107)) {
        for fi25 := 0; fi25 < 11; fi25++ {
            if v5 {
                v5 = (255 == 32767)
            }
            for fi26 := 0; fi26 < 15; fi26++ {
                v2 = 42
                v27 := GenS5{F0: -1, F1: v4}
                _ = fi26
                _ = v27
            }
            var v28 int32 = v2
            _ = fi25
            _ = v28
        }
    }
    fmt.Println(v3)
    var v29 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
