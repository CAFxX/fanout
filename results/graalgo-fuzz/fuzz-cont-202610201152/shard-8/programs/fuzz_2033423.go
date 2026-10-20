package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16; F1 int32 }

func main() {
    var v1 int32 = 704
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]uint32)
    var v5 float64 = 3.14159
    v6 := []string{}
    v1 = (v1 | 672)
    fmt.Println(v1)
    var v7 int32 = (^(v1))
    if true {
        v8 := GenS1{F0: 2147483648, F1: 314, F2: v7}
        var v9 uint = 2
        fmt.Println(v5)
        _ = v8
        _ = v9
    } else {
        var v10 float64 = (v5 / v5)
        v11 := GenS1{F0: 1, F1: 0, F2: v7}
        v5 = float64(1.5)
        _ = v10
        _ = v11
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS2{F0: 1, F1: (v7 | (v7 & -1))}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        delete(v3, fi14)
        _ = fi14
    }
    v7 = (v7 | (v1 * v7))
    if (18446744073709551615 != 0) {
        fmt.Println(v7)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
