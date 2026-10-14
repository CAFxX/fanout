package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint; F1 int }

func main() {
    var v1 uint64 = 148
    var v2 uint64 = v1
    var v3 int64 = -1
    if true {
        v3 = 9223372036854775807
    } else {
        var v4 float64 = ((0.0 * float64(1e10)) / float64((3.14159 / (3.14159 - (float32(float64(float32(0.0))) * 1e10)))))
        v5 := GenS1{F0: true, F1: -2147483648}
        _ = v4
        _ = v5
    }
    v3 -= v3
    v6 := []int{-24, 220, 0, -1}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint64 = (v1 | 42)
        v8 = (v2 % ((v8) | 1))
        _ = fi7
        _ = v8
    }
    var v9 uint = 255
    fmt.Println(v1)
    for fi10 := 0; fi10 < 13; fi10++ {
        v9 += v9
        _ = fi10
    }
    var v11 int32 = -584
    v12 := GenS2{F0: 18446744073709551615, F1: -1}
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
