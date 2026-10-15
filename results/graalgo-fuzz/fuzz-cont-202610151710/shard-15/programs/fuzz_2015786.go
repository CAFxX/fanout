package main

import (
    "fmt"
)


func main() {
    var v1 float64 = float64(3.14159)
    v2 := []int{-945}
    var v3 int16 = -688
    v4 := make(map[string]uint32)
    v5 := []uint32{65535, 154}
    v6 := make(map[int]bool)
    v6[32768] = (1.5 != (v1 / v1))
    v3 = v3
    v3 += 42
    v3 = (v3 + 1)
    v4["a"] = 417
    v3 -= 0
    fmt.Println(v1)
    v6[32767] = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
