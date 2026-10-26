package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0(p0 uint64) uint16 {
    if (-1 != 0) {
        p0 = p0
        v7 := []int{-9223372036854775808}
        _ = v7
    }
    v8 := GenS1{F0: -801, F1: 18446744073709551615}
    v9 := []int{128}
    if false {
        fmt.Println(p0)
        fmt.Println(p0)
    }
    _ = p0
    _ = v8
    _ = v9
    return uint16(p0)
}
func genFn1(p0 bool, p1 uint16, p2 uint8) int {
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            fmt.Println(fi11, p1, p2)
            if (!false) {
                fmt.Println(fi11, fi10)
                p1 = 997
                if (!p0) {
                    switch fi11 {
                    case 971:
                        v12 := GenS1{F0: 32767, F1: 42}
                        _ = v12
                    case 42:
                        p1 = p1
                    }
                    var v13 int32 = -965
                    if (v13 != 0) {
                        var v14 int = fi11
                        var v15 float32 = ((((1e10 * float32((-2.25 / 1e10))) / 1e10) / 3.14159) * 3.14159)
                        _ = v14
                        _ = v15
                    } else {
                        v16 := []int{(fi10 >> ((4294967295) & 63)), fi11}
                        var v17 int8 = -128
                        var v18 float32 = (((-2.25 / (float32(1.5) - (float32((1e10 + ((0.1 * 1e10) + 1e10))) / (1.5 - 1e10)))) * (1.5 / 1e10)) - 3.14159)
                        if p0 {
                            fmt.Println(fi10, v18)
                        } else {
                            var v19 bool = p0
                            switch fi11 {
                            case 0:
                                if ((v13 ^ v13) != (^(v13))) {
                                    var v20 float64 = 1.5
                                    _ = v20
                                } else {
                                    var v21 uint32 = 402
                                    var v22 int32 = v13
                                    _ = v21
                                    _ = v22
                                }
                            default:
                                fmt.Println(p2, v17, v18)
                            }
                            _ = v19
                        }
                        v23 := GenS1{F0: 1000, F1: 820}
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v23
                    }
                    _ = v13
                } else {
                    v24 := []int{(fi10 % (((fi11 * (fi11 | fi10))) | 1)), -527, 0}
                    _ = v24
                }
                p1 += genFn0(0)
                var v25 float64 = 3.14159
                v26 := []int{(fi10 - -26)}
                v27 := GenS1{F0: 0, F1: 2}
                _ = v25
                _ = v26
                _ = v27
            } else {
                p1 -= (p1 + genFn0(493))
            }
            _ = fi11
        }
        _ = fi10
    }
    fmt.Println(p2)
    var v28 uint8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 141
    v3 := make(map[int]uint32)
    var v4 float64 = 1e10
    var v5 uint16 = 481
    var v6 int16 = 0
    v4 = float64(((-2.25 / float32(v4)) + 0.0))
    v4 = (v4 * v4)
    fmt.Println(v4, v5, v2)
    fmt.Println(v6, v5, v5)
    var v29 int32 = -243
    v29 = -1
    var v30 string = ("foo bar" + "\n\t")
    v31 := GenS1{F0: -925, F1: 729}
    v32 := GenS1{F0: v6, F1: 460}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
