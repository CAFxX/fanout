package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 string }

func main() {
    var v1 int64 = 42
    v2 := make(map[string]bool)
    var v3 int32 = 2147483647
    var v4 float32 = float32(-2.25)
    v5 := []uint32{}
    v6 := []uint32{42, 611}
    fmt.Println(v1, v1)
    v1 += (v1 | 1)
    var v7 string = "\n\t"
    v8 := GenS1{F0: 18446744073709551615, F1: false, F2: v7}
    delete(v2, v7)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 float64 = float64(v4)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(v4, v4)
        _ = fi9
        _ = v10
    }
    v6 = append(v6, 4294967295)
    _ = len(v2)
    if false {
        var v11 string = "\n\t"
        v2["a"] = false
        _ = v11
    }
    if false {
        var v12 uint16 = 359
        var v13 int64 = v1
        _ = v12
        _ = v13
    }
    v3 += (-(v3))
    for fi14 := 0; fi14 < 8; fi14++ {
        v2[("foo bar" + (v7 + "hello"))] = true
        _ = fi14
    }
    v4 -= v4
    var v15 float64 = (((0.1 + (3.14159 / (1e10 / (3.14159 * 1.5)))) - 1.5) + (0.1 / 1.0))
    _ = len(v5)
    switch -1 {
    case 2147483647:
        var v16 uint = 2
        _ = v16
    case -824:
        fmt.Println(v15)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
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
