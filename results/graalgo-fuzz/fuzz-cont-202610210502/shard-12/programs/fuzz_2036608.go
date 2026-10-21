package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint64 }
type GenS2 struct { F0 int32; F1 int8 }

func main() {
    var v1 uint16 = 737
    v2 := []int{42, -2147483648}
    var v3 bool = false
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(v3, fi5)
        var v6 float32 = 1e10
        var v7 int8 = -128
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        v3 = (2147483647 == 42)
        _ = fi8
    }
    _ = len(v2)
    v3 = ((127 > uint8(v1)) || v3)
    v1 -= v1
    if v3 {
        var v9 uint64 = 75
        _ = v9
    }
    var v10 int32 = 140
    switch "\n\t" {
    case "\n\t":
        switch 42 {
        case -9223372036854775808:
            switch "\n\t" {
            case "\n\t":
                var v11 int16 = 464
                _ = v11
            case "foo bar":
                var v12 int64 = -9223372036854775808
                _ = v12
            }
        }
    case "hello":
        v4[42] = -917
    case "":
        var v13 uint64 = 0
        _ = v13
    default:
        fmt.Println(v3, v1)
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v3 = (!v3)
        v3 = (v3 && v3)
        _ = fi14
    }
    var v15 int = -2147483648
    v16 := GenS1{F0: 128, F1: -481, F2: 553}
    v17 := GenS2{F0: v10, F1: 77}
    var v18 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
