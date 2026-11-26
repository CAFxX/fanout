package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 42
    var v3 int8 = -67
    var v4 int8 = v3
    v5 := []int{-335, 4294967295, -1, 854}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v3, v3, v2)
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    _ = len(v5)
    v1 -= (v1 * v1)
    v2 = v2
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 9223372036854775807}
        _ = fi6
        _ = v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
