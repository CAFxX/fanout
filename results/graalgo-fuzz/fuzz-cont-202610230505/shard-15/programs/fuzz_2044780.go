package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int; F1 bool; F2 bool }
type GenS7 struct { F0 int; F1 int; F2 uint8 }
func genFn0(p0 uint8, p1 int64, p2 uint64) int8 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 float64 = 0.1
        _ = fi5
        _ = v6
    }
    v7 := []int{-661, -351}
    v8 := GenS1{F0: 1000, F1: p2, F2: -128}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = ((((((0.0 * 1e10) / (float64(1.5) - ((3.14159 - (((1.5 + ((float64(0.0) / 0.1) / (-2.25 * 3.14159))) / (0.0 - 1e10)) + -2.25)) + -2.25))) + 0.1) - 1.5) != 0.0) || (false && false))
    var v3 string = "\n\t"
    var v4 int32 = -2147483648
    switch 128 {
    case -234:
        fmt.Println(v2)
    case -211:
        var v9 int8 = 0
        _ = v9
    default:
        v4 = v4
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint16 = 32767
        v12 := GenS2{F0: 42, F1: 1}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int8 = genFn0(139, 205, 42)
    v3 = v3
    v4 = v4
    v3 = v3
    if (v3 <= "") {
        v14 := GenS3{F0: 192}
        v13 += (v13 * (v13 * 100))
        if (0.0 <= (1.5 - 0.0)) {
            v15 := GenS4{F0: "foo bar", F1: 711}
            switch 617 {
            case 42:
                if false {
                    fmt.Println(v2)
                    for fi16 := 0; fi16 < 4; fi16++ {
                        var v17 float64 = 0.1
                        v18 := GenS5{F0: 42, F1: 4294967295}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v19 := GenS6{F0: 32767, F1: (369 >= 18446744073709551615), F2: (v4 != (v4 - (v4 ^ 927)))}
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        v1[fi20] = 65535
                        _ = fi20
                    }
                }
            case -1:
                var v21 int64 = 289
                _ = v21
            case -984:
                fmt.Println(v13, v13, v4)
            }
            _ = v15
        }
        _ = v14
    }
    v2 = v2
    v22 := GenS7{F0: -9223372036854775808, F1: 0, F2: 127}
    var v23 float64 = 3.14159
    fmt.Println(v23)
    var v24 int8 = v13
    for fi25 := 0; fi25 < 20; fi25++ {
        v24 -= genFn0(128, 0, 0)
        fmt.Println(v24, v23, v4)
        _ = fi25
    }
    v24 -= (-(v13))
    var v26 uint8 = 1
    for fi27 := 0; fi27 < 19; fi27++ {
        if (!(v3 <= v3)) {
            v13 -= v24
        } else {
            v28 := []int{}
            var v29 bool = true
            _ = v28
            _ = v29
        }
        v1[fi27] = 2147483647
        _ = fi27
    }
    if (v2 || v2) {
        v24 = v24
        var v30 int16 = 42
        _ = v30
    } else {
        var v31 int = 32767
        v24 -= 1
        v32 := GenS2{F0: 9223372036854775807, F1: v26}
        _ = v31
        _ = v32
    }
    v33 := []int{int(v4), 42, 127, -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v33
    _ = v4
}
