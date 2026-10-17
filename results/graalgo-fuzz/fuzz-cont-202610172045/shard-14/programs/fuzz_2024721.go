package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int16; F1 string }

func main() {
    var v1 int64 = 42
    var v2 uint16 = 1
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    switch -2147483648 {
    case 198:
        var v5 int8 = -77
        _ = v5
    case 0:
        v6 := GenS1{F0: 938, F1: -839}
        _ = v6
    default:
        v7 := []int{898, 128, 1, 0}
        _ = v7
    }
    var v8 int = 127
    var v9 int16 = 32767
    v4["foo bar"] = ""
    var v10 int64 = 167
    v11 := GenS2{F0: v9, F1: "a"}
    var v12 int = (v8 | -2147483648)
    var v13 uint32 = 4294967295
    var v14 uint32 = 65
    v2 = 2
    switch (v8 / ((-724) | 1)) {
    case -1:
        fmt.Println(v14, v2)
    }
    fmt.Println(v1)
    var v15 float32 = 0.0
    var v16 float64 = (1e10 * (((0.1 / ((0.1 / -2.25) + 0.0)) - (float64((v15 - v15)) / 0.1)) * -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
