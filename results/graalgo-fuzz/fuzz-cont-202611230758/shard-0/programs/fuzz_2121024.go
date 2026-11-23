package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }

func main() {
    var v1 int16 = 816
    v2 := []int{52, -898, 1, 32768}
    var v3 uint64 = 165
    var v4 float64 = ((((3.14159 + 0.0) - -2.25) / 1.5) / 1.0)
    if (!true) {
        var v5 uint64 = v3
        _ = v5
    }
    v2 = append(v2, 32767)
    fmt.Println(v1, v3)
    v6 := GenS1{F0: 0, F1: 255, F2: 127}
    var v7 uint = 0
    var v8 int32 = 1
    var v9 float32 = 0.0
    v10 := GenS1{F0: 70, F1: 127, F2: 208}
    var v11 float64 = v4
    var v12 int8 = -108
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
