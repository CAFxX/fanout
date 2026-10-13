package main

import (
    "fmt"
)


func main() {
    var v1 int64 = -213
    var v2 int64 = v1
    var v3 uint32 = 65535
    v4 := make(map[uint32]string)
    var v5 uint32 = (v3 / (((v3 | v3)) | 1))
    var v6 uint8 = 255
    var v7 bool = ((^(v3)) != (v5 >> ((42) & 63)))
    var v8 uint8 = v6
    v7 = false
    v3 -= (^(v3))
    _ = len(v4)
    v9 := []int{}
    v2 = -457
    v3 += (^(v3))
    v8 += 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
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
