package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int = 4294967295
    var v2 uint = 4294967295
    var v3 int64 = -9223372036854775808
    v4 := make(map[int]int)
    fmt.Println(v3, v3, v1)
    v4[(^(v1))] = (v1 | 215)
    fmt.Println(v1, v3, v1)
    v5 := v4[v1]
    v6 := GenS1{F0: (0.1 > (1e10 / 1.0))}
    v7 := []int{440, (v1 * (^(v1)))}
    var v8 uint8 = 1
    var v9 int = v5
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint8 = (v8 / ((v8) | 1))
        _ = fi10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
