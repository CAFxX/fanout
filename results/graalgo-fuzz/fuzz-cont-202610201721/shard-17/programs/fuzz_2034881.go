package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }

func main() {
    var v1 uint32 = 0
    v2 := []int{2147483647}
    v3 := make(map[uint32]uint32)
    var v4 int64 = 391
    if false {
        fmt.Println(v1, v4)
    }
    v5 := GenS1{F0: 42, F1: -489}
    v6 := GenS1{F0: -128, F1: 210}
    var v7 int = 9223372036854775807
    if (false && false) {
        if (!(!(!(170 == 2)))) {
            var v8 uint = 773
            _ = v8
        }
        switch ("" + ("\n\t" + (("" + "foo bar") + "\n\t"))) {
        case "":
            v1 += v1
        case "foo bar":
            v9 := GenS1{F0: 5, F1: v7}
            _ = v9
        default:
            v1 -= (v1 << ((2) & 63))
        }
    }
    var v10 uint64 = 0
    fmt.Println(v7)
    v4 += -507
    if (1.5 < 0.0) {
        v1 += v1
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v2 = append(v2, v7)
        }
    } else {
        v10 += (v10 | v10)
        _ = len(v3)
        var v11 int = (v7 & v7)
        _ = v11
    }
    v7 = (-(v7))
    v1 -= (v1 - v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
