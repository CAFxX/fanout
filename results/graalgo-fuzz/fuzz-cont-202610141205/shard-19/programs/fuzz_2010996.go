package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32) uint16 {
    if false {
        p0 -= 0
        p0 = p0
    }
    p0 = p0
    p0 = (^(p0))
    var v7 float64 = float64(float32(0.1))
    if (false || false) {
        v8 := []int{128, 23, 191}
        _ = len(v8)
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(v7, fi9, p0)
            var v10 uint64 = 4294967295
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 int8 = 81
                v13 := GenS1{F0: 1}
                v14 := GenS2{F0: 179}
                switch fi9 {
                case 955:
                    if true {
                        fmt.Println(v7)
                        fmt.Println(v7)
                        var v15 int64 = 220
                        _ = v15
                    } else {
                        var v16 bool = (((127 != 128) && true) || (!(v10 > v10)))
                        fmt.Println(fi9, fi11, v7)
                        var v17 float32 = 3.14159
                        _ = v16
                        _ = v17
                    }
                case 0:
                    _ = len(v8)
                default:
                    v12 -= 1
                }
                var v18 int = (-(fi9))
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
                _ = v18
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 uint64, p1 string) uint {
    for fi19 := 0; fi19 < 9; fi19++ {
        if false {
            p0 -= (p0 | p0)
            var v20 uint64 = 1
            _ = v20
        } else {
            var v21 int = (fi19 + (fi19 * fi19))
            p0 -= (p0 >> ((18446744073709551615) & 63))
            v22 := GenS1{F0: 1}
            _ = v21
            _ = v22
        }
        p0 = (p0 << ((255) & 63))
        var v23 int64 = 42
        _ = fi19
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        if (!(-1 < 42)) {
            p0 += p0
            var v25 int8 = -112
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 string = p1
                _ = fi26
                _ = v27
            }
            _ = v25
        } else {
            if false {
                var v28 uint16 = 51
                var v29 uint16 = 830
                var v30 uint64 = (p0 % ((4294967295) | 1))
                _ = v28
                _ = v29
                _ = v30
            }
        }
        var v31 float32 = -2.25
        var v32 uint32 = 1
        _ = fi24
        _ = v31
        _ = v32
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 string = ""
    v2 := []string{}
    var v3 int = -9223372036854775808
    v4 := []int{32768, v3, (-(v3)), (^(v3))}
    var v5 int32 = 42
    var v6 int = (-(v3))
    v5 -= 42
    if len(v4) > 0 {
        v4[0] = (v6 + 643)
    }
    v3 += (v3 & -1)
    if false {
        v4 = append(v4, (v6 * v3))
    }
    switch (v3 & v3) {
    case 2147483647:
        if len(v2) > 0 {
            v2[0] = "a"
        }
    case -9223372036854775808:
        v3 = v6
    }
    v6 += v3
    if len(v4) > 0 {
        v4[0] = 1
    }
    fmt.Println(v3)
    fmt.Println(v5, v3, v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
