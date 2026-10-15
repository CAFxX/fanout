package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    v3[0] = 65535
    delete(v5, 2)
    v1 = v1
    fmt.Println(v1)
    fmt.Println(v1)
    v1 -= v1
    v6 := GenS1{F0: 989}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
