package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 string }

func main() {
    v1 := []uint32{1, 2, 42, 2147483648}
    v2 := make(map[int]int)
    var v3 float32 = float32(-2.25)
    v4 := []int{1, -437, 510}
    var v5 int8 = -128
    var v6 uint8 = 1
    v5 += 100
    delete(v2, 4294967295)
    var v7 uint64 = 4294967295
    fmt.Println(v7)
    var v8 uint32 = 1
    var v9 int = int(v7)
    v7 = v7
    v10 := GenS1{F0: "", F1: true, F2: ("\n\t" + "hello")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
