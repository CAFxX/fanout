package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int16; F2 bool }
type GenS5 struct { F0 int8; F1 int64; F2 int }
func genFn0(p0 string, p1 uint16) int32 {
    switch 4294967295 {
    case 9223372036854775807:
        p1 = (p1 ^ 295)
    case 42:
        p1 += p1
    case 987:
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := []int{fi6, fi6}
            v8 := GenS1{F0: -820}
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := GenS1{F0: 9223372036854775807}
            if true {
                fmt.Println(p1, fi9, fi9)
            }
            _ = fi9
            _ = v10
        }
    }
    if (1 <= 65535) {
        v11 := []int{-2147483648, 430}
        fmt.Println(p0, p0)
        _ = v11
    } else {
        v12 := GenS1{F0: 597}
        var v13 int32 = 42
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1() int32 {
    v14 := GenS2{F0: 42, F1: 567, F2: 419}
    var v15 uint8 = 0
    v16 := []int{-2147483648}
    var v17 string = "foo bar"
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -718
    v3 := []string{"a"}
    var v4 uint16 = 610
    var v5 int64 = 9223372036854775807
    v4 = (^(v4))
    v5 -= (v5 + -811)
    v18 := v1[984]
    v19 := GenS1{F0: (v5 - (^(v5)))}
    delete(v1, -9223372036854775808)
    for fi20 := 0; fi20 < 10; fi20++ {
        switch (fi20 + fi20) {
        case 4294967295:
            _ = len(v3)
        case 128:
            v21 := GenS2{F0: (v18 ^ (v18 & v18)), F1: v4, F2: 310}
            _ = v21
        default:
            fmt.Println(fi20, v5)
        }
        v22 := GenS3{F0: (^(v4))}
        _ = fi20
        _ = v22
    }
    var v23 uint16 = 0
    v3 = append(v3, "a")
    v24 := v1[-497]
    if (2147483647 != 452) {
        var v25 float64 = (1.5 / 1e10)
        var v26 uint8 = 127
        _ = v25
        _ = v26
    } else {
        if (v2 < (-(v2))) {
            v27 := GenS1{F0: (v5 + 42)}
            var v28 int64 = v5
            fmt.Println(v23)
            _ = v27
            _ = v28
        }
    }
    if false {
        v29 := GenS4{F0: (v2 ^ v2), F1: 0, F2: (-2.25 == -2.25)}
        _ = v29
    } else {
        v4 -= v23
        var v30 float64 = float64((float32(0.0) - 0.0))
        v18 = v18
        _ = v30
    }
    var v31 string = "\n\t"
    v32 := GenS5{F0: 103, F1: 9223372036854775807, F2: -945}
    v18 = (^(v18))
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 int16 = 592
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
