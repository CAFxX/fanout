package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 int = -9223372036854775808
    var v2 float64 = 3.14159
    v3 := make(map[string]uint32)
    var v4 uint32 = 42
    var v5 int32 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: false, F1: v1}
        v8 := GenS1{F0: false, F1: (v1 * 42)}
        v5 -= v5
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int = 9223372036854775807
    fmt.Println(v2, v5, v2)
    v10 := v3[("" + (("" + "hello") + "hello"))]
    fmt.Println(v5, v9)
    for fi11 := 0; fi11 < 9; fi11++ {
        fmt.Println(v4, v2, v2)
        _ = fi11
    }
    fmt.Println(v9, v10, v10)
    v2 += float64(-2.25)
    v4 -= v4
    var v12 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
