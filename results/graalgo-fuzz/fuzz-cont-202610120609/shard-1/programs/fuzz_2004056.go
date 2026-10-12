package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint8 }
func genFn0(p0 uint64, p1 string) uint32 {
    v5 := GenS1{F0: 643, F1: 0, F2: 26}
    p0 = (^(p0))
    v6 := []int{53, -9223372036854775808, -556, 2147483647}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 910
}

func main() {
    var v1 int32 = 42
    var v2 int32 = 42
    var v3 float32 = 3.14159
    var v4 uint32 = 2147483648
    fmt.Println(v2)
    if ((0.1 > ((((1e10 * 1e10) / -2.25) / 1.5) - 1e10)) && false) {
        var v7 int16 = 662
        v1 += v1
        _ = v7
    } else {
        v4 -= (v4 & v4)
        v2 += v1
        var v8 uint16 = 65535
        _ = v8
    }
    var v9 float32 = v3
    v9 -= float32(((float64(-2.25) * ((((1.5 * (0.0 + -2.25)) - 1.5) + float64((0.0 + ((-2.25 - v3) * v3)))) * 1e10)) / 1.5))
    v10 := GenS1{F0: 4294967295, F1: -107, F2: 42}
    v1 += v1
    fmt.Println(v1, v4, v1)
    switch 4294967295 {
    case 0:
        var v11 uint = 42
        _ = v11
    }
    var v12 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
