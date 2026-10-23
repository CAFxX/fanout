package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }

func main() {
    var v1 uint8 = 2
    var v2 int16 = -1
    var v3 int = 42
    v4 := make(map[int]int)
    v5 := []int{(v3 | v3), v3}
    var v6 int32 = 1
    v6 += (v6 - v6)
    var v7 uint16 = 0
    var v8 bool = false
    v9 := GenS1{F0: 42, F1: (v3 & v3)}
    v1 = 119
    var v10 uint32 = uint32(v7)
    fmt.Println(v1, v10)
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint64 = 1
        var v13 int64 = 35
        v6 = (v6 >> ((uint(v10)) & 63))
        _ = fi11
        _ = v12
        _ = v13
    }
    if ((-(v2)) >= (v2 >> ((586) & 63))) {
        var v14 uint = 423
        v1 = (v1 - (v1 * v1))
        delete(v4, -2147483648)
        _ = v14
    }
    fmt.Println(v3, v7)
    v5 = append(v5, (v3 & (v3 | v3)))
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
