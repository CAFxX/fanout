package main

import (
    "fmt"
)


func main() {
    v1 := []int{9223372036854775807, 2147483647, -9223372036854775808}
    var v2 uint16 = 1
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    v2 += v2
    v2 += v2
    var v5 uint8 = 1
    v5 -= v5
    v2 = 670
    v6 := v4[128]
    _ = len(v4)
    fmt.Println(v5)
    v2 = 1
    v6 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
