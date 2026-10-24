package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint16 = 163
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 1
    v6 := []int{32768, 32767}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int64 = 493
        _ = fi7
        _ = v8
    }
    fmt.Println(v1)
    var v9 float64 = ((1e10 - float64(float32(3.14159))) + 1.5)
    switch "hello" {
    case "\n\t":
        v5 = 1
    }
    switch "\n\t" {
    case "hello":
        var v10 int8 = 100
        _ = v10
    case "foo bar":
        v9 -= (1e10 - ((0.0 * v9) * (v9 + v9)))
    case "\n\t":
        var v11 uint8 = (v5 % ((v5) | 1))
        _ = v11
    default:
        if true {
            v1 += (^(v1))
            fmt.Println(v9, v5, v1)
        }
    }
    switch "\n\t" {
    case "\n\t":
        v12 := GenS1{F0: v5}
        _ = v12
    default:
        var v13 bool = (1 >= 0)
        _ = v13
    }
    fmt.Println(v5, v1)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
