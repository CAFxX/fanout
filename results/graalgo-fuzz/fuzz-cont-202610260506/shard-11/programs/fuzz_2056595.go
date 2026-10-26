package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint64) uint16 {
    var v4 uint64 = (^(p1))
    var v5 int64 = 517
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 491
}
func genFn1(p0 int64) uint {
    p0 = (p0 & 9223372036854775807)
    p0 -= (p0 & 42)
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]int)
    var v3 int64 = -1
    v3 = (-(v1))
    fmt.Println(v3, v3)
    var v6 float32 = (((((float32(((1e10 / float64(1e10)) / 0.1)) - (float32(1e10) + 1.5)) * -2.25) + (1e10 + 3.14159)) * 1e10) * ((float32(3.14159) / 1.5) + 0.0))
    var v7 uint16 = 42
    v8 := GenS1{F0: 1}
    var v9 float32 = ((v6 + v6) - (float32(0.1) + (1.5 * -2.25)))
    v3 = (v1 - (v3 | 762))
    var v10 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
