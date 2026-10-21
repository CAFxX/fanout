package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 bool }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }
func genFn0(p0 bool, p1 int32) uint64 {
    if (("" == "") || (false && false)) {
        var v4 uint64 = 4294967295
        if ((42 > -59) && p0) {
            v5 := []int{}
            _ = v5
        }
        _ = v4
    } else {
        v6 := GenS1{F0: -1, F1: 0, F2: (206 >= 42)}
        for fi7 := 0; fi7 < 13; fi7++ {
            for fi8 := 0; fi8 < 16; fi8++ {
                p1 = p1
                var v9 int = fi8
                _ = fi8
                _ = v9
            }
            p1 = (p1 - p1)
            _ = fi7
        }
        p1 = 684
        _ = v6
    }
    switch "" {
    case "foo bar":
        if ((2147483647 >= 0) || p0) {
            v10 := GenS1{F0: -73, F1: 9223372036854775807, F2: p0}
            switch "hello" {
            case "\n\t":
                p0 = (0 > -1)
            }
            _ = v10
        } else {
            if p0 {
                v11 := []int{-381, -9223372036854775808, 1}
                _ = v11
            }
            fmt.Println(p0)
            for fi12 := 0; fi12 < 14; fi12++ {
                p0 = false
                p1 = (^(p1))
                var v13 int8 = -1
                _ = fi12
                _ = v13
            }
        }
    case "\n\t":
        p1 = (p1 >> ((841) & 63))
    case "":
        var v14 int = 128
        _ = v14
    }
    _ = p0
    _ = p1
    return 164
}
func genFn1(p0 int64, p1 int8) int {
    var v15 bool = true
    v16 := []int{1, -514, 4294967295}
    if (-853 == 42) {
        p1 = -1
    } else {
        p0 = p0
    }
    fmt.Println(p1, p0)
    switch "" {
    case "foo bar":
        var v17 int16 = 109
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return -2147483648
}

func main() {
    var v1 uint = 65535
    var v2 uint = v1
    v3 := make(map[uint32]int)
    v2 = (^(v2))
    fmt.Println(v1, v2)
    var v18 uint16 = 42
    v19 := GenS1{F0: -786, F1: -702, F2: true}
    var v20 uint64 = 687
    fmt.Println(v2, v18)
    var v21 float64 = float64(-2.25)
    fmt.Println(v2, v20)
    v22 := GenS1{F0: 9223372036854775807, F1: -1, F2: true}
    v20 = genFn0((!true), 2147483647)
    var v23 uint64 = v20
    v21 -= (0.0 / ((-2.25 * 1.5) + v21))
    v18 = (v18 | 255)
    if (!(!((v21 / v21) != (1e10 + float64(float32(v21)))))) {
        v24 := GenS2{F0: 127, F1: v20, F2: 128}
        v21 = (v21 - v21)
        v25 := GenS2{F0: 0, F1: genFn0((false && false), 42), F2: 128}
        _ = v24
        _ = v25
    } else {
        fmt.Println(v20)
    }
    v3[42] = -9223372036854775808
    v3[756] = -2147483648
    v2 = (v2 ^ (v1 << ((v1) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
