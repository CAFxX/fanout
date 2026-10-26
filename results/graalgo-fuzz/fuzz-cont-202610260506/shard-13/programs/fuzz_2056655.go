package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 uint64, p1 int32) uint16 {
    p1 -= 837
    switch 998 {
    case 602:
        var v4 int = 243
        _ = v4
    default:
        if (p0 >= 18446744073709551615) {
            var v5 uint = 402
            _ = v5
        }
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 uint8) int16 {
    p0 = (p0 % (((^(p0))) | 1))
    if (true || (820 < 42)) {
        var v6 uint16 = 2
        _ = v6
    }
    p0 = (p0 ^ (p0 << ((255) & 63)))
    if true {
        fmt.Println(p0)
        var v7 uint16 = 0
        _ = v7
    } else {
        p0 = (p0 / ((p0) | 1))
    }
    if (true && (562 == 1)) {
        for fi8 := 0; fi8 < 7; fi8++ {
            switch -1 {
            case -9223372036854775808:
                p0 = (p0 | p0)
            case 2147483647:
                fmt.Println(p0, p0)
            case 9223372036854775807:
                var v9 float64 = (-2.25 / 1.5)
                _ = v9
            }
            p0 += (^(p0))
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := GenS1{F0: 1, F1: 2}
                v12 := GenS1{F0: 0, F1: 127}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
        }
        v13 := GenS1{F0: 0, F1: 0}
        fmt.Println(p0)
        _ = v13
    }
    _ = p0
    return 0
}

func main() {
    var v1 bool = (!false)
    v2 := make(map[int]string)
    var v3 uint32 = 2147483647
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 int16 = 1
        _ = fi14
        _ = v15
    }
    v1 = (!true)
    var v16 uint8 = 101
    v17 := []int{}
    v17 = append(v17, -9223372036854775808)
    if (!(v1 && (uint32(v16) == (v3 - 359)))) {
        var v18 string = "foo bar"
        for fi19 := 0; fi19 < 17; fi19++ {
            v3 -= 1
            _ = fi19
        }
        _ = v18
    }
    v1 = false
    v1 = v1
    v20 := GenS1{F0: -1, F1: -128}
    var v21 uint16 = 350
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
