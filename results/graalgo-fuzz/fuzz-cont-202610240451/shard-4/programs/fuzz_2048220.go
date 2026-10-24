package main

import (
    "fmt"
)


func main() {
    var v1 int64 = -516
    var v2 uint8 = 23
    v3 := make(map[uint32]bool)
    var v4 int16 = 42
    fmt.Println(v4, v1)
    v2 += 200
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int8 = 1
        var v7 bool = true
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint8 = v2
    fmt.Println(v8, v8)
    v8 += (v2 & (v2 - (v8 / ((v2) | 1))))
    var v9 float32 = 0.0
    fmt.Println(v4, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
