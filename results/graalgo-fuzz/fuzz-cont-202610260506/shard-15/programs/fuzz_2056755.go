package main

import (
    "fmt"
)


func main() {
    var v1 int = 9223372036854775807
    var v2 int = v1
    var v3 bool = true
    var v4 uint16 = 587
    v5 := make(map[uint32]bool)
    v6 := make(map[int]int)
    var v7 string = ("a" + "")
    var v8 uint8 = 240
    var v9 int = v2
    v10 := []int{127}
    v4 = v4
    var v11 int32 = 42
    _ = len(v10)
    var v12 string = v7
    fmt.Println(v9, v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
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
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
