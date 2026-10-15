package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 int16 = 32767
    var v4 uint = 255
    var v5 uint16 = 42
    v6 := make(map[int]int)
    v3 += v3
    v3 = 32767
    if false {
        fmt.Println(v5)
    }
    v3 -= v3
    var v7 uint8 = 2
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: int32(v7), F1: v7}
        _ = fi8
        _ = v9
    }
    v5 += (v5 & 42)
    switch "" {
    case "a":
        _ = len(v6)
    default:
        var v10 int16 = 32767
        _ = v10
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
