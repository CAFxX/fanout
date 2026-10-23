package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }
type GenS2 struct { F0 uint16; F1 uint64; F2 int }

func main() {
    var v1 int8 = 54
    v2 := []uint32{1, 2147483648}
    var v3 bool = true
    fmt.Println(v1, v1)
    switch "hello" {
    case "hello":
        for fi4 := 0; fi4 < 1; fi4++ {
            var v5 float64 = -2.25
            _ = fi4
            _ = v5
        }
    case "":
        v3 = v3
    case "foo bar":
        v6 := GenS1{F0: 255, F1: -2147483648, F2: "\n\t"}
        _ = v6
    }
    switch 520 {
    case -652:
        for fi7 := 0; fi7 < 17; fi7++ {
            v3 = v3
            var v8 float64 = float64(1e10)
            var v9 float32 = (float32((v8 / (v8 + v8))) - 3.14159)
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        if (195 >= 915) {
            v3 = v3
            fmt.Println(v3, v1)
            if len(v2) > 0 {
                v2[0] = 2
            }
        }
    }
    v1 = (v1 | (v1 ^ v1))
    if v3 {
        var v10 uint64 = 42
        _ = len(v2)
        _ = v10
    }
    v3 = v3
    var v11 int = 4294967295
    if v3 {
        var v12 uint64 = 1
        var v13 int = (v11 - v11)
        v14 := GenS2{F0: 751, F1: v12, F2: v13}
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 uint = 42
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 int32 = 511
        var v18 bool = (!v3)
        var v19 int8 = v1
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    _ = len(v2)
    var v20 uint32 = uint32(v15)
    var v21 uint64 = 18446744073709551615
    var v22 int32 = 0
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 int32 = int32(v11)
        _ = fi23
        _ = v24
    }
    var v25 bool = v3
    v26 := GenS2{F0: 255, F1: (^(v21)), F2: (v11 ^ -1)}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
}
