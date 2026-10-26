package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint16; F2 int8 }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    var v3 int = 4294967295
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 string = "a"
        _ = len(v1)
        v5 = v5
        _ = fi4
        _ = v5
    }
    v3 -= 128
    v3 += 0
    var v6 uint64 = 0
    var v7 float32 = (0.0 / 1e10)
    _ = len(v2)
    if false {
        for fi8 := 0; fi8 < 16; fi8++ {
            v6 = v6
            fmt.Println(v3, v6)
            _ = fi8
        }
    }
    v9 := GenS1{F0: 42}
    v10 := GenS2{F0: -179, F1: 938, F2: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 int8 = -128
    if false {
        v11 = -1
    }
    var v12 int64 = 1
    var v13 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
