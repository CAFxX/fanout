package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 string; F2 int16 }
func genFn0(p0 uint) uint16 {
    var v6 uint16 = 2
    fmt.Println(p0, p0)
    var v7 float64 = -2.25
    _ = p0
    _ = v6
    _ = v7
    return 483
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{2147483648, 1}
    v3 := []uint32{}
    var v4 uint = 255
    var v5 int = -2147483648
    _ = len(v1)
    var v8 int64 = 42
    var v9 int16 = 42
    v4 += 42
    v8 = (v8 / (((v8 * (v8 % ((1) | 1)))) | 1))
    v10 := GenS1{F0: 688}
    v11 := GenS2{F0: 18446744073709551615, F1: "\n\t", F2: (v9 + (-(v9)))}
    switch "foo bar" {
    case "foo bar":
        v12 := GenS2{F0: 468, F1: "", F2: (v9 << ((20) & 63))}
        _ = v12
    case "\n\t":
        delete(v1, v5)
    default:
        var v13 float64 = 0.1
        _ = v13
    }
    switch "foo bar" {
    case "a":
        v8 = 1
    case "hello":
        v14 := GenS1{F0: 312}
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint = v4
        _ = fi15
        _ = v16
    }
    var v17 int64 = v8
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 uint16 = 211
        fmt.Println(fi18, v4)
        var v20 string = "\n\t"
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
