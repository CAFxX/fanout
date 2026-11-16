package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint32 }
func genFn0(p0 int16, p1 uint16, p2 int) uint32 {
    v5 := GenS1{F0: 9223372036854775807, F1: 0, F2: 4294967295}
    var v6 bool = false
    if v6 {
        fmt.Println(p2, v6)
        var v7 int = p2
        _ = v7
    }
    fmt.Println(p1, v6, p1)
    v8 := []int{(-(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v8
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 1
    var v3 float64 = ((float64(-2.25) + 1.5) + 3.14159)
    v4 := make(map[string]int)
    var v9 int16 = -945
    v3 = (float64(0.0) * 0.1)
    v9 -= int16(v2)
    var v10 uint16 = 1
    v1[2] = 1
    v2 += uint8(v10)
    v2 += v2
    fmt.Println(v10)
    v2 += v2
    v11 := GenS1{F0: 1, F1: 42, F2: 545}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
