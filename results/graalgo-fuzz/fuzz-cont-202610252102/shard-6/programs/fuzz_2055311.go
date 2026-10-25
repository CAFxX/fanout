package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint; F1 int64; F2 int16 }
type GenS3 struct { F0 uint; F1 int8; F2 uint8 }
func genFn0(p0 int32) uint {
    p0 += 528
    v5 := GenS1{F0: 162, F1: 571}
    v6 := GenS1{F0: 0, F1: 2}
    _ = p0
    _ = v5
    _ = v6
    return 255
}
func genFn1() uint64 {
    var v7 int64 = -926
    v7 -= 42
    v8 := []int{}
    _ = v7
    _ = v8
    return 760
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 42
    var v3 int = 32767
    var v4 uint8 = 128
    if (-1 > 823) {
        var v9 uint = 18446744073709551615
        var v10 int64 = (-(v2))
        if false {
            v11 := []int{}
            switch "a" {
            case "":
                v12 := GenS2{F0: genFn0(-2147483648), F1: (v10 | v2), F2: 42}
                _ = v12
            default:
                var v13 int8 = 53
                _ = v13
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := GenS3{F0: 843, F1: 0, F2: 115}
                v4 += v4
                _ = fi14
                _ = v15
            }
            _ = v11
        }
        _ = v10
        _ = v9
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint = 42
        _ = fi16
        _ = v17
    }
    v3 -= (v3 >> ((4294967295) & 63))
    var v18 float64 = (1e10 / ((((3.14159 - (0.0 * (0.0 + -2.25))) * 1.5) - (1e10 * 1e10)) / 0.1))
    var v19 float64 = v18
    v20 := []int{748, v3, -295}
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 float64 = v18
        if false {
            fmt.Println(v19, v3)
            v2 += v2
        }
        var v23 int16 = 630
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v19)
    switch (v3 * (v3 - (v3 * v3))) {
    case -2147483648:
        v3 = v3
    case 128:
        switch v3 {
        case 42:
            var v24 int16 = -922
            _ = v24
        case 790:
            fmt.Println(v3, v19)
        case -453:
            v25 := GenS1{F0: v4, F1: 654}
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 uint64 = 0
        var v28 int64 = (^(v2))
        _ = fi26
        _ = v27
        _ = v28
    }
    if len(v20) > 0 {
        v20[0] = (v3 & v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
