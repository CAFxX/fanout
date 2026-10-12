package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int8 }

func main() {
    var v1 int64 = -1
    var v2 uint64 = 18446744073709551615
    var v3 int32 = -2147483648
    var v4 int = 978
    v5 := GenS1{F0: 157, F1: (v2 >> ((70) & 63)), F2: -1}
    var v6 uint32 = 564
    var v7 int32 = -1
    for fi8 := 0; fi8 < 13; fi8++ {
        v4 += (v4 ^ v4)
        _ = fi8
    }
    v9 := []int{}
    v6 = v6
    var v10 int16 = 107
    var v11 uint32 = 2147483648
    if false {
        v12 := GenS1{F0: 128, F1: (v2 * (^(v2))), F2: 127}
        if ((v4 >> ((22) & 63)) >= 4294967295) {
            v6 = 2147483647
            v1 -= 42
        }
        var v13 int8 = -96
        _ = v12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
