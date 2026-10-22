package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint32) uint16 {
    v4 := []int{4294967295}
    if len(v4) > 0 {
        v4[0] = -1
    }
    _ = p0
    _ = p1
    _ = v4
    return 61
}
func genFn1(p0 int8, p1 uint32) int {
    p0 += p0
    switch "a" {
    case "\n\t":
        fmt.Println(p0, p1)
    default:
        v5 := GenS1{F0: 1000, F1: 0}
        _ = v5
    }
    fmt.Println(p1, p0)
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 bool = true
        fmt.Println(p0)
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint8 = 1
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint16 = 876
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := []int{fi12, fi8, (fi12 << ((2) & 63)), fi12}
            if (0.1 == ((1e10 / 0.1) / 1.0)) {
                fmt.Println(v9, fi12, v9)
                if true {
                    if (p1 >= 2147483648) {
                        v13 = append(v13, (fi8 / ((fi8) | 1)))
                        if (false || (false || false)) {
                            fmt.Println(p1, p1)
                        }
                        var v14 int = int(v9)
                        _ = v14
                    } else {
                        v15 := GenS1{F0: 42, F1: 701}
                        fmt.Println(p0, fi12, fi8)
                        p1 = (p1 % ((p1) | 1))
                        _ = v15
                    }
                    if true {
                        p1 -= uint32(v9)
                    } else {
                        p0 += (p0 + (p0 >> ((719) & 63)))
                    }
                    v16 := GenS1{F0: -134, F1: 18446744073709551615}
                    _ = v16
                }
            }
            _ = fi12
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 float64 = ((1.5 - (float64((-2.25 / 1.0)) + (1.5 / 1.5))) / float64(1.5))
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    v17 := v3[606]
    var v18 int8 = 55
    v19 := GenS1{F0: -40, F1: 2}
    v1 -= v1
    v20 := []int{-9223372036854775808, genFn1((v18 >> ((246) & 63)), 1)}
    var v21 float32 = -2.25
    var v22 int32 = 0
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS2{F0: 769}
        _ = fi23
        _ = v24
    }
    fmt.Println(v21, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
