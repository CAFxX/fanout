package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 int, p1 int8, p2 uint8) int64 {
    fmt.Println(p1)
    v4 := GenS1{F0: 0}
    var v5 uint64 = uint64(p2)
    var v6 bool = ((("\n\t" + ("foo bar" + "foo bar")) != "foo bar") || (!false))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return -614
}
func genFn1(p0 uint64, p1 uint16, p2 string) int64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := []int{}
            p0 -= p0
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 uint32 = 101
    var v11 float64 = float64(float32(1.5))
    if ((1e10 - 1e10) != (((v11 - ((float64(float32((-2.25 / v11))) + -2.25) * 3.14159)) / 1e10) / -2.25)) {
        var v12 int32 = 1
        var v13 string = "a"
        v14 := []int{-9223372036854775808, -2147483648, -954, 9223372036854775807}
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return genFn0(int(p1), 100, 0)
}

func main() {
    var v1 bool = false
    v2 := []uint32{42}
    var v3 uint8 = 2
    v15 := GenS1{F0: 18446744073709551615}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!v1) {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 int = int(v3)
            v18 := GenS2{F0: ((!v1) && v1), F1: fi16}
            v1 = (0 != 670)
            _ = fi16
            _ = v17
            _ = v18
        }
        switch "foo bar" {
        case "a":
            var v19 uint64 = 456
            _ = v19
        case "":
            v20 := GenS1{F0: 1}
            _ = v20
        case "foo bar":
            v1 = (3.14159 < 0.0)
        default:
            v1 = (!(255 == 65535))
        }
        fmt.Println(v3)
    }
    fmt.Println(v3, v3)
    var v21 uint8 = v3
    v21 -= v3
    var v22 float64 = 0.1
    v2 = append(v2, 65535)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
