package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 int64 = 42
    var v2 uint32 = 272
    v3 := []int{349}
    var v4 uint32 = (v2 + v2)
    v5 := GenS1{F0: 1, F1: 1}
    v4 -= (^(v4))
    v6 := GenS1{F0: 830, F1: 2}
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    v2 = (v4 ^ 0)
    _ = len(v3)
    var v7 int16 = 42
    switch "a" {
    case "a":
        fmt.Println(v4)
    case "hello":
        for fi8 := 0; fi8 < 2; fi8++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v4 -= 4294967295
            _ = fi8
        }
    case "\n\t":
        fmt.Println(v2)
    default:
        var v9 int32 = 2147483647
        _ = v9
    }
    var v10 float64 = float64(3.14159)
    v11 := GenS1{F0: 1, F1: 18446744073709551615}
    var v12 uint16 = 778
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
