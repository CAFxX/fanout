package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }

func main() {
    var v1 float64 = (float64((1e10 - float32((1.5 * (((float64(3.14159) - 0.0) + -2.25) / 1.0))))) * 1e10)
    var v2 uint8 = 131
    v3 := make(map[uint32]uint32)
    var v4 float32 = ((1e10 - ((0.1 + ((1.5 / ((1e10 / (0.1 + float32((v1 - 0.1)))) * (float32(v1) + (0.1 * 3.14159)))) * (-2.25 + 0.1))) / (3.14159 - (float32((v1 * float64(float32(v1)))) / -2.25)))) / 1e10)
    var v5 uint16 = 1
    var v6 uint = 400
    var v7 int64 = 1
    fmt.Println(v1, v6, v2)
    if false {
        if true {
            var v8 bool = (((!false) && ((true && (!false)) && false)) || true)
            v2 += (v2 % ((v2) | 1))
            _ = v8
        }
    }
    var v9 string = ("" + "foo bar")
    var v10 int32 = 42
    var v11 float64 = ((1.5 / ((v1 - (3.14159 * 1.5)) + float64(((v4 + v4) / 0.1)))) / v1)
    v12 := GenS1{F0: -128, F1: -375, F2: true}
    var v13 uint16 = (v5 | v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
