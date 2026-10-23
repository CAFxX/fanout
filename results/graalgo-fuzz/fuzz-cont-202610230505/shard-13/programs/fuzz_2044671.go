package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    var v1 float64 = 1.5
    var v2 uint = 65535
    var v3 uint = (v2 % ((v2) | 1))
    var v4 int8 = 127
    var v5 float64 = (((v1 / ((v1 * (v1 * -2.25)) * v1)) + v1) / v1)
    v6 := make(map[string]string)
    v1 += float64((0.0 + float32(((v1 - float64(1.5)) + v1))))
    v6["\n\t"] = "hello"
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{-1, -1, fi8}
        v10 := v6[""]
        _ = fi8
        _ = v10
        _ = v9
    }
    v6["a"] = "hello"
    v11 := GenS1{F0: 155, F1: v4}
    v7 -= (v7 | 32767)
    v4 = v4
    v1 -= v5
    v1 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
