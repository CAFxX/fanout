package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int64 = 1
    v2 := []int{4294967295, 4294967295, 4294967295, -665}
    var v3 int8 = -123
    var v4 int32 = 0
    if (false || (!(0 > 42))) {
        fmt.Println(v1)
        var v5 uint8 = 2
        _ = v5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: -1, F1: 255}
        var v8 uint16 = 357
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 474}
    fmt.Println(v4, v1)
    _ = len(v2)
    var v10 uint64 = 18446744073709551615
    var v11 uint8 = 255
    var v12 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
