package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint, p1 bool) int16 {
    for fi4 := 0; fi4 < 7; fi4++ {
        p0 = 42
        v5 := []int{fi4, (fi4 * fi4)}
        p0 = (p0 / ((18446744073709551615) | 1))
        _ = fi4
        _ = v5
    }
    var v6 int64 = 1
    if p1 {
        for fi7 := 0; fi7 < 3; fi7++ {
            fmt.Println(fi7, fi7, fi7)
            for fi8 := 0; fi8 < 15; fi8++ {
                fmt.Println(fi7, p0, v6)
                var v9 int = fi7
                p0 -= 963
                fmt.Println(p1)
                v6 = -9223372036854775808
                fmt.Println(fi7, v6, p1)
                var v10 uint8 = 231
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := []int{-511, 32767, (fi11 / ((fi11) | 1)), fi11}
            fmt.Println(v6, p0)
            _ = fi11
            _ = v12
        }
        p1 = (1e10 != 0.0)
    }
    _ = p0
    _ = p1
    _ = v6
    return -32768
}

func main() {
    var v1 int = 128
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v13 := GenS1{F0: 1}
    if false {
        v1 += (v1 << ((65535) & 63))
        var v14 int64 = 42
        v15 := GenS1{F0: -209}
        _ = v14
        _ = v15
    } else {
        var v16 string = ("" + "")
        v17 := []int{(v1 - 0), (v1 & (v1 + (v1 | v1))), v1, 4294967295}
        var v18 uint = 4294967295
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 42, F1: 2}
    if false {
        v20 := GenS2{F0: 684, F1: 223}
        var v21 int32 = 2147483647
        _ = v20
        _ = v21
    } else {
        v1 = (^(v1))
        fmt.Println(v1)
        v1 = 42
    }
    switch "foo bar" {
    case "hello":
        for fi22 := 0; fi22 < 12; fi22++ {
            if (false && false) {
                v23 := []int{(v1 >> ((18446744073709551615) & 63)), -2147483648, (fi22 | (fi22 | fi22)), (fi22 << ((18446744073709551615) & 63))}
                switch 128 {
                case 42:
                    var v24 float64 = -2.25
                    _ = v24
                default:
                    var v25 int16 = 1000
                    _ = v25
                }
                var v26 uint = 18446744073709551615
                _ = v23
                _ = v26
            }
            v27 := GenS1{F0: 28}
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 bool = (-1 > -549)
                _ = fi28
                _ = v29
            }
            _ = fi22
            _ = v27
        }
    case "a":
        v1 = v1
    case "":
        for fi30 := 0; fi30 < 4; fi30++ {
            v31 := GenS2{F0: 335, F1: 119}
            _ = fi30
            _ = v31
        }
    }
    var v32 uint32 = 0
    v33 := []int{v1, v1}
    v1 += (v1 >> ((255) & 63))
    v1 += v1
    var v34 uint32 = (v32 | v32)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
