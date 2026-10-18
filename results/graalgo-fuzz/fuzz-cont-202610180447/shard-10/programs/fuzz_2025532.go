package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{}
    var v3 int = 82
    var v4 uint64 = 1
    v5 := []string{"foo bar", "", "\n\t", ""}
    if (false || (!false)) {
        v3 = v3
        v6 := GenS1{F0: 24, F1: 0}
        var v7 uint = 4294967295
        _ = v6
        _ = v7
    } else {
        fmt.Println(v4, v4)
    }
    var v8 uint = 2
    switch v3 {
    case 1:
        v1 -= (v1 << (((v8 - 381)) & 63))
    }
    v4 -= 4294967295
    fmt.Println(v3)
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint8 = uint8(v8)
        _ = fi9
        _ = v10
    }
    v8 = 42
    fmt.Println(v8, v4, v8)
    if len(v2) > 0 {
        v2[0] = 276
    }
    var v11 uint16 = 303
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
