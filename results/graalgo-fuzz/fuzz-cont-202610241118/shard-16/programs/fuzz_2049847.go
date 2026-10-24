package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }

func main() {
    var v1 bool = (!false)
    var v2 int16 = 447
    var v3 uint8 = 42
    if v1 {
        var v4 uint16 = 0
        _ = v4
    }
    v5 := GenS1{F0: 42, F1: -9223372036854775808}
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{127, fi6, int(v2)}
        v2 = (v2 << ((1) & 63))
        for fi8 := 0; fi8 < 6; fi8++ {
            v7 = append(v7, (^(fi6)))
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    v9 := []int{127}
    if true {
        var v10 int64 = 9223372036854775807
        _ = v10
    } else {
        var v11 int32 = -2147483648
        _ = v11
    }
    var v12 uint64 = 18446744073709551615
    v2 = 1000
    switch ("a" + "hello") {
    case "\n\t":
        v1 = ((v2 & (v2 << ((42) & 63))) == (^(v2)))
    default:
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
