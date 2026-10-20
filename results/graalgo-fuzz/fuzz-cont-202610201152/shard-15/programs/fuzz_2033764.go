package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int64; F2 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 0
    v3 := make(map[int]bool)
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 26
    if true {
        fmt.Println(v2, v2)
    }
    switch -1 {
    case 9223372036854775807:
        v6 := GenS1{F0: (v5 | 42)}
        _ = v6
    case 588:
        v7 := []int{32767}
        _ = v7
    default:
        fmt.Println(v2, v2)
    }
    v2 = -9223372036854775808
    var v8 float32 = 0.0
    var v9 uint16 = v5
    v10 := []int{}
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS1{F0: v5}
        var v13 int = -730
        switch "foo bar" {
        case "\n\t":
            if true {
                fmt.Println(v2)
                var v14 int8 = -128
                if len(v10) > 0 {
                    v10[0] = (v13 | fi11)
                }
                _ = v14
            }
        default:
            v15 := GenS2{F0: 100, F1: (v2 * (v2 / ((v2) | 1))), F2: 18446744073709551615}
            _ = v15
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
