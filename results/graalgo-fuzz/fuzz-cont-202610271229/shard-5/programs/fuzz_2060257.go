package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
type GenS2 struct { F0 uint }
func genFn0(p0 int64, p1 int32) int32 {
    p0 += (^(p0))
    v4 := []int{}
    _ = p0
    _ = p1
    _ = v4
    return (-(p1))
}
func genFn1() uint32 {
    var v5 uint8 = 0
    switch 0 {
    case 127:
        v5 -= (v5 / ((42) | 1))
    case -2147483648:
        if (9223372036854775807 != 133) {
            fmt.Println(v5)
        }
    }
    _ = v5
    return 0
}

func main() {
    v1 := []int{}
    var v2 uint64 = 4294967295
    v3 := make(map[string]bool)
    fmt.Println(v2)
    switch "hello" {
    case "hello":
        v2 = 0
    case "a":
        var v6 float32 = -2.25
        _ = v6
    case "":
        for fi7 := 0; fi7 < 14; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                fmt.Println(fi7, fi7, fi8)
                _ = len(v3)
                _ = fi8
            }
            v2 += v2
            _ = fi7
        }
    }
    v2 = v2
    v2 = v2
    v9 := GenS1{F0: 9223372036854775807, F1: 2147483647, F2: 0}
    var v10 int = 4294967295
    fmt.Println(v10, v10)
    v11 := GenS1{F0: 42, F1: 2, F2: 0}
    v12 := v3["foo bar"]
    switch "a" {
    case "\n\t":
        var v13 int8 = -9
        _ = v13
    case "":
        var v14 int64 = 1
        _ = v14
    case "hello":
        v2 = (v2 - (^(v2)))
    default:
        switch (v10 << ((4294967295) & 63)) {
        case 417:
            v1 = append(v1, (v10 >> ((782) & 63)))
        case 4294967295:
            v2 -= (v2 ^ (v2 >> ((1) & 63)))
        case 137:
            for fi15 := 0; fi15 < 11; fi15++ {
                v16 := GenS1{F0: -1, F1: 418, F2: (^(v10))}
                v12 = (-354 <= -252)
                v17 := GenS2{F0: 42}
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    }
    v3[("hello" + ("foo bar" + ""))] = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
