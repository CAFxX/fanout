package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    v3 := []int{}
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    v6 := make(map[int]uint32)
    switch -1 {
    case -12:
        var v7 int16 = 0
        _ = v7
    case -545:
        var v8 uint = 42
        _ = v8
    }
    var v9 uint16 = 2
    var v10 string = "\n\t"
    _ = len(v3)
    v9 += v9
    _ = len(v2)
    v11 := GenS1{F0: 109}
    _ = len(v4)
    var v12 int = 128
    v13 := GenS1{F0: v9}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
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
    _ = v9
}
