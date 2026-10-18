package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int16 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 int32 = 0
    var v5 string = ""
    v6 := []int{42, 1, 127}
    var v7 string = (v5 + "hello")
    var v8 uint32 = 19
    v9 := GenS1{F0: uint16(v1), F1: (v8 ^ (^(v8))), F2: 1}
    var v10 int16 = int16(v4)
    v5 = v7
    fmt.Println(v7, v1, v4)
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
