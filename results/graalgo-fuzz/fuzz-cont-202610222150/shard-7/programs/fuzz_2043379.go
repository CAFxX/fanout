package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 2147483647
    var v2 uint32 = 2147483647
    var v3 uint = 18446744073709551615
    v4 := make(map[string]string)
    v5 := make(map[string]int)
    v6 := make(map[int]bool)
    if true {
        _ = len(v5)
        v3 -= (v3 & (v3 & v3))
        v7 := []int{-996}
        _ = v7
    } else {
        var v8 int = -9223372036854775808
        v2 = v1
        v9 := []int{127, 32767, v8, v8}
        _ = v8
        _ = v9
    }
    var v10 uint8 = 2
    v6[1] = true
    fmt.Println(v2)
    v10 -= (v10 | v10)
    var v11 uint32 = 2147483647
    v11 += v2
    v2 = (v11 * v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
