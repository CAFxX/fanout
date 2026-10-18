package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int16 = 0
    var v2 uint8 = 42
    var v3 uint64 = 918
    var v4 float64 = float64(1e10)
    v5 := make(map[string]string)
    v4 -= (v4 + 0.0)
    v6 := GenS1{F0: 234, F1: v1}
    var v7 uint8 = (v2 ^ 128)
    switch "a" {
    case "":
        var v8 int16 = (v1 | v1)
        _ = v8
    default:
        v1 = 32767
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(v4, v1)
        _ = fi9
    }
    v2 = v7
    v10 := GenS2{F0: v3}
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int = (^(fi11))
        fmt.Println(fi11)
        _ = fi11
        _ = v12
    }
    v13 := []int{32767, 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
