package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }

func main() {
    v1 := []int{42}
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    var v4 int16 = 1
    v5 := make(map[string]uint32)
    var v6 int16 = -275
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3["hello"] = "\n\t"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 = -1
    v7 := GenS1{F0: "\n\t", F1: 1}
    fmt.Println(v6)
    v4 = (v6 / ((v4) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
