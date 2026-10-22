package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 2
    var v2 uint = 42
    v3 := make(map[string]uint32)
    var v4 int = 42
    _ = len(v3)
    v1 += v1
    v5 := []int{}
    v2 = v2
    var v6 int64 = 42
    var v7 int8 = 29
    if (0 <= 42) {
        if len(v5) > 0 {
            v5[0] = 42
        }
    }
    v8 := v3["foo bar"]
    fmt.Println(v6)
    v4 = (^(v4))
    v4 -= 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
