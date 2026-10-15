package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 883
    var v3 string = "a"
    var v4 int8 = 42
    fmt.Println(v4)
    var v5 float64 = 0.1
    var v6 uint64 = 2
    _ = len(v1)
    var v7 bool = ((v2 % ((v2) | 1)) >= v2)
    var v8 uint32 = uint32(v6)
    v9 := GenS1{F0: 2}
    v6 = 360
    fmt.Println(v4)
    v5 += float64(1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
    _ = v9
}
