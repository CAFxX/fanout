package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 uint8 = 42
    v4 := make(map[int]bool)
    var v5 int8 = 127
    v6 := []string{"hello", "foo bar"}
    var v7 uint64 = 18446744073709551615
    v8 := GenS1{F0: (false && (2 < 730)), F1: 0}
    v5 = 127
    fmt.Println(v5, v7)
    v7 -= 0
    fmt.Println(v5, v5, v7)
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint64 = (v7 >> ((65535) & 63))
        v11 := GenS1{F0: false, F1: 235}
        fmt.Println(v5, v10, fi9)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = len(v2)
    v5 = -128
    v5 = -74
    v3 = (v3 * v3)
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
