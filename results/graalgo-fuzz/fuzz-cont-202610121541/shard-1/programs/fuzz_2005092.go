package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0() int8 {
    var v4 float32 = 0.1
    var v5 int = -9223372036854775808
    v6 := GenS1{F0: 819, F1: 468}
    _ = v4
    _ = v5
    _ = v6
    return -128
}

func main() {
    var v1 uint64 = 372
    var v2 float64 = ((((0.0 - 0.0) / float64((3.14159 * ((float32(-2.25) * -2.25) / (1e10 / 1e10))))) + (3.14159 + (1.5 - 0.0))) * (-2.25 + -2.25))
    var v3 uint64 = v1
    var v7 uint = 2
    var v8 float64 = v2
    v8 -= v2
    v2 += 0.1
    v9 := GenS1{F0: v7, F1: 42}
    v10 := []int{128, 984, 127}
    v11 := GenS1{F0: (v7 / (((^(v7))) | 1)), F1: 0}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v12 := GenS1{F0: 775, F1: 2147483648}
    if true {
        fmt.Println(v2, v3, v7)
    } else {
        _ = len(v10)
        v13 := GenS2{F0: (v3 >> (((^(v7))) & 63))}
        _ = v13
    }
    if (42 < 0) {
        var v14 float32 = 0.0
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
