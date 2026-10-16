package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    var v3 int32 = 2147483647
    v4 := []int{32767}
    var v5 int8 = 100
    switch 908 {
    case 32767:
        v6 := GenS1{F0: 0, F1: (989 <= 42), F2: 233}
        _ = v6
    }
    fmt.Println(v5, v5)
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(v5)
        var v8 uint16 = 0
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint32 = 42
            var v11 int8 = 0
            v12 := GenS1{F0: -9223372036854775808, F1: (42 > -1), F2: 630}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    var v13 int8 = v5
    v3 = (v3 + (^(v3)))
    for fi14 := 0; fi14 < 9; fi14++ {
        if false {
            _ = len(v1)
            var v15 float64 = 0.1
            _ = v15
        } else {
            fmt.Println(v13, v13, v13)
            v4 = append(v4, 437)
            v3 = (v3 * v3)
        }
        fmt.Println(v13, v13, fi14)
        v13 = v5
        _ = fi14
    }
    var v16 uint32 = 1
    v16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
