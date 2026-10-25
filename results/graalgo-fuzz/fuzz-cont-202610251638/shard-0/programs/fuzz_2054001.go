package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    var v3 int64 = 441
    v4 := make(map[string]bool)
    var v5 bool = (-2147483648 > -2147483648)
    var v6 int32 = 2147483647
    fmt.Println(v5)
    v7 := GenS1{F0: "\n\t", F1: 585, F2: -676}
    var v8 float64 = (((float64(float32(1.5)) - (3.14159 * 1e10)) / 1.5) / (((float64((3.14159 / (0.1 * (3.14159 - 0.1)))) / ((1.5 / 1e10) / 1.0)) - (3.14159 + 1e10)) * 0.1))
    v3 = v3
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int64 = 587
        var v11 int16 = 550
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int16 = 32767
    fmt.Println(v12, v12, v12)
    _ = len(v2)
    v3 = (v3 + (v3 ^ int64(v6)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
