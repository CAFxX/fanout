package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 int; F1 uint64; F2 int32 }
func genFn0(p0 string, p1 int, p2 uint32) uint {
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 = ((p0 + (p0 + p0)) + "")
        _ = fi6
    }
    var v7 uint16 = 65535
    switch (p1 << ((4294967295) & 63)) {
    case 42:
        var v8 uint = 42
        _ = v8
    case 9223372036854775807:
        fmt.Println(v7, p1)
    }
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 193
}
func genFn1() uint {
    var v9 uint16 = 42
    fmt.Println(v9)
    switch 2147483647 {
    case -1:
        v10 := GenS1{F0: 699, F1: 42, F2: 255}
        _ = v10
    }
    v9 = 375
    _ = v9
    return 994
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-9223372036854775808, -36, 877}
    v3 := []int{-9223372036854775808, 0, 32767}
    var v4 int = 0
    var v5 float64 = float64(0.1)
    v4 -= v4
    v11 := GenS1{F0: 636, F1: 9223372036854775807, F2: 913}
    var v12 bool = true
    var v13 int64 = 1
    v5 = 0.0
    switch (^(v4)) {
    case 893:
        v5 += float64(0.0)
    default:
        switch ("" + "foo bar") {
        case "foo bar":
            for fi14 := 0; fi14 < 17; fi14++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v15 := GenS1{F0: 1, F1: 9223372036854775807, F2: 65535}
                _ = fi14
                _ = v15
            }
        }
    }
    v16 := GenS2{F0: v4, F1: 921, F2: 287}
    for fi17 := 0; fi17 < 16; fi17++ {
        _ = len(v3)
        if len(v2) > 0 {
            v2[0] = v4
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := GenS1{F0: 1, F1: 42, F2: 4294967295}
            fmt.Println(fi18, fi17)
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int8 = 1
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v4, fi22, fi22)
        var v23 float64 = (-2.25 * (v5 - 3.14159))
        _ = fi22
        _ = v23
    }
    var v24 float64 = (float64(0.0) + ((float64(-2.25) - (v5 + float64(1.5))) * v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
