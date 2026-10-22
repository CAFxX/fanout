package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 bool = (false || true)
    var v4 uint = 0
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(v3, v3, fi5)
        _ = fi5
    }
    v4 = v4
    var v6 int64 = -1
    var v7 int64 = (v6 * v6)
    v8 := []int{128, 2147483647, 2147483647}
    var v9 uint32 = 931
    _ = len(v2)
    v6 += (v6 ^ (v7 ^ 673))
    var v10 uint = 2
    var v11 uint8 = 34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
