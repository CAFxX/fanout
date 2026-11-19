package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint; F2 int64 }
type GenS5 struct { F0 int64; F1 bool }
func genFn0() int8 {
    v4 := []int{-803, 9223372036854775807, -2147483648}
    var v5 int16 = 0
    v6 := GenS1{F0: 345, F1: 2}
    v5 = 1
    switch -1 {
    case -203:
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 float64 = (1.5 / 0.1)
            var v9 int64 = -9223372036854775808
            v10 := GenS2{F0: 4294967295}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    case -797:
        if (((0.1 == 1.5) && false) || (-559 < 9223372036854775807)) {
            var v11 string = "foo bar"
            _ = v11
        } else {
            var v12 uint16 = 1
            v5 = -32768
            _ = v12
        }
    case 1:
        if false {
            v5 -= -32768
            v13 := GenS1{F0: 42, F1: 2}
            _ = v13
        } else {
            v14 := GenS1{F0: 2, F1: 255}
            _ = v14
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return 85
}
func genFn1(p0 string, p1 uint64, p2 int64) int8 {
    var v15 bool = false
    if (((p2 | (^(p2))) != (p2 - p2)) && (v15 || (!false))) {
        var v16 int16 = -886
        var v17 uint64 = (p1 - (p1 >> ((0) & 63)))
        _ = v16
        _ = v17
    }
    p2 -= (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return genFn0()
}

func main() {
    var v1 float64 = float64(0.1)
    var v2 bool = (123 >= 1)
    var v3 bool = (128 >= 49)
    fmt.Println(v1, v3)
    fmt.Println(v2, v1)
    v18 := GenS1{F0: 870, F1: 1}
    v19 := GenS1{F0: 313, F1: 18446744073709551615}
    for fi20 := 0; fi20 < 20; fi20++ {
        v2 = v2
        _ = fi20
    }
    if false {
        v3 = false
        v3 = v3
    } else {
        v21 := GenS3{F0: 127, F1: -840}
        fmt.Println(v1)
        if v2 {
            fmt.Println(v1)
            if (255 < 2) {
                v22 := []int{}
                if v2 {
                    fmt.Println(v1, v3)
                }
                _ = v22
            } else {
                var v23 int8 = 60
                if v2 {
                    v24 := []int{32768, 4294967295, -993, -1}
                    if v2 {
                        var v25 uint = 2
                        _ = v25
                    }
                    _ = v24
                }
                _ = v23
            }
        } else {
            v26 := []int{563, -1}
            fmt.Println(v1, v1, v1)
            v27 := GenS4{F0: 65535, F1: 2, F2: -212}
            _ = v26
            _ = v27
        }
        _ = v21
    }
    v1 += ((-2.25 / v1) + (v1 * v1))
    switch 0 {
    case 9223372036854775807:
        v1 += 1.5
    case 32767:
        v28 := []int{32767, 0, -1, 9223372036854775807}
        _ = v28
    default:
        v1 = float64(-2.25)
    }
    v29 := []int{-2147483648}
    v30 := GenS5{F0: 9223372036854775807, F1: v2}
    v3 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
