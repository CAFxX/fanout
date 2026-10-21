package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1.5
    var v3 int8 = -1
    var v4 string = "hello"
    v5 := GenS1{F0: 1, F1: ((v4 + "foo bar") + (v4 + v4)), F2: 4294967295}
    var v6 uint64 = 553
    var v7 int16 = 1000
    fmt.Println(v4, v2, v3)
    var v8 uint16 = 2
    fmt.Println(v3)
    var v9 uint64 = v6
    v10 := []int{1, 32768, 23}
    fmt.Println(v2)
    var v11 uint = 458
    var v12 bool = (2 <= 869)
    v1[2147483648] = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
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
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
