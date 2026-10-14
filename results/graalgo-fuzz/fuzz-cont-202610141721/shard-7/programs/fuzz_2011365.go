package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []int{395}
    v2 := make(map[uint32]bool)
    var v3 int = 2147483647
    v4 := make(map[uint32]uint32)
    v5 := []string{("\n\t" + (("" + "hello") + (("\n\t" + "a") + "a"))), ("" + ("" + "hello"))}
    v4[2] = 140
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 bool = true
        v8 = v8
        _ = fi7
        _ = v8
    }
    fmt.Println(v3)
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int16 = -410
        _ = len(v4)
        v11 := GenS1{F0: false, F1: 1, F2: 2}
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS2{F0: 983}
        var v14 int8 = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v3, v3)
    v4[2147483647] = 101
    v6 = 552
    switch 1 {
    case 128:
        switch (v3 / ((v3) | 1)) {
        case 32768:
            var v15 int32 = 2147483647
            _ = v15
        case -471:
            v3 += (v3 | (v3 * v3))
        default:
            fmt.Println(v6)
        }
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
