package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{2147483647, 847}
    var v2 int32 = -848
    v3 := make(map[string]uint32)
    fmt.Println(v2)
    var v4 int64 = 0
    _ = len(v1)
    v2 = (v2 / ((v2) | 1))
    _ = len(v3)
    v2 += v2
    fmt.Println(v2, v4)
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int64 = (-(v4))
        fmt.Println(v2, v2, v4)
        _ = fi5
        _ = v6
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
