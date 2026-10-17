package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint64 = 309
    var v2 uint16 = 65535
    var v3 uint8 = 1
    var v4 bool = ((563 == 42) && false)
    fmt.Println(v4)
    var v5 uint64 = v1
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 float32 = 3.14159
        v8 := GenS1{F0: 1}
        fmt.Println(v2)
        _ = fi6
        _ = v7
        _ = v8
    }
    v5 += v1
    v9 := []int{128, 1, -713, 4294967295}
    v2 = v2
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
