package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16; F1 uint }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]int)
    var v3 uint8 = 1
    var v4 int64 = 703
    var v5 uint8 = (^(v3))
    v6 := GenS1{F0: 1, F1: -128}
    v7 := []int{32767, 1}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 uint8 = (^(v3))
    v9 := GenS1{F0: 1, F1: -6}
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int = fi10
        var v12 uint = 42
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint16 = uint16(v8)
    v14 := GenS2{F0: -598, F1: 348}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
