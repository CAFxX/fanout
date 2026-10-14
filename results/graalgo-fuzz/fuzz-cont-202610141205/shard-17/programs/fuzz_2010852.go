package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int = 4294967295
    v2 := make(map[uint32]bool)
    var v3 uint64 = 42
    var v4 uint8 = 42
    var v5 uint32 = 523
    v6 := GenS1{F0: 255}
    v3 -= (v3 & (^(v3)))
    var v7 int16 = 32767
    v1 = v1
    var v8 uint64 = v3
    v9 := GenS1{F0: 934}
    v7 = (v7 | v7)
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int8 = 42
        _ = fi10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
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
    _ = v9
}
