package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int8 }

func main() {
    var v1 float64 = 0.0
    var v2 float32 = 0.1
    var v3 float32 = -2.25
    v4 := []string{"a"}
    v5 := []uint32{2147483648, 0}
    if true {
        var v6 uint64 = 4294967295
        if false {
            v6 += v6
            v7 := GenS1{F0: -125}
            var v8 float32 = (-2.25 - (float32(float64((float32((v1 / 1.0)) * v3))) - 3.14159))
            _ = v7
            _ = v8
        }
        _ = v6
    }
    var v9 bool = (18446744073709551615 <= 18446744073709551615)
    if len(v5) > 0 {
        v5[0] = 0
    }
    if v9 {
        v10 := GenS2{F0: -478, F1: -1}
        _ = v10
    }
    var v11 int32 = -2147483648
    v12 := GenS2{F0: v11, F1: -72}
    v13 := GenS1{F0: -1}
    fmt.Println(v2, v2, v2)
    v2 = v3
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
