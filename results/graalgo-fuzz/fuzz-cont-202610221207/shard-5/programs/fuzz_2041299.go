package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int64; F2 bool }

func main() {
    var v1 int32 = -607
    var v2 int64 = 272
    var v3 int8 = 1
    var v4 int = 4294967295
    switch v4 {
    case 1:
        v5 := GenS1{F0: 2147483647}
        _ = v5
    case 128:
        v6 := GenS2{F0: 344}
        _ = v6
    case -375:
        var v7 uint32 = 0
        _ = v7
    }
    v8 := GenS3{F0: 996, F1: v2, F2: ((!true) || (0 < 621))}
    v9 := []int{}
    v2 -= -967
    v9 = append(v9, v4)
    var v10 uint32 = 65535
    switch "" {
    case "\n\t":
        v4 = 42
    }
    _ = len(v9)
    fmt.Println(v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
