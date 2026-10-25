package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }

func main() {
    var v1 int = 145
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int = 127
        _ = fi4
        _ = v5
    }
    fmt.Println(v1)
    fmt.Println(v1)
    var v6 float64 = 3.14159
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint64 = 0
        if false {
            if false {
                if (1 != uint32(v8)) {
                    _ = len(v3)
                }
            }
        }
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (v1 ^ v1), F1: 4294967295, F2: "foo bar"}
    if (83 == -32) {
        var v10 float64 = 0.0
        var v11 int64 = -9223372036854775808
        _ = v10
        _ = v11
    } else {
        fmt.Println(v6, v1)
    }
    v12 := []int{(v1 << ((2) & 63)), (v1 & (v1 * v1)), -2147483648, (-(v1))}
    v1 += -407
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
