package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }

func main() {
    var v1 int64 = 749
    var v2 int8 = 127
    var v3 int64 = -1
    var v4 int32 = 2147483647
    v5 := make(map[uint32]bool)
    var v6 bool = ((v1 + (v3 * v1)) > v3)
    var v7 int32 = -30
    var v8 int32 = (v7 >> ((42) & 63))
    v5[62] = true
    v6 = v6
    v7 += v7
    fmt.Println(v3, v1)
    v9 := GenS1{F0: v1, F1: 374}
    if (-1 > 128) {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 uint32 = 0
            var v12 int64 = v3
            _ = fi10
            _ = v11
            _ = v12
        }
        delete(v5, 42)
        _ = len(v5)
    }
    var v13 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
