package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []int{-1, 42, 2147483647}
    var v2 uint32 = 577
    var v3 int64 = 0
    var v4 uint32 = v2
    var v5 uint32 = (v2 ^ v4)
    var v6 int8 = -1
    var v7 uint64 = 182
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint8 = 128
        _ = fi8
        _ = v9
    }
    var v10 int8 = 42
    v11 := GenS1{F0: v6}
    if len(v1) > 0 {
        v1[0] = 128
    }
    fmt.Println(v10, v6, v7)
    var v12 string = "a"
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := GenS1{F0: v6}
        v2 = v5
        _ = fi13
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
