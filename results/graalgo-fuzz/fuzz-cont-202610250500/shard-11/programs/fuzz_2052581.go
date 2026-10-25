package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 float64 = (1e10 / (0.1 * -2.25))
    var v2 float64 = 0.0
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    var v5 int8 = -1
    var v6 uint8 = 42
    var v7 float64 = 0.1
    v6 = 2
    var v8 int8 = (v5 ^ (v5 / (((v5 | (v5 + -128))) | 1)))
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: (-(v8))}
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(fi11)
            v12 := GenS1{F0: 100}
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    if (float64(((1e10 / -2.25) - (1e10 / float32((v1 - 0.1))))) != v2) {
        for fi13 := 0; fi13 < 1; fi13++ {
            v4[fi13] = (fi13 * fi13)
            fmt.Println(v6)
            _ = fi13
        }
        v14 := GenS1{F0: (v5 / ((v8) | 1))}
        _ = v14
    }
    v8 = (v8 ^ v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
