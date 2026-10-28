package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }

func main() {
    var v1 int16 = 91
    var v2 int32 = 42
    var v3 float64 = 0.0
    var v4 int16 = 1000
    v5 := []int{-793}
    _ = len(v5)
    v1 -= v1
    v6 := GenS1{F0: 2147483648, F1: 4294967295}
    v4 -= v1
    if (1 == 1) {
        for fi7 := 0; fi7 < 19; fi7++ {
            v5 = append(v5, fi7)
            v2 = 42
            _ = fi7
        }
        v5 = append(v5, -884)
    }
    var v8 int16 = v4
    v8 = (v1 ^ (v1 ^ v8))
    fmt.Println(v4, v8)
    _ = len(v5)
    switch 32767 {
    case -2147483648:
        var v9 bool = false
        _ = v9
    case 127:
        v10 := GenS1{F0: 277, F1: 165}
        _ = v10
    case 4294967295:
        var v11 uint = 0
        _ = v11
    default:
        var v12 int8 = -128
        _ = v12
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        fmt.Println(v8)
    default:
        var v13 uint32 = 4294967295
        _ = v13
    }
    fmt.Println(v2, v1)
    if (797 > 0) {
        var v14 int32 = (v2 - v2)
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
