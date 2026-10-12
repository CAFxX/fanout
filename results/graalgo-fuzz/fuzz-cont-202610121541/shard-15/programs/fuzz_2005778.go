package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
func genFn0(p0 int, p1 uint8, p2 int) uint16 {
    switch (p0 + p2) {
    case 128:
        for fi5 := 0; fi5 < 10; fi5++ {
            fmt.Println(fi5, fi5)
            p2 = (p0 / ((fi5) | 1))
            fmt.Println(fi5, p2)
            _ = fi5
        }
    case 2147483647:
        var v6 float32 = 1e10
        _ = v6
    case -9223372036854775808:
        v7 := GenS1{F0: 1, F1: -91, F2: -72}
        _ = v7
    default:
        switch 1 {
        case 0:
            p1 -= 1
        case 1:
            p1 += (p1 * p1)
        }
    }
    if (1 == 42) {
        if (1 <= 0) {
            var v8 uint64 = 250
            _ = v8
        } else {
            fmt.Println(p2)
            v9 := []int{p2, (p0 ^ (p0 % ((-1) | 1))), (p0 + p2), 32767}
            _ = v9
        }
        p2 -= -263
        p2 = (p2 & p0)
    } else {
        if true {
            fmt.Println(p0)
            var v10 uint16 = 0
            _ = v10
        } else {
            p0 -= p0
        }
        v11 := []int{}
        _ = v11
    }
    var v12 string = "hello"
    fmt.Println(p0, p0)
    v13 := []int{9223372036854775807, 9223372036854775807, p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 2
}
func genFn1(p0 bool, p1 string, p2 uint16) int {
    if p0 {
        p2 = genFn0(32767, 153, -9223372036854775808)
        if ((p2 % ((p2) | 1)) != p2) {
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := []int{-530, (fi14 % ((fi14) | 1)), fi14}
                p2 -= (p2 % (((p2 >> ((0) & 63))) | 1))
                var v16 int32 = 2147483647
                _ = fi14
                _ = v15
                _ = v16
            }
            p1 = (p1 + p1)
            fmt.Println(p1, p2, p0)
        }
    }
    var v17 uint8 = 128
    if (("" + p1) >= p1) {
        p2 -= uint16(v17)
        v17 = v17
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 float32 = -2.25
            _ = fi18
            _ = v19
        }
    }
    p0 = (!p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 32767
}

func main() {
    v1 := []string{"foo bar"}
    v2 := []uint32{1, 591, 42}
    var v3 uint16 = 32767
    var v4 int = 70
    for fi20 := 0; fi20 < 3; fi20++ {
        v3 = v3
        fmt.Println(v3)
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 int = 457
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 int64 = 0
    v4 += (v4 / (((v4 / (((v4 & (v4 * genFn1((92 > -128), "foo bar", (v3 | (v3 & v3)))))) | 1))) | 1))
    v23 -= v23
    var v24 uint32 = 412
    var v25 uint8 = 127
    var v26 int16 = 1000
    var v27 int8 = 60
    fmt.Println(v25, v26)
    if (0.1 == float64((1e10 + 3.14159))) {
        v24 = 1
        v26 -= (v26 / ((v26) | 1))
    } else {
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        var v28 string = "hello"
        fmt.Println(v23)
        _ = v28
    }
    v2 = append(v2, v24)
    if ((!(476 != int32(v3))) || (!false)) {
        if (float32(1.5) < ((0.0 - 0.0) / (((3.14159 * 1e10) + 0.0) - 3.14159))) {
            v4 -= (v4 + 4294967295)
            v2 = append(v2, 117)
        }
    } else {
        _ = len(v2)
        v27 = 111
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
