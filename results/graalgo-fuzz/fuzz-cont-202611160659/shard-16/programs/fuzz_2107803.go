package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }
func genFn0(p0 int8) uint {
    p0 = 100
    fmt.Println(p0)
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := []string{"hello", "\n\t"}
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    var v4 int = 262
    var v5 uint64 = 18446744073709551615
    var v6 uint32 = 2
    v7 := GenS1{F0: -817, F1: (-262 < 9223372036854775807)}
    v6 = 182
    fmt.Println(v4)
    v4 += v4
    v5 = 0
    v8 := GenS2{F0: false, F1: 1, F2: 0}
    v4 -= 427
    fmt.Println(v4)
    var v9 int16 = -139
    v10 := GenS1{F0: (v9 * 42), F1: true}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
