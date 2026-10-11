package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint64 = 114
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    v5 := make(map[string]int)
    var v6 uint64 = v1
    fmt.Println(v1)
    v7 := []int{}
    v1 = 0
    v8 := v2[-2147483648]
    v1 = (v1 - (v1 ^ v6))
    v9 := GenS1{F0: -437, F1: 1}
    var v10 int = v8
    switch "\n\t" {
    case "hello":
        v6 = v6
    default:
        v11 := GenS2{F0: -32768, F1: 0}
        _ = v11
    }
    if len(v7) > 0 {
        v7[0] = (v10 ^ (v10 * 9223372036854775807))
    }
    if (4294967295 <= 109) {
        _ = len(v3)
        var v12 float32 = ((1e10 * (0.0 - 3.14159)) - (((((1e10 / 1e10) + float32(1e10)) + 0.0) * float32(3.14159)) - 0.1))
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
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
    _ = v9
}
