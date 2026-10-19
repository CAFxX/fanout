package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }

func main() {
    v1 := []string{"hello", "", "hello", "foo bar"}
    var v2 float64 = 1.5
    var v3 uint64 = 42
    fmt.Println(v3, v2)
    var v4 int8 = -128
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    v4 += v4
    _ = len(v1)
    var v5 int16 = 42
    v3 -= 42
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: 540, F1: (v4 - v4)}
        _ = fi6
        _ = v7
    }
    var v8 int64 = 226
    fmt.Println(v8, v3)
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint = 65535
        v3 -= (v3 % ((v3) | 1))
        var v11 float64 = 3.14159
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 4294967295, F1: (v4 / ((1) | 1))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
