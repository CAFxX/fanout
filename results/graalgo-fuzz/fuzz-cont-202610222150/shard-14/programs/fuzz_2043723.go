package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int {
    var v7 uint32 = 1
    v8 := []int{-750, -9223372036854775808, 128}
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 uint32 = 600
    var v2 float64 = 1.5
    v3 := []uint32{(^(v1)), (v1 * 510)}
    var v4 uint = 270
    var v5 int32 = -1
    v6 := make(map[int]uint32)
    var v9 uint8 = 20
    var v10 int32 = v5
    v9 += v9
    if ((v2 / v2) != v2) {
        v1 -= v1
    }
    var v11 uint64 = 1
    var v12 int64 = 223
    v13 := v6[9223372036854775807]
    v4 += (v4 & v4)
    v2 -= v2
    v14 := GenS1{F0: (v10 & v10)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
