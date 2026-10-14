package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }

func main() {
    var v1 int = 464
    var v2 int16 = int16(v1)
    var v3 int32 = -578
    var v4 int8 = 127
    var v5 uint16 = 1
    v3 -= 323
    v5 += (v5 % ((715) | 1))
    fmt.Println(v4, v3, v3)
    var v6 float64 = (1e10 / (((1e10 - (0.0 + -2.25)) + float64(float32((0.1 * (0.1 / -2.25))))) * 1e10))
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 663, F1: v1}
        fmt.Println(v4, v1)
        _ = fi7
        _ = v8
    }
    v6 = (v6 - v6)
    var v9 string = "hello"
    fmt.Println(v6, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
