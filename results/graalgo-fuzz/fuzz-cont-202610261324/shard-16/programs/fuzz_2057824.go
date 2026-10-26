package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 1
    var v3 uint16 = 0
    v4 := make(map[string]uint32)
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]uint32)
    var v7 string = "hello"
    var v8 uint64 = uint64(v2)
    v3 = v3
    v9 := GenS1{F0: 910, F1: (v8 << ((v2) & 63))}
    v10 := []int{1}
    var v11 string = v7
    if (100 > 0) {
        v12 := v5[2147483647]
        _ = v12
    }
    fmt.Println(v7, v7, v11)
    if true {
        var v13 int16 = 1000
        for fi14 := 0; fi14 < 8; fi14++ {
            switch (-(fi14)) {
            case -788:
                var v15 int16 = 1000
                _ = v15
            default:
                v8 = (v8 - v8)
            }
            v16 := GenS2{F0: (v8 | 501), F1: fi14, F2: v2}
            _ = fi14
            _ = v16
        }
        var v17 bool = (42 != 483)
        _ = v13
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
