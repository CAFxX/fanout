package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 bool }
func genFn0(p0 int, p1 uint16) uint64 {
    if (!(0 < 4294967295)) {
        for fi7 := 0; fi7 < 9; fi7++ {
            p0 = (p0 | p0)
            if ((^(p1)) >= 42) {
                for fi8 := 0; fi8 < 13; fi8++ {
                    fmt.Println(fi8, fi7)
                    _ = fi8
                }
            }
            p1 += p1
            _ = fi7
        }
    } else {
        v9 := []int{}
        _ = v9
    }
    v10 := GenS1{F0: 18446744073709551615, F1: p0, F2: (42 > 100)}
    v11 := []int{(-(p0)), p0}
    p0 += (p0 | 4294967295)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 124
}

func main() {
    var v1 int16 = 0
    var v2 float64 = 1e10
    v3 := []uint32{2147483648, 4294967295, 105, 2147483648}
    v4 := make(map[uint32]uint32)
    var v5 int8 = 100
    var v6 int64 = 1
    var v12 float64 = (0.0 * ((float64(0.0) / v2) / v2))
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(fi13)
            var v15 string = "\n\t"
            _ = fi14
            _ = v15
        }
        var v16 float64 = (v2 * 3.14159)
        for fi17 := 0; fi17 < 3; fi17++ {
            fmt.Println(v12, fi13, v6)
            if (true || true) {
                var v18 int64 = v6
                _ = v18
            } else {
                var v19 uint32 = 2147483648
                fmt.Println(fi17, v12)
                v1 += 42
                _ = v19
            }
            var v20 int = (fi13 / ((-2147483648) | 1))
            _ = fi17
            _ = v20
        }
        _ = fi13
        _ = v16
    }
    v5 -= (v5 << ((0) & 63))
    v5 -= (v5 * (^(v5)))
    fmt.Println(v12, v2)
    v12 += v12
    var v21 int32 = 42
    fmt.Println(v5, v5)
    v1 -= (^(v1))
    v5 += (v5 & v5)
    fmt.Println(v21)
    if len(v3) > 0 {
        v3[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
