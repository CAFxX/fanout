package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 4294967295
    var v3 uint16 = 32767
    v4 := []int{2147483647, -2147483648, 9223372036854775807, 32767}
    var v5 bool = (!false)
    var v6 float32 = (3.14159 + 0.1)
    v6 = (1e10 / v6)
    v7 := GenS1{F0: 32767, F1: 100}
    v3 += 0
    var v8 uint16 = 931
    v8 -= (v3 % ((v3) | 1))
    fmt.Println(v3, v8, v2)
    var v9 float64 = 3.14159
    v10 := GenS1{F0: 32767, F1: -59}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    delete(v1, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
