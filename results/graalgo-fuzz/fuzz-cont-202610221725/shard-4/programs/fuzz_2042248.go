package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint16, p1 uint) string {
    p0 -= p0
    v6 := GenS1{F0: 255, F1: 1}
    for fi7 := 0; fi7 < 7; fi7++ {
        switch (fi7 << (((p1 % ((p1) | 1))) & 63)) {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 3; fi8++ {
                p1 += 2
                p0 -= p0
                fmt.Println(fi7)
                var v9 uint = p1
                v9 += (^(p1))
                _ = fi8
                _ = v9
            }
        }
        _ = fi7
    }
    var v10 int16 = -621
    v10 += (-(v10))
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return ""
}

func main() {
    var v1 float64 = (float64(1e10) + (1e10 / 1.5))
    var v2 uint8 = 1
    var v3 int32 = -871
    var v4 uint32 = 42
    v5 := []int{-9223372036854775808, 634, -2147483648, -75}
    switch -851 {
    case 128:
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    case 893:
        fmt.Println(v2)
    case 2147483647:
        v2 += v2
    }
    v11 := GenS1{F0: (v2 * 1), F1: v4}
    v2 = 1
    var v12 uint64 = 18446744073709551615
    switch -586 {
    case -2147483648:
        var v13 uint32 = 0
        _ = v13
    default:
        v5 = append(v5, 1)
    }
    v5 = append(v5, 9223372036854775807)
    var v14 uint16 = 893
    v15 := GenS1{F0: 34, F1: 583}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
