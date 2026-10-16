package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
func genFn0(p0 int8, p1 int32, p2 int8) int {
    v5 := GenS1{F0: 2, F1: 174, F2: 32767}
    if (!false) {
        v6 := GenS1{F0: 42, F1: 42, F2: -668}
        var v7 uint8 = 1
        p2 += p0
        _ = v6
        _ = v7
    }
    v8 := []int{911, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return 127
}

func main() {
    var v1 int16 = -1
    var v2 float64 = 0.1
    var v3 bool = ((4294967295 > 255) && true)
    var v4 float32 = 1e10
    v2 -= 0.0
    if v3 {
        v3 = ((0.0 / (v2 + 0.0)) != ((1.5 * float64(v4)) / (1.5 - v2)))
        var v9 uint = 18446744073709551615
        _ = v9
    }
    v4 -= 1.5
    var v10 int16 = v1
    var v11 int = -2147483648
    v10 = 951
    var v12 float32 = ((1e10 * v4) + v4)
    v1 += -478
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
