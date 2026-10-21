package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0() uint32 {
    var v5 float32 = 1e10
    v6 := GenS1{F0: 100}
    v7 := []int{32768, 32768}
    _ = len(v7)
    _ = v5
    _ = v6
    _ = v7
    return 2147483648
}
func genFn1(p0 int) uint {
    v8 := []int{(p0 % ((p0) | 1)), p0, (p0 + p0), (-(p0))}
    p0 -= p0
    for fi9 := 0; fi9 < 6; fi9++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
    }
    _ = p0
    _ = v8
    return 4294967295
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]string)
    var v3 int32 = 576
    var v4 uint32 = 869
    for fi10 := 0; fi10 < 17; fi10++ {
        fmt.Println(v3)
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS1{F0: 42}
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v3 -= -2147483648
    fmt.Println(v1, v3, v1)
    v13 := []int{1, 127, -9223372036854775808}
    if (!false) {
        v14 := GenS1{F0: 127}
        _ = v14
    }
    if (!true) {
        v15 := v2[865]
        v1 = (v1 + 3.14159)
        _ = v15
    } else {
        var v16 float32 = (3.14159 - (float32((1.5 / ((1e10 + -2.25) - (0.0 / ((0.1 + -2.25) - v1))))) * 0.0))
        var v17 uint8 = 145
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 1, F1: genFn1(32768)}
    v1 = float64(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
