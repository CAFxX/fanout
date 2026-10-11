package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int64 }

func main() {
    var v1 float32 = ((0.1 + (((0.0 - 0.1) * 1.5) + (3.14159 - ((0.0 + (1.5 / 3.14159)) / (1e10 + ((3.14159 * (0.1 - 3.14159)) / 3.14159)))))) - -2.25)
    var v2 uint16 = 0
    var v3 int32 = -1
    fmt.Println(v1, v1)
    fmt.Println(v2, v1, v2)
    v4 := GenS1{F0: 2, F1: -570, F2: 0}
    v5 := []int{-1, -2147483648, 32768}
    var v6 uint32 = 42
    v6 = (v6 ^ v6)
    var v7 int8 = -111
    fmt.Println(v6)
    v8 := GenS1{F0: 777, F1: 9223372036854775807, F2: 193}
    v7 -= v7
    fmt.Println(v3, v3)
    v9 := GenS1{F0: 717, F1: -1, F2: -650}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
