package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int32 }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0(p0 int8, p1 int, p2 bool) uint32 {
    v4 := []int{p1, (-(p1))}
    var v5 int32 = 220
    v6 := GenS1{F0: "a", F1: 42, F2: v5}
    if (false && p2) {
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    if (!false) {
        var v7 int16 = 0
        _ = v7
    } else {
        v8 := []int{-977, 379, 0, -2147483648}
        var v9 uint8 = 125
        _ = v8
        _ = v9
    }
    var v10 float64 = 1e10
    v11 := GenS2{F0: -701, F1: 214}
    var v12 uint16 = 2
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint8 = 3
        switch fi13 {
        case 307:
            _ = len(v1)
        }
        v1[42] = (fi13 % ((fi13) | 1))
        _ = fi13
        _ = v14
    }
    var v15 int8 = 58
    v15 += v15
    var v16 uint16 = v12
    var v17 bool = ((true || (!(1 == 218))) && false)
    var v18 uint = 65535
    for fi19 := 0; fi19 < 13; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(v16, v12, v17)
            fmt.Println(v12)
            if (65535 < genFn0((v15 / (((v15 * (v15 + v15))) | 1)), fi20, v17)) {
                if (!v17) {
                    var v21 int64 = 42
                    if ((0.1 / (((0.0 - ((v10 + v10) / (-2.25 + 1e10))) + 1e10) * (v10 * v10))) != float64((3.14159 / (1e10 + (0.1 / ((1.5 - (1.5 / 1e10)) + -2.25)))))) {
                        v2[2] = v17
                        var v22 uint64 = 2
                        v23 := GenS1{F0: "", F1: (fi19 - -2147483648), F2: int32(fi19)}
                        _ = v22
                        _ = v23
                    } else {
                        v16 = (v16 ^ 65535)
                        v17 = v17
                        if v17 {
                            v12 -= 575
                        }
                        if v17 {
                            v24 := GenS2{F0: -109, F1: uint8(v16)}
                            _ = v24
                        }
                        v16 = (v16 | (^(v12)))
                    }
                    v25 := []int{(fi20 >> ((v18) & 63))}
                    _ = v21
                    _ = v25
                } else {
                    fmt.Println(v15)
                    v12 -= v12
                }
                v18 += v18
                var v26 uint8 = 92
                fmt.Println(v26)
                _ = v26
            }
            v27 := []int{fi20, -2147483648, 0}
            if true {
                var v28 uint = (v18 * (v18 & 572))
                v1[uint32(v12)] = 9223372036854775807
                _ = v28
            } else {
                v10 = v10
                v16 += v16
                v17 = (-1 != 42)
            }
            _ = fi20
            _ = v27
        }
        _ = fi19
    }
    v29 := GenS1{F0: ("a" + "hello"), F1: -1, F2: 970}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
}
