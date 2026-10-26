package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int }

func main() {
    var v1 int16 = -999
    var v2 int8 = 62
    var v3 uint = 42
    v4 := GenS1{F0: 1000, F1: v3, F2: 127}
    fmt.Println(v2)
    var v5 int16 = (v1 & (^(v1)))
    v6 := []int{32768, 32767, 9223372036854775807}
    var v7 uint8 = 1
    var v8 int16 = (v1 | (v1 & v1))
    fmt.Println(v2, v1)
    fmt.Println(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
