package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }

func main() {
    var v1 float64 = ((0.0 + 1.5) / -2.25)
    v2 := make(map[uint32]bool)
    var v3 uint = 630
    v4 := []uint32{2147483648, 42}
    var v5 uint64 = 1
    var v6 int32 = 42
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int16 = -328
        _ = fi7
        _ = v8
    }
    v2[2147483648] = false
    var v9 uint = v3
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(v1, v6)
        _ = fi10
    }
    v11 := GenS1{F0: 115, F1: "hello"}
    for fi12 := 0; fi12 < 16; fi12++ {
        v6 = v6
        v3 = (v3 + v9)
        _ = fi12
    }
    switch "foo bar" {
    case "\n\t":
        var v13 int16 = -220
        _ = v13
    case "foo bar":
        v3 = v9
    case "hello":
        v3 += 320
    default:
        var v14 float64 = ((v1 / ((1e10 * v1) + (0.0 / 0.1))) + -2.25)
        _ = v14
    }
    v5 = (v5 ^ v5)
    v9 -= 4294967295
    if false {
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(v5)
            _ = fi15
        }
    }
    v2[766] = false
    var v16 uint16 = 2
    v17 := v2[769]
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
