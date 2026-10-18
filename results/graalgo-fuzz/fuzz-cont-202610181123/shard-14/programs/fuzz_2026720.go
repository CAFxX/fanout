package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    v2 := []string{"", "", "hello"}
    var v3 uint32 = 1
    var v4 int32 = 145
    var v5 int32 = v4
    var v6 uint64 = 42
    var v7 int64 = -1
    var v8 int = 42
    v8 = (v8 << ((0) & 63))
    _ = len(v1)
    v6 = (^(v6))
    fmt.Println(v4, v4, v5)
    var v9 uint32 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
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
