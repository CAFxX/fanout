package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    v4 := []uint32{}
    var v5 uint = 98
    v6 := GenS1{F0: 42, F1: 2}
    if len(v4) > 0 {
        v4[0] = 65535
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v5 = 42
            _ = fi8
        }
        var v9 bool = false
        _ = fi7
        _ = v9
    }
    var v10 int = 595
    v1 -= v1
    switch v10 {
    case -465:
        if true {
            var v11 float32 = 1e10
            var v12 float32 = 1e10
            _ = v11
            _ = v12
        }
    case 42:
        v13 := GenS1{F0: 1, F1: 128}
        _ = v13
    case 1:
        fmt.Println(v1, v1, v5)
    }
    v14 := GenS1{F0: -1, F1: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
