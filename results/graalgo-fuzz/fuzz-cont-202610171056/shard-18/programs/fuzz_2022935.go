package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 157
    var v3 int64 = 491
    var v4 bool = true
    v2 += v2
    v5 := GenS1{F0: v3, F1: -570, F2: ""}
    fmt.Println(v3)
    var v6 float32 = (1e10 * 1.5)
    v4 = false
    v3 -= v3
    var v7 uint64 = 42
    switch -809 {
    case 502:
        var v8 float32 = ((((float32(1e10) + (float32(-2.25) * v6)) / v6) / 0.1) * v6)
        _ = v8
    default:
        if true {
            if (-137 == -304) {
                var v9 float32 = 1.5
                delete(v1, 4294967295)
                _ = v9
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
