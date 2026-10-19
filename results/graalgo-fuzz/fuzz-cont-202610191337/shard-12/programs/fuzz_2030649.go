package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 127
    v3 := make(map[int]uint32)
    var v4 float32 = ((1.5 - (float32(1e10) * (0.0 * (float32((0.0 * (1e10 / float64(-2.25)))) - 3.14159)))) - (0.1 * 0.0))
    v5 := []string{}
    var v6 uint16 = 0
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            _ = len(v1)
            _ = fi8
        }
        v2 = (v2 % (((v2 % ((0) | 1))) | 1))
        _ = fi7
    }
    v6 = v6
    var v9 int16 = -32768
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            switch "\n\t" {
            case "foo bar":
                v12 := GenS1{F0: 907, F1: v2}
                _ = v12
            case "hello":
                v13 := v3[-2147483648]
                _ = v13
            }
            v9 -= v9
            _ = fi11
        }
        v9 += (v9 ^ v9)
        _ = len(v5)
        _ = fi10
    }
    v14 := GenS1{F0: 803, F1: 0}
    var v15 int32 = 42
    var v16 uint = 857
    v17 := GenS2{F0: -585, F1: 42}
    var v18 int = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
