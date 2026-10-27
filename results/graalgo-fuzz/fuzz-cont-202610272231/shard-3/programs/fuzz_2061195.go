package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{1, -1}
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 int64 = 9223372036854775807
    v6 := GenS1{F0: 1, F1: 985}
    fmt.Println(v5)
    var v7 int = 66
    v5 += 42
    v8 := GenS1{F0: 40, F1: 0}
    fmt.Println(v7, v7)
    var v9 uint32 = 2147483647
    var v10 string = ("hello" + ("foo bar" + "\n\t"))
    var v11 string = v10
    var v12 uint = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
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
