package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }

func main() {
    var v1 int8 = 100
    var v2 uint32 = 0
    var v3 int16 = 1
    v4 := []int{-1, -9223372036854775808, -273, -2147483648}
    if (!(!false)) {
        var v5 float64 = 0.1
        _ = v5
    } else {
        switch -1 {
        case 952:
            if len(v4) > 0 {
                v4[0] = 0
            }
        }
    }
    v6 := GenS1{F0: -2147483648, F1: 540}
    v2 -= (v2 | 678)
    switch "foo bar" {
    case "\n\t":
        v7 := GenS1{F0: -92, F1: 1}
        _ = v7
    default:
        v2 -= v2
    }
    _ = len(v4)
    var v8 uint64 = 18446744073709551615
    v8 += v8
    for fi9 := 0; fi9 < 11; fi9++ {
        v2 -= (v2 | (v2 * v2))
        _ = fi9
    }
    v10 := GenS1{F0: -2147483648, F1: 0}
    fmt.Println(v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, 543)
    for fi11 := 0; fi11 < 20; fi11++ {
        _ = len(v4)
        _ = fi11
    }
    fmt.Println(v2)
    var v12 uint32 = v2
    var v13 uint8 = 139
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}
