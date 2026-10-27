package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }

func main() {
    var v1 int32 = 88
    var v2 int32 = (v1 / (((^(v1))) | 1))
    var v3 int16 = -32768
    var v4 uint32 = 694
    v4 = 436
    v5 := GenS1{F0: 65535, F1: 42, F2: true}
    v2 = 0
    var v6 uint64 = 502
    fmt.Println(v3, v1)
    v7 := GenS1{F0: 2, F1: 18446744073709551615, F2: false}
    fmt.Println(v4)
    v8 := []int{0, 9223372036854775807}
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 float32 = 0.1
        var v11 uint16 = 165
        _ = fi9
        _ = v10
        _ = v11
    }
    if (127 < 0) {
        var v12 string = "a"
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
