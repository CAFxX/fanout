package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint) uint16 {
    v7 := GenS1{F0: -32768, F1: -889, F2: 18446744073709551615}
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := []int{(fi9 * (fi9 & 127))}
            if true {
                v11 := GenS2{F0: 32767, F1: (fi9 % ((fi8) | 1)), F2: -1}
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS3{F0: 100}
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: -1, F1: 210, F2: (^(p0))}
    _ = p0
    _ = v14
    _ = v7
    return 42
}
func genFn1(p0 int8, p1 string, p2 int64) string {
    p1 = p1
    for fi15 := 0; fi15 < 20; fi15++ {
        switch (fi15 | fi15) {
        case -9223372036854775808:
            fmt.Println(p0, fi15, p2)
        case 2147483647:
            var v16 int = fi15
            _ = v16
        case -646:
            p0 = p0
        default:
            v17 := []int{fi15, 32768}
            _ = v17
        }
        v18 := []int{32767, fi15, 1}
        v19 := GenS2{F0: 576, F1: 9223372036854775807, F2: -32768}
        _ = fi15
        _ = v18
        _ = v19
    }
    var v20 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return p1
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]uint32)
    var v3 uint64 = 4294967295
    var v4 int32 = -418
    var v5 float64 = 3.14159
    v6 := make(map[string]bool)
    if (42 != 42) {
        v21 := []int{-580}
        for fi22 := 0; fi22 < 9; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                v4 += (^(v4))
                var v24 int8 = 100
                _ = fi23
                _ = v24
            }
            var v25 uint16 = uint16(v3)
            switch "hello" {
            case "foo bar":
                fmt.Println(fi22, v3)
            case "\n\t":
                var v26 bool = false
                _ = v26
            }
            _ = fi22
            _ = v25
        }
        _ = len(v21)
        _ = v21
    }
    v27 := []int{}
    fmt.Println(v1)
    v27 = append(v27, -130)
    var v28 int32 = v4
    var v29 uint32 = 498
    switch 127 {
    case 9223372036854775807:
        v28 = 594
    case -1:
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 uint16 = 653
            v3 = 0
            _ = fi30
            _ = v31
        }
    case -595:
        v5 = 1e10
    default:
        var v32 uint64 = v3
        _ = v32
    }
    var v33 int32 = (-(v28))
    if (!(v3 <= 1)) {
        if ((false || (2 >= 1)) || false) {
            fmt.Println(v1)
            v29 += 2
        }
    }
    switch int(v1) {
    case -1:
        var v34 int64 = -853
        _ = v34
    case 9223372036854775807:
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 int64 = -9223372036854775808
            fmt.Println(v29, v33, v1)
            v3 += (^(v3))
            _ = fi35
            _ = v36
        }
    case 42:
        v37 := GenS2{F0: 1, F1: 127, F2: v1}
        _ = v37
    }
    if (!(219 != 128)) {
        var v38 uint = 4294967295
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
