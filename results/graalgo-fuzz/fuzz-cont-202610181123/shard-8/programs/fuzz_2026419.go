package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int32 = -1
    var v2 float32 = 0.1
    var v3 int32 = v1
    v4 := []int{32768, 2147483647, -1}
    v1 = v1
    v2 = float32((float64(v2) + 3.14159))
    v3 = 186
    v5 := GenS1{F0: 0}
    if true {
        var v6 uint8 = 0
        _ = v6
    }
    fmt.Println(v2)
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v7 := GenS1{F0: 164}
    v3 -= v3
    var v8 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
