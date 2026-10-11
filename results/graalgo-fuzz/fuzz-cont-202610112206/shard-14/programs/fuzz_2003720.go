package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]bool)
    var v3 float64 = (0.0 - 0.0)
    var v4 float64 = 1.5
    v5 := []string{"", ""}
    var v6 uint64 = 920
    fmt.Println(v4, v4)
    v6 -= 1
    fmt.Println(v4, v3)
    var v7 int64 = 42
    v1 = v1
    _ = len(v5)
    v6 = (v6 & v6)
    for fi8 := 0; fi8 < 12; fi8++ {
        v7 = (-(v7))
        _ = fi8
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
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
