package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint }
type GenS3 struct { F0 int8; F1 uint8; F2 int }
type GenS4 struct { F0 int8; F1 uint; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string }
func genFn0(p0 uint32, p1 uint, p2 bool) int16 {
    v7 := []int{599}
    v8 := GenS1{F0: 0, F1: 128}
    switch -601 {
    case 1:
        for fi9 := 0; fi9 < 4; fi9++ {
            p2 = true
            p1 = p1
            _ = fi9
        }
    case 127:
        var v10 bool = ((-788 <= 1000) || (int64(p0) != 0))
        _ = v10
    case 712:
        _ = len(v7)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -57
}
func genFn1(p0 uint64, p1 uint, p2 uint8) string {
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS2{F0: 0, F1: genFn0(343, (p1 << ((255) & 63)), ((p1 * p1) >= (p1 >> ((p1) & 63)))), F2: p1}
        p1 = 42
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        if (true || false) {
            p0 -= p0
            v14 := GenS1{F0: 0, F1: p2}
            v15 := GenS2{F0: 42, F1: genFn0(4294967295, p1, true), F2: (p1 / (((p1 << ((0) & 63))) | 1))}
            _ = v14
            _ = v15
        }
        v16 := GenS3{F0: 119, F1: p2, F2: fi13}
        var v17 int64 = 42
        _ = fi13
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 int8 = 127
    var v5 int = 32768
    var v6 uint8 = 255
    v4 = v4
    v6 += v6
    v6 = v6
    var v18 bool = true
    v1 -= (((-2.25 * (v1 / v1)) * (v1 + -2.25)) * float32(float64(-2.25)))
    v19 := []int{(v5 - -9223372036854775808), -9223372036854775808}
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS4{F0: v4, F1: 0, F2: -1}
        switch 32768 {
        case -9223372036854775808:
            _ = len(v2)
        }
        v2[-2147483648] = 425
        _ = fi20
        _ = v21
    }
    if v18 {
        fmt.Println(v1, v5, v5)
    }
    fmt.Println(v1, v18, v4)
    v22 := GenS2{F0: 42, F1: 0, F2: 0}
    if false {
        if v18 {
            v5 += (^(v5))
        }
        var v23 int = v5
        _ = v23
    }
    switch "\n\t" {
    case "hello":
        var v24 int8 = v4
        _ = v24
    case "\n\t":
        if v18 {
            v25 := GenS5{F0: 0}
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            var v26 int16 = -32768
            _ = v25
            _ = v26
        }
    default:
        if len(v19) > 0 {
            v19[0] = (-(v5))
        }
    }
    fmt.Println(v18, v6, v4)
    v27 := GenS6{F0: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
