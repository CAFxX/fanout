package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }

func main() {
    var v1 uint32 = 2
    var v2 uint64 = 1
    var v3 uint32 = v1
    var v4 uint8 = 0
    var v5 uint = 2
    if (18446744073709551615 != 0) {
        fmt.Println(v2)
        v4 -= v4
        var v6 uint64 = 18446744073709551615
        _ = v6
    }
    v3 = v1
    v7 := GenS1{F0: -705, F1: 0, F2: -9223372036854775808}
    if ((-2.25 + 0.0) >= 3.14159) {
        fmt.Println(v1, v2)
    }
    var v8 uint64 = v2
    var v9 float32 = float32((float64((0.0 / (1.5 - 3.14159))) * ((((-2.25 + 0.1) * float64(0.0)) - 1e10) - 1e10)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
