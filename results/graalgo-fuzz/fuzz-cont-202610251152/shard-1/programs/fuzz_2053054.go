package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 172
    v2 := make(map[int]bool)
    v3 := []string{"a"}
    var v4 uint16 = 42
    var v5 int16 = -1
    var v6 float32 = 1.5
    v3 = append(v3, "")
    var v7 int16 = v1
    v6 += v6
    var v8 bool = (2 > 273)
    fmt.Println(v1, v4)
    fmt.Println(v6, v4)
    fmt.Println(v4, v8)
    var v9 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
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
