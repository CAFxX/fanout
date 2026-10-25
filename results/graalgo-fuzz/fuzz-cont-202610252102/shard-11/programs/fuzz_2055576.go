package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    var v1 float64 = -2.25
    v2 := []string{"hello"}
    var v3 int = 2147483647
    var v4 uint32 = 666
    v5 := []uint32{(v4 >> ((18446744073709551615) & 63))}
    var v6 float64 = v1
    v7 := GenS1{F0: 0}
    if (0 <= 318) {
        fmt.Println(v3, v3, v1)
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            v3 = v3
            v10 := GenS1{F0: -128}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v3 -= -9223372036854775808
    v3 = 1
    v11 := GenS2{F0: 127, F1: 119}
    var v12 uint32 = v4
    var v13 uint8 = 0
    var v14 uint = 2
    var v15 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
