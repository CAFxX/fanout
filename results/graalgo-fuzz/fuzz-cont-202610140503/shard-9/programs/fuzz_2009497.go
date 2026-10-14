package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }

func main() {
    var v1 uint = 365
    var v2 uint16 = 183
    var v3 int = -348
    v4 := GenS1{F0: 42, F1: (v3 / (((^(v3))) | 1))}
    var v5 uint64 = 18446744073709551615
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := []int{(fi6 << (((v1 >> (((v1 >> ((v1) & 63))) & 63))) & 63)), (fi6 | (fi6 | -1)), 4294967295}
        _ = fi6
        _ = v7
    }
    if true {
        fmt.Println(v1, v5)
        v8 := []int{(v3 >> ((v1) & 63)), (-(v3))}
        if ((2 >= uint64(v2)) && (342 < 2147483648)) {
            var v9 uint64 = 2
            if (!(!(42 > uint32(v2)))) {
                var v10 uint16 = v2
                _ = v10
            }
            _ = v9
        }
        _ = v8
    }
    fmt.Println(v5)
    v1 -= v1
    v11 := []int{}
    v5 -= (v5 << (((^(v1))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
