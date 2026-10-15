package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0(p0 bool, p1 int16, p2 uint) uint8 {
    var v4 float64 = 1e10
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}
func genFn1(p0 uint8, p1 int) uint16 {
    if (genFn0(true, 32767, 4) <= p0) {
        p1 += p1
    }
    p0 += (p0 & (p0 >> ((255) & 63)))
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS1{F0: "foo bar", F1: 247, F2: 42}
        _ = fi5
        _ = v6
    }
    switch ("" + "a") {
    case "\n\t":
        var v7 uint8 = p0
        _ = v7
    case "":
        v8 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: 18446744073709551615}
        _ = v8
    default:
        var v9 uint32 = 2
        _ = v9
    }
    if (!false) {
        v10 := GenS2{F0: 300, F1: 42}
        for fi11 := 0; fi11 < 12; fi11++ {
            p0 -= (p0 / ((p0) | 1))
            switch (fi11 | (fi11 | -9223372036854775808)) {
            case -9223372036854775808:
                var v12 uint16 = 857
                _ = v12
            case 127:
                var v13 uint = 2
                _ = v13
            default:
                var v14 uint8 = 194
                _ = v14
            }
            p1 = (fi11 * p1)
            _ = fi11
        }
        v15 := GenS1{F0: "", F1: 42, F2: 2}
        _ = v10
        _ = v15
    }
    _ = p0
    _ = p1
    return 32767
}

func main() {
    v1 := []uint32{}
    v2 := []int{-9223372036854775808, -430, 2147483647}
    v3 := make(map[int]uint32)
    if ((!false) && true) {
        var v16 int16 = 0
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v16
    }
    var v17 int8 = 1
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 uint32 = 1
        _ = fi18
        _ = v19
    }
    var v20 int8 = -11
    v17 = -128
    if ((v20 - v20) >= (v20 / ((0) | 1))) {
        var v21 uint64 = 4294967295
        fmt.Println(v21, v20)
        _ = v21
    }
    var v22 float64 = 0.0
    v23 := GenS1{F0: ("foo bar" + "\n\t"), F1: 40, F2: 42}
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 uint16 = 65535
        _ = fi24
        _ = v25
    }
    var v26 uint64 = 1
    switch "\n\t" {
    case "":
        fmt.Println(v17, v20)
    case "a":
        var v27 int8 = v17
        _ = v27
    case "\n\t":
        v28 := GenS2{F0: 32767, F1: v20}
        _ = v28
    default:
        var v29 float64 = (v22 - v22)
        _ = v29
    }
    v30 := GenS1{F0: "", F1: v26, F2: v26}
    var v31 uint64 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (128 == genFn0(false, 0, 4294967295)) {
        fmt.Println(v22, v20)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v30
    _ = v31
}
