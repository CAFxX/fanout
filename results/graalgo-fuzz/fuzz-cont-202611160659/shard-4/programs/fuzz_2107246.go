package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]bool)
    var v2 int = 128
    v3 := []uint32{57}
    v4 := make(map[uint32]string)
    var v5 uint = 883
    var v6 bool = (255 < 2)
    var v7 uint8 = 38
    var v8 int64 = 9223372036854775807
    fmt.Println(v6, v7, v5)
    var v9 int = v2
    fmt.Println(v2, v2)
    v3 = append(v3, 32)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
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
