package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int64; F2 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    var v3 int = 9223372036854775807
    v1["foo bar"] = true
    var v4 string = "hello"
    var v5 uint64 = 0
    v5 = (v5 * v5)
    v6 := GenS1{F0: (v5 - v5)}
    v5 += v5
    v4 = v4
    var v7 int = 557
    v3 -= 2147483647
    var v8 string = v4
    v9 := GenS2{F0: 100, F1: 758, F2: 181}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
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
    _ = v9
}
