package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 float32 = 1.5
    var v2 uint32 = 2
    v3 := make(map[int]bool)
    v4 := []int{}
    v2 -= v2
    v1 = v1
    var v5 uint32 = 42
    if (!true) {
        v5 -= v2
    } else {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 uint = 235
            _ = fi6
            _ = v7
        }
        v2 += v5
        var v8 int = 128
        _ = v8
    }
    if true {
        v9 := GenS1{F0: 65535}
        _ = v9
    }
    var v10 float64 = -2.25
    var v11 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
