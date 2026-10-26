package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 bool, p1 uint64) bool {
    v7 := GenS1{F0: -1, F1: 670}
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(fi8, p1, p0)
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 bool = p0
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return true
}
func genFn1(p0 uint16) uint {
    v12 := GenS2{F0: 18446744073709551615, F1: 18446744073709551615}
    v13 := []int{-1, 242, 2147483647, -389}
    if (-1 != 0) {
        var v14 uint16 = p0
        _ = v14
    }
    switch "foo bar" {
    case "a":
        fmt.Println(p0)
    case "hello":
        fmt.Println(p0)
    case "foo bar":
        if len(v13) > 0 {
            v13[0] = -1
        }
    }
    if (true && ((false && false) || (-223 == -2147483648))) {
        var v15 uint8 = 2
        _ = len(v13)
        _ = v15
    }
    _ = p0
    _ = v12
    _ = v13
    return 65535
}

func main() {
    var v1 int64 = 646
    var v2 float64 = (1.5 / (1.5 + -2.25))
    var v3 uint64 = 2
    var v4 int32 = 920
    v5 := make(map[int]uint32)
    v6 := []int{127, 340}
    for fi16 := 0; fi16 < 16; fi16++ {
        _ = len(v5)
        if len(v6) > 0 {
            v6[0] = (fi16 - fi16)
        }
        v3 += v3
        _ = fi16
    }
    v2 -= (v2 / 1e10)
    if false {
        v3 += 42
    } else {
        if true {
            v4 -= v4
        } else {
            v2 -= float64(1e10)
            v2 -= -2.25
            var v17 string = "a"
            _ = v17
        }
        v18 := GenS1{F0: 1, F1: 6}
        _ = v18
    }
    switch "foo bar" {
    case "\n\t":
        v2 -= 1e10
    case "":
        switch "foo bar" {
        case "foo bar":
            if (4294967295 > v3) {
                var v19 int16 = -1
                _ = v19
            }
        case "\n\t":
            for fi20 := 0; fi20 < 7; fi20++ {
                v4 -= v4
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 int8 = -1
                    _ = fi21
                    _ = v22
                }
                fmt.Println(fi20, v4, v1)
                _ = fi20
            }
        default:
            if true {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        }
    case "hello":
        switch "hello" {
        case "a":
            var v23 int32 = (v4 ^ v4)
            _ = v23
        }
    }
    fmt.Println(v4, v3, v2)
    v24 := GenS2{F0: 65535, F1: 840}
    v25 := GenS1{F0: 42, F1: 1}
    var v26 string = "foo bar"
    v4 = 113
    if (0 != 142) {
        _ = len(v6)
        fmt.Println(v26, v26)
        v5[-270] = 2147483648
    }
    v4 -= 454
    if len(v6) > 0 {
        v6[0] = -27
    }
    var v27 float64 = (v2 * float64(3.14159))
    v28 := GenS1{F0: 1, F1: 42}
    var v29 uint32 = 2
    var v30 int8 = -128
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 int = 128
        _ = fi31
        _ = v32
    }
    v29 -= v29
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
