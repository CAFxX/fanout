package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 uint = 65535
    v4 := make(map[uint32]int)
    v5 := make(map[int]bool)
    fmt.Println(v3)
    v6 := []int{-85, 32767, 9223372036854775807, 127}
    fmt.Println(v3)
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int64 = 42
        v8 = int64(fi7)
        _ = len(v6)
        _ = fi7
        _ = v8
    }
    var v9 int = 4294967295
    v9 -= (v9 << (((v3 >> ((v3) & 63))) & 63))
    v10 := GenS1{F0: 958, F1: 0}
    var v11 uint32 = 1
    v11 = v11
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 float32 = 0.1
        var v14 int32 = 360
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 = 943
    v11 -= 0
    var v15 float32 = (1e10 + (float32(0.0) / 3.14159))
    v11 += 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
