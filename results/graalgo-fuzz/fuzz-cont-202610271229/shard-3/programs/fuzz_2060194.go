package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = ((float64(3.14159) / -2.25) - (0.0 * ((0.1 / (3.14159 + 1.5)) * -2.25)))
    v3 := make(map[int]string)
    v4 := []uint32{2147483647, 308}
    var v5 int64 = -1
    v6 := []int{0, 613, 0}
    v7 := GenS1{F0: -152}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 float32 = ((0.1 + 3.14159) + 3.14159)
        _ = fi8
        _ = v9
    }
    if (!false) {
        var v10 uint64 = 42
        v10 = 175
        _ = v10
    }
    v5 = (v5 & v5)
    var v11 uint8 = 133
    v5 = (v5 % (((v5 * int64(v11))) | 1))
    var v12 bool = (0 <= 2)
    v12 = v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
