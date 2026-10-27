package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 int64 = 42
    v5 := GenS1{F0: 42, F1: v4, F2: 0}
    v6 := GenS1{F0: 1, F1: (v4 & v4), F2: 32767}
    var v7 int64 = v4
    v4 += (v4 + -1)
    var v8 float32 = 3.14159
    v7 += (v7 - (^(v7)))
    if ((!true) || false) {
        var v9 int64 = (v4 + 0)
        v1[2147483647] = 9223372036854775807
        _ = v9
    }
    v7 -= v7
    var v10 uint32 = 690
    v10 = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
