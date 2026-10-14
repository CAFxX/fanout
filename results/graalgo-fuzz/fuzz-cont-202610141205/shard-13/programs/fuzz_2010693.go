package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint64 }
func genFn0(p0 bool) uint {
    p0 = (p0 && (18446744073709551615 == 8))
    var v4 int16 = 1000
    _ = p0
    _ = v4
    return 0
}

func main() {
    var v1 float64 = 0.1
    v2 := []uint32{2}
    var v3 bool = false
    v5 := GenS1{F0: 128, F1: 303, F2: 449}
    fmt.Println(v3)
    fmt.Println(v3, v3)
    var v6 int8 = 42
    v7 := GenS1{F0: 127, F1: 0, F2: 655}
    v1 += ((v1 - v1) - ((((v1 - v1) + -2.25) + v1) - 1.5))
    for fi8 := 0; fi8 < 3; fi8++ {
        if (fi8 == fi8) {
            fmt.Println(v1, v6, v1)
        }
        var v9 bool = false
        _ = fi8
        _ = v9
    }
    v1 = 0.0
    if v3 {
        v10 := GenS1{F0: 0, F1: 910, F2: 0}
        _ = v10
    }
    v3 = v3
    var v11 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
