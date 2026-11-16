package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int8) uint8 {
    v4 := GenS1{F0: -32768, F1: 1}
    p0 += p0
    p0 -= (^(p0))
    _ = p0
    _ = v4
    return 42
}
func genFn1(p0 int64, p1 int32, p2 int) int8 {
    p0 += 589
    v5 := []int{127, p2, (-(p2)), 2147483647}
    var v6 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 42
    var v3 int32 = -2147483648
    v3 = (v3 + -746)
    v7 := []int{9223372036854775807, 0, 127, -892}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 float32 = (3.14159 - 1.5)
        _ = fi8
        _ = v9
    }
    v3 = -836
    var v10 int32 = 2147483647
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 int16 = 0
            v1[-1] = (fi11 / ((fi12) | 1))
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    switch 32768 {
    case 276:
        for fi14 := 0; fi14 < 14; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                if true {
                    v3 = (v10 / ((1) | 1))
                } else {
                    var v16 uint = 878
                    var v17 int64 = 1
                    _ = v16
                    _ = v17
                }
                _ = fi15
            }
            _ = fi14
        }
    }
    v3 -= (v10 - 2147483647)
    v18 := GenS2{F0: 91}
    if (true || true) {
        fmt.Println(v2)
        var v19 uint = (^(v2))
        _ = v19
    }
    if len(v7) > 0 {
        v7[0] = -554
    }
    v3 -= (v10 & -913)
    var v20 int8 = 127
    if (-452 < 1000) {
        var v21 string = "foo bar"
        if true {
            v20 -= genFn1(28, (v10 & (v3 % (((v3 + (v10 << (((v2 & (v2 ^ 2))) & 63)))) | 1))), 327)
            var v22 int32 = v10
            _ = v22
        }
        _ = v21
    } else {
        var v23 float64 = 3.14159
        _ = len(v7)
        switch 42 {
        case 4294967295:
            v10 = v3
        default:
            v7 = append(v7, 2147483647)
        }
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
