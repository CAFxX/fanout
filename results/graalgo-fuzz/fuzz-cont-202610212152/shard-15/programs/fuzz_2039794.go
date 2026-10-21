package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int32, p1 uint32) uint16 {
    fmt.Println(p1)
    v5 := []int{}
    var v6 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 944
}

func main() {
    var v1 int8 = 79
    var v2 float32 = 0.1
    var v3 bool = ((1000 != -962) || true)
    var v4 int16 = -5
    fmt.Println(v3, v1, v1)
    v7 := []int{26}
    v4 = (v4 + v4)
    v3 = (3.14159 != v2)
    var v8 uint8 = 93
    fmt.Println(v3, v3)
    _ = len(v7)
    var v9 uint64 = 0
    v10 := GenS1{F0: v1}
    var v11 bool = (1 <= -107)
    var v12 uint64 = (v9 & v9)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
