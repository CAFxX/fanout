package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint16; F1 bool; F2 int32 }
func genFn0(p0 uint8) int8 {
    var v5 int8 = 42
    v5 += v5
    if false {
        p0 = (p0 & (^(p0)))
        p0 += 128
        if (!true) {
            v6 := []int{2147483647, 331, -567}
            p0 = (^(p0))
            _ = v6
        }
    }
    _ = p0
    _ = v5
    return (v5 * (v5 & (v5 << ((uint(p0)) & 63))))
}
func genFn1(p0 int16, p1 uint32) int8 {
    var v7 string = "foo bar"
    if (554 != 42) {
        v8 := []int{32767, 128, 127, 4294967295}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = v8
    } else {
        fmt.Println(p1)
        p1 -= (p1 / ((p1) | 1))
        v9 := GenS1{F0: (p1 + p1), F1: 18446744073709551615, F2: 26}
        _ = v9
    }
    v7 = "hello"
    _ = p0
    _ = p1
    _ = v7
    return -96
}

func main() {
    var v1 uint32 = 42
    var v2 uint32 = (^(v1))
    var v3 float32 = (0.0 - -2.25)
    v4 := []int{0, -9223372036854775808, -9223372036854775808}
    v3 -= 0.0
    var v10 int64 = 9223372036854775807
    for fi11 := 0; fi11 < 19; fi11++ {
        v1 += v2
        if (false && (311 > 1)) {
            v1 = v2
        } else {
            v4 = append(v4, fi11)
            var v12 uint16 = 255
            _ = v12
        }
        v10 -= (v10 + v10)
        _ = fi11
    }
    v3 = v3
    if len(v4) > 0 {
        v4[0] = 1
    }
    switch -9223372036854775808 {
    case -2147483648:
        if (473 <= 410) {
            v10 = v10
            v3 = (-2.25 + 0.0)
            var v13 float64 = (3.14159 / 1.0)
            _ = v13
        } else {
            fmt.Println(v10, v1, v2)
            var v14 bool = false
            _ = v14
        }
    case 127:
        v15 := GenS1{F0: v2, F1: 2, F2: 100}
        _ = v15
    case -477:
        if (-1 <= genFn0(84)) {
            fmt.Println(v1, v10)
            v16 := GenS2{F0: 2, F1: true, F2: -2147483648}
            _ = v16
        }
    }
    v2 = (v1 * 766)
    v17 := GenS2{F0: 1, F1: true, F2: 1}
    fmt.Println(v3)
    for fi18 := 0; fi18 < 12; fi18++ {
        if false {
            var v19 uint64 = 1
            _ = v19
        }
        _ = fi18
    }
    _ = len(v4)
    var v20 uint16 = 562
    fmt.Println(v2, v2, v20)
    v3 = (((1e10 / v3) * ((v3 * (v3 - 1e10)) + v3)) * (v3 + (0.0 / (1e10 / v3))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
