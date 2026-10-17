package main

import (
    "fmt"
)

func genFn0(p0 uint) uint8 {
    var v6 int16 = -782
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 128
}

func main() {
    var v1 int64 = 359
    v2 := make(map[int]int)
    v3 := []uint32{2147483647}
    var v4 bool = true
    var v5 uint64 = 4294967295
    var v7 uint8 = 127
    switch "\n\t" {
    case "foo bar":
        for fi8 := 0; fi8 < 8; fi8++ {
            v5 += (v5 | (v5 / ((973) | 1)))
            v4 = true
            _ = fi8
        }
    case "a":
        _ = len(v2)
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v1 = (v1 * (v1 % ((9223372036854775807) | 1)))
        _ = fi9
    }
    v1 += (v1 - v1)
    v1 -= v1
    fmt.Println(v5, v4, v1)
    if false {
        fmt.Println(v7, v7, v5)
        var v10 uint = 4294967295
        v4 = true
        _ = v10
    }
    if v4 {
        fmt.Println(v7, v7, v7)
        fmt.Println(v1)
    } else {
        var v11 int32 = int32(v7)
        _ = v11
    }
    var v12 int16 = 130
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
