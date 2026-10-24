package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint16 = 0
    var v2 uint16 = 65535
    v3 := []int{}
    v4 := GenS1{F0: 42}
    fmt.Println(v2)
    var v5 float64 = 0.0
    v5 -= (v5 * -2.25)
    if false {
        var v6 uint64 = 545
        fmt.Println(v2, v6, v2)
        _ = v6
    }
    var v7 int = 1
    v8 := GenS1{F0: 0}
    if true {
        v5 += v5
        v1 = (v2 / ((v1) | 1))
    } else {
        v3 = append(v3, v7)
        var v9 float32 = (3.14159 - float32(0.1))
        _ = v9
    }
    v10 := GenS1{F0: -1}
    var v11 uint32 = 2147483647
    var v12 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
