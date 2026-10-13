package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }

func main() {
    var v1 int = 1
    var v2 float64 = 1e10
    v3 := []int{128, v1}
    var v4 string = "hello"
    var v5 float32 = (float32(v2) * 0.0)
    fmt.Println(v2, v1, v1)
    var v6 int16 = -56
    v4 = v4
    v4 = v4
    var v7 string = "\n\t"
    var v8 uint64 = 418
    var v9 string = ""
    v10 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: 10}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
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
