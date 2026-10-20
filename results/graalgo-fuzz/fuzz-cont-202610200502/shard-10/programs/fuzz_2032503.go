package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int; F2 uint32 }

func main() {
    v1 := []string{}
    var v2 uint16 = 0
    v3 := make(map[int]int)
    var v4 int64 = 9223372036854775807
    v4 = v4
    var v5 uint64 = 18446744073709551615
    v6 := GenS1{F0: 204, F1: 255}
    fmt.Println(v5, v5, v5)
    var v7 uint64 = 0
    switch -9223372036854775808 {
    case 4294967295:
        v5 -= (^(v5))
    case 921:
        _ = len(v3)
    case 32767:
        v8 := GenS2{F0: 955, F1: v5}
        _ = v8
    default:
        v4 -= -9223372036854775808
    }
    v9 := GenS2{F0: 2, F1: 2}
    fmt.Println(v5)
    v10 := GenS3{F0: 107, F1: 32767, F2: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
