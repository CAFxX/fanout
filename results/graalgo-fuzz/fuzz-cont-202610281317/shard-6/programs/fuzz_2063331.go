package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    v4 := GenS1{F0: -1, F1: 42, F2: -1}
    if (161 < 127) {
        v5 := GenS1{F0: -490, F1: -9223372036854775808, F2: 127}
        var v6 string = ("\n\t" + "a")
        _ = v5
        _ = v6
    }
    switch "a" {
    case "\n\t":
        var v7 float32 = 1.5
        _ = v7
    case "a":
        var v8 int64 = -9223372036854775808
        _ = v8
    }
    v9 := GenS1{F0: 395, F1: -1, F2: 0}
    v1 = float32((0.1 * 0.1))
    var v10 uint = 2
    var v11 int32 = -1
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
