package main

import (
    "fmt"
)


func main() {
    var v1 float32 = float32((-2.25 - 3.14159))
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 int32 = 1
    var v5 int64 = 42
    var v6 uint32 = 1
    v4 = int32(v5)
    v2[-325] = ("" + ("" + "\n\t"))
    v7 := []int{2147483647, 32767}
    var v8 int64 = v5
    fmt.Println(v6, v5)
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    v1 = -2.25
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
