package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 1
    var v2 int8 = v1
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 int16 = 32767
    var v6 int8 = (v1 / (((v1 & v1)) | 1))
    var v7 bool = (376 == 0)
    v6 = v6
    v1 = (v6 ^ 1)
    v2 -= v6
    v6 = (-(v6))
    fmt.Println(v6, v6, v7)
    v8 := []int{2147483647, -119, 127}
    var v9 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
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
