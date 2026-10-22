package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 32767
    v2 := []int{0, -1}
    v3 := make(map[string]string)
    var v4 int8 = -128
    var v5 uint16 = (v1 + v1)
    fmt.Println(v1, v1, v1)
    var v6 int64 = 1
    v1 = 32767
    v1 = (v5 ^ (v1 ^ v1))
    var v7 int32 = 42
    v1 = (v1 * v5)
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint8 = 127
        _ = fi8
        _ = v9
    }
    var v10 int64 = v6
    v6 -= (^(v10))
    var v11 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
