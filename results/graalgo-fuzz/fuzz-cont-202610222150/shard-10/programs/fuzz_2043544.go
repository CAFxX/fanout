package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }

func main() {
    v1 := []string{"\n\t"}
    var v2 int8 = 100
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    var v5 int64 = 0
    var v6 int64 = v5
    fmt.Println(v2, v2)
    var v7 int32 = 748
    v8 := GenS1{F0: 209, F1: 2, F2: (false || false)}
    v9 := GenS1{F0: 1, F1: 4294967295, F2: true}
    var v10 uint = 65535
    var v11 float32 = (0.1 + float32((3.14159 / 0.1)))
    v12 := GenS1{F0: 4294967295, F1: 42, F2: false}
    v6 = 0
    v1 = append(v1, "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
