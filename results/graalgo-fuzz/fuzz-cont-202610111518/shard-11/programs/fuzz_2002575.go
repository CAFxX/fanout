package main

import (
    "fmt"
)


func main() {
    var v1 int = 9223372036854775807
    v2 := []uint32{2147483647, 340}
    v3 := make(map[uint32]uint32)
    var v4 int32 = 0
    v4 -= v4
    var v5 float32 = 0.0
    v5 -= (v5 - 3.14159)
    var v6 uint8 = 127
    var v7 uint = 255
    v6 += v6
    var v8 uint8 = v6
    fmt.Println(v1, v6, v1)
    var v9 uint8 = (^(v6))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
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
