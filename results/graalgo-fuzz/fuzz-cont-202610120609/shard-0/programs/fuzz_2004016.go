package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := []uint32{2147483648, 1, 2147483647, 42}
    var v4 uint16 = 2
    var v5 int8 = 127
    v4 = v4
    v6 := GenS1{F0: 736, F1: true}
    v4 = v4
    fmt.Println(v4)
    v4 += v4
    fmt.Println(v5, v5)
    var v7 int16 = -128
    fmt.Println(v5, v4)
    switch "" {
    case "a":
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint16 = v4
            var v10 uint8 = 180
            var v11 uint16 = 928
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(v4, fi12, v5)
        var v13 int32 = -2147483648
        v5 = 42
        _ = fi12
        _ = v13
    }
    v4 += (v4 * 32767)
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 int8 = int8(v4)
        v16 := v2[fi14]
        _ = fi14
        _ = v15
        _ = v16
    }
    v5 += 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
