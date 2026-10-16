package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 float32 = ((float32(0.0) * (1.5 / (1e10 - 3.14159))) / 3.14159)
    var v2 int32 = 449
    var v3 int16 = -396
    var v4 float32 = 3.14159
    v5 := make(map[string]string)
    var v6 float32 = float32(0.1)
    v2 += v2
    v7 := GenS1{F0: 1, F1: 255}
    var v8 uint8 = 0
    var v9 float64 = -2.25
    v3 -= (-(v3))
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(v8, v6)
        _ = fi10
    }
    fmt.Println(v4, v9, v2)
    if (!true) {
        fmt.Println(v3, v8)
        var v11 string = ("" + "a")
        v9 = (-2.25 / v9)
        _ = v11
    }
    v1 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
