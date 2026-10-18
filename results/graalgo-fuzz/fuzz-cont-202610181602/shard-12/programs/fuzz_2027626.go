package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := []string{""}
    var v4 uint64 = 0
    v5 := make(map[string]int)
    v6 := make(map[string]uint32)
    var v7 uint16 = 514
    var v8 int64 = 0
    var v9 int16 = -1
    var v10 uint8 = 1
    v10 -= (^(v10))
    v11 := GenS1{F0: 1}
    v10 = uint8(v4)
    fmt.Println(v10)
    v12 := GenS1{F0: 0}
    var v13 uint16 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
