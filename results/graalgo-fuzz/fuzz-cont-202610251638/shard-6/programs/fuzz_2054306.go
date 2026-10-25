package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{0, -686, -3}
    v3 := make(map[int]bool)
    var v4 uint16 = 598
    var v5 bool = false
    v6 := GenS1{F0: 1}
    v4 = v4
    var v7 float64 = 1e10
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: 1}
        var v10 uint64 = 0
        var v11 uint32 = 2147483648
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 float64 = v7
    var v13 uint32 = 631
    v14 := GenS1{F0: 9223372036854775807}
    v7 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
