package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{583, 4294967295, 0}
    var v3 float64 = 1.5
    var v4 float32 = 1e10
    var v5 int64 = 689
    var v6 uint = 683
    switch "hello" {
    case "\n\t":
        switch 0 {
        case 9223372036854775807:
            v5 = v5
        case 127:
            var v7 float32 = 1e10
            _ = v7
        case -1:
            fmt.Println(v4, v5, v3)
        }
    default:
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                v5 -= (v5 & (v5 + (^(v5))))
                if len(v2) > 0 {
                    v2[0] = 847
                }
                _ = fi9
            }
            delete(v1, 0)
            _ = fi8
        }
    }
    var v10 float64 = float64((v4 * v4))
    v4 = (float32(-2.25) / 0.1)
    v11 := GenS1{F0: 42, F1: "hello"}
    v12 := GenS1{F0: 65535, F1: "a"}
    fmt.Println(v3, v4, v6)
    v3 = v10
    fmt.Println(v5, v5)
    if len(v2) > 0 {
        v2[0] = 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
