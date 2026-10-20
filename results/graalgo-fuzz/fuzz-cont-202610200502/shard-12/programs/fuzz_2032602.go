package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint64 }

func main() {
    var v1 uint8 = 193
    var v2 uint8 = v1
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: -306, F1: 0}
    v5 := GenS1{F0: -302, F1: 1}
    v6 := []int{32768, 494, -654}
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(v2, v1)
        _ = fi7
    }
    fmt.Println(v2, v1)
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS2{F0: 704, F1: -111, F2: 2}
        v10 := GenS2{F0: uint16(v1), F1: 1, F2: 42}
        switch ("a" + "foo bar") {
        case "\n\t":
            delete(v3, 2)
        default:
            var v11 int16 = -32768
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v12 int64 = -9223372036854775808
    var v13 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
