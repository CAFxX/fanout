package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 1000
    v3 := make(map[int]bool)
    var v4 bool = true
    v2 = 1
    var v5 int64 = 116
    v5 = 1
    v6 := []int{0, 957, 42, -759}
    var v7 uint = 0
    delete(v3, 0)
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 bool = (0 <= 0)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi8
        _ = v9
    }
    v5 += (v5 ^ v5)
    var v10 uint64 = 1
    switch "foo bar" {
    case "foo bar":
        var v11 uint32 = 279
        _ = v11
    default:
        var v12 uint8 = 1
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
