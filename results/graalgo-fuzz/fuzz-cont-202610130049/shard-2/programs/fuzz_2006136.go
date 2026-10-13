package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint }
func genFn0(p0 int, p1 uint16) bool {
    fmt.Println(p1)
    v4 := GenS1{F0: 6}
    p0 += (p0 * (p0 << ((4294967295) & 63)))
    var v5 float64 = 1e10
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (4294967295 <= 433)
}
func genFn1(p0 uint16, p1 int8, p2 bool) int16 {
    var v6 uint16 = (p0 << ((1) & 63))
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 float32 = float32(-2.25)
    var v2 bool = false
    var v3 float32 = (v1 + 1e10)
    var v7 uint16 = 150
    v8 := []int{-38, 2147483647, 789, 1}
    var v9 uint16 = 32767
    var v10 int = 128
    if (-1 != -2147483648) {
        v11 := GenS1{F0: 946}
        var v12 float64 = 1.5
        if (2 == 0) {
            var v13 float64 = float64(v3)
            _ = v13
        }
        _ = v11
        _ = v12
    } else {
        v8 = append(v8, v10)
        for fi14 := 0; fi14 < 13; fi14++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 uint8 = 0
                if ((-1 >= -2147483648) || true) {
                    v1 -= v3
                    var v17 bool = v2
                    fmt.Println(fi15, v16)
                    v18 := GenS2{F0: 9223372036854775807, F1: 1}
                    _ = v17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
            fmt.Println(v9, v3, v3)
            _ = fi14
        }
    }
    var v19 uint16 = 527
    var v20 uint32 = 626
    var v21 int = v10
    fmt.Println(v2)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v22 := GenS2{F0: 9223372036854775807, F1: 18446744073709551615}
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 uint8 = 127
        if v2 {
            v25 := GenS1{F0: 1}
            _ = v25
        }
        fmt.Println(v3, v3, v3)
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
