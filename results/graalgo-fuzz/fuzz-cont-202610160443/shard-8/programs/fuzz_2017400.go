package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }
func genFn0(p0 uint32, p1 int64) uint8 {
    var v6 int = 42
    p0 += 2147483647
    p0 += 280
    _ = p0
    _ = p1
    _ = v6
    return 42
}
func genFn1(p0 int64, p1 int) int8 {
    var v7 uint32 = 4294967295
    v7 = v7
    v7 += v7
    v8 := GenS1{F0: 292, F1: (p1 % ((2147483647) | 1)), F2: "\n\t"}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 37
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{65535, 148}
    var v3 int = 32767
    v4 := make(map[int]bool)
    var v5 float64 = (((0.1 * (float64(1.5) - 1e10)) - 1.5) / 1.0)
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint32 = 2
        _ = fi9
        _ = v10
    }
    v1[198] = 2147483647
    fmt.Println(v5, v5)
    var v11 bool = false
    v5 += (0.0 / float64((float32(v5) - (-2.25 + 1e10))))
    v11 = (v11 || (!v11))
    var v12 int8 = 127
    var v13 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
