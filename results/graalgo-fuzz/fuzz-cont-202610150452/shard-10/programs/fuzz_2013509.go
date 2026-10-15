package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int16 }

func main() {
    var v1 int32 = 0
    v2 := []string{}
    var v3 uint16 = 32767
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 bool = true
        var v6 uint64 = 449
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint16 = v3
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v3 = (v7 | v3)
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    if false {
        var v8 int8 = -1
        v3 += v7
        _ = v8
    }
    _ = len(v2)
    if true {
        var v9 float32 = (3.14159 + (float32(((0.0 * 1e10) / -2.25)) / 1e10))
        v1 -= -883
        _ = v9
    }
    var v10 int8 = -1
    switch "\n\t" {
    case "\n\t":
        v11 := GenS1{F0: 1, F1: true, F2: 1000}
        _ = v11
    default:
        var v12 uint8 = 138
        _ = v12
    }
    v13 := GenS1{F0: 451, F1: (255 != 255), F2: 42}
    v14 := GenS1{F0: 42, F1: (0.1 < (0.0 - float64(-2.25))), F2: -325}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
}
