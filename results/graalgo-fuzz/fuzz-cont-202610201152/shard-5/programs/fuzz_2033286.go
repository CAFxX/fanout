package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int64; F1 string; F2 uint8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint64; F1 string; F2 uint }
type GenS5 struct { F0 int32 }
func genFn0() int8 {
    var v7 int16 = 42
    fmt.Println(v7)
    v8 := GenS1{F0: ((v7 % (((v7 | v7)) | 1)) != 62), F1: 1}
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 string, p1 int8, p2 string) uint {
    var v9 float64 = (0.0 / 1e10)
    v10 := GenS1{F0: true, F1: 18446744073709551615}
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := []int{(-(fi11)), 9223372036854775807}
        p2 = (p2 + "a")
        _ = fi11
        _ = v12
    }
    var v13 int8 = (p1 | (p1 ^ p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    _ = v9
    return 255
}

func main() {
    var v1 string = ""
    var v2 float64 = 0.1
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    var v5 int32 = 2147483647
    v6 := []uint32{}
    v2 += ((1.5 * (0.1 + v2)) * ((0.1 / v2) - 0.1))
    fmt.Println(v1, v1)
    v14 := GenS2{F0: 0, F1: v1, F2: 255}
    var v15 int64 = -1
    switch (v1 + v1) {
    case "hello":
        switch 2147483647 {
        case 32768:
            v5 = -2147483648
        case -1:
            var v16 uint8 = 0
            _ = v16
        case 42:
            _ = len(v3)
        }
    case "foo bar":
        var v17 int64 = (v15 + v15)
        _ = v17
    default:
        var v18 int8 = 127
        _ = v18
    }
    v6 = append(v6, 4294967295)
    var v19 int = 9223372036854775807
    if false {
        var v20 uint8 = 255
        _ = v20
    }
    v1 = v1
    _ = len(v6)
    v21 := GenS1{F0: true, F1: 854}
    _ = len(v4)
    fmt.Println(v19, v5)
    for fi22 := 0; fi22 < 9; fi22++ {
        if true {
            v23 := GenS2{F0: (-(v15)), F1: v1, F2: 129}
            v24 := GenS3{F0: (v15 * 1), F1: genFn0()}
            v25 := GenS4{F0: 4294967295, F1: v1, F2: 85}
            _ = v23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    if (false && (307 > 1)) {
        _ = len(v6)
    } else {
        v1 = "\n\t"
        v5 = 1
        v15 += v15
    }
    v26 := GenS5{F0: (v5 << ((0) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
