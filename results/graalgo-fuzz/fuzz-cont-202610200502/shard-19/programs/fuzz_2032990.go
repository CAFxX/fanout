package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 int }
func genFn0(p0 uint32, p1 string, p2 int32) int64 {
    v6 := []int{969, -930}
    for fi7 := 0; fi7 < 10; fi7++ {
        v6 = append(v6, 2147483647)
        switch p1 {
        case "foo bar":
            p0 = (p0 | (p0 - 4294967295))
        default:
            _ = len(v6)
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 925
}
func genFn1(p0 int16, p1 int8, p2 int64) uint32 {
    p1 += p1
    fmt.Println(p0, p0)
    for fi8 := 0; fi8 < 6; fi8++ {
        p2 = (p2 * 9223372036854775807)
        _ = fi8
    }
    var v9 float64 = 3.14159
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 62
}

func main() {
    var v1 int64 = -432
    v2 := []uint32{4294967295, 2}
    var v3 uint64 = 4294967295
    var v4 int32 = 0
    v5 := make(map[int]string)
    var v10 int32 = (-(v4))
    v4 -= (v10 & (-(v10)))
    var v11 float64 = (3.14159 / -2.25)
    _ = len(v2)
    var v12 uint32 = 0
    var v13 uint64 = 826
    if ((v11 + v11) != 1.5) {
        v13 -= (v13 % ((v3) | 1))
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            if len(v2) > 0 {
                v2[0] = 1
            }
            v4 += (v10 % ((v10) | 1))
            v11 += (v11 + 0.0)
            _ = fi15
        }
        var v16 uint = 0
        v3 -= 349
        _ = fi14
        _ = v16
    }
    v17 := GenS1{F0: 42}
    v2 = append(v2, genFn1(32767, 100, v1))
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS1{F0: -32768}
        v3 += (v3 - 0)
        _ = fi18
        _ = v19
    }
    fmt.Println(v10)
    if (false && true) {
        v20 := GenS1{F0: 212}
        v21 := GenS2{F0: 128, F1: 32768, F2: 2147483647}
        v10 += 2147483647
        _ = v20
        _ = v21
    } else {
        v11 = 3.14159
    }
    if (false || true) {
        if false {
            var v22 int = 1
            for fi23 := 0; fi23 < 4; fi23++ {
                if false {
                    var v24 uint64 = (v13 * uint64(v12))
                    if false {
                        v22 = (v22 | (v22 ^ (v22 ^ (fi23 - (fi23 ^ v22)))))
                    }
                    var v25 int16 = 1
                    _ = v24
                    _ = v25
                }
                _ = fi23
            }
            fmt.Println(v11, v13)
            _ = v22
        }
        v3 += v13
    }
    fmt.Println(v1, v11)
    v3 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
