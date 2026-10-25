package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }

func main() {
    var v1 uint16 = 0
    var v2 uint = 0
    var v3 int64 = -9223372036854775808
    var v4 bool = (!(0.1 > (0.1 + -2.25)))
    var v5 float32 = -2.25
    var v6 uint = v2
    v4 = v4
    fmt.Println(v4, v2, v3)
    v5 = 1.5
    fmt.Println(v2)
    v7 := GenS1{F0: v6, F1: 0, F2: 9223372036854775807}
    v8 := []int{42, 9223372036854775807, 2147483647, 326}
    v9 := GenS1{F0: v6, F1: 18, F2: 127}
    for fi10 := 0; fi10 < 20; fi10++ {
        v2 = v6
        _ = fi10
    }
    if len(v8) > 0 {
        v8[0] = 32768
    }
    var v11 float64 = ((0.0 + 3.14159) - float64((((0.1 * -2.25) * v5) + 0.1)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
