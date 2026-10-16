package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
func genFn0(p0 int, p1 uint8) uint {
    v6 := GenS1{F0: 2147483647, F1: 32767}
    p1 = (p1 >> ((255) & 63))
    v7 := GenS1{F0: 42, F1: -327}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 bool) int32 {
    if (!false) {
        p0 = p0
        p0 = (42 < 42)
        var v8 bool = false
        _ = v8
    }
    p0 = (65535 > genFn0(1, 1))
    _ = p0
    return 42
}

func main() {
    var v1 uint = 1
    var v2 uint8 = 1
    v3 := []int{999, 32768}
    v4 := []int{}
    v5 := []string{}
    switch 128 {
    case 32767:
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 bool = (-2.25 > ((0.0 / 1e10) - (float32(float64(0.0)) / 3.14159)))
            v10 = v10
            _ = fi9
            _ = v10
        }
    }
    if (false || false) {
        if (!false) {
            for fi11 := 0; fi11 < 14; fi11++ {
                v1 = (^(v1))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi11
            }
            v1 -= (v1 & v1)
            v1 = (v1 & (v1 >> ((2) & 63)))
        }
        v12 := GenS1{F0: 218, F1: 0}
        var v13 uint8 = v2
        _ = v12
        _ = v13
    }
    var v14 int = 2147483647
    v15 := GenS1{F0: -1, F1: 0}
    var v16 int8 = 82
    var v17 uint64 = 2
    if (42 == 78) {
        var v18 float64 = ((3.14159 * ((float64(3.14159) + 0.0) - 0.1)) / (3.14159 / 1.5))
        var v19 uint16 = 1
        var v20 int8 = int8(v19)
        _ = v18
        _ = v19
        _ = v20
    }
    _ = len(v3)
    v17 = v17
    v14 -= v14
    v2 = (v2 & 128)
    if len(v5) > 0 {
        v5[0] = "a"
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (v14 >> ((4294967295) & 63)) {
    case -770:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 114:
        v3 = append(v3, (v14 * v14))
    case 1:
        v1 = (v1 << ((v1) & 63))
    }
    var v21 uint64 = (v17 % ((v17) | 1))
    for fi22 := 0; fi22 < 13; fi22++ {
        switch "\n\t" {
        case "":
            v2 = (v2 | 155)
        default:
            v14 -= (fi22 << ((v1) & 63))
        }
        var v23 bool = true
        _ = fi22
        _ = v23
    }
    v14 = (v14 << ((uint(v21)) & 63))
    v4 = append(v4, (v14 << ((898) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
