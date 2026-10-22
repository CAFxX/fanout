package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int32 }

func main() {
    v1 := []int{-2147483648}
    v2 := []int{-734, -9223372036854775808}
    v3 := make(map[uint32]bool)
    var v4 int = 314
    var v5 uint = 949
    v4 -= v4
    switch "foo bar" {
    case "":
        var v6 uint = (v5 % (((v5 * v5)) | 1))
        _ = v6
    case "hello":
        v5 += 786
    case "foo bar":
        var v7 float64 = 1.5
        _ = v7
    default:
        if len(v1) > 0 {
            v1[0] = 813
        }
    }
    fmt.Println(v5, v4)
    switch ("a" + "hello") {
    case "":
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint64 = 542
            _ = fi8
            _ = v9
        }
    case "foo bar":
        var v10 uint16 = 2
        _ = v10
    }
    var v11 float64 = 0.0
    if (!false) {
        var v12 int8 = 1
        v12 += v12
        _ = v12
    } else {
        v13 := GenS1{F0: v5}
        var v14 int32 = 13
        v5 = v5
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: (v5 % ((450) | 1))}
    if false {
        v5 -= v5
        switch "\n\t" {
        case "a":
            var v16 int = (v4 + (v4 & 1))
            _ = v16
        case "hello":
            v5 = 689
        case "foo bar":
            var v17 string = ("foo bar" + ("\n\t" + "hello"))
            _ = v17
        default:
            fmt.Println(v11)
        }
        v4 = v4
    } else {
        v18 := GenS2{F0: 0, F1: -932}
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
