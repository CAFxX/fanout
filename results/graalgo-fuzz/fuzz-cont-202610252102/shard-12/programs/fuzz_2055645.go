package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
func genFn0() uint {
    var v4 float64 = -2.25
    var v5 int = -622
    var v6 float64 = v4
    _ = v4
    _ = v5
    _ = v6
    return 659
}

func main() {
    var v1 int8 = 100
    v2 := []int{}
    var v3 int64 = 1
    var v7 float32 = 1.5
    var v8 string = "hello"
    var v9 int = -265
    fmt.Println(v1)
    _ = len(v2)
    v10 := GenS1{F0: 1, F1: 905, F2: 2}
    var v11 int8 = v1
    var v12 float64 = (1.5 * 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
