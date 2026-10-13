package main

import (
    "fmt"
)


func main() {
    var v1 int8 = -38
    v2 := make(map[uint32]bool)
    var v3 int32 = 828
    var v4 int32 = v3
    v1 = (v1 ^ 57)
    v1 = (v1 * v1)
    var v5 uint32 = 117
    var v6 uint32 = v5
    v1 -= (v1 + v1)
    v6 = v5
    v1 = 1
    var v7 int = 128
    var v8 int = v7
    var v9 uint8 = 128
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int64 = -17
        fmt.Println(v3)
        _ = fi10
        _ = v11
    }
    var v12 int64 = 597
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
