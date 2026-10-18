package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    v1 := []string{("\n\t" + "\n\t")}
    var v2 float32 = (1e10 - 1.5)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 830
    var v5 uint32 = (v4 << ((2) & 63))
    var v6 uint32 = v5
    v4 += 1
    if (uint8(v5) <= 2) {
        var v7 uint16 = 250
        var v8 int = 2147483647
        _ = v7
        _ = v8
    }
    var v9 float64 = 0.0
    var v10 uint32 = v6
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int32 = 696
        v3[v6] = true
        _ = fi11
        _ = v12
    }
    var v13 float32 = 0.1
    switch 4294967295 {
    case -2147483648:
        v4 = (v10 % (((v5 << ((18446744073709551615) & 63))) | 1))
    }
    v14 := GenS1{F0: -1, F1: 32768}
    _ = len(v1)
    var v15 float64 = v9
    for fi16 := 0; fi16 < 8; fi16++ {
        v1 = append(v1, "a")
        _ = fi16
    }
    var v17 int64 = 42
    var v18 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
