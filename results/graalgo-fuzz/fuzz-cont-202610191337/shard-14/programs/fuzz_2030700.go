package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    v1 := []int{}
    var v2 uint16 = 282
    var v3 uint = 18446744073709551615
    var v4 uint8 = 0
    var v5 int32 = int32(v4)
    var v6 float32 = ((-2.25 + -2.25) * (1e10 + 0.0))
    var v7 uint = uint(v4)
    _ = len(v1)
    v8 := GenS1{F0: 69, F1: 4294967295}
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint16 = 1
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v11 uint = 4294967295
        _ = fi9
        _ = v10
        _ = v11
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v6 = 1e10
    for fi12 := 0; fi12 < 5; fi12++ {
        v1 = append(v1, fi12)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi12
    }
    if (false || false) {
        var v13 float64 = float64(1e10)
        var v14 int16 = 623
        _ = v13
        _ = v14
    }
    var v15 int64 = 9223372036854775807
    fmt.Println(v2, v2, v4)
    v5 += 42
    fmt.Println(v6, v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
