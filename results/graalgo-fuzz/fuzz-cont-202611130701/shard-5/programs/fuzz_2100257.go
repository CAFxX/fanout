package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    v3 := make(map[int]int)
    var v4 uint32 = 2
    var v5 float32 = float32(0.1)
    var v6 uint = 2
    var v7 bool = (-169 != 1)
    _ = len(v2)
    var v8 int8 = -128
    var v9 bool = (1 == 2)
    v5 = v5
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v10 uint = v6
    if true {
        fmt.Println(v5, v6, v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
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
