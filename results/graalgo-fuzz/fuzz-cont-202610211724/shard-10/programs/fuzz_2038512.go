package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 bool }
func genFn0() uint32 {
    var v4 float64 = ((1e10 / ((float64((0.1 / 3.14159)) / 1.5) - 0.0)) * 1.5)
    v4 = ((0.0 * v4) / 3.14159)
    v5 := GenS1{F0: (2 >= 18446744073709551615), F1: 42, F2: false}
    _ = v4
    _ = v5
    return 605
}
func genFn1(p0 uint8, p1 int8, p2 int64) bool {
    var v6 int = 4294967295
    p1 = -128
    v7 := GenS1{F0: false, F1: 571, F2: (!(true || (-1 < -746)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ((1e10 / ((0.1 / 1.0) + -2.25)) <= 0.0)
}

func main() {
    var v1 float32 = (((3.14159 + 0.1) - 1.5) * (0.1 + 3.14159))
    var v2 int = 0
    var v3 int32 = 42
    fmt.Println(v3)
    v8 := []int{v2, 696, (v2 << ((493) & 63))}
    fmt.Println(v3)
    v9 := GenS1{F0: true, F1: 65535, F2: false}
    fmt.Println(v3, v1)
    if (false || ((0.1 * (v1 - (0.0 * v1))) >= 0.0)) {
        if len(v8) > 0 {
            v8[0] = (-(v2))
        }
        v10 := GenS1{F0: false, F1: 18446744073709551615, F2: true}
        _ = v10
    } else {
        var v11 uint64 = 2
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = v11
    }
    var v12 uint16 = 0
    var v13 float64 = (float64(v1) - 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
