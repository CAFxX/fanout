package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    var v1 uint64 = 721
    var v2 bool = true
    var v3 int8 = 1
    var v4 uint8 = 129
    fmt.Println(v3, v2, v2)
    var v5 float32 = 0.1
    v3 = (v3 ^ (v3 << ((42) & 63)))
    v4 = v4
    var v6 string = "hello"
    var v7 uint16 = 133
    for fi8 := 0; fi8 < 17; fi8++ {
        if v2 {
            fmt.Println(fi8, v6, v5)
        }
        _ = fi8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := []int{-1, fi9, (^(fi9))}
        var v11 float64 = 0.0
        if v2 {
            var v12 uint = 255
            v13 := GenS1{F0: 65535, F1: (fi9 & fi9)}
            var v14 int32 = -391
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v15 := []int{9223372036854775807}
    fmt.Println(v5, v4, v7)
    v15 = append(v15, 32768)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
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
}
