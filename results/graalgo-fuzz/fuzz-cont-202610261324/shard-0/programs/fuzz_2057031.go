package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{4294967295, 952}
    var v3 int16 = -1
    var v4 uint64 = 338
    var v5 float64 = (3.14159 - 1.5)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int32 = 517
        _ = fi6
        _ = v7
    }
    v8 := v1[127]
    v9 := GenS1{F0: 1, F1: 1}
    fmt.Println(v8, v5, v4)
    delete(v1, 32768)
    v10 := GenS1{F0: 0, F1: -712}
    for fi11 := 0; fi11 < 5; fi11++ {
        v2 = append(v2, 127)
        _ = fi11
    }
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
