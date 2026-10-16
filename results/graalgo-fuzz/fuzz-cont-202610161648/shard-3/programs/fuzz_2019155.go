package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int64 = -213
    v2 := make(map[string]string)
    var v3 int8 = 1
    fmt.Println(v1)
    v4 := []int{950, 127}
    _ = len(v4)
    var v5 int64 = (-(v1))
    var v6 uint8 = 128
    v7 := GenS1{F0: -565}
    if true {
        fmt.Println(v1, v1, v6)
    }
    v8 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
