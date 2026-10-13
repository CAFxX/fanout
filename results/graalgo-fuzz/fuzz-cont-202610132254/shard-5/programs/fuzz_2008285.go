package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []int{42, 371, 2147483647, -9223372036854775808}
    v2 := make(map[uint32]string)
    var v3 uint8 = 2
    v4 := make(map[int]int)
    v5 := GenS1{F0: false}
    v6 := GenS1{F0: (!(1e10 > 0.1))}
    fmt.Println(v3)
    v7 := GenS1{F0: true}
    if len(v1) > 0 {
        v1[0] = 128
    }
    v3 = 0
    fmt.Println(v3)
    fmt.Println(v3)
    var v8 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
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
