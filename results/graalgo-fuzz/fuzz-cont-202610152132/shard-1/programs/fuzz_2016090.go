package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }

func main() {
    v1 := []int{32768, 128}
    var v2 int16 = -1
    v3 := []string{"\n\t", "\n\t"}
    var v4 uint8 = 128
    var v5 float32 = float32(float64(0.0))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = append(v1, 897)
    switch "" {
    case "":
        v6 := GenS1{F0: -128, F1: 0, F2: 1000}
        _ = v6
    }
    var v7 string = "\n\t"
    var v8 uint8 = (v4 & 0)
    fmt.Println(v4, v5, v7)
    var v9 uint32 = 0
    v10 := GenS1{F0: int8(v8), F1: 588, F2: v2}
    var v11 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
