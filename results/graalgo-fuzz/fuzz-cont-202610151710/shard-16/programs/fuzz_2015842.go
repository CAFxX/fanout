package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }

func main() {
    v1 := []string{}
    var v2 float64 = (float64(float32(float64(-2.25))) + (1e10 * (0.0 * -2.25)))
    v3 := make(map[int]string)
    var v4 uint32 = 4294967295
    var v5 uint16 = 257
    var v6 int8 = 42
    v4 += v4
    _ = len(v3)
    v4 = (^(v4))
    v3[32768] = "\n\t"
    v5 = 65535
    var v7 bool = ((true && false) || (!((((0.0 - ((-2.25 / 1.0) * (1e10 / (3.14159 / ((0.1 / 1.5) - float32(((v2 * 1e10) / v2))))))) + float32(v2)) - 0.0) >= (3.14159 + 0.1))))
    v4 -= (v4 ^ (v4 / ((v4) | 1)))
    v8 := GenS1{F0: 2147483647, F1: -198}
    if v7 {
        if true {
            v2 = -2.25
            fmt.Println(v6)
            var v9 int8 = (v6 ^ 100)
            _ = v9
        }
    }
    var v10 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
