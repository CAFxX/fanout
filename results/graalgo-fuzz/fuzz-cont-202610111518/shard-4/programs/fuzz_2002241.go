package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 100
    var v2 float32 = 3.14159
    v3 := make(map[string]int)
    v4 := []int{9223372036854775807, 4294967295}
    var v5 int64 = -1
    _ = len(v3)
    var v6 float64 = (-2.25 + 0.0)
    var v7 int8 = v1
    delete(v3, "")
    var v8 uint16 = 0
    if len(v4) > 0 {
        v4[0] = 1
    }
    var v9 int64 = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
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
