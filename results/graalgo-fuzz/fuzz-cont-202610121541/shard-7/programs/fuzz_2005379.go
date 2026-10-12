package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0(p0 int64, p1 uint64) uint8 {
    p0 = (p0 | p0)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint = 4294967295
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return uint8(p1)
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 18446744073709551615
    v3 := []uint32{2}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    v6 := []uint32{2}
    v2 = 4294967295
    if (float64(float32(1e10)) == (0.1 / 3.14159)) {
        if ((1 < 496) || (-128 != 100)) {
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 float64 = 1.5
                _ = fi9
                _ = v10
            }
            if (((2 == 42) && false) || false) {
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 int64 = 1
                    _ = fi11
                    _ = v12
                }
            }
        } else {
            v2 = v2
        }
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        if (1 <= -32768) {
            var v14 uint64 = 0
            if (!(!(42 > 101))) {
                v15 := GenS1{F0: 255, F1: 0, F2: 1}
                v16 := GenS1{F0: 255, F1: -1, F2: 1}
                v1[fi13] = (("\n\t" + "a") + "\n\t")
                _ = v15
                _ = v16
            }
            fmt.Println(v14, v14)
            _ = v14
        } else {
            if len(v3) > 0 {
                v3[0] = 375
            }
        }
        _ = fi13
    }
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    v2 -= (v2 & v2)
    if false {
        var v17 uint32 = 65535
        _ = v17
    }
    var v18 uint16 = 42
    fmt.Println(v18)
    switch 127 {
    case 32767:
        v2 -= v2
    case 0:
        var v19 uint8 = 2
        _ = v19
    default:
        var v20 bool = (((v2 % ((v2) | 1)) > v2) || true)
        _ = v20
    }
    v21 := GenS1{F0: (v18 * 485), F1: 1, F2: -2147483648}
    if (false || true) {
        v2 += (v2 * (v2 ^ (^(v2))))
        v22 := GenS1{F0: 2, F1: int32(v2), F2: -911}
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
