package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int; F1 string }

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    v4 := GenS1{F0: 255, F1: -9223372036854775808}
    for fi5 := 0; fi5 < 8; fi5++ {
        v1 -= v1
        v1 = fi5
        _ = fi5
    }
    if false {
        v1 -= v1
        _ = len(v3)
    } else {
        v6 := GenS1{F0: 318, F1: 9223372036854775807}
        switch "a" {
        case "a":
            v1 = 148
        case "":
            v7 := GenS2{F0: v1, F1: "foo bar"}
            _ = v7
        default:
            var v8 int64 = 9223372036854775807
            _ = v8
        }
        v1 = v1
        _ = v6
    }
    v1 = (v1 >> ((1) & 63))
    v9 := []int{(v1 % ((v1) | 1))}
    if false {
        v1 = v1
        var v10 uint32 = 65535
        if false {
            v9 = append(v9, v1)
            v1 = (v1 | v1)
        } else {
            if len(v9) > 0 {
                v9[0] = -9223372036854775808
            }
            var v11 int16 = -32768
            _ = v11
        }
        _ = v10
    }
    v1 = (v1 + (v1 >> ((690) & 63)))
    fmt.Println(v1)
    v12 := GenS1{F0: 42, F1: 42}
    var v13 int32 = 1
    if (0 <= 42) {
        _ = len(v9)
    }
    if ((65535 >= 42) || (497 >= 93)) {
        var v14 float32 = (1.5 / ((1e10 / 1.5) / 3.14159))
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
