package main

import (
    "fmt"
)


func main() {
    var v1 string = "hello"
    v2 := []int{-1, 2147483647, 127}
    v3 := make(map[string]bool)
    var v4 string = ("" + "a")
    _ = len(v2)
    var v5 uint64 = 4294967295
    var v6 string = v1
    fmt.Println(v6, v6, v6)
    var v7 int = 9223372036854775807
    var v8 float32 = 0.0
    v4 = (v1 + v6)
    v4 = "hello"
    v8 += 3.14159
    var v9 float64 = 1e10
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
