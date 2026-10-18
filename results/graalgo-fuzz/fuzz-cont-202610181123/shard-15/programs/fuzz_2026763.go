package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }

func main() {
    var v1 int8 = -56
    v2 := []uint32{}
    var v3 int64 = -122
    v4 := make(map[string]bool)
    fmt.Println(v3)
    v3 -= v3
    _ = len(v2)
    switch -1 {
    case 1:
        var v5 uint32 = 42
        _ = v5
    case 2147483647:
        var v6 uint32 = 394
        _ = v6
    }
    var v7 uint64 = 42
    if false {
        if true {
            var v8 int = 32768
            v1 = (^(v1))
            _ = v8
        }
    }
    fmt.Println(v7)
    v3 += v3
    for fi9 := 0; fi9 < 17; fi9++ {
        v3 = int64(v1)
        v10 := GenS1{F0: 1, F1: 599, F2: 280}
        var v11 int64 = v3
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
