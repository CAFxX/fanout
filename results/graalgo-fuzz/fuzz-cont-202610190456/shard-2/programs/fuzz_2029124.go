package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint8 }

func main() {
    var v1 float64 = 1e10
    v2 := []uint32{4294967295, 975, 2147483648, 888}
    v3 := []uint32{2, 2, 129, 65535}
    var v4 int64 = 42
    v5 := make(map[int]int)
    switch "" {
    case "":
        var v6 uint32 = 777
        _ = v6
    }
    v4 = v4
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 float64 = v1
        _ = fi7
        _ = v8
    }
    if (false || (3.14159 <= 0.0)) {
        var v9 int8 = 100
        _ = v9
    } else {
        var v10 string = "a"
        _ = v10
    }
    _ = len(v5)
    _ = len(v5)
    var v11 uint64 = 385
    if true {
        var v12 int = -303
        _ = v12
    } else {
        v1 += v1
        v13 := GenS1{F0: 1, F1: 42, F2: 127}
        var v14 uint32 = 2147483647
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
