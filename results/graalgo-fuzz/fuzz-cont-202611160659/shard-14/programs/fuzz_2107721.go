package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }

func main() {
    v1 := []uint32{109}
    var v2 int32 = 0
    var v3 uint16 = 2
    var v4 uint64 = 697
    var v5 int8 = 1
    if (65535 <= 4294967295) {
        v3 = (v3 ^ 65535)
        var v6 string = "a"
        _ = v6
    }
    v7 := GenS1{F0: 740, F1: 9223372036854775807}
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            _ = len(v1)
            _ = fi9
        }
        var v10 uint16 = (v3 & (v3 + (v3 * (v3 ^ 793))))
        _ = fi8
        _ = v10
    }
    v1 = append(v1, 399)
    v2 -= -73
    for fi11 := 0; fi11 < 18; fi11++ {
        _ = len(v1)
        _ = fi11
    }
    v4 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
