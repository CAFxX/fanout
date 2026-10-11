package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }

func main() {
    v1 := []string{}
    var v2 float64 = 0.0
    v3 := []uint32{2147483648, 2147483648, 4294967295}
    _ = len(v3)
    var v4 float64 = v2
    v1 = append(v1, ("\n\t" + "a"))
    var v5 int64 = 415
    _ = len(v1)
    v4 = (0.0 / v4)
    v4 += v2
    var v6 uint8 = 2
    v7 := GenS1{F0: 9223372036854775807, F1: 1}
    var v8 uint32 = 2147483648
    for fi9 := 0; fi9 < 16; fi9++ {
        v2 += (v4 * v2)
        var v10 int = -1
        _ = fi9
        _ = v10
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
