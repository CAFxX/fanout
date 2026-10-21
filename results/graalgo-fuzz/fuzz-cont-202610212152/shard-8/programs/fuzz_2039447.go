package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }

func main() {
    var v1 int32 = 727
    v2 := make(map[uint32]int)
    v3 := []int{773, 4294967295, 42}
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    var v6 int64 = -9223372036854775808
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int16 = 1
        _ = fi7
        _ = v8
    }
    var v9 int = 4294967295
    v9 = v9
    var v10 uint8 = 135
    v11 := GenS1{F0: 0, F1: 42, F2: -23}
    var v12 uint32 = 0
    v13 := v2[v12]
    var v14 int64 = v6
    v9 += v13
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
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
