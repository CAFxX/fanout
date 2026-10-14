package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 255
    var v2 float64 = -2.25
    var v3 float32 = (3.14159 / 1.0)
    v4 := make(map[uint32]string)
    var v5 uint16 = v1
    fmt.Println(v5, v5)
    v2 -= (float64(float32(v2)) * ((v2 * 1.5) + v2))
    var v6 string = "hello"
    v7 := []int{717, 0, 9223372036854775807}
    var v8 int32 = 0
    fmt.Println(v2)
    v3 -= (v3 + 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
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
