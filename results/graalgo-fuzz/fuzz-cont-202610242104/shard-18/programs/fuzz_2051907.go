package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 1e10
    var v3 int = 4294967295
    v4 := []uint32{}
    v3 += (v3 - (v3 & v3))
    v3 += -980
    var v5 int8 = 1
    v2 -= (float64(((0.1 * (0.1 * 1.5)) + 3.14159)) * v2)
    fmt.Println(v5, v2)
    fmt.Println(v3, v2)
    var v6 int16 = 66
    fmt.Println(v3)
    v4 = append(v4, 4294967295)
    var v7 int16 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
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
