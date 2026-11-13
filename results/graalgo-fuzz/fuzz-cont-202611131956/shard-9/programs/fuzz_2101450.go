package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
func genFn0(p0 bool, p1 uint32) uint8 {
    p1 = p1
    var v5 uint16 = 14
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: v5, F1: 33}
        _ = fi6
        _ = v7
    }
    var v8 int16 = 522
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 2
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 0
    var v3 float32 = 1.5
    var v4 float32 = (0.1 / (v3 - ((1.5 / v3) / float32(((0.0 + 1e10) * 1.5)))))
    v1 = append(v1, 2147483647)
    v2 = 0
    v9 := GenS1{F0: (v2 << ((469) & 63)), F1: 1}
    if false {
        v4 = v3
        fmt.Println(v3, v3)
    }
    v2 += 849
    v10 := GenS1{F0: (v2 | 831), F1: 978}
    switch "foo bar" {
    case "":
        var v11 int = -2147483648
        _ = v11
    case "foo bar":
        v2 += v2
    case "a":
        var v12 uint8 = 128
        _ = v12
    default:
        var v13 int32 = -614
        _ = v13
    }
    v4 -= v3
    v3 = v4
    v2 = v2
    if (434 <= 2) {
        v4 += 1.5
    }
    v2 = (v2 + (v2 - (v2 - v2)))
    var v14 uint64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
