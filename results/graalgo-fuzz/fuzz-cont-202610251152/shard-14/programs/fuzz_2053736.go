package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 float32 = 0.0
    var v3 int = -915
    v4 := []int{}
    v5 := GenS1{F0: 2147483647, F1: 880}
    for fi6 := 0; fi6 < 14; fi6++ {
        v4 = append(v4, 42)
        _ = fi6
    }
    v7 := GenS1{F0: (v3 ^ (v3 / (((v3 * 32768)) | 1))), F1: 2147483648}
    v4 = append(v4, v3)
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = (v2 * ((v2 * v2) / 1.0))
    v8 := GenS1{F0: v3, F1: 4294967295}
    switch "\n\t" {
    case "a":
        var v9 float64 = (1.5 / (1.5 * 1e10))
        _ = v9
    default:
        v10 := GenS1{F0: v3, F1: 2}
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
