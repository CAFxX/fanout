package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]uint32)
    var v3 int64 = 537
    var v4 int64 = v3
    var v5 bool = false
    var v6 int8 = 100
    v6 -= v6
    fmt.Println(v5, v3, v1)
    v7 := []int{-574, 1}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 bool = false
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 18446744073709551615, F1: 42}
    v5 = (751 != 127)
    var v11 uint32 = v1
    v12 := GenS1{F0: 42, F1: 127}
    fmt.Println(v5, v6)
    var v13 uint64 = 104
    v4 += 0
    v14 := GenS2{F0: 327}
    for fi15 := 0; fi15 < 1; fi15++ {
        v1 = 1
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
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
    _ = v7
}
