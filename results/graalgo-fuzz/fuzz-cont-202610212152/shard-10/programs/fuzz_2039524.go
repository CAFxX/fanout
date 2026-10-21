package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 int }
func genFn0(p0 int, p1 int64, p2 uint64) uint64 {
    switch "" {
    case "a":
        p0 = 2147483647
    case "foo bar":
        var v7 int8 = 127
        _ = v7
    case "\n\t":
        p2 = (p2 & (p2 / (((p2 ^ (p2 >> ((42) & 63)))) | 1)))
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int = (p0 * (^(p0)))
            p1 = p1
            p1 -= p1
            _ = fi8
            _ = v9
        }
    }
    var v10 int64 = -1
    v11 := []int{p0, p0, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return p2
}

func main() {
    var v1 int = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 int16 = -1
    var v5 uint8 = 171
    v6 := make(map[string]string)
    for fi12 := 0; fi12 < 13; fi12++ {
        if false {
            fmt.Println(v5)
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                v5 = 128
                _ = fi13
            }
            fmt.Println(v4, fi12, v5)
        }
        var v14 int = (v1 + (fi12 + (fi12 ^ 1)))
        var v15 string = "a"
        _ = fi12
        _ = v14
        _ = v15
    }
    fmt.Println(v1, v1, v4)
    v16 := []int{(v1 | v1), (-(v1))}
    fmt.Println(v1, v5)
    var v17 int16 = (v4 ^ v4)
    v4 -= (v17 - -1)
    v18 := GenS1{F0: (v4 ^ v17), F1: v5, F2: true}
    _ = len(v3)
    v19 := GenS1{F0: v17, F1: 152, F2: false}
    v20 := GenS2{F0: v5, F1: 405, F2: v1}
    var v21 int16 = v4
    var v22 float64 = 1e10
    var v23 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
