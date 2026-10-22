package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint8 }

func main() {
    var v1 string = "hello"
    var v2 int = -970
    var v3 int8 = 100
    v3 = (v3 ^ v3)
    switch (v1 + v1) {
    case "a":
        var v4 float64 = (1e10 + (1.5 / 1e10))
        _ = v4
    case "hello":
        v5 := GenS1{F0: 49, F1: -1}
        _ = v5
    default:
        v6 := []int{v2, (v2 >> ((42) & 63)), v2}
        _ = v6
    }
    var v7 uint64 = 4294967295
    var v8 float32 = (0.0 * 0.0)
    v9 := GenS2{F0: 634}
    v10 := GenS3{F0: 18446744073709551615, F1: 2}
    v8 -= float32(0.1)
    var v11 uint32 = 1
    v12 := []int{(v2 + 42)}
    v12 = append(v12, v2)
    for fi13 := 0; fi13 < 6; fi13++ {
        fmt.Println(v1)
        _ = fi13
    }
    v14 := GenS1{F0: v3, F1: -195}
    v15 := GenS1{F0: v3, F1: -2147483648}
    fmt.Println(v8, v3, v7)
    fmt.Println(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
