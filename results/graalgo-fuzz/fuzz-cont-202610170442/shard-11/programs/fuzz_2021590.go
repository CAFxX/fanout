package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }

func main() {
    var v1 uint16 = 836
    v2 := make(map[string]bool)
    var v3 int64 = 0
    v4 := make(map[uint32]string)
    var v5 uint16 = (^(v1))
    v6 := make(map[string]uint32)
    v7 := []int{128}
    v8 := v6["hello"]
    var v9 bool = true
    var v10 float32 = ((-2.25 * ((float32(0.1) - 3.14159) - (3.14159 / (float32(1.5) / 0.1)))) + 1.5)
    var v11 uint = 4294967295
    v12 := GenS1{F0: (v8 * (v8 + v8)), F1: -2147483648}
    var v13 int64 = (v3 / (((v3 >> ((v11) & 63))) | 1))
    v3 = (^(v13))
    v9 = (v5 <= v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
