package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
func genFn0(p0 int8, p1 int, p2 int8) uint32 {
    p1 = p1
    p1 -= p1
    v7 := []int{p1, p1, -687, 4294967295}
    for fi8 := 0; fi8 < 6; fi8++ {
        if ((742 != 0) || true) {
            _ = len(v7)
        } else {
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: 32767}
                var v11 int64 = 0
                fmt.Println(p2)
                _ = fi9
                _ = v10
                _ = v11
            }
            p0 = p2
            for fi12 := 0; fi12 < 7; fi12++ {
                p2 = (p2 | (p0 % ((127) | 1)))
                _ = fi12
            }
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483648
}

func main() {
    v1 := []int{-670, -758, 2147483647, 4294967295}
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    v4 := []int{127, 4294967295}
    var v5 int8 = -126
    var v6 int8 = v5
    v5 = -128
    fmt.Println(v5, v5)
    if false {
        fmt.Println(v5, v6)
    } else {
        v13 := GenS1{F0: -1}
        v5 = v6
        _ = v13
    }
    switch "a" {
    case "foo bar":
        var v14 float32 = 1e10
        _ = v14
    case "a":
        v5 = (v5 ^ v6)
    default:
        if (0 != -32768) {
            var v15 uint32 = genFn0((v5 % (((v6 % ((v6) | 1))) | 1)), -9223372036854775808, v5)
            v6 -= (v6 * (v5 % (((v6 - (v6 + 42))) | 1)))
            _ = v15
        } else {
            fmt.Println(v6, v5)
            switch 128 {
            case 9223372036854775807:
                v16 := GenS1{F0: 0}
                _ = v16
            }
            if false {
                v3[648] = false
                var v17 int32 = 42
                _ = v17
            }
        }
    }
    switch 9223372036854775807 {
    case 128:
        var v18 string = "foo bar"
        _ = v18
    default:
        var v19 float64 = (0.1 * (((float64(1.5) / (3.14159 / -2.25)) * float64((0.1 - 1e10))) * (float64(-2.25) * (1e10 + 1.5))))
        _ = v19
    }
    fmt.Println(v6)
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 bool = (443 < -616)
        var v22 uint = 1
        v21 = (fi20 < (fi20 >> ((v22) & 63)))
        _ = fi20
        _ = v21
        _ = v22
    }
    delete(v3, 0)
    v23 := GenS2{F0: 805, F1: 2147483647, F2: true}
    v5 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
