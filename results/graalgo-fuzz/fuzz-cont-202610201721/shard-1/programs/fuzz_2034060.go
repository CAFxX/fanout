package main

import (
    "fmt"
)


func main() {
    v1 := []string{}
    var v2 int8 = -1
    v3 := make(map[int]int)
    v4 := make(map[uint32]bool)
    var v5 bool = true
    v6 := v4[628]
    fmt.Println(v5, v2, v5)
    v1 = append(v1, "")
    fmt.Println(v2, v5)
    v2 = -31
    var v7 int32 = 0
    v5 = (((v6 && (1e10 <= -2.25)) || v6) && (2 == 2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
