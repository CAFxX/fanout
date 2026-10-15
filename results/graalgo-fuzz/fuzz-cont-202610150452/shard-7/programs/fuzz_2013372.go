package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []uint32{0, 848, 65535, 275}
    v2 := make(map[string]bool)
    var v3 uint8 = 1
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    v6 := GenS1{F0: -798}
    v7 := GenS1{F0: 1}
    var v8 int64 = -398
    switch 32767 {
    case 42:
        v8 += v8
    default:
        for fi9 := 0; fi9 < 6; fi9++ {
            if (fi9 < fi9) {
                var v10 int = (fi9 | fi9)
                _ = v10
            }
            var v11 float64 = (0.0 * float64(float32(0.0)))
            _ = fi9
            _ = v11
        }
    }
    _ = len(v2)
    var v12 int8 = 1
    v12 += (-(v12))
    var v13 int64 = -829
    v14 := GenS1{F0: 181}
    if true {
        v15 := GenS1{F0: -102}
        switch "foo bar" {
        case "\n\t":
            v8 = (v13 % ((v8) | 1))
        }
        _ = v15
    }
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
