package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
func genFn0(p0 int16, p1 int8) int32 {
    v7 := GenS1{F0: 0, F1: 32767}
    p1 -= -1
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = -32768
    v3 := make(map[string]string)
    var v4 int8 = 42
    v5 := []int{1, 789, 127, 490}
    v6 := make(map[int]bool)
    var v8 uint64 = 1
    v8 = 42
    fmt.Println(v8)
    var v9 uint = 10
    fmt.Println(v9, v9)
    var v10 float32 = 1e10
    v11 := GenS1{F0: -1, F1: 32767}
    var v12 uint32 = 4294967295
    v13 := GenS1{F0: 32767, F1: 100}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
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
    _ = v8
    _ = v9
}
