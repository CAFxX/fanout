package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
func genFn0(p0 string) bool {
    v5 := []int{2147483647, -2147483648}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return true
}

func main() {
    var v1 int8 = -128
    var v2 int16 = -32768
    var v3 float32 = (((-2.25 - (((0.0 * (1e10 / 1.0)) + -2.25) / (0.1 + (((float32(float64(1.5)) + 0.1) + (0.1 - 0.0)) + 0.1)))) * 0.1) + 0.0)
    var v4 uint16 = 2
    var v6 bool = false
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int32 = 407
        _ = fi7
        _ = v8
    }
    var v9 float64 = 3.14159
    if (false || false) {
        if ((2 <= 42) && ((v1 & 127) < v1)) {
            fmt.Println(v4)
            fmt.Println(v2, v6, v2)
            if v6 {
                v10 := []int{-695}
                _ = v10
            }
        } else {
            fmt.Println(v3)
        }
        v11 := []int{-2147483648, 32767}
        _ = v11
    } else {
        v12 := GenS1{F0: 215, F1: 0}
        _ = v12
    }
    v13 := []int{}
    fmt.Println(v2)
    if v6 {
        v6 = true
    }
    v9 = (0.0 + (v9 - (1e10 / 1.5)))
    v14 := GenS1{F0: 2, F1: 12}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
