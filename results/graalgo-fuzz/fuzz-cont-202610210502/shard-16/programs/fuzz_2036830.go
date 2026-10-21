package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }

func main() {
    var v1 float32 = float32((3.14159 + 1.5))
    var v2 uint8 = 137
    var v3 float32 = 0.0
    var v4 int16 = 42
    var v5 bool = (3.14159 < 1e10)
    var v6 uint8 = (^(v2))
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 42, F1: fi7}
        switch (^(fi7)) {
        case 786:
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := []int{}
                var v11 uint64 = 0
                v12 := GenS1{F0: 100, F1: (fi9 * fi9)}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        case 2147483647:
            var v13 bool = v5
            _ = v13
        case 127:
            var v14 int8 = 42
            _ = v14
        }
        fmt.Println(v6, v5)
        _ = fi7
        _ = v8
    }
    v15 := GenS1{F0: 100, F1: 128}
    if true {
        v1 = (v3 / v3)
        v16 := []int{}
        _ = v16
    } else {
        v5 = (!v5)
    }
    fmt.Println(v3, v4, v4)
    v17 := GenS1{F0: 0, F1: 128}
    fmt.Println(v3, v1)
    switch -2147483648 {
    case 4294967295:
        var v18 float64 = 1.5
        _ = v18
    default:
        v5 = v5
    }
    var v19 int = 128
    var v20 int64 = -9223372036854775808
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 17; fi22++ {
            v5 = true
            _ = fi22
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            switch fi23 {
            case 433:
                var v24 int16 = (v4 >> ((0) & 63))
                _ = v24
            case 1:
                v4 = (v4 % ((-385) | 1))
            default:
                v25 := GenS1{F0: 1, F1: 1}
                _ = v25
            }
            _ = fi23
        }
        _ = fi21
    }
    if v5 {
        switch "hello" {
        case "\n\t":
            var v26 int16 = (v4 % (((v4 << ((4294967295) & 63))) | 1))
            _ = v26
        }
        fmt.Println(v5, v2)
        v27 := GenS1{F0: -128, F1: v19}
        _ = v27
    }
    v28 := []int{32768, (^(v19))}
    v29 := GenS1{F0: -128, F1: (-(v19))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
