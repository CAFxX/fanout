package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }

func main() {
    var v1 float32 = 1e10
    var v2 int32 = 0
    v3 := make(map[int]uint32)
    var v4 int = 127
    v5 := make(map[int]int)
    v6 := GenS1{F0: -32768, F1: 1}
    var v7 float64 = 3.14159
    for fi8 := 0; fi8 < 18; fi8++ {
        v2 = v2
        v9 := []int{}
        v10 := GenS1{F0: -32768, F1: 0}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v4, v1)
    v4 -= v4
    for fi11 := 0; fi11 < 8; fi11++ {
        v7 = v7
        _ = fi11
    }
    fmt.Println(v2, v2)
    var v12 float64 = -2.25
    switch "hello" {
    case "a":
        var v13 int8 = -112
        _ = v13
    default:
        fmt.Println(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
