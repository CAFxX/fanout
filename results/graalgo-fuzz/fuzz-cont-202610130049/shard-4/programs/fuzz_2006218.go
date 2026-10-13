package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }

func main() {
    var v1 int32 = -1
    v2 := make(map[string]int)
    var v3 float32 = float32(float64(0.1))
    fmt.Println(v3, v3)
    v1 = (v1 | 42)
    var v4 uint = 18446744073709551615
    v3 -= float32(3.14159)
    v5 := v2["\n\t"]
    v3 += 0.1
    v5 += 9223372036854775807
    switch -435 {
    case -9223372036854775808:
        _ = len(v2)
    case 32768:
        v6 := GenS1{F0: 35, F1: 0}
        _ = v6
    case -52:
        v4 = v4
    default:
        v4 = v4
    }
    v1 += (v1 % ((-102) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
