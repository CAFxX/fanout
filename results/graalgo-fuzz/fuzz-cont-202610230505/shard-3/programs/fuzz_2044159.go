package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint64; F1 uint64; F2 int64 }
func genFn0(p0 uint64) int32 {
    if false {
        v7 := GenS1{F0: 65535, F1: 255}
        switch 32768 {
        case 2147483647:
            if true {
                fmt.Println(p0)
                v8 := []int{}
                _ = v8
            }
        case 127:
            p0 += (p0 ^ 42)
        }
        p0 = 42
        _ = v7
    }
    if (1 != int64(p0)) {
        fmt.Println(p0)
    }
    p0 -= (p0 * (p0 & p0))
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 string = "\n\t"
            var v12 float64 = -2.25
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    _ = p0
    return 799
}
func genFn1(p0 uint, p1 uint32) uint16 {
    v13 := GenS1{F0: (^(p1)), F1: 691}
    v14 := []int{4294967295}
    var v15 uint64 = 1
    p0 = 384
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 549
}

func main() {
    var v1 int32 = 864
    v2 := make(map[string]int)
    var v3 uint32 = 2147483648
    var v4 int16 = -8
    v5 := []uint32{864, 0}
    var v6 string = "foo bar"
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int8 = 1
        _ = fi16
        _ = v17
    }
    fmt.Println(v4, v1, v1)
    var v18 int16 = (v4 >> ((4294967295) & 63))
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 int32 = (v1 & int32(v18))
        _ = fi19
        _ = v20
    }
    _ = len(v5)
    var v21 string = v6
    v3 += v3
    var v22 uint32 = v3
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint32 = (v22 + v3)
        _ = fi23
        _ = v24
    }
    var v25 float64 = (0.0 * float64(((0.1 - 0.1) * 0.0)))
    v26 := GenS2{F0: 4294967295, F1: uint64(v22), F2: 9223372036854775807}
    v3 = 2147483647
    v3 -= v3
    fmt.Println(v25)
    for fi27 := 0; fi27 < 7; fi27++ {
        if false {
            v28 := GenS1{F0: v22, F1: 739}
            _ = v28
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
