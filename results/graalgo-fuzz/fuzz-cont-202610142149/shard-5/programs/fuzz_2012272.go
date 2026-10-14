package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[uint32]int)
    var v3 int8 = 100
    var v4 uint32 = 2
    v5 := GenS1{F0: -974}
    fmt.Println(v3, v4)
    v6 := GenS1{F0: 1}
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v4)
        var v8 int64 = 0
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 77
    for fi10 := 0; fi10 < 15; fi10++ {
        v3 -= (v3 ^ v3)
        _ = fi10
    }
    v4 -= v4
    v11 := GenS1{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
