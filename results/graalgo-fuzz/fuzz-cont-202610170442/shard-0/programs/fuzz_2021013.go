package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    var v1 uint = 255
    v2 := make(map[string]bool)
    var v3 uint32 = 229
    v4 := make(map[int]uint32)
    var v5 int = -457
    var v6 uint16 = 42
    for fi7 := 0; fi7 < 17; fi7++ {
        v5 -= fi7
        _ = fi7
    }
    v2["a"] = true
    v8 := GenS1{F0: -1, F1: "a"}
    var v9 int8 = -128
    v10 := []int{(^(v5)), (v5 | (v5 % ((v5) | 1)))}
    v10 = append(v10, (v5 / ((v5) | 1)))
    var v11 int16 = -943
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
