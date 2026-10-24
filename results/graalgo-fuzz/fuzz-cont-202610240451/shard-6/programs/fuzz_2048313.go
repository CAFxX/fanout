package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint32 = 2
    var v3 uint64 = 523
    var v4 int8 = 1
    var v5 float64 = (1.5 - (3.14159 * float64((((0.0 - 3.14159) + 3.14159) + 1e10))))
    v2 -= 4294967295
    v6 := GenS1{F0: v4, F1: (v3 << ((4294967295) & 63)), F2: (2 > 1)}
    switch "foo bar" {
    case "\n\t":
        var v7 int = 42
        _ = v7
    case "":
        fmt.Println(v2)
    case "a":
        for fi8 := 0; fi8 < 5; fi8++ {
            v1 = -9223372036854775808
            var v9 bool = false
            _ = fi8
            _ = v9
        }
    default:
        v10 := []int{}
        _ = v10
    }
    var v11 uint64 = v3
    var v12 uint8 = 99
    switch -9223372036854775808 {
    case -1:
        var v13 int16 = -161
        _ = v13
    case 4294967295:
        var v14 uint64 = 1
        _ = v14
    case -9223372036854775808:
        v1 = (v1 | (v1 % (((-(v1))) | 1)))
    }
    v15 := []int{-2147483648, 32767, 127, 4294967295}
    fmt.Println(v3, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
